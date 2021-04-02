const axios = ({ method, url }) => new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    xhr.setHeader("Access-Control-Allow-Origin", "*");  
    xhr.onreadystatechange = () => {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                resolve(xhr.response)
            } else {
                reject(xhr.status)
            }
        }
    }
    xhr.open(method, url);
    xhr.send();
})

axios.get = url => axios({ method: "GET", url})

axios.post = url => axios({ method: "POST", url})

export default axios
