## 1 区别

- Integer是int的包装类，int则是java的一种基本数据类型
- Integer变量必须实例化后才能使用，而int变量不需要
- Integer实际是对象的引用，当new一个Integer时，实际上是生成一个指针指向此对象；而int则是直接存储数据值
- Integer的默认值是null，int的默认值是0

## 2 ==比较

#### 2.1、由于Integer变量实际上是对一个Integer对象的引用，所以两个通过new生成的Integer变量永远是不相等的（因为new生成的是两个对象，其内存地址不同）。

```
Integer i = new Integer(100);
Integer j = new Integer(100);
System.out.print(i == j); //false
123
```

#### 2.2、Integer变量和int变量比较时，只要两个变量的值是向等的，则结果为true（因为包装类Integer和基本数据类型int比较时，java会自动拆包装为int，然后进行比较，实际上就变为两个int变量的比较）

```
Integer i = new Integer(100);
int j = 100；
System.out.print(i == j); //true
123
```

#### 2.3、非new生成的Integer变量和new Integer()生成的变量比较时，结果为false。（因为非new生成的Integer变量指向的是java常量池中的对象，而new Integer()生成的变量指向堆中新建的对象，两者在内存中的地址不同）

```
Integer i = new Integer(100);
Integer j = 100;
System.out.print(i == j); //false
123
```

#### 2.4、对于两个非new生成的Integer对象，进行比较时，如果两个变量的值在区间-128到127之间，则比较结果为true，如果两个变量的值不在此区间，则比较结果为false

```
Integer i = 100;
Integer j = 100;
System.out.print(i == j); //true
Integer i = 128;
Integer j = 128;
System.out.print(i == j); //false
123456
```

对于第4条的原因：
 java在编译Integer i = 100 ;时，会翻译成为Integer i = Integer.valueOf(100)；，而java API中对Integer类型的valueOf的定义如下：

```
public static Integer valueOf(int i){
    assert IntegerCache.high >= 127;
    if (i >= IntegerCache.low && i <= IntegerCache.high){
        return IntegerCache.cache[i + (-IntegerCache.low)];
    }
    return new Integer(i);
}
1234567
```

java对于-128到127之间的数，会进行缓存，Integer i = 127时，会将127进行缓存，下次再写Integer j = 127时，就会直接从缓存中取，就不会new了

## 3 延伸

#### 3.1、理解自动装箱、拆箱

自动装箱与拆箱实际上算是一种“语法糖”。所谓语法糖，可简单理解为Java平台为我们自动进行了一些转换，保证不同的写法在运行时等价。因此它们是发生在编译阶段的，也就是说生成的字节码是一致的。

对于整数，javac替我们自动把装箱转换为Integer.valueOf()，把拆箱替换为Integer.intValue()。可以通过将代码编译后，再反编译加以证实。

原则上，建议避免无意中的装箱、拆箱行为，尤其是在性能敏感的场合，创建10万个Java对象和10万个整数的开销可不是一个数量级的。当然请注意，只有确定你现在所处的场合是性能敏感的，才需要考虑上述问题。毕竟大多数的代码还是以开发效率为优先的。

顺带说一下，在32位环境下，Integer对象占用内存16字节；在64位环境下则更大。

#### 3.2、值缓存

就像上一讲谈到的String，Java也为Integer提供了值缓存。

Integer i1 = 1;Integer i2 = Integer.valueOf(2);Integer i3 = new Integer(3);

上述代码中第一行与第二行的写法取值使用了值缓存，而第三行的写法则没有利用值缓存。结合刚刚讲到的自动装箱、拆箱的知识，第一行代码用到的自动装箱，等价于调用了Integer.valueOf()。

不仅仅是Integer，Java也为其它包装类提供了值缓存机制，包括Boolean、Byte、Short和Character等。但与String不同的是，默认都只会将绝对值较小的值放入缓存。以Integer为例，默认情况下只会缓存-128到127之间的值。当然如果你愿意也可以通过以下JVM参数进行设置：

-XX:AutoBoxCacheMax=N