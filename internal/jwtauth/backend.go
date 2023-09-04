package jwtauth

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"

	"github.com/statnett/vault-plugin-auth-jwt-auto-roles/pkg/version"
)

const (
	backendHelp = `
The multirole JWT backend plugin allows authentication with multiple roles using JWTs (including OIDC).
`
	vaultClientTimeoutSeconds = 5
)

func Factory(ctx context.Context, c *logical.BackendConfig) (logical.Backend, error) {
	b := backend(c)
	if err := b.Setup(ctx, c); err != nil {
		return nil, fmt.Errorf("failed to setup backend: %w", err)
	}
	return b, nil
}

type multiroleJWTAuthBackend struct {
	*framework.Backend

	l            sync.RWMutex
	cachedConfig *multiroleJWTConfig
	roleIndex    *roleIndex
	client       *vault.Client
}

func backend(_ *logical.BackendConfig) *multiroleJWTAuthBackend {
	var backend multiroleJWTAuthBackend
	backend.Backend = &framework.Backend{
		BackendType: logical.TypeCredential,
		Help:        backendHelp,
		PathsSpecial: &logical.Paths{
			Unauthenticated: []string{"login"},
		},
		Paths: []*framework.Path{
			pathLogin(&backend),
			pathConfig(&backend),
		},
		RunningVersion: version.Version,
	}
	return &backend
}

func (b *multiroleJWTAuthBackend) reset() {
	b.l.Lock()
	defer b.l.Unlock()
	b.cachedConfig = nil
	b.roleIndex = nil
}

func (b *multiroleJWTAuthBackend) getRoleIndex(config *multiroleJWTConfig) (*roleIndex, error) {
	b.l.Lock()
	defer b.l.Unlock()

	if b.roleIndex != nil {
		return b.roleIndex, nil
	}

	index, err := createRoleIndex(config)
	if err != nil {
		return nil, err
	}

	b.roleIndex = index
	return index, nil
}

func (b *multiroleJWTAuthBackend) getClient(config *multiroleJWTConfig) (*vault.Client, error) {
	b.l.Lock()
	defer b.l.Unlock()

	if b.client != nil {
		return b.client, nil
	}

	client, err := vault.New(
		vault.WithAddress(config.JWTAuthHost),
		vault.WithRequestTimeout(vaultClientTimeoutSeconds*time.Second),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create vault client: %w", err)
	}

	b.client = client
	return client, nil
}
