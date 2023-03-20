import "./style.css"
import Navbar from "./Navbar.svelte"
import "./libs/Link.svelte"
import "./libs/locale/locale.js"

const app = new Navbar({
  target: document.getElementById("app")
})

export default app
