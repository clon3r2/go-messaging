let socketUri = 'ws://127.0.0.1:8080/socket/';
let ws = new WebSocket(socketUri);

ws.onmessage = function(evt) {
    console.log("got a new msg ==> ", evt)
    const out = document.getElementById('msgOutput');
    out.innerHTML += `<div class="card-body text-center text-bold">${evt.data}</div>`;

}

ws.onclose = () => {
    console.log("closed !!")
    setInterval(()=>{
        if (ws.readyState !== WebSocket.OPEN){
        ws = new WebSocket(socketUri);} else {
            console.log("connected again!")
        }
        clearInterval(this.id)
    },1000)
}

function SendMessage() {
    inputElem = document.getElementById("msgInput")

    if (inputElem.value !== "") {
        ws.send(inputElem.value)
    }
}