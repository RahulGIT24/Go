const express = require('express')

const app = express()
app.use(express.json())

app.get('/', (req, res) => {
    res.send("Hello From Node")
})

app.get('/go', (req, res) => {
    res.json({
        "name": "Server for go",
        "status": true
    })
})
app.post('/save', (req, res) => {
    let myJson = req.body;
    res.status(200).send(myJson)
})
app.post('/postForm', (req, res) => {
    res.status(200).send(JSON.stringify(req.body))
})

app.listen(3000, () => {
    console.log("Server Listening at PORT 3000")
})