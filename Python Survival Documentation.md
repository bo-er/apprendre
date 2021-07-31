## __init__

```
class MyClass(object):
    i = 123
    def __init__(self):
        self.i = 345
     
a = MyClass()
print(a.i)
print(MyClass.i)

```

执行结果:

```
345
123

```