### Multi-Stage build
This application uses a multi-stage build to both compartmentalize the build soley into the Dockerfile as well as to reduce the image size.  When packed this reduces the image to ~5mb.  

### Makefile
A simple Makefile is included to perform the following: <br/> 
1) Run GO unit tests
2) Build Dockerfile with specified tag
3) Run Docker container locally
4) Push Docker image to DockerHub repository





