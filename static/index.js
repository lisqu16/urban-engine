// init
function $(h) {
    return document.querySelector(h);
}

const ws = new WebSocket("ws://localhost:8080/gateway");
ws.onerror = (err) => console.error(`gamingn't: ${err}`);
ws.onopen = () => console.log("gaming!!!!");

ws.onmessage = (msg) => {
    const pld = JSON.parse(msg.data);
    switch (pld.op) {
        case 1:
            console.log(pld.msg)
            const h = document.createElement("h1");
            h.innerHTML = pld.msg;
            $("#chat").appendChild(h);
            break;
    }
}

// main
$("#send").onclick = () => {
    const pralka = $("#chatinput").value;
    const payload = {
        op: 1,
        msg: pralka,
    };
    $("#chatinput").value = "";
    ws.send(JSON.stringify(payload));
};