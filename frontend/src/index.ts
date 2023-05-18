import "/src/style.css"
import "/src/libs/Link.svelte"
import "/src/libs/locale"
import Navbar from "./Navbar.svelte"

/* 
 * index.html imports this file. 
 * Start svelte here.
 */
const app = new Navbar({
  target: document.getElementById("app")!
})

export default app
