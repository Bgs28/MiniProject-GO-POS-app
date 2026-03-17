async function login() {
  const username = document.getElementById("username").value;
  const password = document.getElementById("password").value;

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

  const data = await response.json();

  if (data.token) {
    localStorage.setItem("token", data.token);

    window.location.href = "/dashboard-page";
  } else {
    alert("login gagal");
  }
}
