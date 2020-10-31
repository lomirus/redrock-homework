let str = "can-enter-volunteer-organization"

function change(str) {
    return str.replace(/(-.)/g, match => match[1].toUpperCase())
}

change(str)