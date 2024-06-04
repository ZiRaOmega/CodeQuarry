```sh
sudo -u postgres psql
```
```sql
CREATE DATABASE db_name_test OWNER db_user;
GRANT ALL PRIVILEGES ON DATABASE db_name_test TO db_user;
```