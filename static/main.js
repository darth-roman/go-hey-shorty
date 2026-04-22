const deleteLinks = document.querySelectorAll(".btn-delete")

document.addEventListener("DOMContentLoaded", () => {
    deleteLinks.forEach(link => {
        link.addEventListener("click", async (e) => {
            e.preventDefault()
            let shortLinkID = e.target.dataset.id
            // console.log(e.target.dataset.id);
            await fetch(`/shorten/${shortLinkID}`, {
                method: "DELETE"
            })

            history.go(0)
        })
    })
})