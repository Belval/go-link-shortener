#!/bin/bash

# This script will install or update a previous installation of go-link-shortener
if [[ $EUID -ne 0 ]]; then
   echo "This script must be run as root" 
   exit 1
fi

if [ -e /etc/go-link-shortener/config.json ]
then
    echo "Pre-existing configuration found. Updating binaries only..."
else
    echo "No pre-existing configuration found. Creating /etc/go-link-shortener/config.json..."
    mkdir /etc/go-link-shortener/
    cp ./go-link-shortener/config.json /etc/go-link-shortener/
fi

echo "Creating database folder..."
mkdir /var/go-link-shortener/

echo "Stopping service if running..."
systemctl stop go-link-shortener

echo "Updating binaries..."
rm /usr/bin/go-link-shortener
cp ./go-link-shortener/go-link-shortener /usr/bin/go-link-shortener

echo "Updating systemctl unit..."
echo "
[Unit]
Description=A link shortening service!
After=network.target

[Service]
Type=simple
ExecStart=/usr/bin/go-link-shortener -config=/etc/go-link-shortener/config.json

[Install]
WantedBy=multi-user.target" > ./go-link-shortener.service

mv ./go-link-shortener.service /etc/systemd/system/

systemctl daemon-reload
systemctl start go-link-shortener



