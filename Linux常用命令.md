

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

### 删除docker中的全部容器

```
docker ps -a | awk '{print $1}' | xargs docker rm -f   

```

其中 `xargs`的作用是


