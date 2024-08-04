let list = document.querySelector('.humb ul'); 
document.querySelector(".fas").addEventListener('click', () => {
    if (list.style.display == "none") {
        list.style.display = "flex"
    } else {
        list.style.display = "none"
    }
})
