setup:
	mkdir tmp_migrate
	curl -sSOL https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz
	mv *tar.gz tmp_migrate && cd tmp_migrate && tar -xvzf *.tar.gz
	sudo mv tmp_migrate/migrate /usr/local/bin
	rm -r tmp_migrate

migrate-up:
	migrate -source file://migrations -database mysql://localhost:3306 up

run:
	./bin/up.sh

