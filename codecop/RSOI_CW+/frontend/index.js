let express = require('express');
let path = require('path');

let app = express();
const port = 8887;

app.use('/admin', express.static(path.join(__dirname, '/admin' ) ));
app.use('/airports', express.static(path.join(__dirname, '/airports' ) ));
app.use('/bonus', express.static(path.join(__dirname, '/bonus' ) ));
app.use('/flights', express.static(path.join(__dirname, '/flights' )));
app.use('/helpers', express.static(path.join(__dirname, '/helpers' )));
app.use('/ticket', express.static(path.join(__dirname, '/ticket' )));
app.use('/user', express.static(path.join(__dirname, '/user' )));


app.get('/', function(req, res) {
    res.sendFile("main.html", { root: path.join(__dirname, '/main' ) });
});

app.get('/main/icon.png', function(req, res) {
    res.sendFile("icon.png", { root: path.join(__dirname, '/main' ) });
});


app.get('/main.css', function(req, res) {
    res.sendFile("main.css", { root: path.join(__dirname, '/main' ) });
});


app.listen(port, () => {
    console.log(`App listening at http://localhost:${port}`)
})
