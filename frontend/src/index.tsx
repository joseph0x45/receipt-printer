/* @refresh reload */
import { render } from 'solid-js/web'
import App from './App.tsx'
import { Router, Route } from "@solidjs/router"
import Test from './pages/Test.tsx'
import "./index.css"

const root = document.getElementById('root')

render(() => (
  <Router>
    <Route path="/" component={App} />
    <Route path="/test" component={Test} />
  </Router>
), root!)
