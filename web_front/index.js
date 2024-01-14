const express = require("express");
const app = express()
const port = 8000

app.use(express.json());
app.use(express.urlencoded({ extended: true }))

app.get('/', (req, res) => {
    res.status(200).send("Welcome to the server")
})

app.get('/get', (req, res) => {
    res.status(200).json({ message: "Hello from rohit" })
})

//for handling raw json data
app.post('/post', (req, res) => {
    let myJson = req.body;      // our JSON

    res.status(200).send(myJson);
})

//for handling form-encoded data
app.post('/postform', (req, res) => {
    res.status(200).send(JSON.stringify(req.body));
})


app.listen(port, () => {
    console.log(`App listening at http://localhost:${port}`)
})