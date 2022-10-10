
###导出
/usr/local/bin/docker exec -i mysql_master_1 /usr/bin/mysqldump --opt --single-transaction --master-data=2 -R -hlocalhost -P23306  -uroot -phiDronedb2020.  hievidence  > ~/data/db_export/hievidence.sql

###导入
docker cp ~/data/db_export/hievidence.sql common_mysql:/root/
docker exec common_mysql bash -c "mysql -uroot -p07b7e8#V5@d1a2K9 -e \"use hievidence;source /root/hievidence.sql; \" "


###导出
/usr/local/bin/docker exec -i mysql_master_1 /usr/bin/mysqldump --opt --single-transaction --master-data=2 -R -hlocalhost -P23306  -uroot -phiDronedb2020.  its_license  > ~/data/db_export/its_license.sql

###导入
docker cp ~/data/db_export/its_license.sql common_mysql:/root/
docker exec common_mysql bash -c "mysql -uroot -p07b7e8#V5@d1a2K9 -e \"use its_license;source /root/its_license.sql; \" "


/usr/local/bin/docker exec -i mysql_master_1 /usr/bin/mysqldump --opt --single-transaction --master-data=2 -R -hlocalhost -P23306  -uroot -phiDronedb2020.  hidrone  > ~/data/db_export/hidrone.sql

docker cp ~/data/db_export/hidrone.sql common_mysql:/root/
docker exec common_mysql bash -c "mysql -uroot -p07b7e8#V5@d1a2K9 -e \"create hidrone; use hidrone;source /root/hidrone.sql; \" "
