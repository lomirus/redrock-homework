class myResponse {
    constructor({ url, status, res }) {
        this.url = url
        this.status = status
        this.json = function () {
            return new Promise(function (resolve) {
                resolve({ code: status, result: [JSON.parse(res)] })
            })
        }
        this.text = function () {
            return new Promise(function (resolve) {
                resolve(JSON.stringify({ code: status, result: [JSON.parse(res)] }))
            })
        }
    }
}
class myRequest {
    constructor({ method = 'GET', headers = {}, data = '', mode }) {
        this.method = method
        this.data = data
        this.headers = headers,
        this.mode = mode
    }
}
function myFetch(url, req = {}) {
    return new Promise(function (resolve, reject) {
        const request = new myRequest(req)
        const xhr = new XMLHttpRequest()
        if (request.mode === 'cors')
            xhr.withCredentials = true
        xhr.onreadystatechange = function () {
            if (xhr.readyState === 4) {
                if ((xhr.status >= 200 && xhr.status < 300) || xhr.status === 304) {
                    resolve(new myResponse({
                        res: xhr.response,
                        status: xhr.status,
                        url: url
                    }))
                } else {
                    reject(xhr.status)
                }
            }
        }
        xhr.open(request.method, url, true)
        for (i in request.headers)
            xhr.setRequestHeader(i, request.headers[i])
        xhr.send(request.data)
    })
}

myFetch('http://musicapi.leanapp.cn/personalized?limit=1', {
    method: 'GET',
})
    .then(value => value.json())
    .then(value => console.log(value))

fetch('http://musicapi.leanapp.cn/personalized?limit=1', {
    method: 'GET',
})
    .then(value => value.json())
    .then(value => console.log(value))
