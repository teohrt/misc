// using Node's event emitter
var Emitter = require('events');

// using a config like this helps avoid misstyping a property
// for better debugging
var eventConfig = require('./config').events;

var emtr = new Emitter();

// create type and listeners
emtr.on(eventConfig.MARCO, function() {
    console.log('POLO!!!!');
});

emtr.on(eventConfig.GREET, function() {
    console.log('Somewhere, someone said, "Hello".');
});

emtr.on(eventConfig.GREET, function() {
    console.log('A greeting occured!');
});


console.log('Hello!');
// manually "emit" each time
emtr.emit(eventConfig.GREET);

console.log('Marco!');
emtr.emit(eventConfig.MARCO);

