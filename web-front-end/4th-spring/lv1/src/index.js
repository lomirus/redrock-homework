import './styles/index.css'
import { moveCallback, movePromise, moveAsync } from './move.js'

const buttons = document.querySelectorAll('button')

buttons[0].addEventListener('click', moveCallback)
buttons[1].addEventListener('click', movePromise)
buttons[2].addEventListener('click', moveAsync)
