var fs = require('fs');

// part of the node core. Allows us to implemnt a gzip file 
var zlib = require('zlib');

// creates a read stream and places chunks
// of the file in a specified buffer
// encoding makes it a string rather than hex
// highWaterMark specpifies the size of the buffer: 32 * 1024 = 32k
var readable = fs.createReadStream(__dirname + '/greet.txt', { encoding: 'utf8', highWaterMark: 16 * 1024});

// create write stream
var writable = fs.createWriteStream(__dirname + '/greetcopy.txt');

// creating a compressing stream. readable and writable. you can pipe the output!
var gzip = zlib.createGzip();

var compressed = fs.createWriteStream(__dirname + '/greet.txt.gz');

// read some of the file, fill the buffer, 
// emit 'data' event, run all the listeners, 
// read some more, etc...
// every time the 'data' event is emitted it passes the data to the listener

// readable.on('data', function(chunk) {
// 	console.log(chunk.length);
// 	// writes data with our write stream to greetcopy.txt
// 	writable.write(chunk);
// });

// read from read stream and write to write stream
// simply put with a pipe:
readable.pipe(writable);

// read from greet.txt
// pipe every chunk to gzip
// pipe gzips output to the writable 'compressed' stream
// stream to stream to stream! chaining baby!
readable.pipe(gzip).pipe(compressed);