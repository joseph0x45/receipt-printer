CONTAINER_NAME=receipt-printer

setup-container:
	@docker run --user $(id -u):$(id -g) -itd --name $(CONTAINER_NAME) --network=host -p 19000:19000 -p 19001:19001 -v $$(pwd):/app -w /app node:lts-alpine

into-container:
	@docker exec -it receipt-printer sh
