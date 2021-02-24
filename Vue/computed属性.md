### 1.如果computed属性需要当做函数使用如何传递参数？

关键字： **回调函数**

```
computed:{
	getId:function(){
		return function(obj){
			return obj.id
		}
	}
}
```

