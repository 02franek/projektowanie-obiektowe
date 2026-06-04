const express = require("express");
const cors = require("cors");

const app = express();
const PORT = 3001;

app.use(cors());
app.use(express.json());

const products = [
  {
    id: 1,
    name: "Klawiatura RGB",
    price: 300,
    description: "Świeci się na wszystkie kolory tęczy",
  },
  {
    id: 2,
    name: "Myszka ergonomiczna",
    price: 500,
    description: "Śmiesznie się ją trzyma, tak pionowo",
  },
  {
    id: 3,
    name: "Słuchawki bezprewodowe",
    price: 100,
    description: "9 na 10 programisów je poleca",
  },
];

app.get("/api/products", (req, res) => {
  console.log("/api/products received GET request ");
  res.json(products);
});

app.post("/api/payments", (req, res) => {
  const { amount } = req.body;

  console.log("/api/payments received POST request");

  if (amount && amount >= 0) {
    res
      .status(200)
      .json({ success: true, message: "Płatność zrealizowana pomyślnie" });
  } else {
    res.status(400).json({ success: false, message: "Błąd płatności" });
  }
});

app.listen(PORT, () => {
  console.log(`Server availalbe at http://localhost:${PORT}`);
});
