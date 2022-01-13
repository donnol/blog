// server.js
let express = require('express');
let app = express();
app.use(express.static(__dirname));
app.
    get('favicon.ico', function (req, res) {
        console.log("favicon")
        res.send('./favicon.ico', { root: __dirname });
    }).
    get('/', function (req, res) {
        res.send('./index.html', { root: __dirname });
    })
let port = 8203;
console.log("listen: ", port);
app.listen(port);
