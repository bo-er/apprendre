​     

## Docker Compose Up With Postgres Quick Tips

*February 20th 2020*

I found myself trying to spin up a docker container recently to  containerize a Postgres database and ran into a few issues that took me  awhile to flesh out and thought I would share.  We are going to assume  that you have a version of Docker Engine and Docker Compose installed on your computer and/or the Docker Toolbox (which includes engine and  compose).  You can easily pull the docker postgres image and get the  container started by running:

```
docker run --name pg-docker --rm -p 5400:5432 -e POSTGRES_PASSWORD=docker -e POSTGRES_USER=docker -d postgres
```

This will pull the image and run a container (in detached mode) and set the environment variable `POSTGRES_PASSWORD=docker` and `POSTGRES_USER=docker` which is used to create the  `docker` role which will have Superuser privileges.  We are also mapping the  host port 5400 to the port 5432 on the container, and the postgres image exposes 5432 already so we can test out login by running:

```
psql -h localhost -p 5400 -U docker postgres
```

You will be prompt for the password we created: `docker` and once logged in you can see that we have one user by typing: `\du`

```
                                   List of roles
 Role name |                         Attributes                         | Member of 
-----------+------------------------------------------------------------+-----------
 docker    | Superuser, Create role, Create DB, Replication, Bypass RLS | {}
```

That was easy enough, but once we stop our container, the database  will not persist our data. Speaking of stopping, let’s stop our  container: `docker stop pg-docker`. We need a way to mount  some volumes, add some configuration, and import/seed our database.  Docker Compose to the rescue! Let’s setup a simple project structure: `mkdir -p docker-postgres/data && cd docker-postgres && touch docker-compose.yml`. So we have a simple project folder with a blank docker-compose.yml file. While we’re in here we can copy the `postgres.conf` from the docker image to our project folder so we can do some configuration inside our container when we run it:

```
docker run -i --rm postgres cat /usr/share/postgresql/postgresql.conf.sample > postgres.conf
```

This runs the `postgres` container quick in interactive mode -i and copies the sample config into our directory.  Now we can dive into our `docker-compose.yml` file to start our definitions:

```
version: '3'

services:
  postgresql:
    image: postgres
    container_name: pg-docker
    ports:
      - "5400:5432"
    environment:
      - POSTGRES_USER=docker
      - POSTGRES_PASSWORD=docker
    volumes:
      - ./postgres.conf:/etc/postgresql/postgresql.conf
    command: postgres -c config_file=/etc/postgresql/postgresql.conf
    restart: always
```

So we have a service that is called postgresql that creates a  container from the postgres image and it defines our environment  variables. It also mounts the postgres.conf file that we copied into our project folder and then runs the command  `config_file=/etc/postgresql/postgresql.conf` which defines the path to the config file we mounted to the container.  We can test by running:

```
docker-compose up -d
```

And as before we can login with:

```
psql -h localhost -p 5400 -U docker postgres
```

Another interesting thing happens when we run this container is that  there is a new docker network that is created.  To see what I mean you  can type: `docker network ls` and you will see:

```
NETWORK ID          NAME                DRIVER              SCOPE
71d9a521ce99        postgres_default    bridge              local
```

We can also inspect our container and see the network it is by running: `docker inspect pg-docker`

```
  "Networks": {
    "postgres_default": {
      "IPAMConfig": null,
      "Links": null,
      "Aliases": [
        "postgresql",
        "2cf746153a09"
      ],
      "NetworkID": "3a03ed57ca73d4ca55d516f1e5043af61d9d2c59dff09a4a7ff6df21700229d8",
      "EndpointID": "b83b7d2b31049ccc8f8b5576f4d6193a356fae8a4a9be069ad3963debbad3f17",
      "Gateway": "192.168.48.1",
      "IPAddress": "192.168.48.2",
      "IPPrefixLen": 20,
      "IPv6Gateway": "",
      "GlobalIPv6Address": "",
      "GlobalIPv6PrefixLen": 0,
      "MacAddress": "02:42:c0:a8:30:02",
      "DriverOpts": null
    }
  }
```

