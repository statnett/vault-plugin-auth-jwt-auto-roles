package jwtauth

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

const (
	configPath string = "config"

	pathConfigHelpSyn = `
Configures the multirole JWT authentication backend.
`
	pathConfigHelpDesc = `
The multirole JWT authentication backend uses configured bound claims of another
jwt auth backends's roles to quickly deduce which roles an incoming jwt unlocks.
`
)

type multiroleJWTConfig struct {
	// Roles map role names to bound claims. Bound claims must be of type
	// map[string][]string.
	Roles       map[string]any `json:"roles"`
	JWTAuthHost string         `json:"jwt_auth_host"`
	JWTAuthPath string         `json:"jwt_auth_path"`
}

func pathConfig(backend *multiroleJWTAuthBackend) *framework.Path {
	return &framework.Path{
		Pattern: `config`,
		Fields: map[string]*framework.FieldSchema{
			"roles": {
				Type: framework.TypeMap,
				DisplayAttrs: &framework.DisplayAttributes{
					Name: "List of role names",
				},
			},
			"jwt_auth_host": {
				Type: framework.TypeString,
				DisplayAttrs: &framework.DisplayAttributes{
					Name: "Hostname of the vault instance",
				},
			},
			"jwt_auth_path": {
				Type: framework.TypeString,
				DisplayAttrs: &framework.DisplayAttributes{
					Name: "Path of the default jwt auth plugin (without 'auth' or 'login')",
				},
			},
		},

		Operations: map[logical.Operation]framework.OperationHandler{
			logical.ReadOperation: &framework.PathOperation{
				Callback: backend.pathConfigRead,
				Summary:  "Read the current multirole JWT authentication backend configuration.",
			},
			logical.UpdateOperation: &framework.PathOperation{
				Callback: backend.pathConfigWrite,
				Summary:  "Configure the multirole JWT authentication backend.",
			},
		},

		HelpSynopsis:    pathConfigHelpSyn,
		HelpDescription: pathConfigHelpDesc,
	}
}

func (b *multiroleJWTAuthBackend) config(ctx context.Context, storage logical.Storage) (*multiroleJWTConfig, error) {
	if b.cachedConfig != nil {
		return b.cachedConfig, nil
	}

	b.l.Lock()
	defer b.l.Unlock()

	entry, err := storage.Get(ctx, configPath)
	if err != nil {
		return nil, err
	}
	if entry == nil {
		return nil, nil
	}

	config := &multiroleJWTConfig{}
	if err := entry.DecodeJSON(config); err != nil {
		return nil, err
	}

	b.cachedConfig = config
	return config, nil
}

func (b *multiroleJWTAuthBackend) pathConfigWrite(
	ctx context.Context, req *logical.Request, d *framework.FieldData,
) (*logical.Response, error) {
	config := multiroleJWTConfig{
		Roles:       d.Get("roles").(map[string]any),
		JWTAuthHost: d.Get("jwt_auth_host").(string),
		JWTAuthPath: d.Get("jwt_auth_path").(string),
	}

	_, err := parseRoles(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse roles: %w", err)
	}

	entry, err := logical.StorageEntryJSON(configPath, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create storage: %w", err)
	}
	if err := req.Storage.Put(ctx, entry); err != nil {
		return nil, fmt.Errorf("failed to read storage: %w", err)
	}

	b.reset()
	return nil, nil
}

func (b *multiroleJWTAuthBackend) pathConfigRead(
	ctx context.Context, req *logical.Request, d *framework.FieldData,
) (*logical.Response, error) {
	config, err := b.config(ctx, req.Storage)
	if err != nil {
		return nil, err
	}
	if config == nil {
		return nil, nil
	}

	return &logical.Response{
		Data: map[string]any{
			"roles":         config.Roles,
			"jwt_auth_host": config.JWTAuthHost,
			"jwt_auth_path": config.JWTAuthPath,
		},
	}, nil
}
