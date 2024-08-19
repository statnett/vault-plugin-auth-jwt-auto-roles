package jwtauth

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/vault/sdk/logical"
)

func TestConfigRolesRefresh_Write(t *testing.T) {
	t.Parallel()

	configData := testConfig()
	backend, storage := createTestBackend(t)

	req := &logical.Request{
		Operation: logical.UpdateOperation,
		Path:      configPath,
		Storage:   storage,
		Data:      configData,
	}

	resp, err := backend.HandleRequest(context.Background(), req)
	if err != nil || (resp != nil && resp.IsError()) {
		t.Fatalf("err:%s resp:%#v\n", err, resp)
	}

	newRoles := map[string]any{
		"foo": map[string]any{
			"project_path": []any{"foo", "bar"},
		},
	}
	var roleClient fakeRoleFetcher = func(_ context.Context, _ *jwtAutoRolesConfig, vaultToken string) (map[string]any, error) {
		return newRoles, nil
	}
	backend.roleClient = roleClient

	req = &logical.Request{
		Operation: logical.UpdateOperation,
		Path:      configRolesRefreshPath,
		Storage:   storage,
		Data: map[string]any{
			"vault_token": "secret",
		},
	}

	resp, err = backend.HandleRequest(context.Background(), req)
	if err != nil || (resp != nil && resp.IsError()) {
		t.Fatalf("err:%s resp:%#v\n", err, resp)
	}

	conf, err := backend.config(context.Background(), storage)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(newRoles, conf.Roles) {
		t.Fatalf("expected did not match actual: expected %#v\n got %#v\n", newRoles, conf.Roles)
	}
}

type fakeRoleFetcher func(context.Context, *jwtAutoRolesConfig, string) (map[string]any, error)

func (c fakeRoleFetcher) roles(ctx context.Context, config *jwtAutoRolesConfig, vaultToken string) (map[string]any, error) {
	return c(ctx, config, vaultToken)
}
