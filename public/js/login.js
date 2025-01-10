document.getElementById("loginForm").addEventListener("submit", async (e) => {
    e.preventDefault();

    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    try {
        const response = await fetch("http://localhost:3000/auth/loginUser", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ email, password }),
            credentials: "include", // Penting untuk mengirim dan menerima cookie
        });

        if (response.ok) {
            const data = await response.json();
            alert("Login successful: " + data.message);
            window.location.href = "/protected/dashboard";
        } else {
            const error = await response.json();
            alert("Login failed: " + error.error);
        }
    } catch (err) {
        console.error("Error:", err);
        alert("Something went wrong!");
    }
});
