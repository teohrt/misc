function mapForEach(arr, fn) {

    var newArr = [];
    for (i = 0; i < arr.length; i++) {
        newArr.push(
            fn(arr[i])
        );
    }

    return newArr;
}

//this example of functional programming elminiates the need
//to write a lines of code looping through a copy of a new array
//to do individual arbitrary work for each problem

var arr1 = [1, 2, 3];
console.log(arr1);

//simply calling the mapForEach function and passing the 
//work that you want done in a function object is enough!
console.log(mapForEach(arr1, function(number) {
    return number * 2;
}));

console.log(mapForEach(arr1, function(number) {
    return number > 2;
}));

var checkPastLimit = function(limiter, item) {
    return item > limiter;
}

//passes a function that takes two arguments
//when the mapForEch function is expecting a function that only takes one
//this is done by making the first argument a default
//using the bind keyword!
//effectively checking if each number is greater than the number 1
var arr2 = mapForEach(arr1, checkPastLimit.bind(this, 1));
console.log(arr2);

var checkPastLimitSimplified = function (limiter) {
    return function(limiter, item) {
        return item > limiter;
    }.bind(this, limiter)
};

var arr3 = mapForEach(arr1, checkPastLimitSimplified(2));
console.log(arr3);