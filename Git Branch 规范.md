## Git Branch 规范

### 

### 分支管理必要性

#### 

#### 分支管理

- 代码提交在应该提交的分支
- 随时可以切换到线上稳定版本代码
- 多个版本的开发工作同时进行

#### 

#### 提交记录的可读性

- 准确的提交描述，具备可检索性
- 合理的提交范围，避免一个功能就一笔提交
- 分支间的合并保有提交历史，且合并后结果清晰明了
- 避免出现过多的分叉

#### 

#### 团队协作

- 明确每个分支的功用，做到对应的分支执行对应的操作
- 合理的提交，每次提交有明确的改动范围和规范的提交信息
- 使用 Git 管理版本迭代、紧急线上 bug fix、功能开发等任务

### 

### 分支说明和操作

#### 

#### master/main 分支

- 主分支，永远处于稳定状态，对应当前线上版本
- 以 tag 标记一个版本，因此在 master 分支上看到的每一个 tag 都应该对应一个线上版本
- 不允许在该分支直接提交代码

#### 

#### develop 分支

- 开发分支，包含了项目最新的功能和代码，所有开发都依赖 develop 分支进行
- develop 分支作为开发的主分支，也不允许直接提交代码。小改动也应该以 feature 分支提 merge request 合并，目的是保证每个改动都经过了强制代码 review，降低代码风险

#### 

#### feature 分支

- 功能分支，开发新功能的分支
- 开发新的功能或者改动较大的调整，从 develop 分支切换出 feature 分支，分支名称为 feature/xxx
- 开发完成后合并回 develop 分支并且删除该 feature/xxx 分支

#### 

#### release 分支

- 发布分支，新功能合并到 develop 分支，准备发布新版本时使用的分支
- 当 develop 分支完成功能合并和部分 bug fix，准备发布新版本时，切出一个 release 分支，来做发布前的准备，分支名约定为release/xxx
- 发布之前发现的 bug 就直接在这个分支上修复，确定准备发版本就合并到 master 分支，完成发布，同时合并到 develop 分支
- product正式环境

#### 

#### fixbug 分支

- 修复已有功能bug的分支，发现解决了bug立即merge到develop进行详细测试

#### 

#### hotfix 分支

- 紧急修复线上 bug 分支
- 当线上版本出现 bug 时，从 master 分支切出一个 hotfix/xxx 分支，完成 bug 修复，然后将 hotfix/xxx 合并到 master 和 develop 分支(如果此时存在 release 分支，则应该合并到 release 分支)，合并完成后删除该  hotfix/xxx 分支

#### 

#### 总结

- master 分支: 线上稳定版本分支
- develop 分支: 开发分支，衍生出 feature 分支和 release 分支
- release 分支: 发布分支，准备待发布版本的分支，存在多个，版本发布之后删除
- feature 分支: 功能分支，完成特定功能开发的分支，存在多个，功能合并之后删除
- hotfix 分支: 紧急热修复分支，存在多个，紧急版本发布之后删除

### 

### 分支间操作规范

#### 

#### 同一分支 git pull 使用 rebase

默认的 `git pull` 使用的是 merge 行为，当你更新代码时，如果本地存在未推送到远程的提交，就会产生一个这样的 merge 提交记录。因此在同一个分支上更新代码时推荐使用 `git pull --rebase`。

下面这张图展示了默认的 `git pull` 和 `git pull --rebase` 的结果差异，使用 git pull --rebase 目的是修整提交线图，使其形成一条直线。 提交前

```
                 A---B---C  remotes/origin/master
                /
           D---E---F---G  master
```

git pull

```
                 A---B---C remotes/origin/master
                /         \
           D---E---F---G---H master
```

git pull --rebase

```
                       remotes/origin/master
                           |
           D---E---A---B---C---F'---G'  master
```

默认的 git pull 行为是 merge，可以进行如下设置修改默认的 git pull 行为：

```
# 为某个分支单独设置，这里是设置 dev 分支
git config branch.dev.rebase true
# 全局设置，所有的分支 git pull 均使用 --rebase
git config --global pull.rebase true
git config --global branch.autoSetupRebase always
```

`git pull --rebase` 能够得到一条很清晰的提交直线图，方便查看提交记录和 `code review`，但是由于 `rebase` 会改变提交历史，也存在一些不好的影响。

#### 

#### 分支合并使用 --no-ff

```
# 例如当前在 develop 分支，需要合并 feature/xxx 分支
  git merge --no-ff feature/xxx
```

在解释这个命令之前，先解释下 Git 中的 fast-forward： 举例来说，开发一直在 develop 分支进行，此时有个新功能需要开发，新建一个 `feature/a` 分支，并在其上进行一系列开发和提交。当完成功能开发时，此时回到 develop 分支，此时 develop 分支在创建 `feature/a` 分支之后没有产生任何的 commit，那么此时的合并就叫做 fast-forward。

