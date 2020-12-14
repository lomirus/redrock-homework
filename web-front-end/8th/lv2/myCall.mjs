import lib from './lib.mjs'
Function.prototype.myCall = function (newThis, ...args) {
    let tempThis = lib.deepCopy(newThis)
    tempThis.func = this
    tempThis.func(...args)
}
function sayInfo(greeting) {
    console.log("name:" + this.name)
    console.log("age:" + this.age)
    console.log(greeting)
}
var user = {
    name: "07*",
    age: 19
}
sayInfo.call(user, "Hello World")
sayInfo.myCall(user, "Hello World")
// name:07*
// age:19
// Hello World