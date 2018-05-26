//POLYFILL
//code that adds a feature which the engine MAY lack
//this checks to see if Object.create is available
//and handles it if it's not

if (!Object.create) {
    Object.create = function(o) {
        if (arguments.length > 1) {
            throw new Error('Object.create implementation only accepts the first parameter.');
        }
        function F() {};
        F.prototype = o;
        return new F();
    };
}

var person = {
    firstname: 'Default',
    lastname: 'Default',
    greet: function() {
        return 'Hi ' + this.firstname;
    }
}

//for "newer" browsers
//pure prototypal inheritance
var john = Object.create(person);
//simply override default values
john.firstname = 'John';
john.lastname - 'Doe';
console.log(john.greet());