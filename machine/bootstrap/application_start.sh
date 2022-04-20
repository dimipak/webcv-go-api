#!/bin/bash

cd /home/ubuntu/admin

sudo rm -rf base

sudo aws s3 cp s3://webcv-configurations/backend/go/.env.production /home/ubuntu/admin/.env

sudo systemctl start gapi

sudo rm .env