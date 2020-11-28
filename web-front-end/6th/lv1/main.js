const main = document.querySelector('div#main')
const text = document.querySelector('input')
const button = document.querySelector('button')
const xhr = new XMLHttpRequest()
xhr.withCredentials = true
xhr.addEventListener('readystatechange', function(){
    if (xhr.readyState === 4 && xhr.status === 200) {
        initHeader()
        const songs = JSON.parse(xhr.response).result.songs
        for (i in songs) {
            const id = document.createElement('div')
            const name = document.createElement('div')
            const artists = document.createElement('div')
            const album = document.createElement('div')
            id.innerText = parseInt(i) + 1
            name.innerText = songs[i].name
            artists.innerText = songs[i].album.name
            for (let j = 0; j < songs[i].artists.length; j++) {
                album.innerText += songs[i].artists[j].name
                if (j != songs[i].artists.length - 1)
                    album.innerText += '，'
            }
            main.appendChild(id)
            main.appendChild(name)
            main.appendChild(artists)
            main.appendChild(album)
        }
    }
})
button.addEventListener('click',  () => search(text.value))
text.addEventListener('change',  () => search(text.value))

function search(keywords) {
    main.innerText = 'Searching...'
    xhr.open('GET', `http://musicapi.leanapp.cn/search?keywords=${keywords}`)
    xhr.send()
}
function initHeader() {
    main.innerHTML = 
        `<div><b>ID</b></div>
        <div><b>NAME</b></div>
        <div><b>ARTISTS</b></div>
        <div><b>ALBUM</b></div>`
}