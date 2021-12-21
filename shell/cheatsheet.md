1. 变量与引号

- 变量

```
NAME="John"
echo $NAME
echo "$NAME"
echo "${NAME}!"
```

- 字符串引号

```
NAME="John"
echo "Hi $NAME"  #=> Hi John
echo 'Hi $NAME'  #=> Hi $NAME
```

- 常见疑问

单引号跟双引号跟反引号的区别: 单引号直接解释，双引号翻译变量

2. \$\#, \$\@, \$?

这些是 bash-shell 的内置变量

```bash
Example:

file:test.sh
#! /bin/sh
echo '$#' $#
echo '$@' $@
echo '$?' $?

*If you run the above script as*

> ./test.sh 1 2 3

You get output:
$#  3
$@  1 2 3
$?  0

*You passed 3 parameters to your script.*

$# = number of arguments. Answer is 3
$@ = what parameters were passed. Answer is 1 2 3
$? = was last command successful. Answer is 0 which means 'yes'
```
