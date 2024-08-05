let list = document.querySelector(".humb ul");
document.querySelector(".fas").addEventListener("click", () => {
  if (list.style.display == "none") {
    list.style.display = "flex";
  } else {
    list.style.display = "none";
  }
});

document.addEventListener("DOMContentLoaded", function () {
  // Check if the script has run before in this session
  if (!sessionStorage.getItem("scriptRan")) {
    // Automatically load the "Trending" category
    fetchCategory("trending");

    // Mark the script as having run
    sessionStorage.setItem("scriptRan", "true");
  }
});

function fetchCategory(categoryName) {
  fetch(`/category?category=${categoryName}`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((response) => response.text())
    .then((html) => {
      // Update the content of the page
      document.querySelector(".pages").innerHTML = html;
    })
    .catch((error) => {
      console.error("Error fetching category:", error);
    });
}

function ShowErrorMessage() {
  let message = document.querySelector(".login-error");

  // Ensure the message exists before attempting to show it
  if (message) {
    message.style.display = "inline-block";

    // Hide the message after 3 seconds (3000 milliseconds)
    setTimeout(() => {
      message.style.display = "none";
    }, 3000);
  }
}
