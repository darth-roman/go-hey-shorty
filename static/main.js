const deleteLinks = document.querySelectorAll(".btn-delete")

document.addEventListener("DOMContentLoaded", () => {
    deleteLinks.forEach(link => {
        link.addEventListener("click", async (e) => {
            e.preventDefault()
            // Use currentTarget to get the <a> element, not the clicked child (<i>)
            await fetch(`${e.currentTarget.href}`, {
                method: "DELETE"
            })

            history.go(0)
        })
    })
})