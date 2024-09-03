# Stop the application
sudo docker-compose stop app
#Stop nginx
sudo systemctl stop nginx
# Start the application again
sudo certbot delete --cert-name domain.name --non-interactive
sudo certbot certonly --standalone --expand -d domain.name -d www.domain.name -d domain.name -d domain.name --email email@adress.gg --agree-tos --no>
sudo cp -r /etc/letsencrypt/archive/domain.name ./cert/
sudo docker-compose start app

#Restart nginx
sudo systemctl start nginx
