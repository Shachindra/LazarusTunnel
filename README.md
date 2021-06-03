# TunnelServicesAPI
Tunnel Services APIs to facilitate Tunnel configuration for Lazarus Network, Functions implemented :
* [Create Tunnel](#create-tunnel)
* [Fetch Tunnel](#fetch-tunnel) 
* [Fetch All Tunnels](#fetch-all-tunnels)
* [Delete Tunnel](#delete-tunnel)

# Installation Notes

After placing the .path and .service files in /etc/systemd/system, Run:
1. ```sudo systemctl daemon-reload```

2. ```sudo systemctl enable caddy-watcher.path && sudo systemctl start caddy-watcher.path```

Created symlink /etc/systemd/system/multi-user.target.wants/caddy-watcher.path → /etc/systemd/system/caddy-watcher.path.

3. ```sudo systemctl enable caddy-watcher.service && sudo systemctl start caddy-watcher.service```

Created symlink /etc/systemd/system/multi-user.target.wants/caddy-watcher.service → /etc/systemd/system/caddy-watcher.service.

4. ```sudo systemctl enable nginx-watcher.path && sudo systemctl start nginx-watcher.path```

Created symlink /etc/systemd/system/multi-user.target.wants/nginx-watcher.path → /etc/systemd/system/nginx-watcher.path.

5. ```sudo systemctl enable nginx-watcher.service && sudo systemctl start nginx-watcher.service```

Created symlink /etc/systemd/system/multi-user.target.wants/nginx-watcher.service → /etc/systemd/system/nginx-watcher.service.

6. ```sudo systemctl status caddy-watcher.path```

7. ```sudo systemctl status caddy-watcher.service```

6. ```sudo systemctl status nginx-watcher.path```

7. ```sudo systemctl status nginx-watcher.service```

# Summary of Functions Implemented

## Caddy API

### Create Tunnel
Creates Web Tunnel with Caddyfile <br>
POST - /api/v1.0/caddy <br>
Request - name <br>
Response - {"message":{"name":"name1","port":"port1", "createdAt":"timeStamp", domain":"caddyDomain"},"status":200} <br>

### Fetch Tunnel
Fetches Web Tunnel configuration and states the status of Port alloted<br>
GET - /api/v1.0/caddy/:name <br>
Request - name <br>
Response - {"message":{"name":"name1","port":"port1", "createdAt":"timeStamp", domain":"caddyDomain", "status":"inactive"},"status":200} <br>

### Fetch All Tunnels
Fetches All Web Tunnel configurations <br>
GET - /api/v1.0/caddy <br>
Request - N/A <br>
Response - {"message":[{"name":"name1","port":"port1", "createdAt":"timeStamp1", domain":"caddyDomain"}, {"name":"name2","port":"port2", "createdAt":"timeStamp2", domain":"caddyDomain"}],"status":200} <br>

### Delete Tunnel
Deletes Web Tunnel configuration from Caddyfile <br>
DELETE - /api/v1.0/caddy/:name <br>
Request - name <br>
Response - {"message": ”Deleted Tunnel $name”,"status":200} <br>

## NGINX API

### Create Tunnel
Creates SSH Tunnel nginx <br>
POST - /api/v1.0/nginx <br>
Request - name <br>
Response - {"message":{"name":"name1","port":"port1", "createdAt":"timeStamp", domain":"nginxDomain"},"status":200} <br>

### Fetch Tunnel
Fetches SSH Tunnel configuration <br>
GET - /api/v1.0/nginx/:name <br>
Request - name <br>
Response - {"message":{"name":"name1","port":"port1"},"status":200} <br>

### Fetch All Tunnels
Fetches All SSH Tunnel configurations <br>
GET - /api/v1.0/nginx <br>
Request - N/A <br>
Response - {"message":[{"name":"name1","port":"port1, "createdAt":"timeStamp1", domain":"nginxDomain""}, {"name":"name2","port":"port2", , "createdAt":"timeStamp2", domain":"nginxDomain"}],"status":200} <br>

### Delete Tunnel
Deletes SSH Tunnel configuration <br>
DELETE - /api/v1.0/nginx/:name <br>
Request - name <br>
Response - {"message": ”Deleted Tunnel $name”,"status":200} <br>

## References

## Note