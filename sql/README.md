# Langtools SQL

## How to run the SQL server locally

In order to run the server locally, run the following command, install
PostgreSQL and a client to connect to it.

Once done, connect to your database with the Postgre user.

Then, create a database for langtools with the following command:

```sql
CREATE DATABASE langtools;
CREATE USER langtools_user WITH ENCRYPTED PASSWORD 'langtools_password';
GRANT ALL PRIVILEGES ON DATABASE langtools TO langtools_user;
ALTER DATABASE langtools OWNER TO langtools_user;
```

Once done, connect to the database with the `langtools_user` user, and run the
create_database script.

## Production

For this section, we will assume that you are using a Debian server with an
amd64 architecture on AWS EC2.

Connect to the server with SSH, and run the following command:

```shell
sudo apt -y install postgresql
```

Then, use the psql command to run the following SQL statement:

**_PLEASE MODIFY `langtools_password` with a stronger password!_**
```shell
CREATE DATABASE langtools;
CREATE USER langtools_user WITH ENCRYPTED PASSWORD 'langtools_password';
GRANT ALL PRIVILEGES ON DATABASE langtools TO langtools_user;
ALTER DATABASE langtools OWNER TO langtools_user;
```

Connect with psql and the `langtools_user`, to run the script in the
`create_database.sql` file.

Then, edit `/etc/postgresql/15/main/pg_hba.conf`, to add the following line:

```shell
host    all             all             0.0.0.0/0               scram-sha-256
```

Then, edit `/etc/postgresql/15/main/postgresql.conf` to uncomment the listen
part with the following:

```shell
listen_addresses = '*'
```

All you have to do now is to restart the server:

```shell
sudo systemctl restart postgresql
```

**_PLEASE BE AWARE THAT THIS CONFIGURATION IS NOT SECURE, AN ISSUE WILL MAKE IT
MORE SECURE ONCE DONE_**