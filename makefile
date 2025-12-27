CI_RAW_URL = https://raw.githubusercontent.com/yeencloud/dpl-ci/refs/heads/main

update:
	curl -O $(CI_RAW_URL)/makefile \
         -O $(CI_RAW_URL)/.golangci.yml

lint:
	golangci-lint run ./...

test:
	go test -v ./...

graphql:
	go get -tool github.com/99designs/gqlgen
	go tool gqlgen generate

air:
	@printf "\033]0;%s\007" "$(notdir $(CURDIR))"
	@go tool air --build.cmd "go build -o ./tmp/$(notdir $(CURDIR))exe cmd/main.go" \
	  --build.entrypoint "./tmp/$(notdir $(CURDIR))exe" \
	  --build.include_file ".env" \
	  --build.include_dir "../../libs"
