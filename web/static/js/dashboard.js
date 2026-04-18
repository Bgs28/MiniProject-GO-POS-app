async function loadDashboard() {
  const token = localStorage.getItem("token");

  // if token unavailable
  if (!token) {
    window.location.href = "/login-page";
  }

  const response = await fetch("/dashboard", {
    headers: {
      Authorization: "Bearer " + token,
    },
  });

  const data = await response.json();

  document.getElementById("products").innerText = data.total_product;
  document.getElementById("transactions").innerText = data.total_transactions;
  document.getElementById("sales").innerText = data.today_sales;
}

loadDashboard();

// logout
document.addEventListener("DOMContentLoaded", function () {
  document.getElementById("btnLogout").addEventListener("click", function () {
    localStorage.removeItem("token");

    window.location.href = "/login-page";
  });
});
