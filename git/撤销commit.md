# Temporarily switch to a different commit

If you want to temporarily go back to it, fool around, then come back to where you are, all you have to do is check out the desired commit:

```
# This will detach your HEAD, that is, leave you with no branch checked out:

git checkout 0d1d7fc32
```

Or if you want to **make commits while you're there**, go ahead and make a new branch while you're at it:

```
git checkout -b old-state 0d1d7fc32
```

To go back to where you were, just check out the branch you were on again. (If you've made changes, as always when switching branches, you'll have to deal with them as appropriate. You could reset to throw them away; you could stash, checkout, stash pop to take them with you; you could commit them to a branch there if you want a branch there.)
Hard delete unpublished commits

If, on the other hand, you want to really get rid of everything you've done since then, there are two possibilities. One, if you haven't published any of these commits, simply reset:

```
# This will destroy any local modifications.

# Don't do it if you have uncommitted work you want to keep.

git reset --hard 0d1d7fc32

# 如果想要撤销最后一次提交并且把文件放回未提交的状态:

git reset --soft HEAD~1


--soft indicates that the uncommitted files should be retained as working files opposed to

--hard which would discard them.

HEAD~1 is the last commit. If you want to rollback 3 commits you could use HEAD~3. If you want to rollback to a specific revision number, you could also do that using its SHA hash.


# Alternatively, if there's work to keep:

git stash
git reset --hard 0d1d7fc32
git stash pop

# This saves the modifications, then reapplies that patch after resetting.

# You could get merge conflicts, if you've modified things which were

# changed since the commit you reset to.
```

If you mess up, you've already thrown away your local changes, but you can at least get back to where you were before by resetting again.
Undo published commits with new commits

使用下面的命令来查看历史提交:

```
git reflog
```

然后找到刚刚被删除的 commit 再次**git reset --hard + SHA1**即可

On the other hand, if you've published the work, you probably don't want to reset the branch, since that's effectively rewriting history. In that case, you could indeed revert the commits. With Git, revert has a very specific meaning: create a commit with the reverse patch to cancel it out. This way you don't rewrite any history.

```
# This will create three separate revert commits:

git revert a867b4af 25eee4ca 0766c053

# It also takes ranges. This will revert the last two commits:

git revert HEAD~2..HEAD

#Similarly, you can revert a range of commits using commit hashes (non inclusive of first hash):
git revert 0d1d7fc..a867b4a

# Reverting a merge commit

git revert -m 1 <merge_commit_sha>

# To get just one, you could use `rebase -i` to squash them afterwards

# Or, you could do it manually (be sure to do this at top level of the repo)

# get your index and work tree into the desired state, without changing HEAD:

git checkout 0d1d7fc32 .


# Then commit. Be sure and write a good message describing what you just did

git commit
```
