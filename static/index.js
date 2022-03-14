const ws = new WebSocket("ws://localhost:8080/ws");
ws.onerror = (err) => console.error(`gamingn't: ${err}`);
ws.onopen = () => console.log("gaming!!!!");