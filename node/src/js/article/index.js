const Axios = require('axios')

class Article {
  openItem(event) {
    console.log("itemId", event.target.itemId)
  }

  invokeSave() {
    let content = document.getElementById("content");

    Axios({
      method: 'post',
      url: '/parse',
      data: content.value
    }).then((response) => {
      let render = document.getElementById("render");
      render.innerHTML = response.data
    })
  }
}

module.exports = Article