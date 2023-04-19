import "./style.css"
import "./libs/Link.svelte"
import "./libs/locale/locale"
import Navbar from './Navbar.svelte'

/* 
 * index.html imports this file. 
 * Start svelte here.
 */
const app = new Navbar({
  target: document.getElementById("app")!
})

export default app
