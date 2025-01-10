// Ambil token dari cookie
console.log("Cookies : ", document.cookie);
const token = document.cookie.replace("token=", "");
console.log("Token from cookie : ", token);
let username = "user1"; // Username default jika token tidak ada

// Jika ada token, decode untuk mendapatkan username (email)
if (token) {
    try {
        const decodedToken = jwt_decode(token);  // Decode JWT
        console.log("Decoded Token : ", decodedToken);
        if(decodedToken.Username){
            username = decodedToken.Username
        }else if(decodedToken.Email){
            username = decodedToken.Email
        }
    } catch (err) {
        console.error("Error decoding JWT:", err);
    }
}else{
    alert("Please login first...");
}

console.log("Using username: ", username); // Debugging username yang digunakan

// Menghubungkan ke WebSocket dengan username dinamis
const socket = new WebSocket("ws://localhost:3000/chat/ws?username=" + username + "&token=" + token);

socket.onopen = function() {
    console.log("WebSocket connection established as " + username);
};

socket.onmessage = function(event) {
    const data = JSON.parse(event.data);
    const messages = document.getElementById("messages");
    const messageElement = document.createElement("div");
    messageElement.textContent = `${data.from}: ${data.message}`;
    messages.appendChild(messageElement);
};

socket.onerror = function(error) {
    console.error("WebSocket Error: " + error);
};

// Kirim pesan ke WebSocket
document.getElementById("send-button").addEventListener("click", () => {
    const input = document.getElementById("message-input");
    if (input.value.trim() !== "") {
        socket.send(JSON.stringify({ message: input.value }));
        input.value = "";  // Kosongkan input setelah mengirim pesan
    }
});
