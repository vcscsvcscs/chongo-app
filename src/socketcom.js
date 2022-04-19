import websocketStore from "svelte-websocket-store";
console.log("ws:"+document.location.host+"/sync")
const initialValue = { };
export const mainsocket = websocketStore(("ws://"+document.location.host+"/sync"), initialValue);