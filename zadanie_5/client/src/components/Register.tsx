import { useState } from "react";

const Register = () => {
  const [formData, setFormData] = useState({
    username: "",
    email: "",
    password: "",
  });

  const [errors, setErrors] = useState({
    username: "",
    email: "",
    password: "",
  });

  const [successMessage, setSuccessMessage] = useState("");

  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
    setErrors((prev) => ({ ...prev, [name]: "" }));
  };

  const validateForm = () => {
    let isValid = true;
    const newErrors = { username: "", email: "", password: "" };

    if (!formData.username.trim()) {
      newErrors.username = "Nazwa użytkownika jest wymagana.";
      isValid = false;
    }

    if (!formData.email.trim()) {
      newErrors.email = "Adres e-mail jest wymagany.";
      isValid = false;
    } else if (!emailRegex.test(formData.email)) {
      newErrors.email = "Wprowadzono niepoprawny format adresu e-mail.";
      isValid = false;
    }

    if (!formData.password.trim()) {
      newErrors.password = "Hasło jest wymagane.";
      isValid = false;
    }

    setErrors(newErrors);
    return isValid;
  };

  const handleSubmit = (e: React.SubmitEvent) => {
    e.preventDefault();
    setSuccessMessage("");

    if (validateForm()) {
      setSuccessMessage("Rejestracja przebiegła pomyślnie!");
      setFormData({ username: "", email: "", password: "" });
    }
  };

  return (
    <div style={{ maxWidth: "400px", margin: "0 auto", padding: "20px" }}>
      <h2>Rejestracja</h2>

      {successMessage && (
        <div
          style={{ color: "green", marginBottom: "15px", fontWeight: "bold" }}
        >
          {successMessage}
        </div>
      )}

      <form onSubmit={handleSubmit} noValidate>
        <div style={{ marginBottom: "15px" }}>
          <label htmlFor="username">Nazwa użytkownika:</label>
          <br />
          <input
            type="text"
            id="username"
            name="username"
            value={formData.username}
            onChange={handleInputChange}
            style={{ width: "100%", padding: "8px" }}
          />
          {errors.username && (
            <span style={{ color: "red", fontSize: "14px" }}>
              {errors.username}
            </span>
          )}
        </div>

        <div style={{ marginBottom: "15px" }}>
          <label htmlFor="email">Adres e-mail:</label>
          <br />
          <input
            type="email"
            id="email"
            name="email"
            value={formData.email}
            onChange={handleInputChange}
            style={{ width: "100%", padding: "8px" }}
          />
          {errors.email && (
            <span style={{ color: "red", fontSize: "14px" }}>
              {errors.email}
            </span>
          )}
        </div>

        <div style={{ marginBottom: "15px" }}>
          <label htmlFor="password">Hasło:</label>
          <br />
          <input
            type="password"
            id="password"
            name="password"
            value={formData.password}
            onChange={handleInputChange}
            style={{ width: "100%", padding: "8px" }}
          />
          {errors.password && (
            <span style={{ color: "red", fontSize: "14px" }}>
              {errors.password}
            </span>
          )}
        </div>

        <button
          type="submit"
          style={{ padding: "10px 20px", cursor: "pointer" }}
        >
          Zarejestruj się
        </button>
      </form>
    </div>
  );
};

export default Register;