fast-forward 合并的结果如下图所示，这种 merge 的结果就是一条直线了，无法明确看到切出一个新的 feature 分支，并完成了一个新的功能开发，因此此时比较推荐使用 `git merge --no-ff`，得到的结果就很明确知道，新的一系列提交是完成了一个新的功能，如果需要对这个功能进行 `code review`，那么只需要检视叉的那条线上的提交即可。

```
                                        B
                                        | \
        A1              A1              |   A1
        |               |               |   |
        A2              A2              |   A2
        |               |               |   |
        A3              A3              |   A3
      /                 |               |  /
    B1                  B1              B1
    |                   |               |
    B2                  B2              B2
    |                   |               |
    B3                  B3              B3
before merge    merge fast-forward    merge --no--ff
```

### 

### 流程示例

这部分内容结合日常项目的开发流程，涉及到开发新功能、分支合并、发布新版本以及发布紧急修复版本等操作，展示常用的命令和操作。

#### 

#### 切到 develop 分支，更新 develop 最新代码

```
git checkout develop
git pull --rebase
```

#### 

#### 新建 feature 分支，开发新功能

```
git checkout -b feature/xxx
...
git add <files>
git commit -m "feat(xxx): commit a"
git commit -m "feat(xxx): commit b"
# 其他提交
...
```

如果此时 develop 分支有一笔提交，影响到你的 feature 开发，可以 rebase develop 分支，前提是 该 feature 分支只有你自己一个在开发，如果多人都在该分支，需要进行协调：

```
# 切换到 develop 分支并更新 develop 分支代码
git checkout develop
git pull --rebase

# 切回 feature 分支
git checkout feature/xxx
git rebase develop

# 如果需要提交到远端，且之前已经提交到远端，此时需要强推(强推需慎重！需要权限)
git push --force
```

#### 

#### 完成 feature 分支，合并到 develop 分支

```
# 切到 develop 分支，更新下代码
git check develop
git pull --rebase

# 合并 feature 分支
git merge feature/xxx --no-ff

# 删除 feature 分支
git branch -d feature/xxx

# 推到远端
git push origin develop
```

#### 

#### 准备发布新版本，提交测试并进行 bug fix

当某个版本所有的 feature 分支均合并到 develop 分支，就可以切出 release 分支。

```
# 当前在 develop 分支
git checkout -b release/xxx

# 在 release/xxx 分支进行 bug fix
git commit -m "fix(xxx): xxxxx"
...
```

#### 

#### 所有 bug 修复完成，准备发布新版本

```
# master 分支合并 release 分支并添加 tag
git checkout master
git merge --no-ff release/xxx --no-ff
# 添加版本标记，这里可以使用版本发布日期或者具体的版本号
git tag 1.0.0

# develop 分支合并 release 分支
git checkout develop
git merge --no-ff release/xxx

# 删除 release 分支
git branch -d release/xxx
```

至此，一个新版本发布完成。

#### 

#### 线上出现 bug，需要紧急发布修复版本

```
# 当前在 master 分支
git checkout master

# 切出 hotfix 分支
git checkout -b hotfix/xxx

... 进行 bug fix 提交

# master 分支合并 hotfix 分支并添加 tag(紧急版本)
git checkout master
git merge --no-ff hotfix/xxx --no-ff
# 添加版本标记，这里可以使用版本发布日期或者具体的版本号
git tag 0.0.6

# develop 分支合并 hotfix 分支(如果此时存在 release 分支的话，应当合并到 release 分支)
git checkout develop
git merge --no-ff hotfix/xxx

# 删除 hotfix 分支
git branch -d hotfix/xxx
```

至此，紧急版本发布完成。

## 

## Git Commit 约定式提交 (Conventional Commits)

git commit 格式 如下：

```
<type>(<scope>): <subject>
// 空一行
<body>
// 空一行
<footer>
```

### 

### Commit 分类 (Type) 有：

- feat: 新功能
- fix: BUG 修复
- docs: 文档变更
- style: 文字格式修改(不影响代码运行的变动)
- refactor: 代码重构(既不增加新功能，也不是修复bug)
- perf: 性能改进
- test: 测试代码
- chore: 工具自动生成
- revert: 回退
- build: 打包

### 

### Scope 修改范围

主要是这次修改涉及到的部分，简单概括，例如 login、opcua

### 

### Subject 修改的描述

具体的修改描述信息，是 commit 目的的简短描述，不超过50个字符。

### 

### Body

Body 部分是对本次 commit 的详细描述，可以分成多行，每行尽量不超过72个字符。

### 

### Footer

Footer 部分只用于两种情况

- 不兼容变动：如果当前代码与上一个版本不兼容，则 Footer 部分以BREAKING CHANGE开头，后面是对变动的描述、以及变动理由和迁移方法。
- 关闭 Issue：如果当前 commit 针对某个issue，那么可以在 Footer 部分关闭这个 issue 。

