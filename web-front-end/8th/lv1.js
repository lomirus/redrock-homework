function makeClosures(arr, fn){
    let result = new Array()
    for(let i in arr){
        result[i] = fn.bind(global, arr[i])
    }
    return result
}

let array = [1, 2, 3]
let foo = (x) => {
    console.log(x)
}
let fun = makeClosures(array, foo)
fun[1]() //2