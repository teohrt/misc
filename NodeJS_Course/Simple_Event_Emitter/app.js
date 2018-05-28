var Emitter = require('./emitter');

var emtr = new Emitter();

// create type and listeners
emtr.on('marco', function() {
    console.log('POLO!!!!');
});

emtr.on('greet', function() {
    console.log('Somewhere, someone said, "Hello".');
});

emtr.on('greet', function() {
    console.log('A greeting occured!');
});


console.log('Hello!');
// manually "emit" each time
emtr.emit('greet');

console.log('Marco!');
emtr.emit('marco');

