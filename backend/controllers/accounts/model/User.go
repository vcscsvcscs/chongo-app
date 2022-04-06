package model

// Only required to break import cycle

type User struct {
	Username       string `json:"username" bson:"username"`
	Name           string `json:"name" bson:"name"`
	Email          string `json:"email" bson:"email"`
	Password       string `json:"password" bson:"password"`
	Ip             string `json:"ip" bson:"ip"`
	DeletedAccount int64  `json:"deleted" bson:"deleted"`
}
