const path = require('path')
const HtmlWebpackPlugin = require('html-webpack-plugin')
const TestPlugin = require('./plugins/plugin.js')

module.exports = {
    mode: "development",
    entry: path.resolve('src', 'index.js'),
    output: {
        path : path.resolve('dist'),
        filename: 'bundle.js'
    },
    plugins: [ new HtmlWebpackPlugin(), new TestPlugin({
        test: "asd"
    })],
    devServer: {
        contentBase: path.resolve('dist'),
        open: true
    }
}