let socketUri = 'ws://127.0.0.1:8081/socket/';
let ws = makeConnection()
let inputElem = document.getElementById("msgInput")
let sendBtn = document.getElementById("send-btn")
sendBtn.disabled=true
inputElem.addEventListener("input", ()=>{
    sendBtn.disabled = inputElem.value === ""
})


function makeConnection(outputElem) {
    let conn = new WebSocket(socketUri);
    conn.onmessage = (evt) => {
        console.log("got a new msg ==> ", evt)
        outputElem = document.getElementById("msgOutput")
        outputElem.innerHTML += `<div class="card-body text-center text-bold bg-primary mb-2">${evt.data}</div>`;
    }
    conn.onclose = () => {
        console.log("connection lost !")
        //TODO: auto-reconnect before send new msg somehow
    }
    return conn
}

function SendMessage() {
    console.log("azaval")
    if (ws.readyState === WebSocket.CLOSED) {
        console.log("no connection available, created new connection.")
        ws = makeConnection()
    }
    if (ws.readyState !== WebSocket.CONNECTING) {
        ws.send(inputElem.value)
        console.log("msg sent successfully.")
        inputElem.value = ""
        sendBtn.disabled = true
    } else {
        console.log("connection pending, retrying ... ")
        setTimeout(SendMessage, 500)
    }
}
