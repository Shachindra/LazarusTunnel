# CaddyAPI
Caddy Services APIs to facilitate Tunnel configuration for Lazarus Network, Functions implemented :
* [Create Tunnel](#create-tunnel)
* [Fetch Tunnel](#fetch-tunnel) 
* [Fetch All Tunnels](#fetch-all-tunnels)
* [Delete Tunnel](#delete-tunnel)

# Summary of Functions Implemented

## Create Tunnel
Creates a ssh Tunnel with Updating the content of Caddyfile and Dynamically assigning the Port<br>
POST - /api/v1.0/admin <br>
Request - name <br>
Response - {"message":{"name":"name","port":"port"},"status":200} <br>

## Fetch Tunnel
Fetches Tunnel configuration from Caddyfile <br>
GET - /api/v1.0/admin/:name <br>
Request - name <br>
Response - {"message":{"name":"name","port":"port"},"status":200} <br>

## Fetch All Tunnels
Fetches All Tunnel configurations from Caddyfile <br>
GET - /api/v1.0/admin <br>
Request - N/A <br>
Response - {"message":[{"name":"name1","port":"port1"}, {"name":"name2","port":"port2"}],"status":200} <br>

## Delete Tunnel
Deletes Tunnel configuration from Caddyfile <br>
DELETE - /api/v1.0/admin/:name <br>
Request - name <br>
Response - {"message": ”Deleted Tunnel $name”,"status":200} <br>

## References

## Note