all: start

GOFUMPT_VERSION ?= v0.7.0
GORELEASER_VERSION ?= 2.4.8

.PHONY: fmt
fmt:
	go fmt $$(go list ./...)
	go run mvdan.cc/gofumpt@$(GOFUMPT_VERSION) -l -w .

.PHONY: test
test:
	go test -v ./...

.PHONY: build
build:
	mkdir -p build/plugins
	go run github.com/goreleaser/goreleaser/v2@v$(GORELEASER_VERSION) \
		build --clean --snapshot --single-target \
		--output build/plugins/vault-plugin-auth-jwt-auto-roles

.PHONY: start
start: build
	vault server -dev -dev-root-token-id=root -dev-plugin-dir=./build/plugins

.PHONY: enable
enable:
	vault auth enable -path=jwt-auto-roles vault-plugin-auth-jwt-auto-roles

.PHONY: disable
disable:
	vault auth disable jwt-auto-roles
	vault plugin deregister auth vault-plugin-auth-jwt-auto-roles

.PHONY: configure
configure:
	vault write auth/jwt-auto-roles/config @config.json

.PHONY: readconfig
readconfig:
	vault read --format=json auth/jwt-auto-roles/config

.PHONY: refreshroles
token=""
refreshroles:
	vault write auth/jwt-auto-roles/config/roles/refresh vault_token="$(token)"

.PHONY: clean
clean:
	rm -rf ./bin ./build ./dist
