import classes from './main.css'

const axios = require('axios')

function exit(event) {
    console.log(event)

    axios({
        method: 'get',
        url: '/exit'
    })
}

export default() => {
    let exitElem = document.getElementById('exit')
    exitElem.onclick = exit

    axios({
        method: 'get',
        url: '/tester'
    }).then((response) => {
        console.log(response.body)
    })
}