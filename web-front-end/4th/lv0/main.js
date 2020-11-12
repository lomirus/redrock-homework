const person = {
    name: 'Boson',
    age: 20,
    address: {
      city: 'Chongqing',
      area: 'Nanan'
    }
}

const {name, age, address:{city, area}} = person

console.log(name)
console.log(age)
console.log(city)
console.log(area)