```
Closes #234
Closes #123, #245, #992
```

### 

### Revert

还有一种特殊情况，如果当前 commit 用于撤销以前的 commit，则必须以revert:开头，后面跟着被撤销 Commit 的 Header。

```
revert: feat(pencil): add 'graphiteWidth' option
This reverts commit 667ecc1654a317a13331b17617d973392f415f02.
```

- feat(detail): 详情页修改样式
- fix(login): 登录页面错误处理
- test(list): 列表页添加测试代码

### 

### 说明

- 控制每笔提交改动的文件尽可能少且集中，避免一次很多文件改动或者多个改动合成一笔。
- 避免重复的提交信息，如果发现上一笔提交没改完整，可以使用 `git commit --amend` 指令追加改动，尽量避免重复的提交信息。

## 

## Git Tag 规范

有 3 位，形如 1.0.0，分别代表：

- 主版本号：当你做了不兼容的 API 修改 ;
- 次版本号：当你做了向下兼容的功能性新增 ;
- 修订号：当你做了向下兼容的问题修正。

## 

## Git emoji

| Emoji | Raw Emoji Code           | Description                                                  |
| ----- | ------------------------ | ------------------------------------------------------------ |
| 🎨     | `:art:`                  | when improving the format/structure of the code              |
| 📰     | `:newspaper:`            | when creating a new file                                     |
| 📝     | `:pencil:`               | when performing minor changes/fixing the code or language    |
| 🐎     | `:racehorse:`            | when improving performance                                   |
| 📚     | `:books:`                | when writing docs                                            |
| 🐛     | `:bug:`                  | when reporting a bug, with @FIXMEComment Tag                 |
| 🚑     | `:ambulance:`            | when fixing a bug                                            |
| 🐧     | `:penguin:`              | when fixing something on Linux                               |
| 🍎     | `:apple:`                | when fixing something on Mac OS                              |
| 🏁     | `:checkered_flag:`       | when fixing something on Windows                             |
| 🔥     | `:fire:`                 | when removing code or files, maybe with @CHANGED Comment Tag |
| 🚜     | `:tractor:`              | when change file structure. Usually together with 🎨          |
| 🔨     | `:hammer:`               | when refactoring code                                        |
| ☔️     | `:umbrella:`             | when adding tests                                            |
| 🔬     | `:microscope:`           | when adding code coverage                                    |
| 💚     | `:green_heart:`          | when fixing the CI build                                     |
| 🔒     | `:lock:`                 | when dealing with security                                   |
| ⬆️     | `:arrow_up:`             | when upgrading dependencies                                  |
| ⬇️     | `:arrow_down:`           | when downgrading dependencies                                |
| ⏩     | `:fast_forward:`         | when forward-porting features from an older version/branch   |
| ⏪     | `:rewind:`               | when backporting features from a newer version/branch        |
| 👕     | `:shirt:`                | when removing linter/strict/deprecation warnings             |
| 💄     | `:lipstick:`             | when improving UI/Cosmetic                                   |
| ♿️     | `:wheelchair:`           | when improving accessibility                                 |
| 🌐     | `:globe_with_meridians:` | when dealing with globalization/internationalization/i18n/g11n |
| 🚧     | `:construction:`         | WIP(Work In Progress) Commits, maybe with @REVIEW Comment Tag |
| 💎     | `:gem:`                  | New Release                                                  |
| 🥚     | `:egg:`                  | New Release with Python egg                                  |
| 🎡     | `:ferris_wheel:`         | New Release with Python wheel package                        |
| 🔖     | `:bookmark:`             | Version Tags                                                 |
| 🎉     | `:tada:`                 | Initial Commit                                               |
| 🔈     | `:speaker:`              | when Adding Logging                                          |
| 🔇     | `:mute:`                 | when Reducing Logging                                        |
| ✨     | `:sparkles:`             | when introducing New Features                                |
| ⚡️     | `:zap:`                  | when introducing Backward-InCompatible Features, maybe with @CHANGED Comment Tag |
| 💡     | `:bulb:`                 | New Idea, with @IDEA Comment Tag                             |
| ❄️     | `:snowflake:`            | changing Configuration, Usually together with 🐧 or 🎀 or 🚀    |
| 🎀     | `:ribbon:`               | Customer requested application Customization, with @HACK Comment Tag |
| 🚀     | `:rocket:`               | Anything related to Deployments/DevOps                       |
| 🐘     | `:elephant:`             | PostgreSQL Database specific (Migrations, Scripts, Extensions, ...) |
| 🐬     | `:dolphin:`              | MySQL Database specific (Migrations, Scripts, Extensions, ...) |
| 🍃     | `:leaves:`               | MongoDB Database specific (Migrations, Scripts, Extensions, ...) |
| 🏦     | `:bank:`                 | Generic Database specific (Migrations, Scripts, Extensions, ...) |
| 🐳     | `:whale:`                | Docker Configuration                                         |
| 🤝     | `:handshake:`            | when Merge files                                             |
| 🍒     | `:cherries:`             | when Commit Arise from one or more Cherry-Pick Commit(s)     |





