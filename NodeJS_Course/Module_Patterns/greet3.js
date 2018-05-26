function Greetr() {

    this.greeting = "Hello world!!";
    this.greet = function() {
        console.log(this.greeting);
    }
}

// returns a new object created from the constructor
module.exports = new Greetr();