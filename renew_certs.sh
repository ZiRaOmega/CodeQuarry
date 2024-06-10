#!/bin/bash

# Stop the application
sudo docker-compose stop app

# Start the application again
sudo certbot certonly --standalone --expand -d domain_name -d domain_name_2 -d domaine_name_3 -d domaine_name_4 --email email_address --agree-tos --non-interactive --keep
sudo cp -r /etc/letsencrypt/archive/domain_name ./cert/
sudo docker-compose start app
