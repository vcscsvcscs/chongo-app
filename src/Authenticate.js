export function UserLogedin() {
    let sessiontoken = window.sessionStorage.getItem("sessiontoken");
    if(sessiontoken === null){
        sessionStorage.removeItem('sessiontoken');
        return false
    }else{
        return sessiontoken;
    }
}
export function Register(name,username,mail,password){
        let baseUrl = window.location.origin;
        let data = {
            username: username,
            name: name,
            email: mail,
            password: password,
        };
        let xhr = new XMLHttpRequest();
        xhr.open("POST", baseUrl + "/register", true);
        xhr.setRequestHeader("Content-Type", "application/json");
        xhr.onreadystatechange = function () {
            if (this.readyState != 4) return;
            if (this.status == 200) {
                data = JSON.parse(this.responseText);
                window.sessionStorage.setItem("sessiontoken", data.token);
                console.log(data.token)
                window.location.href = baseUrl + "/home";
            } else {
                data = JSON.parse(this.responseText);
                return data.message;
            }
            // end of state change: it can be after some time (async)
        };
        //console.log(JSON.stringify(data));
        xhr.send(JSON.stringify(data));
}