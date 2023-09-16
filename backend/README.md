# Langtools backend

## How to run the backend service locally

### Build the server

```shell
go build server.go
```

### Run the server

```shell
LANGTOOLS_BACKEND_PORT=":8080" ./server
```

Alternatively, `LANGTOOLS_BACKEND_PORT` can be set as an environment variable of
the host.

## How to deploy the backend service

### Initial Setup

For this part, we will assume that you are running on an amd64 server which has
systemd on it.

First, let's create a bash script which will run the server:

```shell
cat >> /home/admin/server.sh << EOF
#!/usr/bin/env bash

LANGTOOLS_BACKEND_PORT=":80" /home/admin/server
EOF
```

Make sure the bash script can be executed:

```shell
chmod +x /home/admin/server.sh
```

Create a systemd service for the server:

```shell
sudo nano /etc/systemd/system/langtools-backend.service
```

```
[Unit]
Description=Backend server for langtools
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

### Update

Stop the running service :

```shell
sudo systemctl stop langtools-backend.service
```

Remove the previous version of the server

```shell
rm /home/admin/server
```

Leave the console, and run the following commands on your local machine

```
GOOS=linux GOARCH=amd64 go build server.go
scp -i private-cert.pem ./server admin@server-ip:/home/admin/server
```

Go back to the console in ssh, then run the following command:

```shell
sudo systemctl start langtools-backend.service
```