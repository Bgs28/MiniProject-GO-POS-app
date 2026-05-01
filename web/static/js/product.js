async function loadProducts() {
  const token = localStorage.getItem("token");
  if (!token) {
    window.location.href = "/login-page";
    return;
  }

  const res = await fetch("/products", {
    headers: {
      Authorization: "Bearer " + token,
    },
  });

  if (!res.ok) {
    console.error("Gagal fetch products, status:", res.status);
    return;
  }

  const products = await res.json();
  console.log("Jumlah produk:", products.length);

  const tbody = document.querySelector("#productTable tbody");
  tbody.innerHTML = "";

  products.forEach((p) => {
    console.log("Render produk:", p.name); // cek apakah forEach jalan
    const row = `
      <tr>
        <td>${p.name}</td>
        <td>${p.price}</td>
        <td>${p.stock}</td>
        <td>
          <button class="btn btn-warning btn-sm"
          data-id = "${p.id}"
          data-name ="${p.name}"
          data-price ="${p.price}"
          data-stock ="${p.stock}"
          onclick="editProduct(this)">Edit</button>
          <button class="btn btn-danger btn-sm" onclick="deleteProduct(${p.id})">Delete</button>
        </td>
      </tr>
    `;
    tbody.innerHTML += row;
  });
}
function openAddForm() {
  document.getElementById("modalTitle").innerText = "Add Product";

  document.getElementById("productId").value = "";
  document.getElementById("name").value = "";
  document.getElementById("price").value = "";
  document.getElementById("stock").value = "";

  document.getElementById("productModal").style.display = "block";
}

function editProduct(btn) {
  document.getElementById("modalTitle").innerText = "Edit Product";

  const id = btn.getAttribute("data-id");
  const name = btn.getAttribute("data-name");
  const price = btn.getAttribute("data-price");
  const stock = btn.getAttribute("data-stock");

  document.getElementById("productId").value = id;
  document.getElementById("name").value = name;
  document.getElementById("price").value = price;
  document.getElementById("stock").value = stock;

  document.getElementById("productModal").style.display = "block";
}

async function saveProduct() {
  const id = document.getElementById("productId").value;
  const name = document.getElementById("name").value;
  const price = document.getElementById("price").value;
  const stock = document.getElementById("stock").value;

  const token = localStorage.getItem("token");

  const payLoad = {
    name: name,
    price: parseInt(price),
    stock: parseInt(stock),
  };

  if (id) {
    await fetch("/products?id=" + id, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + token,
      },
      alert: alert("Data Updated"),
      body: JSON.stringify(payLoad),
    });
  } else {
    // CREATE
    await fetch("/products", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + token,
      },
      alert: alert("data added"),
      body: JSON.stringify(payLoad),
    });
  }
  closeModal();
  loadProducts();
}

async function deleteProduct(id) {
  const token = localStorage.getItem("token");

  await fetch("/products?id=" + id, {
    method: "DELETE",
    headers: {
      Authorization: "Bearer " + token,
    },
  });

  alert("Data deleted");
  loadProducts();
}

function closeModal() {
  document.getElementById("productModal").style.display = "none";
}

loadProducts();
