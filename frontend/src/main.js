import "./style.css"
import "./libs/Link.svelte"
import "./libs/locale/locale.js"
import Navbar from "./Navbar.svelte"

const app = new Navbar({
  target: document.getElementById("app")
})

export default app
