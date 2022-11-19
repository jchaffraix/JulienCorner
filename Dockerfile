### Build image.

FROM alpine as build

# TODO: Split the build from the final container to avoid extra deps.
RUN apk add --no-cache go

# Create a user with no priviledge
RUN adduser -S -D -H builder

RUN mkdir /src
RUN chown builder /src
COPY src/ /src
COPY go.mod /src
COPY go.sum /src

USER builder
# For building go without a home directory.
ENV GOPATH=/src/mod/
ENV GOCACHE=/src/go-cache

WORKDIR /src
RUN go build main.go

### Final image.

FROM alpine

# Create a user with no priviledge
RUN addgroup -S server
RUN adduser -S -D -H server

RUN mkdir -p /var/www/html
RUN chown -R server:server /var/www

# Copy the content to render.
COPY html /var/www/html/
COPY --from=build --chown=server:server /src/main /var/www
WORKDIR /var/www

USER server
# Comment this line to allow running a shell in the container.
ENTRYPOINT ["/var/www/main"]
