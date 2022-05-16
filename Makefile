
run:
	docker-compose build && docker-compose run --rm app

gen:
	docker-compose run grpc protoc --proto_path=proto proto/hello.proto --go_out=. --go-grpc_out=.
