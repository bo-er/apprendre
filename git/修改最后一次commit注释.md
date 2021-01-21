修改commit注释

## 1.1、修改最后一次commit注释

通过`git log`查看提交历史信息：
![1573523143139](https://segmentfault.com/img/remote/1460000022926069)
输入命令：

```
git commit --amend
```

进入修改注释界面：
![1573522695253](https://segmentfault.com/img/remote/1460000022926067)

第一行就是最后一次commit的注释信息，按`i`键进行编辑状态，修改注释信息后按`Esc`后再按`:wq`保存并退出

再次通过`git log`查看，注释信息由**add test.txt**修改为**新增test.txt**：
![1573523236683](https://segmentfault.com/img/remote/1460000022926068)

## 1.2、修改多次commit注释

命令：

```
# n：需要修改的最近n此commit
git rebase -i HEAD~n
```

比如我想要修改最近3次注释信息就使用`git rebase -i HEAD~3 `，显示下面内容：
![1573523831856](https://segmentfault.com/img/remote/1460000022926071)

> 这上面一行就是一次commit历史，按照提交的顺序进行排序，最下面的一行为最后一次commit

按`i`进行编辑，需要修改那个注释，就将其前面的`pick`修改为`edit`：
![1573524177642](https://segmentfault.com/img/remote/1460000022926070)

> 上面为修改第1行和第3行的注释信息

然后按`Esc`后再按`:wq`保存并退出
此时输入一下命令编辑第1条commit注释：

```
git commit --amend
```

编辑注释信息(按`i`进入编辑状态，按`Esc`和`:wq`保存并退出)，此时分支变为`master|REBASE-i 1/3`；再输入下面信息进行保存：

```
git rebase --continue
```

此时分支变为`master|REBASE-i 3/3`，现在只修改完第1条commit
再通过`git commit --amend`和`git rebase --continue`修改第3条后分支状态变回`master`并提示`Successfully rebased and updated refs/heads/master.`说明已修改完成

# 2、提交到远程仓库

```
# 强制更新到远程仓库
git push -f remote branch
```