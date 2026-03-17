async function loadDashboard() {
  const token = localStorage.getItem("token");

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
