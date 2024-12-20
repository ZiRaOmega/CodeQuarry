# Stop the application
sudo docker-compose stop app
#Stop nginx
sudo systemctl stop nginx
# Start the application again
sudo certbot delete --cert-name codequarry.dev --non-interactive
sudo certbot certonly --standalone --expand -d codequarry.dev -d www.codequarry.dev -d codequarry.ovh -d codequarry.ovh --email maxime.diet76@gmail.com --agree-tos --non-interactive --keep
sudo cp -r /etc/letsencrypt/archive/codequarry.dev ./cert/
sudo docker-compose start app

#Restart nginx
sudo systemctl start nginx
