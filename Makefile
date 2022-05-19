
run:
	docker-compose build &&	docker-compose up

gen:
	docker-compose -f docker-compose.tools.yml run --rm grpc -d proto/ -l go -o internal

test:
	docker-compose -f docker-compose.tools.yml run test_grpc -plaintext localhost:8080 $(endpoint)

test_payload:
	docker-compose -f docker-compose.tools.yml run test_grpc -d '$(payload)' localhost:8080 $(endpoint)

