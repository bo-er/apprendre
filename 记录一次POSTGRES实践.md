```yml
version: "3.0"

services:
  db:
    image: postgres:12
    restart: always
    environment:  
      - POSTGRES_PASSWORD=swsiot
    volumes:
      - ./initdb.sql:/docker-entrypoint-initdb.d/init.sql
      - ./data/postgres:/var/lib/postgresql/data
    ports:
      - "15432:5432"

  swsiot:
    image: sws/swsiot:0.0.4
    ports:
      - "10088:8000"
    volumes:
      - ../configs/:/build/configs/
    restart: always
    depends_on:
      - db    

  swsiot-frontend:
    image: sws/swsiot-frontend:latest
    ports:
      - "18809:80"
    restart: always
    depends_on:
      - swsiot  

volumes:
  postgresql_data:

```

DockerFile:

```dockerfile
FROM golang:1.14.1
WORKDIR /build
COPY . .
RUN apt-get install -y make
RUN make buildvendor
EXPOSE 10088
CMD ["/build/cmd/admin/swsiot-admin", "swsiotadmin", "-c", "./configs/config.toml", "-m", "./configs/model.conf", "--menu", "./configs/menu.yaml"]
```

然后initdb这么写

Initdb.sql :

```
CREATE ROLE swsiot WITH LOGIN PASSWORD 'swsiot';
CREATE DATABASE "swsiot" OWNER = swsiot;
GRANT ALL PRIVILEGES ON DATABASE "swsiot" TO swsiot;
```

Initsb.sh :

```
#!/bin/sh
set -e 

psql --variable=ON_ERROR_STOP=1 --username "postgres" <<-EOSQL
    CREATE ROLE swsiot WITH LOGIN PASSWORD 'swsiot';
    CREATE DATABASE "swsiot" OWNER = swsiot;
    GRANT ALL PRIVILEGES ON DATABASE "swsiot" TO swsiot;
EOSQL
```

