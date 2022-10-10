### master创建用于同步的用户，查看binlog文件信息
```bigquery
docker exec storage_mysql_master bash -c "mysql -uroot -p07b7e8#V5@d1a2K9 -e \"
CREATE USER 'slave'@'%' IDENTIFIED BY '07b7e8#V5@d1a2K9';
GRANT SELECT, REPLICATION SLAVE, REPLICATION CLIENT ON *.* TO 'slave' @'%';
FLUSH PRIVILEGES;
SHOW MASTER STATUS;
select * from information_schema.processlist as p where p.command = 'Binlog Dump'; \" "
```

###slave设置连接到master和复制的位置，启动复制
```bigquery
docker exec storage_mysql_slave1 bash -c "mysql -uroot -p07b7e8#V5@d1a2K9 -e \" 
CHANGE MASTER TO 
MASTER_HOST='192.168.80.242',
MASTER_USER='slave',
MASTER_PASSWORD='07b7e8#V5@d1a2K9',
MASTER_LOG_FILE='mysql-bin.000003',
MASTER_LOG_POS=0, MASTER_PORT=13306;
START SLAVE;flush privileges;\" "
```

### slave启用查看同步状态
```bigquery
docker exec storage_mysql_slave1 bash -c "mysql -uroot -p07b7e8#V5@d1a2K9 -e \"
SHOW SLAVE STATUS;\" "
```


###容器内无法通过主机ip进行连接
``
sysctl -w net.inet.ip.forwarding=1
sysctl -p
sudo sysctl -a | grep forward
``
