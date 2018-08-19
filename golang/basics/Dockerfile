# refer https://docs.docker.com/develop/develop-images/multistage-build/
# first stage just builds the binary, the second copies the binary built in the first into the second.
# This essentially lets us build a small final `production` image with no crud from the building stages
FROM golang:latest as builder
WORKDIR /go/src/app
# probably will need some go get commands here to get dependency packages
# since the current application is just a simple thing, nothing needed here
COPY an_app .
# go build with no cgo linkage
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


# alpine is a minimal image
FROM alpine:latest
WORKDIR /root/
# note the --from=builder, since we named the earlier stage we can reference it by name. yay!
COPY --from=builder /go/src/app/app .
CMD ["./app"]
