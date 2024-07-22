package jwtauth

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"

	"github.com/statnett/vault-plugin-auth-jwt-auto-roles/pkg/version"
)

const (
	backendHelp = `
The JWT auto roles auth plugin allows automatic authentication with all roles
matching a JWT (or OIDC) token.
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

type jwtAutoRolesAuthBackend struct {
	*framework.Backend

	l            sync.RWMutex
	cachedConfig *jwtAutoRolesConfig
	roleIndex    *roleIndex
	policyClient policyFetcher
}

type policyFetcher interface {
	policies(ctx context.Context, request schema.JwtLoginRequest) ([]string, error)
}

func backend(_ *logical.BackendConfig) *jwtAutoRolesAuthBackend {
	var backend jwtAutoRolesAuthBackend
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

func (b *jwtAutoRolesAuthBackend) reset() {
	b.l.Lock()
	defer b.l.Unlock()
	b.cachedConfig = nil
	b.roleIndex = nil
}

func (b *jwtAutoRolesAuthBackend) getRoleIndex(config *jwtAutoRolesConfig) (*roleIndex, error) {
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

func (b *jwtAutoRolesAuthBackend) policyFetcher(config *jwtAutoRolesConfig) (policyFetcher, error) {
	b.l.Lock()
	defer b.l.Unlock()

	if b.policyClient != nil {
		return b.policyClient, nil
	}

	client, err := vault.New(
		vault.WithAddress(config.JWTAuthHost),
		vault.WithRequestTimeout(vaultClientTimeoutSeconds*time.Second),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create vault client: %w", err)
	}

	b.policyClient = &vaultClient{
		Client:    client,
		mountPath: config.JWTAuthPath,
	}
	return b.policyClient, nil
}

type vaultClient struct {
	*vault.Client
	mountPath string
}

func (c *vaultClient) policies(ctx context.Context, request schema.JwtLoginRequest) ([]string, error) {
	r, err := c.Client.Auth.JwtLogin(ctx, request, vault.WithMountPath(c.mountPath))
	if err != nil {
		return nil, fmt.Errorf("vault error: %w", err)
	}
	return r.Auth.Policies, nil
}
