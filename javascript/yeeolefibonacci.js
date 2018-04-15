function fibGen() {
    resultMap = {1: 1, 2: 1, 3: 2, 4: 3, 5: 5}
    return function fib(num) {
        let inMap = resultMap[num];
        if (inMap) {
            return inMap;
        } else {
            result = fib(num - 1) + fib(num - 2);
            resultMap[num] = result;
            return result;
        }
    }
}

let f = fibGen();
console.log(f(23));
console.log(f(15));
console.log(f(12));
