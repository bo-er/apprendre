### 使用 rsa key

使用下面的命令生成 rsa key,默认会让你输入一个密码,可以试着输入一个空密码

```sh
ssh-keygen -t rsa
```

如果上面没有使用空密码
输入下面的命令修改密码:

```sh
ssh-keygen -p
```

帮助信息提示 ssh-keygen -p 会依次让你确定:

1. 密钥文件
2. 格式
3. 新密码
4. 旧密码

```sh
usage: ssh-keygen [-q] [-b bits] [-C comment] [-f output_keyfile] [-m format]
                  [-N new_passphrase] [-t dsa | ecdsa | ed25519 | rsa]
       ssh-keygen -p [-f keyfile] [-m format] [-N new_passphrase]
                   [-P old_passphrase]
       ssh-keygen -i [-f input_keyfile] [-m key_format]
       ssh-keygen -e [-f input_keyfile] [-m key_format]
       ssh-keygen -y [-f input_keyfile]
       ssh-keygen -c [-C comment] [-f keyfile] [-P passphrase]
       ssh-keygen -l [-v] [-E fingerprint_hash] [-f input_keyfile]
       ssh-keygen -B [-f input_keyfile]
       ssh-keygen -D pkcs11
       ssh-keygen -F hostname [-lv] [-f known_hosts_file]
       ssh-keygen -H [-f known_hosts_file]
       ssh-keygen -R hostname [-f known_hosts_file]
       ssh-keygen -r hostname [-g] [-f input_keyfile]
       ssh-keygen -G output_file [-v] [-b bits] [-M memory] [-S start_point]
       ssh-keygen -f input_file -T output_file [-v] [-a rounds] [-J num_lines]
                  [-j start_line] [-K checkpt] [-W generator]
       ssh-keygen -I certificate_identity -s ca_key [-hU] [-D pkcs11_provider]
                  [-n principals] [-O option] [-V validity_interval]
                  [-z serial_number] file ...
       ssh-keygen -L [-f input_keyfile]
       ssh-keygen -A [-f prefix_path]
       ssh-keygen -k -f krl_file [-u] [-s ca_public] [-z version_number]
                  file ...
       ssh-keygen -Q -f krl_file file ...
       ssh-keygen -Y check-novalidate -n namespace -s signature_file
       ssh-keygen -Y sign -f key_file -n namespace file ...
       ssh-keygen -Y verify -f allowed_signers_file -I signer_identity
       		-n namespace -s signature_file [-r revocation_file]
```

使用下面的命令来传输文件

```sh
scp -i ~/.ssh/id_rsa FILENAME USER@SERVER:/home/USER/FILENAME
```

### 通过命令达到不输入密码的效果

```sh
echo "password" | sudo -S command
```

其中 `sudo -S`是指从标准输入中获取密码:

```sh
 -S, --stdin                   read password from standard input
```

### 删除 docker 中的全部容器

```
docker ps -a | awk '{print $1}' | xargs docker rm -f

```

其中 `xargs`的作用是

### awk

1. 打印第一列

```
docker ps -a | awk '{print $1}'

```

2. 指定分隔符
   awk 将空格作为列的分隔符，可以通过`-F`来指定分割符号

获取全部用户:

```
awk -F ":" '{print $1}' /etc/passwd
```

如果要获取多行, 修改指令为'{print $1 $2 ...}', 直接这么用输出会连到一起，所以要这么修改:

```
awk -F ":" '{print $1 " " $2 " " $3}' /etc/passwd
```

在输出元素中间加上空格

3. 正则表达式匹配

使用'/^\//'来搜索以`\`开头的 column

```
awk -F "/" '/^\// {print $NF}' /etc/shells

```

/etc/shells 的内容为:

```
/bin/sh
/bin/bash
/usr/bin/sh
/usr/bin/bash
```

将输出:

```
sh
bash
sh
bash
```

4. 使用`| unique`来过滤重复的内容

下面的命令:

```
awk -F "/" '/^\// {print $NF}' /etc/shells | unique

```

将得到输出:

```
sh
bash
```

5. 过滤内容的长度

```
awk 'length($0) > 7'  /etc/shells
```

6. 使用 if 查找 `./bin/umc`

```
ps -ef | awk '{  if($NF == "./bin/umc") print $0}'
```

7. 打印包含`bash`的行

```
awk  'match($0, /bash/) {print $0 }' /etc/shells
```

### 查询某个文件夹下最近 60 分钟被修改过的文件

"modified file's data" (mtime) means that the content was modified. This date can be manually changed (e.g. with touch).

```
find . -mtime -1  (最近1天)
```

"changed file's status" (ctime) means that either the content was changed, or the file's metadata (permission, owner, etc). This can not be changed manually.

```
find . -mmin -60 (最近60分钟)
```

8. 使用 cut 命令(https://shapeshed.com/unix-cut/)

使用场景:

下面是 Glibc 的版本信息:

```
ldd (GNU libc) 2.28
```

现在需要获取 28 这个小版本号,可以使用`cut -d`:

man:

```
-d, --delimiter=DELIM
              use DELIM instead of TAB for field delimiter
```

```
ldd --version | grep ldd | cut -d' ' -f4 | cut -d. -f2
```
