/* @refresh reload */
import { render } from 'solid-js/web'
import { Router, Route } from "@solidjs/router"
import "./index.css"
import Auth from './pages/Auth.tsx'

const root = document.getElementById('root')

render(() => (
  <Router>
    <Route path="/" component={Auth} />
  </Router>
), root!)
