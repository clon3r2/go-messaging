let socketUri = 'ws://127.0.0.1:8000/socket/';

var ws = new WebSocket(socketUri)

ws.onmessage = function(evt) {
    const out = document.getElementById('output');
    out.innerHTML += 'new msg = ' + evt.data + '<br>';

}

ws.onclose = (s) => {
    console.log("close !!")
    setInterval(()=>{ws.send("ping")},1000)
}

setInterval(sendMessage, 1000);

function sendMessage () {
    if (ws.readyState === WebSocket.OPEN){
        ws.send('Hello, Server!');
    } else {
        console.log("connection is closed. cant send message")
        console.log(`staatus = ${ws.readyState}`)
    }
}