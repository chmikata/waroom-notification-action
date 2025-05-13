# Define the image name and tag
IMAGE_NAME = ghcr.io/chmikata/incident-notification
IMAGE_TAG = 0.2.0

# Declare phony targets
.PHONY: build push

# Default target
build_push:
	docker buildx build --platform linux/amd64,linux/arm64 --push --provenance false -t $(IMAGE_NAME):$(IMAGE_TAG) -t $(IMAGE_NAME):latest .
