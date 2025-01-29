#  test
```
# run
docker run --name mysql8-dev -p 33061:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql:8.0.29-debian --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci

#clean
docker stop mysql8-dev && docker rm -f mysql8-dev

# exec
docker exec -it mysql8-dev sh
```

# todo 处理时间零值
```
docker exec -it mysql8-dev sh
mysql -uroot -p123456

# zero
SELECT @@GLOBAL.sql_mode;
SET GLOBAL sql_mode = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
```