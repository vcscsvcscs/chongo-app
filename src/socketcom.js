import websocketStore from "svelte-websocket-store";
console.log()
const initialValue = { };
export const mainsocket = websocketStore("wss:"+document.location.host+"", initialValue);

// send JSON to websocket server
//$myStore = { content: "to be saved", other_values: "all" };

// receive JSON from server (push)
//let response = $myStore;