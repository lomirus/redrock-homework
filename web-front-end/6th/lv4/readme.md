## Readme
这里选择的题目是用 `Promise` 实现 `fetch`

不过`fetch`参数太多，所以就主要实现了`url`,`method`和`data`，还有`headers`和`mode`这几个算半完成。我之所以说`headers`是半完成是因为我也不确定有没有实现成功，我是用学长给的那个网易云音乐的api做测试的，而它似乎不允许在请求时携带`headers`(我试着用原生的`fetch`传了个`headers:{'Content-type':'application/json'}`也会报错提示不被允许)，所以说不太好验证。

除了`fetch()`和`Headers`之外，在[MDN文档关于 Fetch 的接口](https://developer.mozilla.org/zh-CN/docs/Web/API/Fetch_API)中还提到了`Request`和`Response`。所以也顺便实现了`Request`和`Response`中的一些常用属性和方法，譬如`Response`里面的`url`, `status`, `json()`和 `text()`，至于一些其他的不常用属性方法就没有去实现。

不过大体上还算是实现了`fetch`的基础功能，如以下代码所示：

```javascript
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

```
