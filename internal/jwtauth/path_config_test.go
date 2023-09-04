package jwtauth

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/go-test/deep"
	log "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/vault/sdk/helper/logging"
	"github.com/hashicorp/vault/sdk/logical"
)

func TestConfig_Write(t *testing.T) {
	t.Parallel()
	backend, storage := createTestBackend(t)

	data := testConfig()
	req := &logical.Request{
		Operation: logical.UpdateOperation,
		Path:      configPath,
		Storage:   storage,
		Data:      data,
	}

	resp, err := backend.HandleRequest(context.Background(), req)
	if err != nil || (resp != nil && resp.IsError()) {
		t.Fatalf("err:%s resp:%#v\n", err, resp)
	}

	conf, err := backend.(*multiroleJWTAuthBackend).config(context.Background(), storage)
	if err != nil {
		t.Fatal(err)
	}

	expected := &multiroleJWTConfig{
		Roles:       data["roles"].(map[string]any),
		JWTAuthHost: "http://localhost:8200",
		JWTAuthPath: "foo/jwt",
	}

	if !reflect.DeepEqual(expected, conf) {
		t.Fatalf("expected did not match actual: expected %#v\n got %#v\n", expected, conf)
	}
}

func TestConfig_Read(t *testing.T) {
	t.Parallel()
	backend, storage := createTestBackend(t)

	configData := testConfig()
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

	req = &logical.Request{
		Operation: logical.ReadOperation,
		Path:      configPath,
		Storage:   storage,
		Data:      nil,
	}

	resp, err = backend.HandleRequest(context.Background(), req)
	if err != nil || (resp != nil && resp.IsError()) {
		t.Fatalf("err:%s resp:%#v\n", err, resp)
	}

	if diff := deep.Equal(resp.Data, configData); diff != nil {
		t.Fatalf("Expected did not equal actual: %v", diff)
	}
}

func createTestBackend(t *testing.T) (logical.Backend, logical.Storage) {
	config := &logical.BackendConfig{
		Logger: logging.NewVaultLogger(log.Trace),

		System: &logical.StaticSystemView{
			DefaultLeaseTTLVal: time.Hour * 12,
			MaxLeaseTTLVal:     time.Hour * 24,
		},
		StorageView: &logical.InmemStorage{},
	}
	b, err := Factory(context.Background(), config)
	if err != nil {
		t.Fatalf("unable to create backend: %v", err)
	}

	return b, config.StorageView
}

func testConfig() map[string]any {
	return map[string]any{
		"roles": map[string]any{
			"foo": map[string]any{
				"project_path": []any{"foo", "bar"},
			},
			"bar": map[string]any{
				"namespace_path": []any{"c"},
				"ref":            []any{"master", "main"},
			},
		},
		"jwt_auth_host": "http://localhost:8200",
		"jwt_auth_path": "foo/jwt",
	}
}
