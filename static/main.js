const deleteLinks = document.querySelectorAll(".btn-delete")
const linkTexts = document.querySelectorAll(".in-link")

let toast = new Notyf({
    position: {x: "center", y: "top"},
    dismissibleL: true
})

document.addEventListener("DOMContentLoaded", () => {
    deleteLinks.forEach(link => {
        link.addEventListener("click", async (e) => {
            e.preventDefault()
            // Use currentTarget to get the <a> element, not the clicked child (<i>)
            await fetch(`${e.currentTarget.href}`, {
                method: "DELETE"
            })
            toast.success(`Link deleted`)
            setTimeout(() => {
                history.go(0)

            }, 1000)

        })
    })

    linkTexts.forEach(link => {
        link.addEventListener("click", async (e) => {
            e.preventDefault()
            try {
                await navigator.clipboard.writeText(e.target.textContent)
                toast.success(`Link is copied to clipboard`)
            } catch (error) {
                console.error(error);
                toast.success(`Failed to copy link to clipboard`)
            }
        })
    })
})