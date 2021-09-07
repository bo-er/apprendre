GTID Sets
& Related Variables

- Rather than just a single transaction_id, an interval is given
- A range of transactions - 3b22b14c-fe7e-11eb-9992-eab7c2e3693f:1-3787
- Two ranges with a gap - 3b22b14c-fe7e-11eb-9992-eab7c2e3693f:1-37,980-1230
- Commonly used variables related to GTID
  - server_uuid
  - enforce_gtid_consistency
  - gtid_mode
  - gtid_next

传统的方式是通过文件，以及文件的位置来获取 transaction,但是有了 GTID 出现了一个新的概念叫做:
`MASTER_AUTO_POSITION` 允许 MySQL 自己决定需要从哪里 pull 日志。

Traditional replication coordinates

- MASTER_LOG_FILE
- MASTER_LOG_POS

GTID replication coordinates

- MASTER_AUTO_POSITION

Enabling GTIDs in MySQL 5.6

- Gurantee master and replica(s) are in sync by enabling read_only on the master and allow replication to catch up.
- Shutdown the master and replica(s)
- And the following to my.cnf
  - enforce-gtid-consistency
  - gtid-mode = ON
  - skip-slave-start
  - log-slave-updates
  - read-only=1
- Start the master and replica(s)
- Start replication

CHANGE MASTER TO MASTER_HOST='server1', MASTER_AUTO_POSITION=1;START SLAVE;

- Disable read_only on the master and remove read_only from the my.cnf.

Retrieved_Gtid_Set

- All GTIDs received from the master
- Reset on:
  - CHANGE MASTER
  - RESET SLAVE
  - server restart(if relay_log_recovery is on)

Executed_GTID_Set

- All GTIDs have been written to binary log
- Same value seen in:
  - SHOW MASTER STATUS
  - SHOW SLAVE STATUS
  - gtid_executed variable

## Maintenance

- Determine the currently writing master

  Show slave status can be misleading. Master_UUID can be misleading as to who is
  writing.
  如果架构是主-中继-从，如果在从上执行`SHOW SLAVE STATUS`看到的 Master_UUID 将会是中继的 UUID,
  但是实际上写入的是主而不是中继。

- GTID set gaps

  Gaps within a GTID set occur when

  - slave_parrallel_workers > 1
  - A transaction is missing

  下面是 GTID 丢失的例子:
  Executed_Gtid_Set
  b22b14c-fe7e-11eb-9992-eab7c2e3693f:1-111:113-120(GTID 112 丢失)

- Finding transactions in the binary logs

  使用 mysqlbinlog 工具

  - include_gtids
  - exclude_gtids
  - Beware transactions without gtid_next(gtid_mode=OFF)

- Fixing transactions with gtid_next
  - Accidental writes on a replica happen(set sql_log_bin=0)
  - Apply DDL to replicas for very large tables then promote
  - Realign GTID sets to match the recorded direct write to the replica
- Faking and skipping transactions
  set gtid_nect='xxx_gtid_xxx:nnn';BEGIN; COMMIT;

## Advance Concepts

- GTID variables
  - gtid_executed
    Same information as seen in SHOW MASTER/SLAVE STATUS
  - gtid_purged
    Subset of gtid_executed that are no longer in the binary logs
- binlog_gtid_simple_recovery
- GTID Set functions
- START SLAVE UNTIL...
- SHOW SLAVE STATUS NONBLOCKING

## Definitions

### Command

- RESET MASTER

For a server where binary logging is enabled (log_bin is ON), RESET MASTER deletes all existing binary log files and resets the binary log index file, resetting the server to its state before binary logging was started. A new empty binary log file is created so that binary logging can be restarted.

For a server where GTIDs are in use (gtid_mode is ON), issuing RESET MASTER resets the GTID execution history. The value of the gtid_purged system variable is set to an empty string (''), the global value (but not the session value) of the gtid_executed system variable is set to an empty string, and the mysql.gtid_executed table is cleared (see mysql.gtid_executed Table). If the GTID-enabled server has binary logging enabled, RESET MASTER also resets the binary log as described above. Note that RESET MASTER is the method to reset the GTID execution history even if the GTID-enabled server is a replica where binary logging is disabled; RESET SLAVE has no effect on the GTID execution history. For more information on resetting the GTID execution history, see Resetting the GTID Execution History.

### Variables

- gtid_purged

  gtid_purged holds the GTIDs of all transactions that have been applied on the server,
  but do not exist on any binary log file on the server. gtid_purged is a subset of gtid_executed. 
  The following categories of GTIDs are in gtid_purged:
  - GTIDs of replicated transactions that were committed with binary logging disabled on the replica.
  - GTIDs of transactions that were written to a binary log file that has now been purged.
  - GTIDs that were added explicitly to the set by the statement SET @@GLOBAL.gtid_purged.