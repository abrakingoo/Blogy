let list = document.querySelector(".humb ul");
document.querySelector(".fas").addEventListener("click", () => {
  if (list.style.display == "none") {
    list.style.display = "flex";
  } else {
    list.style.display = "none";
  }
});

// error message function
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

document.addEventListener("DOMContentLoaded", (event) => {
  let linkContainers = document.querySelectorAll(".categs");

  linkContainers.forEach((container) => {
    container.addEventListener("click", (event) => {
      let link = container.querySelector("a");
      let url = link.getAttribute("href");

      // Remove "active" class from all link containers
      linkContainers.forEach((container) => {
        container.classList.remove("active");
      });

      // Add "active" class to the clicked link container
      container.classList.add("active");
      history.pushState(null, "", url);
    });
  });
});

///editor
document.addEventListener("DOMContentLoaded", function () {
  var toolbarOptions = [
    ["bold", "italic", "underline", "strike"],
    ["blockquote", "code-block"],
    [{ header: 1 }, { header: 2 }],
    [{ list: "ordered" }, { list: "bullet" }],
    [{ script: "sub" }, { script: "super" }],
    [{ indent: "-1" }, { indent: "+1" }],
    [{ direction: "rtl" }],
    [{ size: ["small", false, "large", "huge"] }],
    [{ header: [1, 2, 3, 4, 5, 6, false] }],
    [{ color: [] }, { background: [] }],
    [{ font: ["Arial"] }],
    [{ align: [] }],
    ["clean"],
  ];

  var quill = new Quill("#editor-container", {
    modules: {
      toolbar: toolbarOptions,
    },
    theme: "snow",
  });

  // Function to get Quill content and set it to hidden input
  document.querySelector(".post-form").addEventListener("submit", function () {
    var hiddenInput = document.getElementById("hidden-input");
    hiddenInput.value = quill.root.innerHTML;
  });
});
