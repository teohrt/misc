var express = require('express');
var bodyParser = require('body-parser');
var app = express();

var port = process.env.PORT || 3000;

var urlencodedParser = bodyParser.urlencoded({ extended: false });
var jsonParser = bodyParser.json();

app.use('/assets', express.static(__dirname + '/public'));

app.set('view engine', 'ejs');

app.use('/', function (req, res, next) {
	console.log('Request Url:' + req.url);
	next();
});

app.get('/', function(req, res) {
	res.render('index');
});

// "localhost:3000/person/Trace?qstr=123"
// this generates an http request including 
// the querry string in the header
// which was parsed and made availble in express 
// by adding the .query object on the as a property on the
// query object
app.get('/person/:id', function(req, res) {
	res.render('person', { ID: req.params.id, qstr: req.query.qstr });
});

// grabbing a form post using the body-parser middleware
app.post('/person', urlencodedParser, function(req, res) {
	res.send('Hey, thanks!');
	console.log(req.body.firstname);
	console.log(req.body.lastname);
});

// handling json data
app.post('/person/json', jsonParser, function(req, res) {
	res.send('Thanks for the JSON data!');
	console.log(req.body.firstname);
	console.log(req.body.lastname);
})

app.get('/api', function(req, res) {
	res.json({ firstname: 'John', lastname: 'Doe' });
});

app.listen(port);