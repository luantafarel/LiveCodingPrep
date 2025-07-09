function prime(N) {
    let isPrime = true
    let startPrime = 2
    const primeArray = []
    for (let i = 0; i < N;) {
        for (let j = startPrime - 1; j>1 ; j--){
            if (startPrime % j === 0) {
                isPrime = false;
            }
        }
        if (isPrime === true) {
            primeArray.push(startPrime);
            i++
        }
        isPrime = true;
        startPrime++;
    }
    return primeArray;
}

function isPrime(num) {
    if (num < 2) return false;
    for (let i = 2; i * i <= num; i++) {
        if (num % i === 0) return false;
    }
    return true;
}

function primeSqrt(N) {
    const primeArray = [];
    let candidate = 2;
    while (primeArray.length < N) {
        if (isPrime(candidate)) {
            primeArray.push(candidate);
        }
        candidate++;
    }
    return primeArray;
}

console.time('prime');
console.log(prime(2050));
console.timeEnd('prime');

console.time('primeSqrt');
console.log(primeSqrt(2050));
console.timeEnd('primeSqrt');
const primes1 = prime(2050);
const primes2 = primeSqrt(2050);
console.log("prime and primeSqrt outputs are equal:", JSON.stringify(primes1) === JSON.stringify(primes2));