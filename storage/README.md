
####修改密码
```bash
SET PASSWORD = '123456';
SET PASSWORD FOR 'root' = PASSWORD('123456');
SET PASSWORD FOR 'root'@'localhost'=PASSWORD('123456');
https://www.cnblogs.com/leaderjs/p/13206915.html
```


### 修改redis密码
```bash
config get requirepass
config set requirepass 123123
```

#### 导入sql timestarmp=null出错
```bash
set global explicit_defaults_for_timestamp = ON;
https://www.jianshu.com/p/523a0bf27095
```


###导出
```bigquery
/usr/local/bin/docker exec -i mysql_master_1 /usr/bin/mysqldump --opt --single-transaction --master-data=2 -R -h192.168.80.94 -P23306  -uroot -phiDronedb2020.  hievidence  > ~/data/db_export/hievidence.sql
```

###导入
```bigquery
docker cp hievidence.sql mysql_its_license:/root/
docker exec mysql_its_license bash -c "mysql -uroot -phiDronedb2020. -e \"use hievidence;source /root/hievidence.sql; \" "
```

