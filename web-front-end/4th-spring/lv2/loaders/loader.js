const { getOptions } = require('loader-utils');

module.exports = function(source) {
    const { from, to } = getOptions(this)
    return source.replace(from, to)
}