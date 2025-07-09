function anagramChecker(s1, s2){
    const str1 = s1.split("").sort()
    const str2 = s2.split("").sort()
    return str1.join("") ===  str2.join("")
}
console.log(anagramChecker("jndbiosa", "dbiosajn"))