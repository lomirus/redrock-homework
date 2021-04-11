const path = require('path')
const htmlWebpackPlugin = require('html-webpack-plugin')

module.exports = {
    mode: "production",
    entry: path.resolve('src', 'index.js'),
    output: {
        path : path.resolve('dist'),
        filename: 'bundle.js'
    },
    module: {
        rules: [{
            test: /\.js$/,
            use: [ {
                loader: path.resolve('loaders', 'loader.js'),
                options: {
                    from: "Redrock",
                    to: "Kouiwa"
                }
            }]
        }]
    },
    plugins: [ new htmlWebpackPlugin() ],
    devServer: {
        contentBase: path.resolve('dist'),
        open: true,
        hot: true
    }
}