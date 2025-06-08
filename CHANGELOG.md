# Changelog

## [0.3.6](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.3.5...v0.3.6) (2025-06-08)


### Bug Fixes

* bump to Go 1.24 ([#125](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/125)) ([211b73f](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/211b73f3fd6ca7535f094547b6742a4d3a2aa5ef))
* **deps:** update module github.com/hashicorp/vault/api to v1.20.0 ([#123](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/123)) ([7af0d3a](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/7af0d3aef81681a6a7579140758e4bb7394a4b3d))
* **deps:** update module github.com/hashicorp/vault/sdk to v0.18.0 ([#122](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/122)) ([5fbf48f](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/5fbf48f813c912eeb338728f0c4ba862ba453073))

## [0.3.5](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.3.4...v0.3.5) (2025-03-21)


### Bug Fixes

* **deps:** update module github.com/hashicorp/vault/sdk to v0.15.2 ([#106](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/106)) ([1d8d23f](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/1d8d23f98dbbf4db39d0538861716c74edec16ac))

## [0.3.4](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.3.3...v0.3.4) (2025-02-24)


### Bug Fixes

* **deps:** update module github.com/google/go-cmp to v0.7.0 ([#103](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/103)) ([ba5303c](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/ba5303cc41da91813ece52f47ceba262bfe62787))
* **deps:** update module github.com/hashicorp/vault/api to v1.16.0 ([#101](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/101)) ([d84abe9](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/d84abe939fa9589e56c7700af0fbbfe0964b7ce7))
* **deps:** update module github.com/hashicorp/vault/sdk to v0.15.0 ([#102](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/102)) ([21dc29e](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/21dc29e5c3600177e736630e458dfa5b325dcb2a))

## [0.3.3](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.3.2...v0.3.3) (2025-01-09)


### Bug Fixes

* **deps:** update module github.com/hashicorp/vault/sdk to v0.14.1 ([#98](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/98)) ([d25b20e](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/d25b20e6c8816834becab57e9ea1e196c71d13e4))

## [0.3.2](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.3.1...v0.3.2) (2024-12-09)


### Bug Fixes

* various upgrades ([#96](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/96)) ([b9879b3](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/b9879b3b977dbc8878cd6ba3a8be95456482fb9d))

## [0.3.1](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.3.0...v0.3.1) (2024-09-10)


### Bug Fixes

* **deps:** update module github.com/hashicorp/vault/api to v1.15.0 ([#90](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/90)) ([6b366fd](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/6b366fd2c869f74977cb31f6cb449f471e669064))
* **deps:** update module github.com/hashicorp/vault/sdk to v0.14.0 ([#91](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/91)) ([a20568c](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/a20568c78c87106ce19f3a58385554a07ef88359))

## [0.3.0](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.2.16...v0.3.0) (2024-07-31)


### ⚠ BREAKING CHANGES

* fail when unable to login to any roles ([#88](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/88))

### Features

* fail when unable to login to any roles ([#88](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/88)) ([9bf3f87](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/9bf3f87893af77f0c2fbab71dcb7b683c292a2d9))

## [0.2.16](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.2.15...v0.2.16) (2024-07-23)


### Bug Fixes

* clear cached clients on reset ([#83](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/83)) ([96a9adf](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/96a9adf5367878d7c98bd6a5743f1268616eac15))

## [0.2.15](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.2.14...v0.2.15) (2024-07-23)


### Features

* remove custom timeout ([#81](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/81)) ([105067c](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/105067c236257c2dbb936f8fca688648175dbe66))

## [0.2.14](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.2.13...v0.2.14) (2024-07-22)


### Features

* dynamically fetch roles ([#78](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/78)) ([1638a16](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/1638a161b4f833f1aa0d8b28e91e54c8ae14df56))


### Bug Fixes

* **deps:** update module github.com/go-test/deep to v1.1.1 ([#76](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/76)) ([f4255b3](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/f4255b3b48db179a539006f78a532dde35593ba8))
* **deps:** update module github.com/hashicorp/vault/api to v1.14.0 ([#72](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/72)) ([e72ab2c](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/e72ab2c39d5c7d2141d4765bf6bfb13825c86162))
* **deps:** update module github.com/hashicorp/vault/sdk to v0.13.0 ([#80](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/80)) ([0c90158](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/0c90158f4467731b7162ac53444cb15d4f848249))

## [0.2.13](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.2.12...v0.2.13) (2024-04-21)


### Bug Fixes

* **deps:** update module github.com/hashicorp/vault/api to v1.13.0 ([#70](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/70)) ([891d31c](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/891d31c2300232f1adf3f274be96093ce612b75d))
* **deps:** update module github.com/hashicorp/vault/sdk to v0.12.0 ([#67](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/67)) ([eebf9de](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/eebf9de8dbb699687197aec779e54aab07244a78))

## [0.2.12](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.2.11...v0.2.12) (2024-03-19)


### Features

* ability to set JWT claim used for alias ([#60](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/60)) ([c03c9c9](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/c03c9c95363d18d2a33295b4b1feef0ce8d23ee0))


### Bug Fixes

* **deps:** update module github.com/hashicorp/vault/api to v1.12.1 ([#56](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/56)) ([79c2a57](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/79c2a57eeaa1380904f3f5381c0ac54b9f5e05f4))
* **deps:** update module github.com/hashicorp/vault/api to v1.12.2 ([#61](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/61)) ([f65b4e9](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/f65b4e9145626967809483e4ea854abef5b413ab))
* **deps:** update module github.com/hashicorp/vault/sdk to v0.11.1 ([#57](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/57)) ([b6f5203](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/b6f5203448796b0cecf88169e948d7a10f8eaafb))

## [0.2.11](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.2.10...v0.2.11) (2024-03-12)


### Bug Fixes

* **deps:** update module github.com/golang-jwt/jwt/v5 to v5.2.1 ([#54](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/54)) ([e439d3f](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/e439d3f554cbde446375f26ea66fab993fb317c7))
* **deps:** update module github.com/hashicorp/vault/api to v1.11.0 ([#50](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/50)) ([b59d3bd](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/b59d3bdd09f02dfe01730380dab96023c96351dc))
* **deps:** update module github.com/hashicorp/vault/api to v1.12.0 ([#52](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/52)) ([190949d](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/190949dd0b02692d214c99dd102cbcd32a0ff9d8))
* **deps:** update module github.com/hashicorp/vault/sdk to v0.11.0 ([#53](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/53)) ([b1e9d8c](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/b1e9d8c9222840abb7f4255c508a0cec33b4ac9e))

## [0.2.10](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.2.9...v0.2.10) (2023-12-28)


### Bug Fixes

* **deps:** update module github.com/hashicorp/vault-client-go to v0.4.3 ([#48](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/48)) ([facb43e](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/facb43e798c5b2720660a2bbab7f4ed575ba1429))

## [0.2.9](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.2.8...v0.2.9) (2023-12-13)


### Bug Fixes

* **deps:** update module github.com/hashicorp/go-hclog to v1.6.2 ([#44](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/44)) ([97687ba](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/97687baab5875cdaae8016ca1be336c08dc82c6e))

## [0.2.8](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.2.7...v0.2.8) (2023-12-04)


### Bug Fixes

* **deps:** update module github.com/hashicorp/go-hclog to v1.6.1 ([#40](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/40)) ([2ff8a39](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/2ff8a39f36e08d80ef0eb0398da6e60d28f60cf1))

## [0.2.7](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.2.6...v0.2.7) (2023-12-02)


### Bug Fixes

* **deps:** update module github.com/golang-jwt/jwt/v5 to v5.2.0 ([#38](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/38)) ([a1d65d0](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/a1d65d084a402a2265f4121e7c27132a29996a8a))

## [0.2.6](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.2.5...v0.2.6) (2023-11-23)


### Bug Fixes

* **deps:** bump github.com/go-jose/go-jose/v3 from 3.0.0 to 3.0.1 ([#36](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/36)) ([ab03321](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/ab03321b2801a4932f3bd0a7d3c7e52634c65f58))
* **deps:** update module github.com/golang-jwt/jwt/v5 to v5.1.0 ([#34](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/34)) ([a31ad02](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/a31ad02c41aa69b898cf15c63003419c25e473e6))

## [0.2.5](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.2.4...v0.2.5) (2023-10-31)


### Bug Fixes

* **deps:** bump github.com/docker/docker from 24.0.5+incompatible to 24.0.7+incompatible ([#32](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/32)) ([c989031](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/c98903162970f1adcbc4c43f41817c454dac89d5))

## [0.2.4](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.2.3...v0.2.4) (2023-10-25)


### Bug Fixes

* **deps:** bump google.golang.org/grpc from 1.57.0 to 1.57.1 ([#30](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/30)) ([936bdec](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/936bdec14d2fa25c3d81a2fd9dee8957f3fa4cdc))

## [0.2.3](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.2.2...v0.2.3) (2023-10-22)


### Bug Fixes

* **deps:** update module github.com/hashicorp/vault-client-go to v0.4.2 ([#28](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/28)) ([f43ba80](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/f43ba806b167815bdd8d83f79f3e6f2a91280647))

## [0.2.2](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.2.1...v0.2.2) (2023-10-11)


### Bug Fixes

* **deps:** update module github.com/google/go-cmp to v0.6.0 ([#24](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/24)) ([e504db6](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/e504db69754458240ede6eb45b17f59e1f444ff8))
* **deps:** update module github.com/hashicorp/vault/sdk to v0.10.2 ([#21](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/21)) ([56af930](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/56af93021bb05c67df8d32f6c1c6417432e3355a))

## [0.2.1](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.2.0...v0.2.1) (2023-10-06)


### Bug Fixes

* **deps:** update module github.com/hashicorp/vault-client-go to v0.4.1 ([#12](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/12)) ([56b1958](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/56b1958891bc6ef71c41921eb5805a41edaf58b2))
* **deps:** update module github.com/hashicorp/vault/api to v1.10.0 ([#13](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/13)) ([32f323e](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/32f323e3302f531e418cc6f3a9946eca16b2d01e))
* **deps:** update module github.com/hashicorp/vault/sdk to v0.10.0 ([#15](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/15)) ([f10b69d](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/f10b69d5b1c6f86f6107d1a07df5bae1298c5ded))

## [0.2.0](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.1.0...v0.2.0) (2023-09-26)


### Features

* ability to delete config ([#6](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/6)) ([5a73eb6](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/5a73eb6ab8d72655669b58d130ee220e4bdff56f))


### Bug Fixes

* plugin name in code and doc ([#9](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/9)) ([7c1e218](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/7c1e218aece02399f6c58eef68e83a74f3558ff1))

## [0.1.0](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/compare/v0.0.0...v0.1.0) (2023-09-21)


### Features

* initial plugin ([#1](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/issues/1)) ([b18cfe0](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/commit/b18cfe07c6129e011b85ac680f60f2c02d328dbb))
