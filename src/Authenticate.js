export function UserLogedin() {
    let sessiontoken = window.sessionStorage.getItem("sessiontoken");
    if(sessiontoken === null){
        sessionStorage.removeItem('sessiontoken');
        return false
    }else{
        return sessiontoken;
    }
}