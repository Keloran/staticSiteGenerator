const Axios = require('axios')
const Display = require('../display')

class Settings {
  showPage() {
    let display = new Display()
    display.show("settingsPage")
  }

  checkSavedSettings(callback) {
    Axios({
      method: 'get',
      url: '/settings'
    }).then((response) => {
      return callback(false, response.saved)
    }).catch((error) => {
      return callback(error)
    })
  }
}

module.exports = Settings