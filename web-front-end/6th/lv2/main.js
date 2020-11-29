const main = document.querySelector('div#main')
const text = document.querySelector('input')
const button = document.querySelector('button')

button.addEventListener('click', () => search(text.value))
text.addEventListener('change', () => search(text.value))

function search(keywords) {
    main.innerText = 'Searching...'
    ajax({
        method: 'GET',
        url: `http://musicapi.leanapp.cn/search?keywords=${keywords}`,
        async: true
    }, updateGrid)
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

function ajax({ method, url, data, async }, handleRes) {
    const xhr = new XMLHttpRequest()
    xhr.withCredentials = true
    xhr.addEventListener('readystatechange', function () {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                handleRes(xhr.response)
            } else {
                console.error(`Failed: xhr.status: ${xhr.status}`)
            }
        }
    })
    xhr.open(method, url, async)
    xhr.send(data)
}

function updateGrid(res) {
    createHeader()
    createLines(JSON.parse(res))
}