import './style.css'
import Navbar from './Navbar.svelte'

const app = new Navbar({
  target: document.getElementById('app')
})

export default app
