## 更新镜像

查找yaml文件
```
kubectl get pod umc-5cd64d57d-cnsxp -o yaml -n actiontech-dmp
```

编辑yaml文件

```
kubectl  edit deploy umc -n actiontech-dmp
```