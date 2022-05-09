import socketIoStore from "svelte-socketio-store";
console.log("ws:"+document.location.host+"/sync")
const initialValue = { };
export const mainsocket = socketIoStore(("ws://"+document.location.host+"/sync"), initialValue);