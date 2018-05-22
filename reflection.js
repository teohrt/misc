var person = {
    firstname: 'Default',
    lastname: 'Default',
    getFullName: function() {
        return this.firstname + ' ' + this.lastname;
    }
}

var john = {
    firstname: 'John',
    lastname: 'Doe'
}

//THIS IS BAD. JUST EXPEREIMENTING
john.__proto__ = person;

//reflecting on the john object. Checking out metadata and logging it
for (var prop in john) {
    if (john.hasOwnProperty(prop)) {
        console.log(prop + ': ' + john[prop]);
    }
}

console.log(john.getFullName());