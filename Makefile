subscriber:
	cd search-service/ && go run subscriber/main.go

search-service:
	cd search-service/ && go run main.go

anime-service:
	cd anime-service/ && go run server/server.go

es-mapping:
	cd search-service/ && go run elastic/main.go

run-simple:
	cd simple/ && go run main.go

