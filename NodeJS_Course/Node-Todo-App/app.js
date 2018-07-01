var express = require('express');
var app = express();

// if on production we'll have an environment variable for the port number
// otherwise we'll default to localhost:3000
var port = process.env.PORT || 3000;

// set up public assets folder to be delivered straight to the browser
app.use('/assets', express.static(__dirname + '/public'));

// View engine for cool templating for the serverside. 
// Probably use Angular.js for clientside
app.set('view engine', 'ejs');

app.listen(port);