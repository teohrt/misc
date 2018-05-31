var EventEmitter = require('events');
var util = require('util');

function Greetr() {
    // Super constructor: the constructor I'm inheriting from!
    // This gives things the object has outside of it's prototype
    EventEmitter.call(this);
    this.greeting = 'Hello world!';
}

// any new objects created from Greetr will 
// also get access to methods and properties on the
// prototype of EventEmitter!
// JUST PROTOTYPE
util.inherits(Greetr, EventEmitter);

// an example of passing data through events 
Greetr.prototype.greet = function(data) {
    console.log(this.greeting + ': ' + data);
    this.emit('greet', data);
    // emits to any listeners
};

var greeter1 = new Greetr();

greeter1.on('greet', function(data) {
    console.log('Someone greeted!: ' + data);
});

// this emits an event and triggers our listener while passing data!
greeter1.greet('Trace Ohrt');