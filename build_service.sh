IMAGE_NAME=social-todo-service
CACHE_BUILD=$1

if [[ -n "$CACHE_BUILD" ]]; then
  echo "Docker building with cache"
  docker rmi ${IMAGE_NAME}-cached ${IMAGE_NAME}
  docker build -t ${IMAGE_NAME}-cached -f Dockerfile-cache .
fi

echo "Docker building main image..."
docker build -t ${IMAGE_NAME}:latest -f Dockerfile-multi-stage .

echo "Done!!"
