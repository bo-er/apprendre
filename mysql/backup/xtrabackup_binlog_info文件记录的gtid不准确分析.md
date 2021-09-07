1. 什么是`xtrabackup_binlog_info`文件

为了方便建立从库，Xtrabackup 在备份完成后会将 binlog position 与 GTID 的相关信息保存于 xtrabackup_binlog_info 文件中。
