# go_friendly_chat

### サーバーの起動
```
bee run -downdoc=true -gendoc=true
```

### swaggerのurl
```
http://127.0.0.1:8080/swagger/
```

### データベース(mysql)の操作
```
# ログイン
mysql --user root --password

# データベースの作成
mysql> CREATE DATABASE chat_db;

# データベースの確認
mysql> show databases;

# データベースの利用
mysql> use chat_db;

# テーブルの確認
mysql> show tables;

# カラムの確認
mysql> SHOW COLUMNS FROM chats;
```