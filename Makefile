dev:
	docker compose rm -svf && docker compose up
db:
	docker exec -it mysql mysql -h 127.0.0.1 -P 3306 -u root --password wallet