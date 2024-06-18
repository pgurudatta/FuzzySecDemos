const express = require('express')
var os = require("os");
var hostname = os.hostname();
var networkInterfaces = os.networkInterfaces();
var username = "abc";
var password = "AVvEV4Ovf4saFi7UxJTq";
const app = express()
const port = 8000
var name = "abc";
var pwd = "AVvEV4Ovf4saFi7UxJTf";

app.get('/', (req, res) => {
  res.send(`Hello World 777, hostname=${hostname}, IP=${networkInterfaces.eth0[0].address}`)
})

app.listen(port, () => {
  console.log(`App is running on port ${port}.`)
})
