## 赋予权限

1. MySQL 8.0 之后

```
grant all privileges on *.* to universe_op@'%' with grant option;

```

2. MySQL 8.0 之前

```
GRANT ALL PRIVILEGES ON *.* TO 'universe_op'@'%' IDENTIFIED BY '123'  WITH GRANT OPTION;

```
