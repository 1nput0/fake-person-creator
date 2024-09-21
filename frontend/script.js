document
  .getElementById("generateButton")
  .addEventListener("click", function () {
    fetch("http://localhost:8080/api/male-person")
      .then((response) => response.json())
      .then((data) => {
        document.getElementById(
          "name"
        ).textContent = `${data.FirstName} ${data.LastName}`;
        document.getElementById("email").textContent = data.Email;
        document.getElementById("country").textContent = data.Country;
        document.getElementById("state").textContent = data.State;
        document.getElementById("address").textContent = data.Address;

        document.getElementById("result").classList.remove("hidden");
      })
      .catch((error) => console.error("Error", error));
  });
