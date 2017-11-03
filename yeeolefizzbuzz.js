
function fizzBuzz(num) {
    var s = ""
    for (var i = 0; i < num; i ++) {
        if (i % 5 == 0) {
            s += "fizz";
        }
        if (i % 3 == 0) {
            s += "buzz";
        }
        if (s.length > 0) {
            console.log(s);
        } else {
            console.log(i)
        }
        s = ""
    }
}

fizzBuzz(100);