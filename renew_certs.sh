#!/bin/bash

# Stop the application
sudo docker-compose stop app

# Run Certbot to renew the certificate
sudo docker-compose run --rm certbot certonly --standalone -d codequarry.dev --non-interactive --agree-tos --email maxime.diet76@gmail.com

# Start the application again
screen -dmS codequarry sh -c 'sudo docker-compose start app'
