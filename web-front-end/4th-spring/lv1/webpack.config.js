const path = require('path')
const htmlWebpackPlugin = require('html-webpack-plugin')
module.exports = {
    mode: "development",
    entry: path.resolve('src', 'index.js'),
    output: {
        path: path.resolve('dist'),
        filename: "bundle.js"
    },
    module: {
        rules: [{
            test: /\.css$/,
            use: ["style-loader", "css-loader"]
        }]
    },
    plugins: [ new htmlWebpackPlugin({
        template: path.resolve('src', 'templates', 'index.html'),
    })],
    devServer: {
        contentBase: path.resolve('dist'),
        open: true,
    }
}