let arr = [[1, 2], 3, [4, [5, [6]], 7]]

function even(arr) {
    let newArr = new Array()
    const reEven = arr => arr.forEach(item => {
        if (item instanceof Array)
            reEven(item)
        else
            newArr.push(item)
    })
    reEven(arr)
    return newArr
}

console.log(even(arr))   // [1,2,3,4,5,6,7]