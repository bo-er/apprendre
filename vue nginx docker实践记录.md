1. 主要的配置是docker-compose.yml跟nginx.conf
2. nginx.conf可以通过volume被容器访问到而不是使用COPY，否则如果配置文件需要修改只能重新生成镜像或者进入容器修改十分不方便
3. 就算使用COPY要注意要把nginx.conf拷贝到docker中的： /etc/nginx/nginx.conf

Docker-compose.yml

```yml
version: "2"

services:
  postgres-serv:
    image: postgres:latest
    container_name: postgres
    environment:
      - POSTGRES_USER=swsiot
      - POSTGRES_PASSWORD=swsiot
      - POSTGRES_DB=swsiot
    user: postgres
    ports:
      - 18286:5432
    networks:
      - frontend

  swsiot-serv:
    image: swsiot:v0.0.6
    container_name: swsiot
    volumes:
      - ./configs:/app/configs/
    ports:
      - 10088:10088      
    depends_on:
      - postgres-serv
    networks:
      - frontend
      - backend

  frontend-serv:
    image: sws/iot-frontend:v0.0.20
    container_name: iot-frontend
    environment:
      - API_URL=http://117.50.31.249:18809
    ports:
      - 18809:80
    depends_on:
      - postgres-serv
      - swsiot-serv
    volumes:
      - ./configs/nginx.conf:/etc/nginx/nginx.conf  
    networks:
      - frontend
      

networks:
  frontend:
    name: iotfrontend
  backend:
    name: iotbackend

```

前端的BASE_URL必须是nginx所在的container对外暴露的地址以及端口，否则nginx不能监听到来自浏览器的请求。

容器启动成功后在ngin(vue)容器中ping swsiot-serv或者 swsiot都可以解析到dns获取到ip

因此nginx.conf这样配置:

```nginx
events {
    worker_connections 1024;
}



http{
     

    server {
        listen 80;
        server_name localhost;   
        root /usr/share/nginx/html;
        index index.html index.htm;   

        location / {
        try_files $uri $uri/ @router; # 配置使用路由
        }

        # 路由配置信息
        location @router {
        rewrite ^.*$ /index.html last;
        }
        location /api/ {
        proxy_pass http://swsiot-serv:10088;
        }
    }
}
```

 这样浏览器访问localhost:18809其实就访问了容器内的80端口，请求能被Nginx拦截到