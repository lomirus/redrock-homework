import axios from 'schoolwork-1'
axios.get('https://anonym.ink/api/home/sections')
    .then(data => {
        JSON.parse(data).data.map(section => {
            section.List.map(v => {
                console.log("Title: ", v.Title)
                console.log("Description: ", v.Description)
                console.log('')
            })
        })
    })
    .catch(err => {
        console.log(err)
    })