function Greetr() {

    this.greeting = "Hello world!!!";
    this.greet = function() {
        console.log(this.greeting);
    }
}

// returns the constructor function to give the ability to create a new object
// this way when require caches this value, all subsequent requires don't reference the same object!
module.exports = Greetr;