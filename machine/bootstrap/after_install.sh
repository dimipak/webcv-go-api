#!/bin/bash

sudo aws s3 cp s3://webcv-configurations/backend/go/.env.production /home/ubuntu/admin/.env

sudo aws s3 cp s3://webcv-configurations/backend/go/dist/admin /home/ubuntu/admin/admin

sudo chmod 755 admin