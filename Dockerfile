# You can manually build the image by running:
#  docker build . -f Dockerfile -t hub.docker.com/duckladydinh/fibonacci-go-gateway:1.0-SNAPSHOT
#
# You can manually run the image by running:
#  docker run hub.docker.com/duckladydinh/fibonacci-go-gateway:1.0-SNAPSHOT --grpc-server-endpoint=localhost:10001 --http-server-endpoint=localhost:10002
#
FROM golang:1.17.3-bullseye

ADD . / app/
RUN (cd app; go install cmd/main.go;)

# The only command to run when an image is being executed.
# ENTRYPOINT (instead of CMD) can be used to append outside flags.
WORKDIR ${GOPATH}/bin
ENTRYPOINT [ "main" ]