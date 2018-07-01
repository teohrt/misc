var express = require('express');
var app = express();
var mongoose = require('mongoose');
var config = require('./config');
var setupController = require('./controllers/setupController');
var apiController = require ('./controllers/apiController');

// If on production we'll have an environment variable for the port number
// otherwise we'll default to localhost:3000
var port = process.env.PORT || 3000;

// Set up public assets folder to be delivered straight to the browser
app.use('/assets', express.static(__dirname + '/public'));

// View engine for cool templating for the serverside. 
// Probably use Angular.js for clientside
app.set('view engine', 'ejs');

// Connect to the DB!
mongoose.connect(config.getDbConnectionString());
// Get a reference to the connection
var db = mongoose.connection;

// Making express aware of our endpoints
setupController(app);
apiController(app);

console.log('Server started and listening...');
app.listen(port);