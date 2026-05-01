async function login() {
  const username = document.getElementById("username").value;
  const password = document.getElementById("password").value;

  if (!username || !password) {
    alert("Username dan password tidak boleh kosong");
    return;
  }

  const response = await fetch("/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      username: username,
      password: password,
    }),
  });

  // Cek status dulu sebelum parse JSON
  if (!response.ok) {
    const errorText = await response.text();
    if (response.status === 401) {
      alert("Username atau password salah");
    } else {
      alert("Login gagal: " + errorText);
    }
    return;
  }

  const data = await response.json();

  if (data.token) {
    localStorage.setItem("token", data.token);
    window.location.href = "/dashboard-page";
  } else {
    alert("Login gagal, token tidak ditemukan");
  }
}
