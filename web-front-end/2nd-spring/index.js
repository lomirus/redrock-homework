const balls = document.querySelectorAll('.ball')
const moveBall = id => balls[id].style.transform = 'translateX(100px)'
const newMovePromise = id => new Promise(resolve => {
    moveBall(id)
    setTimeout(resolve, 1000)
})

function moveCallback() {
    moveBall(0)
    setTimeout(() => {
        moveBall(1)
        setTimeout(() => moveBall(2), 1000)
    }, 1000)
}

function movePromise() {
    newMovePromise(0)
        .then(() => newMovePromise(1))
        .then(() => moveBall(2))
}

async function moveAsync() {
    await newMovePromise(0)
    await newMovePromise(1)
    await newMovePromise(2)
}
