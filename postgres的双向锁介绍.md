## Two phase locking

双向锁所做的无非是 acquire and release, 一旦释放了锁就不能重新获得了

### double booking

两个顾客在同一时间预定了一家餐厅的同一个位置，这就是double booking问题。双相锁就能解决这个问题。

![TPL1](pictures/TPL1.png)

![TPL2](pictures/TPL2.png)

![TPL3](pictures/TPL3.png)