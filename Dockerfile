### Build image.

FROM alpine as build

RUN apk add --no-cache go git make

# Create a user with no priviledge
RUN addgroup -S builder
RUN adduser -S -D -H builder

# Install smu for generating the Markdown files.
RUN git clone https://github.com/Gottox/smu.git smu
RUN cd smu && make && make install && cd ..

COPY --chown=builder:builder . /src

USER builder

## Build Go binary.
# For building go without a home directory.
ENV GOPATH=/src/mod/
ENV GOCACHE=/src/src/go-cache

WORKDIR /src
RUN go build src/main.go

## Generate the Markdown files.
RUN ./tools/generate_markdown.sh

### Final image.

FROM alpine

# Create a user with no priviledge
RUN addgroup -S server
RUN adduser -S -D -H server

RUN mkdir -p /var/www/html
RUN chown -R server:server /var/www

# Copy the content to render.
COPY --from=build --chown=server:server /src/main /var/www/
COPY --from=build --chown=server:server /src/html/ /var/www/html
COPY --from=build --chown=server:server /src/LICENSE /var/www/
COPY robots.txt /var/www
WORKDIR /var/www

USER server
# Comment this line to allow running a shell in the container.
ENTRYPOINT ["/var/www/main"]
