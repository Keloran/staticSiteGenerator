class Display {
  hideAll() {
    let shown = document.getElementsByClassName("visable")
    let show
    for (let i = 0; i < shown.length; i++) {
      show = shown[i]
      show.className = "hidden"
    }
  }

  show(elemId) {
    this.hideAll()

    let elem = document.getElementById(elemId)
    elem.className = "visable"
  }

  invokeResize(event) {
    external.invoke("windowSize:" + window.innerWidth + "|" + window.innerHeight)
  }
}

module.exports = Display