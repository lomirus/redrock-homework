let arr = [1,5,6,7,"8",10]
function getSum(arr){
    let sum = 0
    arr.forEach(function(item){
        sum += Number(item)
    })
    return sum;
}
console.log(getSum(arr))//37