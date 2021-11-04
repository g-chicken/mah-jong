dev-setup:
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz | tar xvz | mv migrate /usr/local/bin

migrate-up:
	migrate -source file://migrations -database mysql://localhost:3306 up

