import classes from '../css/index.less'

const Homepage = require('./homepage')
const Display = require('./display')
const Settings = require('./settings')

export default() => {
  const homepage = new Homepage()
  const display = new Display()
  const settings = new Settings()

  window.onresize = display.invokeResize

  let homePageButton = document.getElementById("homePageButton");
  if (homePageButton) {
    homePageButton.onclick = () => {
      display.show("homePage")
      homepage.getList()
    }

    homePageButton.click()
  }

  let newPageButton = document.getElementById("newPageButton")
  if (newPageButton) {
    newPageButton.onclick = () => {
      display.show("newPage")
    }
  }

  let settingsButton = document.getElementById("settingsButton")
  if (settingsButton) {
    settingsButton.onclick = () => {
      display.show("settingsPage")
    }
  }

  settings.checkSavedSettings((err, saved) => {
    if (err) {
      return external.invoke("error:checkSaveSettings|" + err)
    }

    if (saved) {
      homePageButton.click()
    } else {
      display.show('settings')
      settings.showPage()
    }
  })
}