
# shubham, 6081, 2022-10-03T04:38:19Z
server {
	listen 6000;
	server_name shubham.sshtun.mydomain.com;
	
	location / {
		proxy_pass http://127.0.0.1:6081$request_uri;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header Host $host;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_buffering off;
		proxy_redirect off;
	}
}

# aman, 6887, 2022-10-03T04:44:31Z
server {
	listen 6000;
	server_name aman.sshtun.mydomain.com;
	
	location / {
		proxy_pass http://127.0.0.1:6887$request_uri;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header Host $host;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_buffering off;
		proxy_redirect off;
	}
}

# shnjfjn, 6318, 2022-10-03T05:46:17Z
server {
	listen 6000;
	server_name shnjfjn.sshtun.mydomain.com;
	
	location / {
		proxy_pass http://127.0.0.1:6318$request_uri;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header Host $host;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_buffering off;
		proxy_redirect off;
	}
}

# basant, 6847, 2022-10-03T10:21:12Z
server {
	listen 6000;
	server_name basant.sshtun.mydomain.com;
	
	location / {
		proxy_pass http://127.0.0.1:6847$request_uri;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header Host $host;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_buffering off;
		proxy_redirect off;
	}
}
