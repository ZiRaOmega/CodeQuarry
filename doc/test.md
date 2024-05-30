```sh
sudo -u postgres psql
```
```sql
CREATE DATABASE codequarrytest OWNER codequarry;
GRANT ALL PRIVILEGES ON DATABASE codequarrytest TO codequarry;
```