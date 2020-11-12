//深拷贝
//思路：遍历+递归
function deepCopy(obj) {
    let newObj = {}
    for (i in obj) {
        if (typeof (obj[i]) == 'object')
            newObj[i] = deepCopy(obj[i])
        else
            newObj[i] = obj[i]
    }
    return newObj
}

//浅拷贝
function shallowCopy(obj){
    return obj
}