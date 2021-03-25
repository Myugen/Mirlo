import React from 'react'
import ReactDOM from 'react-dom'
import './index.css'
import axios from 'axios'
import Signup from 'src/routes/signup'
import { GlobalStore } from 'src/stores/global'
import reportWebVitals from './reportWebVitals'

axios.defaults.baseURL = 'http://localhost:8080'

ReactDOM.render(
  <React.StrictMode>
    <GlobalStore>
      <Signup />
    </GlobalStore>
  </React.StrictMode>,
  document.getElementById('root')
)

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals()
