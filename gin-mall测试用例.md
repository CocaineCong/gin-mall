# 说明

go中的mysql需要自己建数据仓库

```bash
mysql> create database mall_db_test charset=utf8mb4;

mysql> show databases;
```

需要手动安装redis

```bash
sudo apt install redis
```

# 注册
```bash
curl --location --request POST "http://localhost:3000/api/v1/user/register" \
--form "nick_name=雨夜之光" \
--form "user_name=uos" \
--form "password=1"    \
--form "key=1111111111111111"
```

# 登录
```bash
curl --location --request POST "http://localhost:3000/api/v1/user/login" --form "user_name=uos" --form "password=1"
```

# 修改头像 Authorization为用户登录时系统生成的token
```bash
curl --location --request POST "http://localhost:3000/api/v1/avatar" \
  --header "Authorization:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ1b3MiLCJhdXRob3JpdHkiOjAsImV4cCI6MTY3ODc4NTUzOSwiaXNzIjoibWFsbCJ9.N5qnhSoN65otzZ5_tCjj64OuImHNgKJ_H4q-Uyzi8zI" \
  --form "file=@/home/uos/Desktop/my_photo.png"
```

# 获取用户金钱 key为用户注册时指定的16位长度的key
```bash
curl --location --request POST "http://localhost:3000/api/v1/money" \
  --header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ1b3MiLCJhdXRob3JpdHkiOjAsImV4cCI6MTY3ODc4NTUzOSwiaXNzIjoibWFsbCJ9.N5qnhSoN65otzZ5_tCjj64OuImHNgKJ_H4q-Uyzi8zI" \
  --form "key=1111111111111111"
```

# 创建商品
```bash
curl --location --request POST "http://localhost:3000/api/v1/product" \
  --header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ1b3MiLCJhdXRob3JpdHkiOjAsImV4cCI6MTY3ODc4NTUzOSwiaXNzIjoibWFsbCJ9.N5qnhSoN65otzZ5_tCjj64OuImHNgKJ_H4q-Uyzi8zI" \
  --form "name=电脑" \
  --form "category_id=1" \
  --form "title=联想品牌电脑" \
  --form "info=Intel i5 16G内存" \
  --form "price=2000" \
  --form "discount_price=1800" \
  --form "file=@/home/uos/Desktop/pc1.jpg" \
  --form "file=@/home/uos/Desktop/pc2.jpg" \
  --form "file=@/home/uos/Desktop/pc3.jpg"
```

# 查询所有商品信息
```bash
curl --location --request GET "http://localhost:3000/api/v1/products"
```

# 展示特定商品信息
```bash
curl --location --request GET "http://localhost:3000/api/v1/product/2"
```

# 搜索特定商品信息
```bash
curl --location --request POST "http://localhost:3000/api/v1/products"
```

# 绑定邮箱操作 step1　
curl --location --request POST "http://localhost:3000/api/v1/user/sending-email" \
  --header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ1b3MiLCJhdXRob3JpdHkiOjAsImV4cCI6MTY3OTEwNjI4MiwiaXNzIjoibWFsbCJ9.SpFMHW91Mpueso0o3_mKfpYRTshYTMPQWuJbYPzD554" \
  --form "email=hqh2010_9@163.com" \
  --form "operation_type=1"

# 验证邮箱 step2
curl --location --request POST "http://localhost:3000/api/v1/user/valid-email" \
  --header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJlbWFpbCI6ImhxaDIwMTBfOUAxNjMuY29tIiwicGFzc3dvcmQiOiJoMTg2Mjc3NzY4MzYiLCJvcGVyYXRpb25fdHlwZSI6MSwiZXhwIjoxNjc5MDIyNzY5LCJpc3MiOiJjbWFsbCJ9.2yUUAKTUFK9YKd5dQ2OWiN4RYYW8ahpjskd8DNHpPpw" \
  --form "operation_type=1"

# 创建商品分类

```bash
curl --location --request POST "http://localhost:3000/api/v1/category" --form "category_id=1" --form "category_name=电子产品"
```

# 查询商品分类 

```bash
curl --location --request GET "http://localhost:3000/api/v1/categories"
```

# 查询购物车
```bash
curl --location --request GET "http://localhost:3000/api/v1/carts/1" \
  --header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ1b3MiLCJhdXRob3JpdHkiOjAsImV4cCI6MTY3ODc4NTUzOSwiaXNzIjoibWFsbCJ9.N5qnhSoN65otzZ5_tCjj64OuImHNgKJ_H4q-Uyzi8zI"
```

