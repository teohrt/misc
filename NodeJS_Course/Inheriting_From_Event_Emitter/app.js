'use strict';
var EventEmitter = require('events');

// syntactic sugar

// sets up the prototype chain. Greetr inherits from EventEmitter
class Greetr extends EventEmitter{
    constructor() {
        super();
        this.greeting = 'Hello world!';
    }

    // This gives the function to Greetr.prototype.greet like before
    greet(data) {
        console.log(this.greeting + ': ' + data);
        this.emit('greet', data);
    }
}

var greeter1 = new Greetr();

greeter1.on('greet', function(data) {
    console.log('Someone greeted!: ' + data);
});

// this emits an event and triggers our listener while passing data!
greeter1.greet('Trace Ohrt');