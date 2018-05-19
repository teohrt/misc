function buildFunctions() {

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

var fs = buildFunctions();

fs[0]();
fs[1]();
fs[2]();

