const save = () => (number = 0) => save.sum ? save.sum += number : save.sum = number
const add = save()

add(100)
add(200)

console.log(add())
