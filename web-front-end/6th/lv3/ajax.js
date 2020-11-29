class Ajax {
    constructor({ data = null, header = {}, success = () => { }, failure = () => { }, async = true } = {}) {
        this.header = header
        this.data = data
        this.success = success
        this.failure = failure
        this.async = async
    }
    get(url, { async, data, header, success, failure } = {}) {
        this.#default('get', url, { async, data, header, success, failure })
    }
    post(url, { async, data, header, success, failure } = {}) {
        this.#default('post', url, { async, data, header, success, failure })
    }
    #default(method, url, {
        header = {},
        async = this.async,
        data = this.data, 
        success = this.success,
        failure = this.failure 
    }) {
        const newxhr = new XMLHttpRequest
        for (i in this.header)
            newxhr.setRequestHeader(i, this.header[i])
        for (i in header)
            newxhr.setRequestHeader(i, header[i])
        newxhr.onreadystatechange = function () {
            if (newxhr.readyState === 4) {
                if ((newxhr.status >= 200 && newxhr.status < 300) || newxhr.status === 304)
                    success(newxhr.response)
                else
                    failure(newxhr.status)
            }
        }
        newxhr.open(method, url, async)
        newxhr.send(data)
    }
}