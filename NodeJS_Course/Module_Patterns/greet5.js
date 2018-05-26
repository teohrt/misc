var greeting = "Hello world!!!!";

function greet() {
    console.log(greeting);
    
}
// REVEALING MODULE PATTERN
//
//only exposing the function. Not the variable
module.exports = {
    greet: greet
}