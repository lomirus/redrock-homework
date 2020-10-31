let arr = ["myfirstactivity","today activity","yourActivity","activitys"]
//my code start
let newArr = new Array()
arr.forEach(function(item){
    if (item.indexOf('activity'))
        newArr.push(item)
})
//my code end
console.log(newArr);