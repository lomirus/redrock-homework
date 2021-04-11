const { exec } = require('child_process')
const fs = require('fs').promises

class ChangeLogPlugin {
    constructor(options = {}) {
        this.target = options.target ? options.target : 'README.md'
    }
    apply(compiler) {
        compiler.hooks.emit.tapAsync('ChangeLogPlugin', async (compilation, callback) => {
            exec('git log', async (err, stdout, stderr) => {
                fs.writeFile(this.target, stdout)
                callback()
            })
        })
    }
}

module.exports = ChangeLogPlugin