#!/bin/bash

cd /home/ubuntu/admin

sudo systemctl start gapi

sleep 10s

sudo rm .env

sudo rm -rf machine