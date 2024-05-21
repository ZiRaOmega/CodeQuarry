#!/bin/bash

# Stop the application
sudo docker-compose stop app

# Start the application again
sudo certbot certonly --standalone -d codequarry.dev --email maxime.diet76@gmail.com --agree-tos --non-interactive --keep; sudo cp -r /etc/letsencrypt/archive/codequarry.dev ./cert && sudo docker-compose build && sudo docker-compose up'