os.(*File).close STEXT dupok nosplit size=26 args=0x18 locals=0x0
        0x0000 00000 (<autogenerated>:1)        TEXT    os.(*File).close(SB), DUPOK|NOSPLIT|ABIInternal, $0-24
        0x0000 00000 (<autogenerated>:1)        FUNCDATA        $0, gclocals·e6397a44f8e1b6e77d0f200b4fba5269(SB)
        0x0000 00000 (<autogenerated>:1)        FUNCDATA        $1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
        0x0000 00000 (<autogenerated>:1)        MOVQ    ""..this+8(SP), AX
        0x0005 00005 (<autogenerated>:1)        MOVQ    (AX), AX
        0x0008 00008 (<autogenerated>:1)        MOVQ    AX, ""..this+8(SP)
        0x000d 00013 (<autogenerated>:1)        XORPS   X0, X0
        0x0010 00016 (<autogenerated>:1)        MOVUPS  X0, "".~r0+16(SP)
        0x0015 00021 (<autogenerated>:1)        JMP     os.(*file).close(SB)
        0x0000 48 8b 44 24 08 48 8b 00 48 89 44 24 08 0f 57 c0  H.D$.H..H.D$..W.
        0x0010 0f 11 44 24 10 e9 00 00 00 00                    ..D$......
        rel 22+4 t=8 os.(*file).close+0
