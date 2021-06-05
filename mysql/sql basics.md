## 基本结构

SELECT

FROM

WHERE

- 使用''来表示带有空格(space)的语句
- 使用''来表示日期

## Comparison

- \>
- \>=
- \<
- \>=
- \=
- \!= 不等号
- \<> 不等号

## operator
- OR
- AND
- NOT

> NOT( A < 1 OR B > 2) == A >= 1 AND B <=2

- IN
    WHERE state IN ('VA','GA','NY')

- NOT IN
- BETWEEN

    使用`BETWEEN`来判断值是否在一个**区间**,对于**日期**、**数字**都可以使用

    > WHERE birth_date BETWEEN '1996-01-01' AND '1996-03-31'

- LIKE

    使用LIKE来匹配`string pattern`
    
    > WHERE name LIKE 'Wu%' 
    过滤name前缀是Wu的名字，`%`表示任何字符,并且LIKE**不区分**前面Wu的**大小写**

    > WHERE name like '%W%'
    过滤名字中带有`W`字符的

    > WHERE last_name like '_u'
    under_core `_` 表示匹配单个字符,这样匹配只能匹配到last_name是两个字符并且最后一个是u的。
    
    > WHERE name like 'W_______G'
    
- REGEXP (regular expression)

    > WHERE last_name REGEXP 'field'
    直接匹配所有包含`field的文本`