.PHONY: all
all: lint

.PHONY: lint
lint:
	# Install: go get -u github.com/golangci/golangci-lint/cmd/golangci-lint@v1.12.5
	@golangci-Lint run \
		-E golint \
		-E stylecheck \
		-E gosec \
		-E interfacer \
		-E unconvert \
		-E dupl \
		-E goconst \
		-E gocyclo \
		-E gofmt \
		-E goimports \
		-E maligned \
		-E depguard \
		-E misspell \
		-E lll \
		-E unparam \
		-E nakedret \
		-E prealloc \
		-E scopelint \
		-E gocritic \
		-E gochecknoinits \
		-E gochecknoglobals