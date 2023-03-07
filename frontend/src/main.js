import './style.css'
import NavBar from './Navbar.svelte'

const app = new NavBar({
  target: document.getElementById('app')
})

export default app
