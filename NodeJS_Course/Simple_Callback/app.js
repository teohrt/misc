function greet(callback) {
    console.log('Hey!');
    var data = {
        name: 'John Doe'
    }
    callback(data);
}

greet(function(data) {
    console.log('The callback was invoked!');
    console.log(data);
});

greet(function(data) {
    console.log('Another callback was invoked!');
    console.log(data.name);
});