# 创建购物车 boss_id 为商品所属用户id product_id 为产品id
```bash
curl --location --request POST "http://localhost:3000/api/v1/carts" \
  --header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ1b3MiLCJhdXRob3JpdHkiOjAsImV4cCI6MTY3ODc4NTUzOSwiaXNzIjoibWFsbCJ9.N5qnhSoN65otzZ5_tCjj64OuImHNgKJ_H4q-Uyzi8zI" \
--form "boss_id=1" \
--form "product_id=1"
```

# 创建地址
```bash
curl --location --request POST "http://localhost:3000/api/v1/addresses" \
  --header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ1b3MiLCJhdXRob3JpdHkiOjAsImV4cCI6MTY3ODc4NTUzOSwiaXNzIjoibWFsbCJ9.N5qnhSoN65otzZ5_tCjj64OuImHNgKJ_H4q-Uyzi8zI" \
  --form "name=胡先生" \
  --form "address=湖北省黄冈市蕲春" \
  --form "phone=18567774836"
```

```bash
curl --location --request POST "http://localhost:3000/api/v1/addresses" \
  --header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ1b3MiLCJhdXRob3JpdHkiOjAsImV4cCI6MTY3ODc4NTUzOSwiaXNzIjoibWFsbCJ9.N5qnhSoN65otzZ5_tCjj64OuImHNgKJ_H4q-Uyzi8zI" \
  --form "name=刘先生" \
  --form "address=湖北省黄冈市麻城" \
  --form "phone=18567774836"
```

# 获取地址 addresses后面的id为地址编号，一个用户可有多个地址
```bash
curl --location --request GET "http://localhost:3000/api/v1/addresses/2" \
  --header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ1b3MiLCJhdXRob3JpdHkiOjAsImV4cCI6MTY3ODc4NTUzOSwiaXNzIjoibWFsbCJ9.N5qnhSoN65otzZ5_tCjj64OuImHNgKJ_H4q-Uyzi8zI"
```

# 获取用户地址列表
```bash
curl --location --request GET "http://localhost:3000/api/v1/addresses" \
  --header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ1b3MiLCJhdXRob3JpdHkiOjAsImV4cCI6MTY3ODc4NTUzOSwiaXNzIjoibWFsbCJ9.N5qnhSoN65otzZ5_tCjj64OuImHNgKJ_H4q-Uyzi8zI"
```

# 创建订单
```bash
curl --location --request POST "http://localhost:3000/api/v1/orders" \
  --header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ1b3MiLCJhdXRob3JpdHkiOjAsImV4cCI6MTY3ODc4NTUzOSwiaXNzIjoibWFsbCJ9.N5qnhSoN65otzZ5_tCjj64OuImHNgKJ_H4q-Uyzi8zI" \
  --form "product_id=1" \
  --form "num=1" \
  --form "address_id=1" \
  --form "money=180" \
  --form "boss_id=1"
```

# 查询订单
```bash
curl --location --request GET "http://localhost:3000/api/v1/orders" \
  --header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ1b3MiLCJhdXRob3JpdHkiOjAsImV4cCI6MTY3ODc4NTUzOSwiaXNzIjoibWFsbCJ9.N5qnhSoN65otzZ5_tCjj64OuImHNgKJ_H4q-Uyzi8zI"
```

# 获取订单详情
```bash
curl --location --request GET "http://localhost:3000/api/v1/orders/2" \
  --header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ1b3MiLCJhdXRob3JpdHkiOjAsImV4cCI6MTY3ODc4NTUzOSwiaXNzIjoibWFsbCJ9.N5qnhSoN65otzZ5_tCjj64OuImHNgKJ_H4q-Uyzi8zI"
```


# 订单支付
```bash
curl --location --request POST "http://localhost:3000/api/v1/paydown" \
  --header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ1b3MiLCJhdXRob3JpdHkiOjAsImV4cCI6MTY3ODc4NTUzOSwiaXNzIjoibWFsbCJ9.N5qnhSoN65otzZ5_tCjj64OuImHNgKJ_H4q-Uyzi8zI" \
  --form "key=1111111111111111" \
  --form "order_id=3" \
  --form "product_id=1" \
  --form "boss_id=1" \
  --form "money=10000"
```

# 支付后查余额
```bash
curl --location --request POST "http://localhost:3000/api/v1/money" \
  --header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ1b3MiLCJhdXRob3JpdHkiOjAsImV4cCI6MTY3ODc4NTUzOSwiaXNzIjoibWFsbCJ9.N5qnhSoN65otzZ5_tCjj64OuImHNgKJ_H4q-Uyzi8zI" \
  --form "key=1111111111111111"
```

# gin-mall 对应的mysql库信息

