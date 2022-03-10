package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/vcscsvcscs/chongo-app/backend/controllers"
	"github.com/vcscsvcscs/chongo-app/backend/controllers/accounts"
	"github.com/vcscsvcscs/chongo-app/backend/controllers/chat"
	"github.com/vcscsvcscs/chongo-app/backend/sessionmanager"
	"github.com/vcscsvcscs/chongo-app/backend/utilities"
	"gopkg.in/mgo.v2"
)

var (
	cert      = flag.String("cert", "./private/keys/cert.pem", "Specify the path of TLS cert")
	key       = flag.String("key", "./private/keys/key.pem", "Specify the path of TLS key")
	httpsPort = flag.String("https", ":443", "Specify port for http secure hosting(example for format :443)")
	httpPort  = flag.String("http", ":80", "Specify port for http hosting(example for format :80)")
	release   = flag.Bool("release", false, "Set true to release build")
)

var MongoDB *mgo.Session

func main() {
	flag.Parse()
	if *release {
		gin.SetMode(gin.ReleaseMode)
		// Logging to a file.
		gin.DisableConsoleColor() // Disable Console Color, you don't need console color when writing the logs to file.
		path := fmt.Sprintf("private/logs/%02dy_%02dm_%02dd_%02dh_%02dm_%02ds.log", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())
		//fmt.Println(path)
		f, logerror := os.Create(path)
		if logerror != nil {
			log.Println("Cant log to file")
		} else {
			gin.DefaultWriter = io.MultiWriter(f)
			// Use the following code if you need to write the logs to file and console at the same time.
			// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
		}
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(gin.DefaultErrorWriter)
	var err error
	MongoDB, err = mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		log.Println(err)
		syscall.Exit(503)
	}
	go controllers.InnitCredentials(MongoDB, "chongo", "users")
	go sessionmanager.InnitSessions(15, MongoDB, "chongo", "sessions")
	//Router and endpoints
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./public", true)))
	router.POST("/register", accounts.Register)
	router.POST("/login", accounts.Login)
	router.POST("/logout", accounts.Logout)
	router.DELETE("/deleteaccount", accounts.DeleteAcc)
	router.GET("/chat/*any", chat.ChatSocket)
	//Server configuration
	var server *http.Server
	if utilities.Exists(*cert) && utilities.Exists(*key) {
		server = &http.Server{
			Addr:         *httpsPort,
			Handler:      router,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		go func() {
			log.Println("Server starts at port ", *httpsPort)
			if err := server.ListenAndServeTLS(*cert, *key); err != nil && errors.Is(err, http.ErrServerClosed) {
				log.Printf("%s\n", err)
			}
		}()
		// Start the HTTP server and redirect all incoming connections to HTTPS
		go log.Fatal(http.ListenAndServe(*httpPort, http.HandlerFunc(controllers.RedirectToHttps)))
	} else {
		server = &http.Server{
			Addr:         *httpPort,
			Handler:      router,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		go func() {
			log.Println("Server starts at port ", *httpPort)
			if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
				log.Printf("%s\n", err)
			}
		}()
	}

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
