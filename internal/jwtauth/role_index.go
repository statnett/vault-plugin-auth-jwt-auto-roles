package jwtauth

import (
	"fmt"

	"github.com/pkg/errors"
)

type roleIndex struct {
	jwtMatcher *node
}

func createRoleIndex(config *multiroleJWTConfig) (*roleIndex, error) {
	roles, err := parseRoles(config)
	if err != nil {
		return nil, err
	}

	return &roleIndex{
		jwtMatcher: roles.createMatchTree(),
	}, nil
}

// claimsRoles returns the roles with bound claims matching token claims.
func (ri *roleIndex) claimsRoles(claims map[string]any) []string {
	return ri.jwtMatcher.findBoundRoleNames(claims)
}

func parseRoles(config *multiroleJWTConfig) (roles, error) {
	roles := make(roles)
	for roleName, bc := range config.Roles {
		boundClaims, err := parseBoundClaims(bc)
		if err != nil {
			return nil, fmt.Errorf("failed to parse bound claims for role '%s': %w", roleName, err)
		}
		roles[roleName] = boundClaims
	}
	return roles, nil
}

func parseBoundClaims(bc any) (map[string][]string, error) {
	switch bc := bc.(type) {
	case map[string]any:
		boundClaims := make(map[string][]string, len(bc))
		for claimKey, v := range bc {
			values, err := parseBoundClaimValues(v)
			if err != nil {
				return nil, fmt.Errorf("failed to parse bound claim values for claim key '%s': %w", claimKey, err)
			}
			boundClaims[claimKey] = values
		}
		return boundClaims, nil
	default:
		return nil, errors.Errorf("bound claims not of expected type: %v", bc)
	}
}

func parseBoundClaimValues(vs any) ([]string, error) {
	switch vs := vs.(type) {
	case []any:
		var values []string
		for _, v := range vs {
			value, err := parseBoundClaimValue(v)
			if err != nil {
				return nil, fmt.Errorf("failed to parse bound claim value '%v': %w", v, err)
			}
			values = append(values, value)
		}
		return values, nil
	default:
		return nil, errors.Errorf("bound claim values not of expected type: %v", vs)
	}
}

func parseBoundClaimValue(v any) (string, error) {
	switch v := v.(type) {
	case string:
		return v, nil
	default:
		return "", errors.Errorf("bound claim value not of expected type: %v", v)
	}
}
