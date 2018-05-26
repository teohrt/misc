function buildFunctions() {
 
    var arr = [];
    
    for (var i = 0; i < 3; i++) {
        
        arr.push(
            function() {
                console.log(i);   
            }
        )
        
    }
    return arr;
}

var fs = buildFunctions();

fs[0]();
fs[1]();
fs[2]();

//logs 3, 3, 3 rather than 0, 1, 2 because the array is holding references of i,
//and not the value of i when it was passed like it might look like


function buildFunctions2() {

    var arr = [];

    for (var i = 0; i < 3; i++) {
        //passing i through the function
        //so it's stored as a value and not a reference
        arr.push(
            (function(j) {
                return function() {
                    console.log(j);
                }
            }(i))
        )
    }

    return arr;
}

var fs = buildFunctions2();

fs[0]();
fs[1]();
fs[2]();

//walaa, 0, 1, 2