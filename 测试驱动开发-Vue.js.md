### 为什么要测试？

没有经过测试的代码直接上生产环境发布是让人感到担心的。。

至少在JAVA中我有一次犯了这么一个错误：

```java
方法A(){
	List<String> list = Collections.singletonList("14038");
	方法B(list);
}

方法B(Collection<String> collection){
	Iterator iterator = collection.iterator();
	...
	iterator.remove();
	...
}
```

上面的代码编译不会产生任何错误，但是运行的时候就会报错，后来跟同事讨论这个的时候发现原因是SingletonList类没有实现Iterator的remove方法。。

很明显如果有测试上面的问题是可以发现的。。

### 如何测试？

为什么不想测试？因为不知道怎么测试。

那既然不知道，测试的第一步大概就是这样做：

1. 写一个测试描述期待的行为，然后运行它等着它报错😄  

2. 然后写一个最简单最笨的测试，这次要让它测试通过。

3. 重构，并且保持测试通过的绿色状态。

   