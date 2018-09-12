const Axios = require('axios')
const Article = require('../article')

class HomePage {
  addItemToList(item) {
    let article = new Article()
    let liveTab = document.getElementById("liveList")
    let listItem, listItemText, listItemContent
    let itemIcon = "file"

    if (item.icon) {
      itemIcon = item.icon;
    }

    listItem = document.createElement("li");

    listItemContent = document.createElement("i");
    listItemContent.className = "articleItem fas fa-" + itemIcon;
    listItemContent.itemId = item.id;
    listItemContent.onclick = article.openItem;

    listItemText = document.createTextNode(" " + item.title);
    listItemContent.appendChild(listItemText);
    listItem.appendChild(listItemContent)

    liveTab.appendChild(listItem);
  }

  getList() {
    this.removeAll()

    Axios({
      method: 'get',
      url: '/list'
    }).then((response) => {
      let item
      for (let i = 0; i < response.data.length; i++) {
        item = response.data[i]
        this.addItemToList(item)
      }
    })
  }

  removeAll() {
    let liveList = document.getElementById("liveList")
    liveList.innerHTML = ""
  }
}

module.exports = HomePage