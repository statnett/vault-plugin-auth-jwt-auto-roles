all: start

GOFUMPT_VERSION ?= v0.5.0
GORELEASER_VERSION ?= 1.20.0

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
	go run github.com/goreleaser/goreleaser@v$(GORELEASER_VERSION) \
		build --clean --snapshot --single-target \
		--output build/plugins/vault-plugin-auth-jwt-auto-roles

.PHONY: start
start: build
	vault server -dev -dev-root-token-id=root -dev-plugin-dir=./build/plugins

.PHONY: enable
enable:
	vault auth enable -path=multirole-jwt vault-plugin-auth-jwt-auto-roles

.PHONY: disable
disable:
	vault auth disable multirole-jwt
	vault plugin deregister auth vault-plugin-auth-jwt-auto-roles

auth=jwt
config:
	for role in $(shell vault list -format=json auth/$(auth)/role | jq -r .[]); \
	do vault read -format json "auth/$(auth)/role/$${role}" | jq "{\"$${role}\":.data.bound_claims}"; done \
	| jq -s add \
	| jq '{jwt_auth_host:"${VAULT_ADDR}",jwt_auth_path:"$(auth)",roles:.}' \
	> config.json

.PHONY: configure
configure:
	vault write auth/multirole-jwt/config @config.json

.PHONY: clean
clean:
	rm -rf ./bin ./build ./dist config.json
