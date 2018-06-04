var http = require('http');
var fs = require('fs');

http.createServer(function(req, res) {
    res.writeHead(200, { 'Content-Type': 'application/json' });

    var obj = {
        firstname: 'John',
        lastname: 'Doe'
    };

    // Serializes the object. AKA Tranforms object into something that
    // can be stored/transfered. Basicly makes it usable.
    res.end(JSON.stringify(obj));

}).listen(1337, '127.0.0.1');