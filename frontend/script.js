document
  .getElementById("generateButton")
  .addEventListener("click", function () {
    const gender = document.getElementById("genderSelect").value;
    let apiUrl = "http://localhost:8080/api/";

    switch (gender) {
      case "male":
        apiUrl += "male-person";
        break;
      case "female":
        apiUrl += "female-person";
        break;
      default:
        apiUrl += "random-person";
    }

    fetch(apiUrl)
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