"".generateClosure STEXT size=145 args=0x8 locals=0x20
        0x0000 00000 (closure.go:5)     TEXT    "".generateClosure(SB), ABIInternal, $32-8
        0x0000 00000 (closure.go:5)     MOVQ    (TLS), CX
        0x0009 00009 (closure.go:5)     CMPQ    SP, 16(CX)
        0x000d 00013 (closure.go:5)     PCDATA  $0, $-2
        0x000d 00013 (closure.go:5)     JLS     135
        0x000f 00015 (closure.go:5)     PCDATA  $0, $-1
        0x000f 00015 (closure.go:5)     SUBQ    $32, SP
        0x0013 00019 (closure.go:5)     MOVQ    BP, 24(SP)
        0x0018 00024 (closure.go:5)     LEAQ    24(SP), BP
        0x001d 00029 (closure.go:5)     FUNCDATA        $0, gclocals·263043c8f03e3241528dfae4e2812ef4(SB)
        0x001d 00029 (closure.go:5)     FUNCDATA        $1, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
        0x001d 00029 (closure.go:6)     LEAQ    type.int(SB), AX
        0x0024 00036 (closure.go:6)     MOVQ    AX, (SP)
        0x0028 00040 (closure.go:6)     PCDATA  $1, $0
        0x0028 00040 (closure.go:6)     CALL    runtime.newobject(SB)
        0x002d 00045 (closure.go:6)     MOVQ    8(SP), AX
        0x0032 00050 (closure.go:6)     MOVQ    AX, "".&sum+16(SP)
        0x0037 00055 (closure.go:8)     LEAQ    type.noalg.struct { F uintptr; "".sum *int }(SB), CX
        0x003e 00062 (closure.go:8)     MOVQ    CX, (SP)
        0x0042 00066 (closure.go:8)     PCDATA  $1, $1
        0x0042 00066 (closure.go:8)     CALL    runtime.newobject(SB)
        0x0047 00071 (closure.go:8)     MOVQ    8(SP), AX
        0x004c 00076 (closure.go:8)     LEAQ    "".generateClosure.func1(SB), CX
        0x0053 00083 (closure.go:8)     MOVQ    CX, (AX)
        0x0056 00086 (closure.go:8)     PCDATA  $0, $-2
        0x0056 00086 (closure.go:8)     CMPL    runtime.writeBarrier(SB), $0
        0x005d 00093 (closure.go:8)     JNE     119
        0x005f 00095 (closure.go:8)     MOVQ    "".&sum+16(SP), CX
        0x0064 00100 (closure.go:8)     MOVQ    CX, 8(AX)
        0x0068 00104 (closure.go:8)     PCDATA  $0, $-1
        0x0068 00104 (closure.go:8)     MOVQ    AX, "".~r0+40(SP)
        0x006d 00109 (closure.go:8)     MOVQ    24(SP), BP
        0x0072 00114 (closure.go:8)     ADDQ    $32, SP
        0x0076 00118 (closure.go:8)     RET
        0x0077 00119 (closure.go:8)     PCDATA  $0, $-2
        0x0077 00119 (closure.go:8)     LEAQ    8(AX), DI
        0x007b 00123 (closure.go:8)     MOVQ    "".&sum+16(SP), CX
        0x0080 00128 (closure.go:8)     CALL    runtime.gcWriteBarrierCX(SB)
        0x0085 00133 (closure.go:8)     JMP     104
        0x0087 00135 (closure.go:8)     NOP
        0x0087 00135 (closure.go:5)     PCDATA  $1, $-1
        0x0087 00135 (closure.go:5)     PCDATA  $0, $-2
        0x0087 00135 (closure.go:5)     CALL    runtime.morestack_noctxt(SB)
        0x008c 00140 (closure.go:5)     PCDATA  $0, $-1
        0x008c 00140 (closure.go:5)     JMP     0
        0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 78 48  eH..%....H;a.vxH
        0x0010 83 ec 20 48 89 6c 24 18 48 8d 6c 24 18 48 8d 05  .. H.l$.H.l$.H..
        0x0020 00 00 00 00 48 89 04 24 e8 00 00 00 00 48 8b 44  ....H..$.....H.D
        0x0030 24 08 48 89 44 24 10 48 8d 0d 00 00 00 00 48 89  $.H.D$.H......H.
        0x0040 0c 24 e8 00 00 00 00 48 8b 44 24 08 48 8d 0d 00  .$.....H.D$.H...
        0x0050 00 00 00 48 89 08 83 3d 00 00 00 00 00 75 18 48  ...H...=.....u.H
        0x0060 8b 4c 24 10 48 89 48 08 48 89 44 24 28 48 8b 6c  .L$.H.H.H.D$(H.l
        0x0070 24 18 48 83 c4 20 c3 48 8d 78 08 48 8b 4c 24 10  $.H.. .H.x.H.L$.
        0x0080 e8 00 00 00 00 eb e1 e8 00 00 00 00 e9 6f ff ff  .............o..
        0x0090 ff                                               .
        rel 5+4 t=17 TLS+0
        rel 32+4 t=16 type.int+0
        rel 41+4 t=8 runtime.newobject+0
        rel 58+4 t=16 type.noalg.struct { F uintptr; "".sum *int }+0
        rel 67+4 t=8 runtime.newobject+0
        rel 79+4 t=16 "".generateClosure.func1+0
        rel 88+4 t=16 runtime.writeBarrier+-1
        rel 129+4 t=8 runtime.gcWriteBarrierCX+0
        rel 136+4 t=8 runtime.morestack_noctxt+0
"".main STEXT size=303 args=0x0 locals=0x70
        0x0000 00000 (closure.go:14)    TEXT    "".main(SB), ABIInternal, $112-0
        0x0000 00000 (closure.go:14)    MOVQ    (TLS), CX
        0x0009 00009 (closure.go:14)    CMPQ    SP, 16(CX)
        0x000d 00013 (closure.go:14)    PCDATA  $0, $-2
        0x000d 00013 (closure.go:14)    JLS     293
        0x0013 00019 (closure.go:14)    PCDATA  $0, $-1
        0x0013 00019 (closure.go:14)    SUBQ    $112, SP
        0x0017 00023 (closure.go:14)    MOVQ    BP, 104(SP)
        0x001c 00028 (closure.go:14)    LEAQ    104(SP), BP
        0x0021 00033 (closure.go:14)    FUNCDATA        $0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
        0x0021 00033 (closure.go:14)    FUNCDATA        $1, gclocals·d527b79a98f329c2ba624a68e7df03d6(SB)
        0x0021 00033 (closure.go:14)    FUNCDATA        $3, "".main.stkobj(SB)
        0x0021 00033 (closure.go:15)    PCDATA  $1, $0
        0x0021 00033 (closure.go:15)    CALL    "".generateClosure(SB)
        0x0026 00038 (closure.go:15)    MOVQ    (SP), DX
        0x002a 00042 (closure.go:15)    MOVQ    DX, "".c+64(SP)
        0x002f 00047 (closure.go:16)    MOVQ    $10, (SP)
        0x0037 00055 (closure.go:16)    MOVQ    (DX), AX
        0x003a 00058 (closure.go:16)    PCDATA  $1, $1
        0x003a 00058 (closure.go:16)    CALL    AX
        0x003c 00060 (closure.go:16)    MOVQ    8(SP), AX
        0x0041 00065 (closure.go:16)    MOVQ    AX, (SP)
        0x0045 00069 (closure.go:16)    CALL    runtime.convT64(SB)
        0x004a 00074 (closure.go:16)    MOVQ    8(SP), AX
        0x004f 00079 (closure.go:16)    XORPS   X0, X0
        0x0052 00082 (closure.go:16)    MOVUPS  X0, ""..autotmp_19+88(SP)
        0x0057 00087 (closure.go:16)    LEAQ    type.int(SB), CX
        0x005e 00094 (closure.go:16)    MOVQ    CX, ""..autotmp_19+88(SP)
        0x0063 00099 (closure.go:16)    MOVQ    AX, ""..autotmp_19+96(SP)
        0x0068 00104 (<unknown line number>)    NOP
        0x0068 00104 ($GOROOT/src/fmt/print.go:274)     MOVQ    os.Stdout(SB), AX
        0x006f 00111 ($GOROOT/src/fmt/print.go:274)     LEAQ    go.itab.*os.File,io.Writer(SB), DX
        0x0076 00118 ($GOROOT/src/fmt/print.go:274)     MOVQ    DX, (SP)
        0x007a 00122 ($GOROOT/src/fmt/print.go:274)     MOVQ    AX, 8(SP)
        0x007f 00127 ($GOROOT/src/fmt/print.go:274)     LEAQ    ""..autotmp_19+88(SP), AX
        0x0084 00132 ($GOROOT/src/fmt/print.go:274)     MOVQ    AX, 16(SP)
        0x0089 00137 ($GOROOT/src/fmt/print.go:274)     MOVQ    $1, 24(SP)
        0x0092 00146 ($GOROOT/src/fmt/print.go:274)     MOVQ    $1, 32(SP)
        0x009b 00155 ($GOROOT/src/fmt/print.go:274)     NOP
        0x00a0 00160 ($GOROOT/src/fmt/print.go:274)     CALL    fmt.Fprintln(SB)
        0x00a5 00165 (closure.go:17)    MOVQ    $5, (SP)
        0x00ad 00173 (closure.go:17)    MOVQ    "".c+64(SP), DX
        0x00b2 00178 (closure.go:17)    MOVQ    (DX), AX
        0x00b5 00181 (closure.go:17)    PCDATA  $1, $0
        0x00b5 00181 (closure.go:17)    CALL    AX
        0x00b7 00183 (closure.go:17)    MOVQ    8(SP), AX
        0x00bc 00188 (closure.go:17)    MOVQ    AX, (SP)
        0x00c0 00192 (closure.go:17)    CALL    runtime.convT64(SB)
        0x00c5 00197 (closure.go:17)    MOVQ    8(SP), AX
        0x00ca 00202 (closure.go:17)    XORPS   X0, X0
        0x00cd 00205 (closure.go:17)    MOVUPS  X0, ""..autotmp_24+72(SP)
        0x00d2 00210 (closure.go:17)    LEAQ    type.int(SB), CX
        0x00d9 00217 (closure.go:17)    MOVQ    CX, ""..autotmp_24+72(SP)
        0x00de 00222 (closure.go:17)    MOVQ    AX, ""..autotmp_24+80(SP)
        0x00e3 00227 (<unknown line number>)    NOP
        0x00e3 00227 ($GOROOT/src/fmt/print.go:274)     MOVQ    os.Stdout(SB), AX
        0x00ea 00234 ($GOROOT/src/fmt/print.go:274)     LEAQ    go.itab.*os.File,io.Writer(SB), CX
        0x00f1 00241 ($GOROOT/src/fmt/print.go:274)     MOVQ    CX, (SP)
        0x00f5 00245 ($GOROOT/src/fmt/print.go:274)     MOVQ    AX, 8(SP)
        0x00fa 00250 ($GOROOT/src/fmt/print.go:274)     LEAQ    ""..autotmp_24+72(SP), AX
        0x00ff 00255 ($GOROOT/src/fmt/print.go:274)     MOVQ    AX, 16(SP)
        0x0104 00260 ($GOROOT/src/fmt/print.go:274)     MOVQ    $1, 24(SP)
        0x010d 00269 ($GOROOT/src/fmt/print.go:274)     MOVQ    $1, 32(SP)
        0x0116 00278 ($GOROOT/src/fmt/print.go:274)     CALL    fmt.Fprintln(SB)
        0x011b 00283 (closure.go:17)    MOVQ    104(SP), BP
        0x0120 00288 (closure.go:17)    ADDQ    $112, SP
        0x0124 00292 (closure.go:17)    RET
        0x0125 00293 (closure.go:17)    NOP
        0x0125 00293 (closure.go:14)    PCDATA  $1, $-1
        0x0125 00293 (closure.go:14)    PCDATA  $0, $-2
        0x0125 00293 (closure.go:14)    CALL    runtime.morestack_noctxt(SB)
        0x012a 00298 (closure.go:14)    PCDATA  $0, $-1
        0x012a 00298 (closure.go:14)    JMP     0
        0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 12  eH..%....H;a....
        0x0010 01 00 00 48 83 ec 70 48 89 6c 24 68 48 8d 6c 24  ...H..pH.l$hH.l$
        0x0020 68 e8 00 00 00 00 48 8b 14 24 48 89 54 24 40 48  h.....H..$H.T$@H
        0x0030 c7 04 24 0a 00 00 00 48 8b 02 ff d0 48 8b 44 24  ..$....H....H.D$
        0x0040 08 48 89 04 24 e8 00 00 00 00 48 8b 44 24 08 0f  .H..$.....H.D$..
        0x0050 57 c0 0f 11 44 24 58 48 8d 0d 00 00 00 00 48 89  W...D$XH......H.
        0x0060 4c 24 58 48 89 44 24 60 48 8b 05 00 00 00 00 48  L$XH.D$`H......H
        0x0070 8d 15 00 00 00 00 48 89 14 24 48 89 44 24 08 48  ......H..$H.D$.H
        0x0080 8d 44 24 58 48 89 44 24 10 48 c7 44 24 18 01 00  .D$XH.D$.H.D$...
        0x0090 00 00 48 c7 44 24 20 01 00 00 00 0f 1f 44 00 00  ..H.D$ ......D..
        0x00a0 e8 00 00 00 00 48 c7 04 24 05 00 00 00 48 8b 54  .....H..$....H.T
        0x00b0 24 40 48 8b 02 ff d0 48 8b 44 24 08 48 89 04 24  $@H....H.D$.H..$
        0x00c0 e8 00 00 00 00 48 8b 44 24 08 0f 57 c0 0f 11 44  .....H.D$..W...D
        0x00d0 24 48 48 8d 0d 00 00 00 00 48 89 4c 24 48 48 89  $HH......H.L$HH.
        0x00e0 44 24 50 48 8b 05 00 00 00 00 48 8d 0d 00 00 00  D$PH......H.....
        0x00f0 00 48 89 0c 24 48 89 44 24 08 48 8d 44 24 48 48  .H..$H.D$.H.D$HH
        0x0100 89 44 24 10 48 c7 44 24 18 01 00 00 00 48 c7 44  .D$.H.D$.....H.D
        0x0110 24 20 01 00 00 00 e8 00 00 00 00 48 8b 6c 24 68  $ .........H.l$h
        0x0120 48 83 c4 70 c3 e8 00 00 00 00 e9 d1 fe ff ff     H..p...........
        rel 5+4 t=17 TLS+0
        rel 34+4 t=8 "".generateClosure+0
        rel 58+0 t=11 +0
        rel 70+4 t=8 runtime.convT64+0
        rel 90+4 t=16 type.int+0
        rel 107+4 t=16 os.Stdout+0
        rel 114+4 t=16 go.itab.*os.File,io.Writer+0
        rel 161+4 t=8 fmt.Fprintln+0
        rel 181+0 t=11 +0
        rel 193+4 t=8 runtime.convT64+0
        rel 213+4 t=16 type.int+0
        rel 230+4 t=16 os.Stdout+0
        rel 237+4 t=16 go.itab.*os.File,io.Writer+0
        rel 279+4 t=8 fmt.Fprintln+0
        rel 294+4 t=8 runtime.morestack_noctxt+0
"".generateClosure.func1 STEXT nosplit size=21 args=0x10 locals=0x0
        0x0000 00000 (closure.go:8)     TEXT    "".generateClosure.func1(SB), NOSPLIT|NEEDCTXT|ABIInternal, $0-16
        0x0000 00000 (closure.go:8)     FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (closure.go:8)     FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (closure.go:8)     MOVQ    8(DX), AX
        0x0004 00004 (closure.go:9)     MOVQ    "".n+8(SP), CX
        0x0009 00009 (closure.go:9)     ADDQ    (AX), CX
        0x000c 00012 (closure.go:9)     MOVQ    CX, (AX)
        0x000f 00015 (closure.go:10)    MOVQ    CX, "".~r1+16(SP)
        0x0014 00020 (closure.go:10)    RET
        0x0000 48 8b 42 08 48 8b 4c 24 08 48 03 08 48 89 08 48  H.B.H.L$.H..H..H
        0x0010 89 4c 24 10 c3                                   .L$..
go.cuinfo.packagename. SDWARFINFO dupok size=0
        0x0000 6d 61 69 6e                                      main
go.info.fmt.Println$abstract SDWARFINFO dupok size=42
        0x0000 04 66 6d 74 2e 50 72 69 6e 74 6c 6e 00 01 01 11  .fmt.Println....
        0x0010 61 00 00 00 00 00 00 11 6e 00 01 00 00 00 00 11  a.......n.......
        0x0020 65 72 72 00 01 00 00 00 00 00                    err.......
        rel 0+0 t=24 type.[]interface {}+0
        rel 0+0 t=24 type.error+0
        rel 0+0 t=24 type.int+0
        rel 19+4 t=29 go.info.[]interface {}+0
        rel 27+4 t=29 go.info.int+0
        rel 37+4 t=29 go.info.error+0
runtime.nilinterequal·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 runtime.nilinterequal+0
runtime.memequal64·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
        0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=16
        0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
type.*interface {} SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 4f 0f 96 9d 08 08 08 36 00 00 00 00 00 00 00 00  O......6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*interface {}-+0
        rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
        0x0000 02                                               .
type.interface {} SRODATA dupok size=80
        0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
        0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        rel 24+8 t=1 runtime.nilinterequal·f+0
        rel 32+8 t=1 runtime.gcbits.02+0
        rel 40+4 t=5 type..namedata.*interface {}-+0
        rel 44+4 t=6 type.*interface {}+0
        rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=18
        0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface 
        0x0010 7b 7d                                            {}
type.*[]interface {} SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 f3 04 9a e7 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[]interface {}-+0
        rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
        0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[]interface {}-+0
        rel 44+4 t=6 type.*[]interface {}+0
        rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}- SRODATA dupok size=19
        0x0000 00 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65  ...*[1]interface
        0x0010 20 7b 7d                                          {}
type.*[1]interface {} SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 bf 03 a8 35 08 08 08 36 00 00 00 00 00 00 00 00  ...5...6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[1]interface {}-+0
        rel 48+8 t=1 type.[1]interface {}+0
type.[1]interface {} SRODATA dupok size=72
        0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
        0x0010 50 91 5b fa 02 08 08 11 00 00 00 00 00 00 00 00  P.[.............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 01 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.nilinterequal·f+0
        rel 32+8 t=1 runtime.gcbits.02+0
        rel 40+4 t=5 type..namedata.*[1]interface {}-+0
        rel 44+4 t=6 type.*[1]interface {}+0
        rel 48+8 t=1 type.interface {}+0
        rel 56+8 t=1 type.[]interface {}+0
""..inittask SNOPTRDATA size=32
        0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
        0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        rel 24+8 t=1 fmt..inittask+0
type..namedata.*struct { F uintptr; sum *int }- SRODATA dupok size=34
        0x0000 00 00 1f 2a 73 74 72 75 63 74 20 7b 20 46 20 75  ...*struct { F u
        0x0010 69 6e 74 70 74 72 3b 20 73 75 6d 20 2a 69 6e 74  intptr; sum *int
        0x0020 20 7d                                             }
type.*struct { F uintptr; "".sum *int } SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 d7 ee 0d 33 08 08 08 36 00 00 00 00 00 00 00 00  ...3...6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*struct { F uintptr; sum *int }-+0
        rel 48+8 t=1 type.noalg.struct { F uintptr; "".sum *int }+0
type..namedata..F- SRODATA dupok size=5
        0x0000 00 00 02 2e 46                                   ....F
type..namedata.sum- SRODATA dupok size=6
        0x0000 00 00 03 73 75 6d                                ...sum
type.noalg.struct { F uintptr; "".sum *int } SRODATA dupok size=128
        0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
        0x0010 9d c9 1b 7e 02 08 08 19 00 00 00 00 00 00 00 00  ...~............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 02 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
        0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0070 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
        rel 32+8 t=1 runtime.gcbits.02+0
        rel 40+4 t=5 type..namedata.*struct { F uintptr; sum *int }-+0
        rel 44+4 t=6 type.*struct { F uintptr; "".sum *int }+0
        rel 48+8 t=1 type..importpath."".+0
        rel 56+8 t=1 type.noalg.struct { F uintptr; "".sum *int }+80
        rel 80+8 t=1 type..namedata..F-+0
        rel 88+8 t=1 type.uintptr+0
        rel 104+8 t=1 type..namedata.sum-+0
        rel 112+8 t=1 type.*int+0
go.itab.*os.File,io.Writer SRODATA dupok size=32
        0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0010 44 b5 f3 33 00 00 00 00 00 00 00 00 00 00 00 00  D..3............
        rel 0+8 t=1 type.io.Writer+0
        rel 8+8 t=1 type.*os.File+0
        rel 24+8 t=1 os.(*File).Write+0
go.itablink.*os.File,io.Writer SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 go.itab.*os.File,io.Writer+0
type..importpath.fmt. SRODATA dupok size=6
        0x0000 00 00 03 66 6d 74                                ...fmt
gclocals·e6397a44f8e1b6e77d0f200b4fba5269 SRODATA dupok size=10
        0x0000 02 00 00 00 03 00 00 00 01 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
        0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·263043c8f03e3241528dfae4e2812ef4 SRODATA dupok size=10
        0x0000 02 00 00 00 01 00 00 00 00 00                    ..........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
        0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
gclocals·d527b79a98f329c2ba624a68e7df03d6 SRODATA dupok size=10
        0x0000 02 00 00 00 05 00 00 00 00 01                    ..........
"".main.stkobj SRODATA size=40
        0x0000 02 00 00 00 00 00 00 00 e0 ff ff ff ff ff ff ff  ................
        0x0010 00 00 00 00 00 00 00 00 f0 ff ff ff ff ff ff ff  ................
        0x0020 00 00 00 00 00 00 00 00                          ........
        rel 16+8 t=1 type.[1]interface {}+0
        rel 32+8 t=1 type.[1]interface {}+0
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
        0x0000 01 00 00 00 00 00 00 00                          ........

