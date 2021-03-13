# Locking Issues

## How MySQL Locks Tables

You can find a discussion about different locking methods in the appendix. See [Section D.4](https://www.oreilly.com/library/view/mysql-reference-manual/0596002653/apds04.html).

All locking in MySQL is deadlock-free, except for **InnoDB** and **BDB** type tables. This is managed by always requesting all needed locks at once at the beginning of a query and always locking the tables in the same order.

**InnoDB** type tables automatically acquire their row locks and **BDB** type tables their page locks during the processing of SQL statements, not at the start of the transaction.

The locking method MySQL uses for **WRITE** locks works as follows:

- If there are no locks on the table, put a write lock on it.
- Otherwise, put the lock request in the write lock queue.

The locking method MySQL uses for **READ** locks works as follows:

- If there are no write locks on the table, put a read lock on it.
- Otherwise, put the lock request in the read lock queue.

When a lock is released, the lock is made available to the threads in the write lock queue, then to the threads in the read lock queue.

This means that if you have many updates on a table, **SELECT** statements will wait until there are no more updates.

To work around this for the case where you want to do many **INSERT** and **SELECT** operations on a table, you can insert rows in a temporary table and update the real table with the records from the temporary table once in a while.

This can be done with the following code:

```
mysql> LOCK TABLES real_table WRITE, insert_table WRITE;
mysql> INSERT INTO real_table SELECT * FROM insert_table;
mysql> TRUNCATE TABLE insert_table;
mysql> UNLOCK TABLES;
```

You can use the **LOW_PRIORITY** options with  **INSERT**, **UPDATE**,  or **DELETE**, or **HIGH_PRIORITY** with **SELECT** if you want to prioritise retrieval in some specific cases.  You can also start **mysqld** with **--low-priority-updates** to get the same behavior.

Using **SQL_BUFFER_RESULT** can also help make table locks shorter. See [Section 6.4.1](https://www.oreilly.com/library/view/mysql-reference-manual/0596002653/ch06s04.html#mysqlref-CHP-6-SECT-4.1).

You could also change the locking code in `mysys/thr_lock.c` to use a single queue.  In this case, write locks and read locks would have the same priority, which might help some applications.

## Table Locking Issues

The table locking code in MySQL is deadlock free.

MySQL uses table locking (instead of row locking or column locking) on all table types, except **InnoDB** and **BDB** tables, to achieve a very high lock speed.  For large tables, table locking is much better than row locking for most applications, but there are, of course, some pitfalls.

For **InnoDB** and **BDB** tables, MySQL only uses table locking if you explicitly lock the table with **LOCK TABLES**. For these table types we recommend that you not use **LOCK TABLES** at all because **InnoDB** uses automatic row-level locking and **BDB** uses page-level locking to ensure transaction isolation.

In MySQL Versions 3.23.7 and above, you can insert rows into **MyISAM** tables at the same time other threads are reading from the table.  Note that currently this only works if there are no holes after deleted rows in the table at the time the insert is made. When all holes have been filled with new data, concurrent inserts will automatically be enabled again.

Table locking enables many threads to read from a table at the same time, but if a thread wants to write to a table, it must first get exclusive access.  During the update, all other threads that want to access this particular table will wait until the update is ready.

As updates on tables normally are considered to be more important than **SELECT**, all statements that update a table have higher priority than statements that retrieve information from a table. This should ensure that updates are not ‘starved’ because one issues a lot of heavy queries against a specific table. (You can change this by using LOW_PRIORITY with the statement that does the update or **HIGH_PRIORITY** with the **SELECT** statement.)

Starting from MySQL Version 3.23.7 one can use the **max_write_lock_count** variable to force MySQL to temporary give all **SELECT** statements that wait  for a table a higher priority after a specific number of inserts on a table.

Table locking is, however, not very good under the following scenario:

- A client issues a **SELECT** that takes a long time to run.
- Another client then issues an **UPDATE** on a used table. This client will wait until the **SELECT** is finished.
- Another client issues another **SELECT** statement on the same table. As **UPDATE** has higher priority than **SELECT**, this **SELECT** will wait for the **UPDATE** to finish.  It will also wait for the first **SELECT** to finish!
- A thread is waiting for something like **full disk**, in which case all threads that want to access the problem table will also be put in a waiting state until more disk space is made available.

Some possible solutions to this problem are:

- Try to get the **SELECT** statements to run faster. You may have to create some summary tables to do this.
- Start **mysqld** with **--low-priority-updates**.  This will give all statements that update (modify) a table lower priority than a **SELECT** statement. In this case the last **SELECT** statement in the previous scenario would execute before the **INSERT** statement.
- You can give a specific **INSERT**, **UPDATE**, or **DELETE** statement lower priority with the **LOW_PRIORITY** attribute.
- Start **mysqld** with a low value for **max_write_lock_count** to give **READ** locks after a certain number of **WRITE** locks.
- You can specify that all updates from a specific thread should be done with low priority by using the SQL command **SET SQL_LOW_PRIORITY_UPDATES=1**. See [Section 5.5.6](https://www.oreilly.com/library/view/mysql-reference-manual/0596002653/ch05s05.html#mysqlref-CHP-5-SECT-5.6).
- You can specify that a specific **SELECT** is very important with the **HIGH_PRIORITY** attribute. See [Section 6.4.1](https://www.oreilly.com/library/view/mysql-reference-manual/0596002653/ch06s04.html#mysqlref-CHP-6-SECT-4.1).
- If you have problems with **INSERT** combined with **SELECT**, switch to use the new **MyISAM** tables, as these support concurrent **SELECT**s and **INSERT**s.
- If you mainly mix **INSERT** and **SELECT** statements, the **DELAYED** attribute to **INSERT** will probably solve your problems. See [Section 6.4.3](https://www.oreilly.com/library/view/mysql-reference-manual/0596002653/ch06s04.html#mysqlref-CHP-6-SECT-4.3).
- If you have problems with **SELECT** and **DELETE**, the **LIMIT** option to **DELETE** may help. See [Section 6.4.6](https://www.oreilly.com/library/view/mysql-reference-manual/0596002653/ch06s04.html#mysqlref-CHP-6-SECT-4.6).