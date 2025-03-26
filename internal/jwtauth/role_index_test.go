package jwtauth

import (
	"reflect"
	"slices"
	"strings"
	"testing"
)

func TestRoleIndexer(t *testing.T) {
	roles := map[string]any{
		"teamb-teamb/subgroup/repoa-dev": map[string]any{
			"project_path": []any{"teamb/subgroup/repoa"},
		},
		"teamb-teamb/subgroup/repoa-test": map[string]any{
			"project_path":  []any{"teamb/subgroup/repoa"},
			"ref":           []any{"main", "master"},
			"ref_type":      []any{"branch"},
			"ref_protected": []any{"true"},
		},
		"teamb-teamb/subgroup/repoa-prod": map[string]any{
			"project_path":  []any{"teamb/subgroup/repoa"},
			"ref":           []any{"main", "master"},
			"ref_type":      []any{"branch"},
			"ref_protected": []any{"true"},
		},
		"teama-teama/subgroup/repoc-dev": map[string]any{
			"project_path": []any{"teama/subgroup/repoc"},
		},
		"teama-teama/subgroup/repoc-test": map[string]any{
			"project_path":  []any{"teama/subgroup/repoc"},
			"ref":           []any{"main", "master"},
			"ref_type":      []any{"branch"},
			"ref_protected": []any{"true"},
		},
		"teama-teama/subgroup/repoc-prod": map[string]any{
			"project_path":  []any{"teama/subgroup/repoc"},
			"ref":           []any{"main", "master"},
			"ref_type":      []any{"branch"},
			"ref_protected": []any{"true"},
		},
		"teama-teama/subgroup-dev": map[string]any{
			"namespace_path": []any{"teama/subgroup"},
		},
		"teama-teama/subgroup-test": map[string]any{
			"namespace_path": []any{"teama/subgroup"},
			"ref":            []any{"main", "master"},
			"ref_type":       []any{"branch"},
			"ref_protected":  []any{"true"},
		},
		"teama-teama/subgroup-prod": map[string]any{
			"namespace_path": []any{"teama/subgroup"},
			"ref":            []any{"main", "master"},
			"ref_type":       []any{"branch"},
			"ref_protected":  []any{"true"},
		},
		"teamb-teama/subgroup/repoc-prod": map[string]any{
			"project_path":  []any{"teama/subgroup/repoc"},
			"ref":           []any{"main", "master"},
			"ref_type":      []any{"branch"},
			"ref_protected": []any{"true"},
		},
		"teamb-custom": map[string]any{
			"project_path": []any{
				"teamb/subgroup/repoa",
				"teamb/subgroup/repob",
			},
			"ref":           []any{"main", "master"},
			"ref_type":      []any{"branch"},
			"ref_protected": []any{"true"},
		},
		"anyrepo": map[string]any{
			"ref":           []any{"main", "master"},
			"ref_type":      []any{"branch"},
			"ref_protected": []any{"true"},
		},
	}

	config := jwtAutoRolesConfig{
		Roles: roles,
	}

	index, err := createRoleIndex(&config)
	if err != nil {
		t.Fatal(err.Error())
	}

	for i, tt := range []struct {
		claims map[string]any
		roles  []string
	}{
		{
			claims: map[string]any{
				"namespace_path": "teama/subgroup",
				"project_path":   "teama/subgroup/repoc",
				"ref":            "main",
				"ref_type":       "branch",
				"ref_protected":  "true",
			},
			roles: []string{
				"teama-teama/subgroup-dev",
				"teama-teama/subgroup-test",
				"teama-teama/subgroup-prod",
				"teama-teama/subgroup/repoc-dev",
				"teama-teama/subgroup/repoc-test",
				"teama-teama/subgroup/repoc-prod",
				"teamb-teama/subgroup/repoc-prod",
				"anyrepo",
			},
		},
		{
			claims: map[string]any{
				"namespace_path": "teamb/subgroup",
				"project_path":   "teamb/subgroup/repoa",
				"ref":            "main",
				"ref_type":       "branch",
				"ref_protected":  "true",
			},
			roles: []string{
				"teamb-teamb/subgroup/repoa-dev",
				"teamb-teamb/subgroup/repoa-test",
				"teamb-teamb/subgroup/repoa-prod",
				"teamb-custom",
				"anyrepo",
			},
		},
		{
			claims: map[string]any{
				"namespace_path": "teamb/subgroup",
				"project_path":   "teamb/subgroup/repob",
				"ref":            "main",
				"ref_type":       "branch",
				"ref_protected":  "true",
			},
			roles: []string{
				"teamb-custom",
				"anyrepo",
			},
		},
		{
			claims: map[string]any{
				"namespace_path": "teama/subgroup",
				"project_path":   "teama/subgroup/grafana",
			},
			roles: []string{
				"teama-teama/subgroup-dev",
			},
		},
	} {
		got := index.claimsRoles(tt.claims)
		slices.Sort(got)
		slices.Sort(tt.roles)
		if !reflect.DeepEqual(got, tt.roles) {
			t.Fatalf("Test case %v,\nexpected:\n[%s]\ngot:\n[%s]",
				i, strings.Join(tt.roles, ",\n"), strings.Join(got, ",\n"))
		}
	}
}
