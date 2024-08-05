# Docker Command

```console
$ docker-compose build
Build the Docker images.

$ docker-compose up
Start the containers.

$ docker-compose up -d
Start the containers in the background (detached mode).

$ docker-compose up --build
Build and start the containers. This will rebuild the images if there are changes before starting the containers.

$ docker-compose down
Stop and remove the containers.

$ docker-compose down -v
Stop and remove the containers and associated volumes.

$ docker-compose down --rmi all -v
Stop and remove the containers, images, and volumes.
```

```console
$ docker images
Show the available images.

$ docker ps
Show the running containers.

$ docker ps -a
Show all containers (running and stopped).

$ docker rm <container_id>
Delete the Docker container.

$ docker rmi <image_id>
Delete the Docker image.
```

```console
$ docker stats
Show real-time statistics of active containers.

$ docker info
Show detailed information about Docker daemon, including the system-wide information.

$ docker inspect <container_id_or_name>
Show detailed configuration and status of a specific container.

$ docker exec -it <service_name> <command>
Run the command inside the running container.
```