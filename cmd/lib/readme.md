

其实API非常的简洁，Value对应的就是JS中的各种数据类型，Callback是JS中的回调函数，其中包括NewCallback，NewEventCallback就能生成回调函数。接下来我们简单的介绍一下API：

js.Global().get: 可以获取window对象上的属性，包括全局变量，DOM API(例如 addEventListener，setInterval)，然后我们就能操作变量，调用函数。
js.Value.Get() 和 js.Value.Set(): 可以获取和设定对象上的属性
js.Value.Index() 和 js.Value.SetIndex(): 可以获取和设定数组上的元素
js.Value.Call(): 可以执行对象上的方法
js.Value.Invoke: 可以执行全局函数，例如addEventListener
还有另外一些，包括：

js.Undefined(): 可以返回JS中的undefined
js.Null(): 可以返回JS中的null
js.ValueOf(): 接受Go的基本类型，返回对应JS类型的值
了解完API，还需要知道Go与JS中的类型对应关系：

```
| Go                     | JavaScript             |
| ---------------------- | ---------------------- |
| js.Value               | [its value]            |
| js.TypedArray          | typed array            |
| js.Callback            | function               |
| nil                    | null                   |
| bool                   | boolean                |
| integers and floats    | number                 |
| string                 | string                 |
| []interface{}          | new array              |
| map[string]interface{} | new object             |
```