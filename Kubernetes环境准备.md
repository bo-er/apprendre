## K8s(Kubernetes) DMP 快速搭建

1. 至少准备两台主机(ubuntu 系统),然后修改主机名称，确保主机名不会重复,使用:

`echo "修改的主机名称" > /etc/hostname` 来修改主机名

2. 安装 kubectl,在全部主机上执行下面的命令:

```
cat <<EOF > /etc/apt/sources.list.d/kubernetes.list
deb http://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main
EOF

curl -s https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg |

sudo apt-key add -
apt-get update
```

3. 到[Rancher](https://10.186.62.103)上创建 cluster, 不需要修改默认配置

4. 进入刚刚在 Rancher 上创建的集群，点击主机，下拉到最下面找到`Node Options`的位置。

- 对于管理节点勾选`etcd` `Controlplane` `Worker` ,复制下面的命令到安装好 docker 的主机上执行命令。

- 对于工作节点只勾选 `Worker`再执行

5. 在 rancher 上获取集群配置信息，并将配置保存到~/.kube/config。

```
mkdir -p $HOME/.kube
vim ~/.kube/config
```

6. 在管理节点上安装`helm3`

```
curl https://get.helm.sh/helm-v3.2.4-linux-amd64.tar.gz --output helm-v3.2.4-linux-amd64.tar.gz
tar -zxvf helm-v3.2.4-linux-amd64.tar.gz
mv linux-amd64/helm /usr/local/bin/helm
```

7. 搭建 NFS Server

a. 在管理节点执行命令:

```
$shell> sudo apt-get update
$shell> sudo apt install nfs-kernel-server

$shell> mkdir -p /mnt/nfs
# 客户端系统上所有组的所有用户都可以访问我们的“共享文件夹”
$shell> sudo chown nobody:nogroup /mnt/nfs
$shell> sudo chmod 777 /mnt/nfs
```

b. 将下面的放到 `/etc/exports`中

```
# /etc/exports: the access control list for filesystems which may be exported
#       to NFS clients.  See exports(5).
#
# Example for NFSv2 and NFSv3:
# /srv/homes       hostname1(rw,sync,no_subtree_check) hostname2(ro,sync,no_subtree_check)
#
# Example for NFSv4:
# /srv/nfs4        gss/krb5i(rw,sync,fsid=0,crossmnt,no_subtree_check)
# /srv/nfs4/homes  gss/krb5i(rw,sync,no_subtree_check)
#
# 10.186.0.0/16网段的客户端可以访问该共享文件夹
/mnt/nfs 10.186.0.0/16(rw,sync,no_subtree_check)
```

c. 导出共享目录

```
$shell> sudo exportfs -a
$shell> sudo systemctl restart nfs-kernel-server
```

d. 在客户端(工作节点)执行下面的命令, `10.186.52.24`需要替换为控制节点的 IP 地址:

```
$shell> sudo apt-get update
$shell> sudo apt-get install nfs-common
$shell> sudo mkdir -p /mnt/nfs_client
$shell> sudo mount -t nfs -o nosuid,noexec,nodev,rw -o bg,soft 10.186.52.24:/mnt/nfs /mnt/nfs_client
```

8. 配置 storage class

在控制节点执行下面的命令,其中`10.186.65.137`是 NFS Server 的 IP 地址,即控制节点的 IP 地址

```
helm repo add apphub https://apphub.aliyuncs.com

helm install nfs-client-provisioner \
--set storageClass.name=nfs-client \
--set storageClass.defaultClass=true \
--set nfs.server=10.186.65.137 \
--set nfs.path=/mnt/nfs/ \
apphub/nfs-client-provisioner
```

---

完成了上面工作那么 K8s 的环境已经配置好了，现在就开始安装 DMP 应用

1. 选择自己的 K8s 集群和项目(项目名称默认为 Default)
2. 进入应用商店界面(选择 actiontech-dmp)
3. 除了 Image.Tag 需要按版本手工调整，如 v4.21.09.0,其他的按需配置
