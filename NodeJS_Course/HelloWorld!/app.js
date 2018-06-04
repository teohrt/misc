// THE HELLO WORLD OF NODE.JS

// getting this from the node.js core
var http = require('http');
var fs = require('fs');

// creating an object and getting a callback
// which is actually an event listener
http.createServer(function(req, res) {

    // the req (request) and res (response) have methods which allow
    // the sending of information back down the stream
    
    // header
    res.writeHead(200, { 'Content-Type': 'text/html' });
    // body
    var html = fs.readFileSync(__dirname + '/index.html')
    res.end(html);

    // node.js takes this and creates 
    // a properly formatted http response

// when information is sent to this port on the server, use this code
}).listen(1337, '127.0.0.1');