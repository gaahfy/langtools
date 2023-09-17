# Langtools Backend

## How to Run the Backend Service Locally

### Build the Server

```shell
go build server.go
```

### Run the Server

```shell
LANGTOOLS_BACKEND_HTTP_PORT=":8080" ./server
```

Alternatively, you can set `LANGTOOLS_BACKEND_HTTP_PORT`` as an environment
variable on the host.

## How to Deploy the Backend Service

For this section, we will assume that you are using a Debian server with an
amd64 architecture on AWS EC2.

Please adapt the following instructions to match your configuration. Key
assumptions are as follows:

* username is `admin`
* the system utilizes `systemd`
* the domain name is `api.langtools.org`

Feel free to modify these details in the following instructions.

### Initial Setup

#### Initialize Daemon

First, let's create a Bash script to run the server:

```shell
cat > /home/admin/server.sh << EOF
#!/usr/bin/env bash

LANGTOOLS_BACKEND_HTTP_PORT=":80" \\
LANGTOOLS_BACKEND_HTTPS_PORT=":443" \\
LANGTOOLS_BACKEND_DOMAIN_NAME="api.langtools.org" \\
LANGTOOLS_SQL_HOST="sql.langtools.org" \\
LANGTOOLS_SQL_PORT="5432" \\
LANGTOOLS_SQL_USERNAME="langtools_user" \\
LANGTOOLS_SQL_PASSWORD="langtools_password" \\
LANGTOOLS_SQL_DATABASE="langtools" \\
LANGTOOLS_BACKEND_IS_PRODUCTION="yes" /home/admin/server
EOF
```

Ensure that the Bash script is executable:

```shell
chmod +x /home/admin/server.sh
```

Create a systemd service for the server:

```shell
sudo nano /etc/systemd/system/langtools-backend.service
```

```
[Unit]
Description=Backend server for Langtools
After=multi-user.target

[Service]
ExecStart=/usr/bin/bash /home/admin/server.sh
Type=simple

[Install]
WantedBy=multi-user.target
EOF
```

Reload systemd

```shell
sudo systemctl daemon-reload
```

Enable your service:

```shell
sudo systemctl enable langtools-backend.service
```

#### Initial Launch (SSL Setup)

First, build the server and upload it to your running server:

```shell
GOOS=linux GOARCH=amd64 go build server.go
scp -i private-cert.pem ./server admin@server-ip:/home/admin/server
```

Then, open two SSH consoles on your server:

In the first one, create a folder named `letsencrypt` in `/home/admin`:

```shell
mkdir -p /home/admin/letsencrypt
```

In the second console, start the server without HTTPS:

```shell
LANGTOOLS_BACKEND_HTTP_PORT=":80" /home/admin/server
```

In the first console, install Certbot:

```shell
sudo apt install -y certbot
```

Next, use Certbot to generate the certificate:

```shell
sudo certbot certonly --webroot
```

Enter your domain name when prompted, and for the webroot, use the following
path: `/home/admin/letsencrypt`.

You can now return to the second console and press `Ctrl+C` to stop the server.

Run the following command to start the full server with HTTPS:

```shell
sudo systemctl start langtools-backend.service
```

### Updating

Stop the running service:

```shell
sudo systemctl stop langtools-backend.service
```

Remove the previous version of the server:

```shell
rm /home/admin/server
```

Exit the console, and run the following commands on your local machine:

```shell
GOOS=linux GOARCH=amd64 go build server.go
scp -i private-cert.pem ./server admin@server-ip:/home/admin/server
```

Return to the SSH console and run the following command:

```shell
sudo systemctl start langtools-backend.service
```

**_PLEASE BE AWARE THAT THIS CONFIGURATION IS NOT SECURE, AN ISSUE WILL MAKE IT
MORE SECURE ONCE DONE_**