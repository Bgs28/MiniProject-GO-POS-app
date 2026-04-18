async function loadProducts() {
  const token = localStorage.getItem("token");
  if (!token) {
    window.location.href("/login-page");
  }

  const res = await fetch("/product-page", {
    headers: {
      Authorization: "Bearer " + token,
    },
  });

  const products = await res.json();

  const tbody = document.querySelector("#productTable tbody");

  products.forEach((p) => {
    const row = `
        <tr>
        <td>${p.name}</td>
        <td>${p.price}</td>
        <td>${p.stock}</td>
        <td>
        <button class="btn btn-warning btn-sm" onclick="editProduct(${p.id}, ${p.name}, ${p.price}, ${p.stock})">Edit</button>
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

function editProduct(id, name, price, stock) {
  document.getElementById("modalTitle").innerText = "Edit Product";

  document.getElementById("ProductId").value = id;
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
    price: price,
    stock: stock,
  };

  if (id) {
    // UPDATE
    await fetch("/products/" + id, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + token,
      },
      body: JSON.stringify(payLoad),
    });
  } else {
    //   CREATE
    await fetch("/products/" + id, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + token,
      },
      body: JSON.stringify(payLoad),
    });
  }
  closeModal();
  loadProducts();
}

async function deleteProduct(id) {
  const token = localStorage.getItem("token");

  await fetch("/products/" + id, {
    method: "DELETE",
    headers: {
      Authorization: "Bearer " + token,
    },
  });

  loadProducts();
}

function closeModal() {
  document.getElementById("productModal").style.display = "none";
}

loadProducts();
