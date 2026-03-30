package jwtauth

import (
	"context"
	"fmt"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
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

func pathLogin(backend *jwtAutoRolesAuthBackend) *framework.Path {
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

func (b *jwtAutoRolesAuthBackend) pathLogin(
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

	var alias *logical.Alias
	if id, ok := claims[config.UserClaim].(string); ok {
		alias = &logical.Alias{
			Name: id,
		}
	}

	roles := roleIndex.claimsRoles(claims)
	if len(roles) == 0 {
		return logical.ErrorResponse("unable to log into any role: no roles matched jwt claims"), nil
	}

	policies, err := b.policies(ctx, config, roles, token)
	if err != nil {
		return logical.ErrorResponse(err.Error()), nil
	}

	if len(policies) == 0 {
		return logical.ErrorResponse("unable to log into any role: roles matched claims but all upstream logins failed"), nil
	}

	return &logical.Response{
		Auth: &logical.Auth{
			Alias:    alias,
			Period:   time.Hour,
			Policies: policies,
		},
	}, nil
}

// policies use the token to log in to the configured auth backend for each
// role, and returns the aggregated policies of successful logins.
func (b *jwtAutoRolesAuthBackend) policies(
	ctx context.Context, config *jwtAutoRolesConfig, roles []string, token string,
) ([]string, error) {
	policies := map[string]struct{}{}
	var errs []string
	for _, role := range roles {
		rolePolicies, err := b.policyClient.policies(ctx, config, schema.JwtLoginRequest{
			Jwt:  token,
			Role: role,
		})
		if err != nil {
			errs = append(errs, fmt.Sprintf("role %q: %s", role, err))
			continue
		}

		for _, p := range rolePolicies {
			policies[p] = struct{}{}
		}
	}

	if len(policies) == 0 && len(errs) > 0 {
		return nil, fmt.Errorf("all upstream role logins failed: %s", strings.Join(errs, "; "))
	}

	r := make([]string, 0, len(policies))
	for p := range policies {
		r = append(r, p)
	}
	return r, nil
}
