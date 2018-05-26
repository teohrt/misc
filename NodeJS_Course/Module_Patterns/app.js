var greet1 = require('./greet1');
// you can simply invoke the variable that's created because require has set it to a function here
greet1();

// this pattern sets the variale more explicitly by stating which object in the module to set
var greet2 = require('./greet2').greet;
greet2();

// this pattern completely replaces the module.exports object with a newly constructed object
var greet3 = require('./greet3');
greet3.greet();

//making a change to the variables members affects all future requires because it's a reference!!!! Be careful.
greet3.greeting = "Changed Hello World because object is a reference";
greet3b = require('./greet3');
greet3b.greet();

// returns the constructor function to give the ability to create a new object
// this way when require caches this value, all subsequent requires don't reference the same object!
var greet4 = require('./greet4');
var grtr = new greet4();
grtr.greet();

// THE REVEALING MODULE PATTERN
// only the function is exposed here. Not the variable.
// The variable is protected
var greet5 = require('./greet5').greet;
greet5();