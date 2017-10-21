#!/bin/bash

# Date: 10/20/2017
# Author(s): Spicer Matthews (spicer@options.cafe)
# Copyright: 2017 Cloudmanic Labs, LLC. All rights reserved.
#

source .env

# Build the backend app within the docker container.
cd ../

echo "Building net-worth-server"
env GOOS=linux GOARCH=amd64 go build -o builds/net-worth-server

cd scripts

# Deploy to backend with Ansible
cd ../ansible
ansible-playbook deploy.yml
cd ../scripts

# Login as myself and build and restart
ssh $SSH_SERVER "cd $SSH_DIR && docker-compose build && docker-compose down && docker-compose up -d"

## TODO: make an api call to papertail and output the current logs just to see if anything went wrong during deploy