This allows us to connect to our pg-docker container over this  network from other containers. To see what I mean, let’s drop into our  docker container and try to connect from there:

```
docker-compose run postgresql bash
```

The `docker-compose run` command runs a command against our service. As per the [Docker docs](https://docs.docker.com/compose/reference/run/):

> Runs a one-time command against a service. For example, the previous command starts the `postgresql` service and runs `bash` as its command.
>
> docs.docker.com

Once logged in to the terminal we can proceed to connect with:

```
psql -h localhost -U docker -d postgres
```

Whoops! That did not work! 😬 

```
psql: error: could not connect to server: could not connect to server: Connection refused
	Is the server running on host "localhost" (127.0.0.1) and accepting
	TCP/IP connections on port 5432?
```

Hmm, well let’s try over the network that was created earlier! We can either use the IP address of our container: `192.168.48.2` or we can use the alias that it was given: `postgresql`

```
psql -h 192.168.48.2 -U docker -d postgres
//or
psql -h postgresql -U docker -d postgres
```

And Viola! We are into our psql terminal!  You can use this alias or IP address to connect from other containers you create.

Wouldn’t it be handy if we could create a new database and define the table schema structure? Luckily we can with **initialization scripts**. Initialization scripts can be placed in the `/docker-entrypoint-initdb.d/` directory of our container and any script that is put in this directory will be run on initialization.

Let’s create a `schema.sql` file in our project directory:

```
-- schema.sql
-- Since we might run the import many times we'll drop if exists
DROP DATABASE IF EXISTS blog;

CREATE DATABASE blog;

-- Make sure we're using our `blog` database
\c blog;

-- We can create our user table
CREATE TABLE IF NOT EXISTS user (
  id SERIAL PRIMARY KEY,
  username VARCHAR,
  email VARCHAR
);

-- We can create our post table
CREATE TABLE IF NOT EXISTS post (
  id SERIAL PRIMARY KEY,
  userId INTEGER REFERENCES user(id),
  title VARCHAR,
  content TEXT,
  image VARCHAR,
  date DATE DEFAULT CURRENT_DATE
);
```

Now all we have to do is mount our `schema.sql` file into our `/docker-entrypoint-initdb.d/` of our container. Let’s go back into docker-compose.yml  file and bind another volume:

```
version: '3'

services:
  postgresql:
    image: postgres
    container_name: pg-docker
    ports:
      - "5400:5432"
    environment:
      - POSTGRES_USER=docker
      - POSTGRES_PASSWORD=docker
    volumes:
      - ./postgres.conf:/etc/postgresql/postgresql.conf
      - ./schema.sql:/docker-entrypoint-initdb.d/schema.sql
    command: postgres -c config_file=/etc/postgresql/postgresql.conf
    restart: always
```

Now when we run `docker-compose up -d` we can login to our `blog` database to see our table structure.  Login via: `psql -h localhost -p 5400 -U docker -d blog` and run `\dt` and you should see all of the tables we defined in our `schema.sql` file.

One last thing that we want to do is bind a local directory to the postgres `/var/lib/postgresql/data` as this will persist and store our data that is created in our database.  We can use our `./data` directory to store the data. All we need to do is add another volume definition:

```
version: '3'

services:
  postgresql:
    image: postgres
    container_name: pg-docker
    ports:
      - "5400:5432"
    environment:
      - POSTGRES_USER=docker
      - POSTGRES_PASSWORD=docker
    volumes:
      - ./postgres.conf:/etc/postgresql/postgresql.conf
      - ./data:/var/lib/postgresql/data
      - ./schema.sql:/docker-entrypoint-initdb.d/schema.sql
    command: postgres -c config_file=/etc/postgresql/postgresql.conf
    restart: always
```

> The `-v ./data:/var/lib/postgresql/data` part of the command mounts the `./data` directory from the underlying host system as `/var/lib/postgresql/data` inside the container, where PostgreSQL by default will write its data files.
>
> We have a our docker container data storage mounted to our host  system to persist the data as we start and stop our container.  Now you  have a good base for your docker-compose file and services to start a  postgres docker container with some configuration options and volumes  mounted! Enjoy exploring all the possibilities. Until next time, stay  curious, stay creative!

​     