FROM alpine

# TODO: Split the build from the final container to avoid extra deps.
RUN apk add --no-cache go

# Create a user with no priviledge
RUN adduser -S -D -H server

RUN mkdir /src
RUN chown server /src
COPY . /src


USER server
# For building go without a home directory.
ENV GOPATH=/src/mod/
ENV GOCACHE=/src/go-cache

WORKDIR /src
RUN go build main.go

ENTRYPOINT ["/src/main"]
