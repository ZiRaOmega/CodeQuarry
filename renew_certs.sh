#!/bin/bash

# Stop the application
sudo docker-compose stop app

# Run Certbot to renew the certificate
sudo certbot certonly --standalone -d codequarry.dev --email maxime.diet76@gmail.com --agree-tos --non-interactive --keep; sudo cp -r /etc/letsencrypt/archive/codequarry.dev /cert

# Start the application again
screen -dmS codequarry sh -c 'sudo docker-compose build && sudo docker-compose up'
