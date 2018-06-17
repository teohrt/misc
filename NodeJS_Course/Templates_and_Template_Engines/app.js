var express = require('express');
var app = express();

var port = process.env.PORT || 3000;

app.use('/assets', express.static(__dirname + '/public'));

// let's you use code inside your html!
// templates baby!
// Tons of flexibility
app.set('view engine', 'ejs')

app.use('/', function (req, res, next) {
	console.log('Request Url:' + req.url);
	next();
});

// point at the name of the file to render
// EJS takes care of the extension above
app.get('/', function(req, res) {
	res.render('index');
});	

// simple way of passing variables for EJS
app.get('/person/:id', function(req, res) {
	res.render('person', { ID: req.params.id });
});

app.get('/api', function(req, res) {
	res.json({ firstname: 'John', lastname: 'Doe' });
});

app.listen(port);