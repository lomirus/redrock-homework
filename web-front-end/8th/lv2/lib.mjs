// 实现 myCall 时需要用到深拷贝，但是之前已经造过一次轮子了，
// 又考虑到深拷贝也不是这次作业的重点，所以就直接网上找了个现成的实现深拷贝的函数，并做了点修改
// Ctrl + C/V START
function deepCopy(data, hash = new WeakMap()) {
    if (typeof data !== 'object' || data === null) {
        throw new TypeError('The parameter passed in is not an object')
    }
    if (hash.has(data)) {
        return hash.get(data)
    }
    let newData = {};
    const dataKeys = Object.keys(data);
    dataKeys.forEach(value => {
        const currentDataValue = data[value];
        if (typeof currentDataValue !== "object" || currentDataValue === null) {
            newData[value] = currentDataValue;
        } else if (Array.isArray(currentDataValue)) {
            newData[value] = [...currentDataValue];
        } else if (currentDataValue instanceof Set) {
            newData[value] = new Set([...currentDataValue]);
        } else if (currentDataValue instanceof Map) {
            newData[value] = new Map([...currentDataValue]);
        } else {
            hash.set(data, data)
            newData[value] = deepCopy(currentDataValue, hash);
        }
    });
    return newData;
}
// Ctrl + C/V END

export default {
    deepCopy
}