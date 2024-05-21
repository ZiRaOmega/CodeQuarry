#!/bin/bash
echo "Certificate renewed, restarting app service..."
sudo docker-compose -f ./docker-compose.yml restart app
