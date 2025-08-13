.PHONY: install run stop

install:
	sudo apt-get update
	sudo apt-get install -y docker.io docker-compose-plugin
	sudo service docker start || true

run:
	docker compose -f docker-compose.nginx.yml up -d

stop:
	docker compose -f docker-compose.nginx.yml down
