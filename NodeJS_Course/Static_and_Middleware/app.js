var express = require('express');
var app = express();

var port = process.env.PORT || 3000;

// whatever static files are sitting on the server
// can be hooked to and streamed automatically
// with some simple middleware like this!
app.use('/assets', express.static(__dirname + '/public'));

// simple middleware logger
// (the first parameter isn't necessary if you want 
// the function run for every route)
app.use('/', function(req, res, next) {
	console.log('Request URL: ' + req.url);
	next();
});

// static file referenced here!
app.get('/', function(req, res) {
	res.send('<html><head></head><link href=assets/style.css type=text/css rel=stylesheet /><body><h1>Hello world!</h1></body></html>');
});

app.get('/person/:id', function(req, res) {
	res.send('<html><head></head><body><h1>What\'s up, ' + req.params.id + '?!</h1></body></html>');
});

app.get('/api', function(req, res) {
	res.json({ firstname: 'John', lastname: 'Doe' });
});

app.listen(port);