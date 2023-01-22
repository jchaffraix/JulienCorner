#!/bin/bash
set -xeu -o pipefail

export LOCATION="us"
export PROJECT="juliencorner"
export REPOSITORY="us.gcr.io"
export IMAGE_NAME="webapp"
export LABEL="${LOCATION}-docker.pkg.dev/${PROJECT}/${REPOSITORY}/${IMAGE_NAME}:latest"

docker build -t "$LABEL" .

# TODO: Can I check that I am logged in here and log in?

# TODO: I should probably push to a staging env first.
docker push "$LABEL"

echo "Pushed image successfully"
echo ""
# TODO: Automate this through some CD hook.
echo "The new container needs to be deployed on https://console.cloud.google.com/run/detail/us-east1/${PROJECT}"
