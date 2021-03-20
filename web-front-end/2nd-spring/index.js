const balls = document.querySelectorAll('.ball')
const moveBall = id => balls[id].style.transform = 'translateX(100px)'

function moveCallback(ball, target, cb) {
    moveBall(0)
    setTimeout(() => {
        moveBall(1)
        setTimeout(() => moveBall(2), 1000)
    }, 1000)
}

function movePromise(ball, target) {
    new Promise(resolve => {
        moveBall(0)
        setTimeout(resolve, 1000)
    }).then(() => new Promise(resolve => {
        moveBall(1)
        setTimeout(resolve, 1000)
    })).then(() => {
        moveBall(2)
    })
}

async function moveAsync() {
    const thisMovePromise = (i) => new Promise(resolve => {
        moveBall(i)
        setTimeout(resolve, 1000)
    })
    await thisMovePromise(0)
    await thisMovePromise(1)
    await thisMovePromise(2)
}
