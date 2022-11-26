#!/bin/sh

export LOCATION="us"
export PROJECT="juliencorner"
export REPOSITORY="us.gcr.io"
export IMAGE_NAME="webapp"
export LABEL="${LOCATION}-docker.pkg.dev/${PROJECT}/${REPOSITORY}/${IMAGE_NAME}:latest"

docker build -t LABEL .

# TODO: Can I check that I am logged in here and log in?

# TODO: I should probably push to a staging env first.
docker push LABEL
