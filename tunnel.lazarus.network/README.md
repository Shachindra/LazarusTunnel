# Lazarus Tunnel
Steps to setup the Lazarus Tunnel on the Remote Server

## Download Necessary Software
sudo apt-get update && sudo apt-get upgrade -y
sudo add-apt-repository ppa:certbot/certbot
sudo apt install python-certbot-nginx
sudo apt install nginx
sudo ufw status
sudo systemctl status nginx

## Configure Nginx
sudo nano /etc/nginx/sites-available/tunnel.lazarus.network
sudo ln -s /etc/nginx/sites-available/tunnel.lazarus.network /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl restart nginx

## Configure the LetsEncrypt Folder
sudo mkdir -p /var/lib/letsencrypt/.well-known
sudo chgrp www-data /var/lib/letsencrypt
sudo chmod g+s /var/lib/letsencrypt
sudo nano /etc/nginx/snippets/letsencrypt.conf
sudo nano /etc/nginx/snippets/ssl.conf
sudo nano /etc/nginx/sites-available/tunnel.lazarus.network
sudo nginx -t
sudo systemctl restart nginx

> In another window: tmux new -s nginx / tmux a -t nginx
tail -f /var/log/nginx/error.log

## Get the certs using Certbot
sudo certbot certonly --agree-tos --email connect@lazarus.network -w /var/lib/letsencrypt/ --server https://acme-v02.api.letsencrypt.org/directory -d *.tunnel.lazarus.network --manual --preferred-challenges dns-01
> Comment all the lines to disable autorenew: sudo nano /etc/cron.d/certbot

## Create the acess user
sudo adduser --home /restricted/access access
sudo chown access:access /restricted/access
sudo chmod 755 /restricted/access

## Enable Gateway Ports
echo "GatewayPorts yes" >> /etc/ssh/sshd_config
sudo service ssh restart

## Execute on client
ssh -vnNT -R (port-given-to-you):localhost:(port-you-want-to-expose) access@tunnel.lazarus.network
Password: (password-given-to-you)