import classes from '../css/index.css'

const axios = require('axios')

function invokeFullScreen(event) {
  let fullscreen = document.getElementById("fullScreen")
  if (fullscreen.getAttribute("full") == "full") {
    external.invoke("window")
    fullscreen.setAttribute("full", "window")
  }

  external.invoke("fullscreen")
  fullscreen.setAttribute("full", "full")
}

function invokeClose(event) {
  external.invoke("close")
}

function invokeInfo(event) {
  external.invoke("info")
}

function invokeResize(event) {
  external.invoke("windowSize:" + window.innerWidth + "|" + window.innerHeight)
}

function invokeSave() {
  let content = document.getElementById("content")

  axios({
    method: 'post',
    url: '/parse',
    data: content.value
  }).then((response) => {
    let render = document.getElementById("render")
    render.innerHTML = response.data
  })
}

export default() => {
  window.onresize = invokeResize

  let fullScreen = document.getElementById("fullScreen")
  if (fullScreen) {
    fullScreen.onclick = invokeFullScreen
  }

  let close = document.getElementById("close")
  if (close) {
    close.onclick = invokeClose
  }

  let info = document.getElementById("info")
  if (info) {
    info.onclick = invokeInfo
  }

  let save = document.getElementById("save")
  if (save) {
    save.onclick = invokeSave
  }
}