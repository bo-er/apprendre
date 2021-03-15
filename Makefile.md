- 如何在 Makefile 中检查 os 是否存在某个命令
  ```makefile
   ifeq (, $(shell which lzop))
  $(error "No lzop in $(PATH), consider doing apt-get install lzop")
  endif
  ```
- 如何在 Makefile 中检查 os 是否存在多个命令

  ```makefile
  EXECUTABLES = ls dd dudu lxop
  K := $(foreach exec,$(EXECUTABLES),\
  $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH")))
  ```

