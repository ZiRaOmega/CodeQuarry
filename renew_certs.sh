#!/bin/bash

# Stop the application
sudo docker-compose stop app

# Start the application again
sudo certbot certonly --standalone --expand -d codequarry.dev -d www.codequarry.dev -d codequarry.ovh --email maxime.diet76@gmail.com --agree-tos --non-interactive --keep
sudo cp -r /etc/letsencrypt/archive/codequarry.dev ./cert/codequarry.dev
sudo docker-compose start app
