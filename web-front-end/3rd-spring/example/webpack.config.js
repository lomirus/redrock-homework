const path = require('path')
const htmlWebpackPlugin = require('html-webpack-plugin')
module.exports = {
    mode: 'development',
    entry: path.resolve(__dirname, 'src', 'index.js'),
    output: {
        path: path.resolve(__dirname, 'output'),
        filename: 'bundle.js'
    },
    plugins: [new htmlWebpackPlugin()],
    devServer: {
        contentBase: "./output",
        open: true
    }
}