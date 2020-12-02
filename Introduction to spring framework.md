Spring是一个轻量级的开源框架，它的诞生极大地提高了基于JAVA的企业级应用的开发效率。

Spring的诞生来自于一本由澳大利亚人Rod Johson写的书：“*Expert One-on-One: J2EE Design and Development*”

Spring能够帮助开发者开发低耦合(loosely coupled)高内聚(highyly cohesive)的系统。低耦合由Spring的控制反转(Inversion Of Control)实现，高内聚由Spring的面向切面编程（Aspect oriented programming (AOP) ）实现。

### 什么是AOP？

面向切面编程是一个编程模式，它的意图是让已经有的代码实现某种功能而不去修改已有代码。这种情况的出现往往因为需要实现的功能不是业务逻辑相关的，比如说LOGGING日志功能。程序中到处都需要记录日志，这种需求在JAVA中被称为横切的需求，因为它一下子切开了好多面。

如果我们需要开发一个记录程序运行时间的功能，最简单的想法可能是这样的：

```
public class BankAccountDAO
{
  public void withdraw(double amount)
  {
    long startTime = System.currentTimeMillis();
    try
    {
      // Actual method body...
    }
    finally
    {
      long endTime = System.currentTimeMillis() - startTime;
      System.out.println("withdraw took: " + endTime);
    }
  }
}
```

很明显上面的代码很糟糕。

首先计时器的开关很不方便，不需要用的时候你甚至需要重写代码把里面的逻辑取出来。。

其次是你的程序中将充满了与业务逻辑无关的代码

解决上面的问题需要用到AOP模块

### 什么是IOC？

一个典型的JAVA应用包含了很多JAVA类(CLASS),每个类都有可能需要用到其他类提供的功能，对于CLASS A来说如果他需要用到CLASS B那么就说A依赖B，通常来说如果A直接依赖B会导致代码间的紧耦合。

SPRING框架能做的是A不再直接依赖于B而是从SPRING容器获取B，这种方式称为控制反转，即不是A直接依赖B而是A直接依赖于由开发者控制的容器。控制反转也称为依赖注入。在容器启动的时候就把依赖注入到A中。

### Spring中的Bean定义

我们经常需要创建一个Employee类，比如下面这样:

```
public class Employee {
	private String userId;
	
	//无参构造器
	public Employee() {
	
	}
	//有参构造器
	public Employee(String userId) {
		this.userId = userId;
	}
	
	//getter
	public String getUserId() {
		return userId;
	}
	//setter
	public void setUserId(String userId) {
		this.userId = userId;
	}
}
```

然后基于XML的spring beans配置文件是这个样子的：

```java
<?xml version="1.0" encoding="UTF-8"?>
<beans xmlns="http://www.springframework.org/schema/beans"
      xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
      xsi:schemaLocation="http://www.springframework.org/schema/beans
      http://www.springframework.org/schema/beans/spring-beans-3.0.xsd">
  
<!-- 使用setter方法注入依赖的例子 -->
<bean id="user1" class="com.chinasws.frameworks.spring.User.Employee">
      <property name="userId">
            <value>14038</value>
      </property>
</bean>
<!-- Below are examples of injecting the dependencies using constructor -->
<bean id="user2" class="com.chinasws.frameworks.spring.User.Employee">
      <constructor-arg>
            <value>14108</value>
      </constructor-arg>
</bean>
<bean id="user3" class="com.chinasws.frameworks.spring.User.Employee">
      <constructor-arg>
            <value>12860</value>
      </constructor-arg>
</bean>   
</beans>
```

然后是基于注释Annotation的spring beans配置：

这里假设所有的BEAN都在BeansConfiguration这个类里面定义，spring默认bean的名字等于方法名，当然如果使用name属性是可以覆盖默认值的。

```
@Configuration
public class BeansConfiguration {
			//使用name=""覆盖默认是方法名的name
      @Bean(name="common_user")
      public Employee user1(){
            Employee employee = new Employee();
            employee.setUserId("14038");
            return employee ;
      }
      @Bean
      public Employee user2(){
            return new Employee("14108");
      }
       @Bean
      public Employee user3(){
            return new Employee("12860");
      }
  
}
```

