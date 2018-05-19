var person = {
    firstname: 'John',
    lastname: 'Doe',
    getFullName: function() {

        var fullname = this.firstname + ' ' + this.lastname;
        return fullname;

    }
}

var logName = function(lang1, lang2) {

    console.log('Logged: ' + this.getFullName());
    console.log('Arguments: ' + lang1 + ' ' + lang2);
    console.log('-------------');
}

//bind changes what the "this" variable points to
var logPersonName = logName.bind(person);
logPersonName('en');

//call calls the function with the arguemnts rather than making a copy
logName.call(person, 'en', 'es');

//apply does the same thing except the arguments are in an array
logName.apply(person, ['en', 'es']);


//function borrowing
var person2 = {
    firstname: 'Jane',
    lastname: 'Doe'
}

console.log(person.getFullName.apply(person2));


//function currying
//taking a function and making a new function out of it with default parameters!
function multiply(a, b) {
    return a*b;
}

var multiplyByTwo = multiply.bind(this, 2);
console.log(multiplyByTwo(4));
