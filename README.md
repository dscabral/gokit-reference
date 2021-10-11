# gokit-reference

To run the container use the command below

To build the docker image, use the command bellow
``` docker build --no-cache --tag=dscabral-microservices . ```

You can choose the expose port 3000 outside the container to the exposed 8080 inside the container
```docker run --publish 3000:8080 dscabral-microservices```
