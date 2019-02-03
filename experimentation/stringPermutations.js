// Recursively prints all of the possible permutations of a string

var stringPermutations = (pref, suf) => {
    if (suf.length === 0) {
        console.log(pref)
    }
    else {
        for (let i = 0; i < suf.length; i++) {
            stringPermutations(pref + suf[i], suf.substring(0, i) + suf.substring(i+1, suf.length))
        }
    }
}

var start = (x) => {
    stringPermutations("", x)
}

var str = "abc";

start(str);