```bash
uos@uos:~/Desktop/gin-mall$ mysql -u root -p
Enter password: 
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 48
Server version: 8.0.32 MySQL Community Server - GPL

Copyright (c) 2000, 2023, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mall_db_test       |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
5 rows in set (0.00 sec)

mysql> use mall_db_test;
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Database changed
mysql> show tables;
+------------------------+
| Tables_in_mall_db_test |
+------------------------+
| address                |
| admin                  |
| carousel               |
| cart                   |
| category               |
| favorite               |
| notice                 |
| order                  |
| product                |
| product_img            |
| skill_goods            |
| user                   |
+------------------------+
12 rows in set (0.00 sec)

mysql> select * from user;
+----+---------------------+---------------------+------------+-----------+-------+--------------------------------------------------------------+--------------+--------+------------+--------------------------+
| id | created_at          | updated_at          | deleted_at | user_name | email | password_digest                                              | nick_name    | status | avatar     | money                    |
+----+---------------------+---------------------+------------+-----------+-------+--------------------------------------------------------------+--------------+--------+------------+--------------------------+
|  1 | 2023-03-13 08:02:11 | 2023-03-13 08:02:11 | NULL       | uos       |       | $2a$12$dve5.FZ6/VTM87.HfDjXwu6RnGX.HHUqUmo3kF3/x6SLqcAC5pmhW | 雨夜之光     | active | avatar.JPG | ASxZV3mccL7i7hB9r24e9g== |
+----+---------------------+---------------------+------------+-----------+-------+--------------------------------------------------------------+--------------+--------+------------+--------------------------+
1 row in set (0.00 sec)

mysql> select * from product;
+----+---------------------+---------------------+------------+--------+-------------+--------------------+--------------------+------------------+-------+----------------+---------+------+---------+-----------+---------------+
| id | created_at          | updated_at          | deleted_at | name   | category_id | title              | info               | img_path         | price | discount_price | on_sale | num  | boss_id | boss_name | boss_avatar   |
+----+---------------------+---------------------+------------+--------+-------------+--------------------+--------------------+------------------+-------+----------------+---------+------+---------+-----------+---------------+
|  1 | 2023-03-14 01:45:51 | 2023-03-14 01:45:51 | NULL       | 电脑   |           1 | 联想品牌电脑       | Intel i5 16G内存   | boss1/电脑.jpg   | 2000  | 1800           |       1 |    0 |       1 | uos       | user1/uos.jpg |
+----+---------------------+---------------------+------------+--------+-------------+--------------------+--------------------+------------------+-------+----------------+---------+------+---------+-----------+---------------+
1 row in set (0.00 sec)

mysql> select * from cart;
+----+---------------------+---------------------+------------+---------+------------+---------+------+---------+-------+
| id | created_at          | updated_at          | deleted_at | user_id | product_id | boss_id | num  | max_num | check |
+----+---------------------+---------------------+------------+---------+------------+---------+------+---------+-------+
|  1 | 2023-03-14 03:29:42 | 2023-03-14 03:29:42 | NULL       |       1 |          1 |       1 |    1 |      10 |     0 |
+----+---------------------+---------------------+------------+---------+------------+---------+------+---------+-------+
1 row in set (0.00 sec)

mysql> select * from address;
+----+---------------------+---------------------+------------+---------+-----------+-------------+--------------------------+
| id | created_at          | updated_at          | deleted_at | user_id | name      | phone       | address                  |
+----+---------------------+---------------------+------------+---------+-----------+-------------+--------------------------+
|  1 | 2023-03-14 06:27:14 | 2023-03-14 06:27:14 | NULL       |       1 | 胡先生    | 18567774836 | 湖北省黄冈市蕲春         |
+----+---------------------+---------------------+------------+---------+-----------+-------------+--------------------------+
1 row in set (0.00 sec)

mysql> select * from `order`;
+----+---------------------+---------------------+------------+---------+------------+---------+------------+------+-------------+------+-------+
| id | created_at          | updated_at          | deleted_at | user_id | product_id | boss_id | address_id | num  | order_num   | type | money |
+----+---------------------+---------------------+------------+---------+------------+---------+------------+------+-------------+------+-------+
|  1 | 2023-03-14 06:50:48 | 2023-03-14 06:50:48 | NULL       |       1 |          1 |       1 |          1 |    1 | 22064122311 |    1 |   100 |
+----+---------------------+---------------------+------------+---------+------------+---------+------------+------+-------------+------+-------+
1 row in set (0.00 sec)
mysql> select * from category;
+----+---------------------+---------------------+------------+---------------+-------------+
| id | created_at          | updated_at          | deleted_at | category_name | category_id |
+----+---------------------+---------------------+------------+---------------+-------------+
|  1 | 2023-03-16 05:49:23 | 2023-03-16 05:49:23 | NULL       | 电子产品      |           1 |
+----+---------------------+---------------------+------------+---------------+-------------+
2 rows in set (0.01 sec)
```

