package jwtauth

import (
	"context"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

const (
	pathLoginHelpSyn = `
Authenticates to one or more Vault roles using a JWT (or OIDC) token.
`
	pathLoginHelpDesc = `
Authenticates JWTs.
`
)

func pathLogin(backend *multiroleJWTAuthBackend) *framework.Path {
	return &framework.Path{
		Pattern: "login",
		Fields: map[string]*framework.FieldSchema{
			"jwt": {
				Type:        framework.TypeString,
				Description: "The signed JWT to validate.",
			},
		},
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.UpdateOperation: &framework.PathOperation{
				Callback: backend.pathLogin,
				Summary:  pathLoginHelpSyn,
			},
		},
		HelpSynopsis:    pathLoginHelpSyn,
		HelpDescription: pathLoginHelpDesc,
	}
}

func (b *multiroleJWTAuthBackend) pathLogin(
	ctx context.Context, req *logical.Request, d *framework.FieldData,
) (*logical.Response, error) {
	token := d.Get("jwt").(string)
	if len(token) == 0 {
		return logical.ErrorResponse("missing token"), nil
	}

	config, err := b.config(ctx, req.Storage)
	if err != nil {
		return nil, err
	}
	if config == nil {
		return logical.ErrorResponse("could not load configuration"), nil
	}
	roleIndex, err := b.getRoleIndex(config)
	if err != nil {
		return nil, err
	}

	claims := make(jwt.MapClaims)
	_, _, err = jwt.NewParser().ParseUnverified(token, claims)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	roles := roleIndex.claimsRoles(claims)
	policies, err := b.policies(ctx, config, roles, token)
	if err != nil {
		return logical.ErrorResponse(err.Error()), nil
	}

	return &logical.Response{
		Auth: &logical.Auth{
			Period:   time.Hour,
			Policies: policies,
		},
	}, nil
}

// policies use the token to log in to the configured auth backend for each
// role, and returns the aggregated policies of successful logins.
func (b *multiroleJWTAuthBackend) policies(
	ctx context.Context, config *multiroleJWTConfig, roles []string, token string,
) ([]string, error) {
	client, err := b.getClient(config)
	if err != nil {
		return nil, err
	}

	policies := map[string]struct{}{}
	for _, role := range roles {
		response, err := client.Auth.JwtLogin(ctx, schema.JwtLoginRequest{
			Jwt:  token,
			Role: role,
		}, vault.WithMountPath(config.JWTAuthPath))
		if err != nil {
			continue
			// TODO: return error if non-403
			// return nil, err
		}

		for _, p := range response.Auth.Policies {
			policies[p] = struct{}{}
		}
	}

	r := make([]string, 0, len(policies))
	for p := range policies {
		r = append(r, p)
	}
	return r, nil
}
