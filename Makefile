app-name = "project-layout-rest-postgres"

# make install evn=init-local
install:
	sql-migrate up -env=$(env) -config=configs/dbconfig.yml

#make run env=dev
run:
	go run cmd/$(app-name)/main.go $(env)

#make build
build:
	git pull origin
	go mod tidy
	go build cmd/$(app-name)/main.go

#make deploy env=apa
deploy:
	pm2 start --silent './main $(env)' --watch --ignore-watch ./logs --name $(app-name)-$(env) --log=logs/app.log

test:
	go test ./...

coverage:
	go test -coverprofile cp.out ./...

coverage-html:
	go tool cover -html=cp.out

mock:
	sh scripts/mock.sh

lint:
	golangci-lint run -D=typecheck
	gosec -exclude=G304 ./...
	revive ./...
	gocritic check ./...
	go-consistent -v ./...

# Database Migration

# make create-migration name="name"
create-migration:
	sql-migrate new -config=configs/dbconfig.yml $(name)

# make migrate env="init-local"
migrate:
	sql-migrate up -env=$(env) -config=configs/dbconfig.yml

# Sonar

sonar-start:
	/Users/default/Apps/sonarqube-8.2.0.32929/bin/macosx-universal-64/sonar.sh start

sonar-log:
	/Users/default/Apps/sonarqube-8.2.0.32929/bin/macosx-universal-64/sonar.sh start

sonar:
	  /Users/default/Apps/sonar-scanner-4.2.0.1873-macosx/bin/sonar-scanner   -Dsonar.projectKey=organization \
                                                                                  -Dsonar.sources=. \
                                                                                  -Dsonar.host.url=http://localhost:9000 \
                                                                                  -Dsonar.login=e8ceb53356694dcf67b0af096be298497358c05d

.PHONY: clean install unittest build docker run stop vendor