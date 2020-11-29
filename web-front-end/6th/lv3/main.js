const main = document.querySelector('div#main')
const text = document.querySelector('input')
const button = document.querySelector('button')

button.addEventListener('click', () => search(text.value))
text.addEventListener('change', () => search(text.value))

function search(keywords) {
    main.innerText = 'Searching...'
    new Ajax().get(`http://musicapi.leanapp.cn/search?keywords=${keywords}`, {
        success: updateGrid,
        failure: (res) => { console.error(`Failed：status code: ${res}`) }
    })
}

function createHeader() {
    main.innerHTML =
        `<div><b>ID</b></div>
        <div><b>NAME</b></div>
        <div><b>ARTISTS</b></div>
        <div><b>ALBUM</b></div>`
}

function createNewLine(id, name, artists, album) {
    const id_dom = document.createElement('div')
    const name_dom = document.createElement('div')
    const artists_dom = document.createElement('div')
    const album_dom = document.createElement('div')
    id_dom.innerText = id
    name_dom.innerText = name
    album_dom.innerText = album
    for (let i = 0; i < artists.length; i++) {
        artists_dom.innerText += artists[i].name
        if (i != artists.length - 1)
            artists_dom.innerText += '，'
    }
    main.appendChild(id_dom)
    main.appendChild(name_dom)
    main.appendChild(artists_dom)
    main.appendChild(album_dom)
}

function createLines(json) {
    const songs = json.result.songs
    for (i in songs)
        createNewLine(parseInt(i) + 1, songs[i].name, songs[i].artists, songs[i].album.name)
}

function updateGrid(res) {
    createHeader()
    createLines(JSON.parse(res))
}