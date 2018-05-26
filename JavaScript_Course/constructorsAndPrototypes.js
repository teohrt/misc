
//this is the function constructor
function Person(firstname, lastname) {
    console.log(this);
    this.firstname = firstname;
    this.lastname = lastname;
    console.log('This function is invoked');
}

//you could just have this function object in the constructor
//but you have to remember that objects can take up significant memory space and every
//"Person" would then necessarily have a copy of it
//adding it to the protoypte is much more efficient because then there's only one
//its better to put methods in the prototype because of this

//something to keep in mind
Person.prototype.getFullName = function() {
    return this.firstname + ' ' + this.lastname;
}

//using the new keyword (it's actually an operator!) it constructs an object using the function
var john = new Person('John', 'Doe');
console.log(john.getFullName());

var jane = new Person('Jane', 'Doe');
console.log(jane.getFullName());


