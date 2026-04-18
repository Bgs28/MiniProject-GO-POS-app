let cart = [];

// load products

async function loadProducts() {
  const token = localStorage.getItem("token");

  if (!token) {
    alert("Please login first");
    window.location.href = "/login-page";
  }

  const res = await fetch("/products", {
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
        <button class="add-btn"
        data-id="${p.id}"
        data-name="${p.name}"
        data-price="${p.price}"
        >
        Add
        </button>
        </td>
        </tr>
        `;

    tbody.innerHTML += row;
  });
}

// add products to cart

function addToCart(id, name, price) {
  id = Number(id);
  const item = cart.find((i) => i.product_id === id);

  if (item) {
    item.quantity++;
  } else {
    cart.push({
      id: id,
      name: name,
      price: price,
      quantity: 1,
    });
  }

  renderCart();
}

// render Cart

function renderCart() {
  const tbody = document.querySelector("#cartTable tbody");

  tbody.innerHTML = "";

  let total = 0;

  cart.forEach((item) => {
    const subtotal = item.price * item.quantity;
    total += subtotal;

    const row = `
<tr>
<td>${item.name}</td>

<td>
<button onclick="decreaseQty(${item.id})">-</button>
${item.quantity}
<button onclick="increaseQty(${item.id})">+</button>
</td>

<td>${item.price}</td>
<td>${subtotal}</td>

<td>
<button onclick="removeItem(${item.id})">Remove</button>
</td>

</tr>
`;

    tbody.innerHTML += row;
  });

  document.getElementById("total").innerText = total;
}

// checkout transaction

async function checkout() {
  const token = localStorage.getItem("token");

  const payload = {
    user_id: 1,
    items: cart.map((i) => ({
      product_id: i.id,
      quantity: i.quantity,
    })),
  };

  await fetch("/transaction", {
    method: "POST",

    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + token,
    },

    body: JSON.stringify(payload),
  });

  alert("Transaction Success");
  printReceipt();
  cart = [];

  renderCart();
}

// add product qty
function increaseQty(id) {
  id = Number(id);
  const item = cart.find((i) => i.id === id);

  if (!item) return;
  item.quantity++;

  renderCart();
}

// decrease product qty

function decreaseQty(id) {
  id = Number(id);
  const item = cart.find((i) => i.id === id);

  if (!item) return;
  item.quantity--;

  if (item.quantity <= 0) {
    cart = cart.filter((i) => i.id !== id);
  }

  renderCart();
}

// rmeove item in cart list
function removeItem(id) {
  id = Number(id);
  cart = cart.filter((i) => i.id !== id);

  renderCart();
}

// Print Receipt
function printReceipt() {
  let receipt = "POS SYSTEM\n\n";

  cart.forEach((item) => {
    const subtotal = item.price * item.quantity;

    receipt += `${item.name}  ${item.quantity}  ${subtotal}\n`;
  });

  receipt += "\n-------------------\n";
  receipt += "TOTAL: " + document.getElementById("total").innerText;

  const win = window.open("", "", "width=300,height=400");

  win.document.write("<pre>" + receipt + "</pre>");

  win.print();
}

loadProducts();

document.addEventListener("click", function (e) {
  if (e.target.classList.contains("add-btn")) {
    const id = Number(e.target.dataset.id);
    const name = e.target.dataset.name;
    const price = Number(e.target.dataset.price);

    addToCart(id, name, price);
  }
});
