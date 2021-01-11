### 文章目录

- - [步骤](https://blog.csdn.net/u013778905/article/details/83501204#_10)
  - - [一、设置git的user name和email](https://blog.csdn.net/u013778905/article/details/83501204#gituser_nameemail_11)
    - [二、检查是否存在SSH Key](https://blog.csdn.net/u013778905/article/details/83501204#SSH_Key_21)
    - [三、获取SSH Key](https://blog.csdn.net/u013778905/article/details/83501204#SSH_Key_46)
    - [四、GitHub添加SSH Key](https://blog.csdn.net/u013778905/article/details/83501204#GitHubSSH_Key_55)
    - [五、验证和修改](https://blog.csdn.net/u013778905/article/details/83501204#_65)



```html
https://github.com/xiangshuo1992/preload.git
git@github.com:xiangshuo1992/preload.git
12
```

这两个地址展示的是同一个项目，但是这两个地址之间有什么联系呢？
 前者是https url 直接有效网址打开，但是用户每次通过git提交的时候都要输入用户名和密码，有没有简单的一点的办法，一次配置，永久使用呢？当然，所以有了第二种地址，也就是SSH URL，那如何配置就是本文要分享的内容。
 **GitHub配置SSH Key的目的是为了帮助我们在通过git提交代码是，不需要繁琐的验证过程，简化操作流程。**

## 步骤

### 一、设置git的user name和email

如果你是第一次使用，或者还没有配置过的话需要操作一下命令，自行替换相应字段。

```git
git config --global user.name "Luke.Deng"
git config --global user.email  "xiangshuo1992@gmail.com"
12
```

说明：git config --list 查看当前Git环境所有配置，还可以配置一些命令别名之类的。

### 二、检查是否存在SSH Key

```git
cd ~/.ssh
ls
或者
ll
//看是否存在 id_rsa 和 id_rsa.pub文件，如果存在，说明已经有SSH Key
12345
```

如下图
 ![在这里插入图片描述](https://img-blog.csdnimg.cn/2018102909093380.png)
 如果没有SSH Key，则需要先生成一下

```git
ssh-keygen -t rsa -C "xiangshuo1992@gmail.com"
1
```

执行之后继续执行以下命令来获取SSH Key

```git
cd ~/.ssh
ls
或者
ll
//看是否存在 id_rsa 和 id_rsa.pub文件，如果存在，说明已经有SSH Key
12345
```

### 三、获取SSH Key

```git
cat id_rsa.pub
//拷贝秘钥 ssh-rsa开头
12
```

如下图
 ![在这里插入图片描述](https://img-blog.csdnimg.cn/20181029091352393.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3UwMTM3Nzg5MDU=,size_12,color_FFFFFF,t_70)

### 四、GitHub添加SSH Key

GitHub点击用户头像，选择setting
 ![在这里插入图片描述](https://img-blog.csdnimg.cn/20181029092207633.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3UwMTM3Nzg5MDU=,size_12,color_FFFFFF,t_70)

新建一个SSH Key
 ![在这里插入图片描述](https://img-blog.csdnimg.cn/20181029092310463.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3UwMTM3Nzg5MDU=,size_12,color_FFFFFF,t_70)

取个名字，把之前拷贝的秘钥复制进去，添加就好啦。

### 五、验证和修改

测试是否成功配置SSH Key

```git
ssh -T git@github.com
//运行结果出现类似如下
Hi xiangshuo1992! You've successfully authenticated, but GitHub does not provide shell access.
123
```

之前已经是https的链接，现在想要用SSH提交怎么办？
 直接修改项目目录下 `.git`文件夹下的`config`文件，将地址修改一下就好了。

git地址获取可以看如下图切换。
 ![在这里插入图片描述](https://img-blog.csdnimg.cn/20181029093141515.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3UwMTM3Nzg5MDU=,size_12,color_FFFFFF,t_70)

END