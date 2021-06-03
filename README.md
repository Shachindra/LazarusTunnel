# TunnelServicesAPI
Tunnel Services APIs to facilitate Tunnel configuration for Lazarus Network, Functions implemented :
* [Create Tunnel](#create-tunnel)
* [Fetch Tunnel](#fetch-tunnel) 
* [Fetch All Tunnels](#fetch-all-tunnels)
* [Delete Tunnel](#delete-tunnel)

# Installation

# Summary of Functions Implemented

## Caddy API

### Create Tunnel
Creates Web Tunnel with Caddyfile <br>
POST - /api/v1.0/caddy <br>
Request - name <br>
Response - {"message":{"name":"name1","port":"port1"},"status":200} <br>

### Fetch Tunnel
Fetches Web Tunnel configuration and states the status of Port alloted<br>
GET - /api/v1.0/caddy/:name <br>
Request - name <br>
Response - {"message":{"name":"name1","port":"port1", "status":"inactive"},"status":200} <br>

### Fetch All Tunnels
Fetches All Web Tunnel configurations <br>
GET - /api/v1.0/caddy <br>
Request - N/A <br>
Response - {"message":[{"name":"name1","port":"port1"}, {"name":"name2","port":"port2"}],"status":200} <br>

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
Response - {"message":{"name":"name1","port":"port1"},"status":200} <br>

### Fetch Tunnel
Fetches SSH Tunnel configuration <br>
GET - /api/v1.0/nginx/:name <br>
Request - name <br>
Response - {"message":{"name":"name1","port":"port1"},"status":200} <br>

### Fetch All Tunnels
Fetches All SSH Tunnel configurations <br>
GET - /api/v1.0/nginx <br>
Request - N/A <br>
Response - {"message":[{"name":"name1","port":"port1"}, {"name":"name2","port":"port2"}],"status":200} <br>

### Delete Tunnel
Deletes SSH Tunnel configuration <br>
DELETE - /api/v1.0/nginx/:name <br>
Request - name <br>
Response - {"message": ”Deleted Tunnel $name”,"status":200} <br>

## References

## Note