const { exec } = require('child_process')
const fs = require('fs').promises

class TestPlugin {
    constructor(options = {}) {
        this.target = options.target ? options.target : 'README.md'
    }
    apply(compiler) {
        compiler.hooks.emit.tapAsync('TestPlugin', async (compilation, callback) => {
            exec('git commit', async (err, stdout, stderr) => {
                fs.writeFile(this.target, stdout)
                callback()
            })
        })
    }
}

module.exports = TestPlugin