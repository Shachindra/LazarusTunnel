
# prince2, 6081, 2022-09-25T07:09:22Z
server {
	listen 6000;
	server_name prince2.sshtun.mydomain.com;
	
	location / {
		proxy_pass http://127.0.0.1:6081$request_uri;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header Host $host;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_buffering off;
		proxy_redirect off;
	}
}

# prince3, 6887, 2022-09-25T07:13:13Z
server {
	listen 6000;
	server_name prince3.sshtun.mydomain.com;
	
	location / {
		proxy_pass http://127.0.0.1:6887$request_uri;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header Host $host;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_buffering off;
		proxy_redirect off;
	}
}

# prince4, 6847, 2022-09-25T07:16:25Z
server {
	listen 6000;
	server_name prince4.sshtun.mydomain.com;
	
	location / {
		proxy_pass http://127.0.0.1:6847$request_uri;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header Host $host;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_buffering off;
		proxy_redirect off;
	}
}

# creta, 6059, 2022-09-26T05:19:57Z
server {
	listen 6000;
	server_name creta.sshtun.mydomain.com;
	
	location / {
		proxy_pass http://127.0.0.1:6059$request_uri;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header Host $host;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_buffering off;
		proxy_redirect off;
	}
}

# creta2, 6318, 2022-09-26T05:21:00Z
server {
	listen 6000;
	server_name creta2.sshtun.mydomain.com;
	
	location / {
		proxy_pass http://127.0.0.1:6318$request_uri;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header Host $host;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_buffering off;
		proxy_redirect off;
	}
}

# creta3, 6425, 2022-09-26T07:31:10Z
server {
	listen 6000;
	server_name creta3.sshtun.mydomain.com;
	
	location / {
		proxy_pass http://127.0.0.1:6425$request_uri;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header Host $host;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_buffering off;
		proxy_redirect off;
	}
}

# magic, 6540, 2022-09-26T08:09:55Z
server {
	listen 6000;
	server_name magic.sshtun.mydomain.com;
	
	location / {
		proxy_pass http://127.0.0.1:6540$request_uri;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header Host $host;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_buffering off;
		proxy_redirect off;
	}
}

# lolop, 6300, 2022-09-26T09:55:32Z
server {
	listen 6000;
	server_name lolop.sshtun.mydomain.com;
	
	location / {
		proxy_pass http://127.0.0.1:6300$request_uri;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header Host $host;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_buffering off;
		proxy_redirect off;
	}
}

# ayuhs, 6694, 2022-09-26T10:59:21Z
server {
	listen 6000;
	server_name ayuhs.sshtun.mydomain.com;
	
	location / {
		proxy_pass http://127.0.0.1:6694$request_uri;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header Host $host;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_buffering off;
		proxy_redirect off;
	}
}

# nitish, 6511, 2022-09-26T11:13:42Z
server {
	listen 6000;
	server_name nitish.sshtun.mydomain.com;
	
	location / {
		proxy_pass http://127.0.0.1:6511$request_uri;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header Host $host;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_buffering off;
		proxy_redirect off;
	}
}

# abhinav, 6162, 2022-09-26T11:31:52Z
server {
	listen 6000;
	server_name abhinav.sshtun.mydomain.com;
	
	location / {
		proxy_pass http://127.0.0.1:6162$request_uri;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header Host $host;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_buffering off;
		proxy_redirect off;
	}
}

# granted, 6089, 2022-09-26T11:39:56Z
server {
	listen 6000;
	server_name granted.sshtun.mydomain.com;
	
	location / {
		proxy_pass http://127.0.0.1:6089$request_uri;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header Host $host;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_buffering off;
		proxy_redirect off;
	}
}

# pushpa, 6728, 2022-09-26T11:41:49Z
server {
	listen 6000;
	server_name pushpa.sshtun.mydomain.com;
	
	location / {
		proxy_pass http://127.0.0.1:6728$request_uri;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header Host $host;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_buffering off;
		proxy_redirect off;
	}
}
