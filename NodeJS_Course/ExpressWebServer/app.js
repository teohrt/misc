var express = require('express');
var app = express();

// useful for if there is a port variable
// on the server this might be running on
var port = process.env.PORT || 3000;

app.get('/', function(req, res) {
    res.send('<html><head><head/><body><h1>Hello World!</h1></body></html>');
});

app.get('/api', function(req, res) {
    res.json({ firstname: 'John', lastname: 'Doe'});
});

app.listen(port);