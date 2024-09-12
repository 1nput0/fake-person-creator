document
  .getElementById("generateButton")
  .addEventListener("click", function () {
    fetch("http://localhost:8080/api/male-person")
      .then((response) => response.json())
      .then((data) => {
        const resultDiv = document.getElementById("result");
        resultDiv.innerHTML = `
                <h2>${data.FirstName}</h2>
                <h2>${data.LastName}</h2>
                <h2>${data.Email}</h2>
                <h2>${data.Country}</h2>
                <h2>${data.State}</h2>
                <h2>${data.Address}</h2>
            `;
      })
      .catch((error) => console.error("Error", error));
  });
