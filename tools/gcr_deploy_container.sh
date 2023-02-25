#!/bin/bash
set -xeu -o pipefail

export REGION="us-east1"
export PROJECT="juliencorner"
export SERVICE_NAME="$PROJECT"
export REPOSITORY="us.gcr.io"
export IMAGE_NAME="webapp"
export TAG=$(date +%Y%m%d-%H%M%S)
export LABEL="${REPOSITORY}/${PROJECT}/${IMAGE_NAME}:${TAG}"

# Check that we pass an authorization for a service account and log into it.
# Else for CLI invokation, we just use any existing credentials.
export SERVICE_ACCNT_FILE="${SERVICE_ACCNT_FILE:-}"
if [[ "$SERVICE_ACCNT_FILE" != "" ]]; then
  gcloud auth activate-service-account "${SERVICE_ACCNT}" --key-file="$SERVICE_ACCNT_FILE"
  gcloud config set project "$PROJECT"
  gcloud auth configure-docker "$REPOSITORY"
fi

docker build -t "$LABEL" .
docker push "$LABEL"

echo "Pushed image successfully"
echo ""

echo "Deploying the new container"
# TODO: I should probably push to a staging env first.
gcloud run deploy "$SERVICE_NAME" \
  --image="$LABEL" \
  --region="$REGION" \
  --revision-suffix="$TAG"

echo "The new container is deployed!"
echo "You can find more detail here: https://console.cloud.google.com/run/detail/${REGION}/${SERVICE_NAME}"
