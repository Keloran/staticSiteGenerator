import classes from '../css/index.less'

const axios = require('axios');

function invokeResize(event) {
  external.invoke("windowSize:" + window.innerWidth + "|" + window.innerHeight)
}

function invokeSave() {
  let content = document.getElementById("content");

  axios({
    method: 'post',
    url: '/parse',
    data: content.value
  }).then((response) => {
    let render = document.getElementById("render");
    render.innerHTML = response.data
  })
}

function openItem(event) {
  console.log("itemId", event.target.itemId)
}

function getList() {
  axios({
      method: 'get',
      url: '/list'
  }).then((response) => {
    let liveTab = document.getElementById("liveTab");

    let listItem;
    let itemIcon = "file";
    let item;
    let listItemText;

    for (let i = 0; i < response.data.length; i++) {
      item = response.data[i];
      itemIcon = item.icon;

      listItem = document.createElement("i");
      listItem.className = "articleItem fas fa-" + itemIcon;
      listItem.itemId = item.id;
      listItem.onclick = openItem;

      listItemText = document.createTextNode(" " + item.title);
      listItem.appendChild(listItemText);

      liveTab.appendChild(listItem);
    }
  })
}

export default() => {
  window.onresize = invokeResize

  let save = document.getElementById("save");
  if (save) {
    save.onclick = invokeSave
  }

  getList()
}