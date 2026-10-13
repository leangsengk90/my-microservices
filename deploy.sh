#!/bin/bash

echo "Enter tag version (e.g. 1.0.0):"
read TAG_VERSION

SERVICES=("auth" "profile" "messenger" "finance" "share")
REGISTRY_URL="songhieng2k25/zongdev"

for SERVICE in "${SERVICES[@]}"; do
  echo "Building $SERVICE..."
  docker build --build-arg APP_VERSION=$TAG_VERSION -t $REGISTRY_URL-$SERVICE:$TAG_VERSION ./Microservices/$SERVICE
  if [ $? -ne 0 ]; then
    echo "Build failed for $SERVICE. Exiting..."
    exit 1
  fi

  echo "Pushing $SERVICE..."
  docker push $REGISTRY_URL-$SERVICE:$TAG_VERSION
done

echo "Updating docker-compose.yml..."
for SERVICE in "${SERVICES[@]}"; do
  sed -i "s|image: $REGISTRY_URL-$SERVICE:.*|image: $REGISTRY_URL-$SERVICE:$TAG_VERSION|" docker-compose.yml
done

echo "Deploying..."
docker-compose down
docker-compose pull
docker-compose up -d

echo "Deployment complete with tag: $TAG_VERSION"
