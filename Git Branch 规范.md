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