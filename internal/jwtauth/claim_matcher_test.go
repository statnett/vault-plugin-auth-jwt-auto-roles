package jwtauth

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestManyAllowedClaimValues(t *testing.T) {
	t.Parallel()

	matcher := roles{
		"a": {
			"foo": {"a", "b", "c"},
			"bar": {"a", "b", "c"},
		},
		"b": {
			"foo": {"a", "b"},
			"bar": {"a", "b"},
		},
	}.createMatchTree()

	for i, tt := range []struct {
		claims map[string]any
		values []string
	}{
		{
			claims: map[string]any{"foo": "a", "bar": "a"},
			values: []string{"a", "b"},
		},
		{
			claims: map[string]any{"foo": "c", "bar": "c"},
			values: []string{"a"},
		},
		{claims: map[string]any{"foo": "a"}, values: nil},
		{claims: map[string]any{"foo": "b"}, values: nil},
		{claims: map[string]any{"foo": "c"}, values: nil},
		{claims: map[string]any{"bar": "a"}, values: nil},
		{claims: map[string]any{"bar": "b"}, values: nil},
		{claims: map[string]any{"bar": "c"}, values: nil},
	} {
		got := matcher.findBoundRoleNames(tt.claims)
		opt := cmpopts.SortSlices(func(a, b string) bool { return a < b })
		if diff := cmp.Diff(tt.values, got, opt); diff != "" {
			t.Fatalf("Test case %v failed with diff:\n%s", i, diff)
		}
	}
}
