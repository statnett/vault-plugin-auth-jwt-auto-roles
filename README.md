# vault-plugin-auth-jwt-auto-roles

A Vault plugin to automatically authenticate with all roles matching a JWT (or
OIDC) token, and grant access to the union of matching roles' policies.

**Problem**: The [builtin _jwt_ auth
plugin](https://github.com/hashicorp/vault-plugin-auth-jwt) requires a role-name
to be specified during authentication. When authenticated, access is only
granted based on the policies of that single role. Even though the JWT satisfies
multiple roles' bound claims. Thus, multiple logins with different role names
are required when policies are split across multiple roles, resulting in
multiple tokens which cannot be used interchangeably.

**Solution**: The _jwt-auto-roles_ plugin automatically determines which roles'
bound claims a JWT matches. It grants access to the union of all the matching
roles' policies. Thus a single login is sufficient to get a token with combined
access, with the JWT as the sole login parameter.

## Working Details

The plugin currently relies on the builtin _jwt_ plugin for JWT verification and
policy configuration. It simply determine the roles matching an incoming JWT's
claims, tries to login to a mount of the builtin _jwt_ plugin for each role, and
grants access to the union of policies returned by the builtin _jwt_ plugin
mount.

## Usage

Download and unzip a
[release](https://github.com/statnett/vault-plugin-auth-jwt-auto-roles/releases)
of the plugin and place it in the [plugin
directory](https://developer.hashicorp.com/vault/docs/configuration#plugin_directory).

[Register](https://developer.hashicorp.com/vault/docs/commands/plugin/register)
the plugin:

`vault plugin register -sha256=<binary sha> auth vault-plugin-auth-jwt-auto-roles`

Enable the plugin at a specific mount path (e.g. `jwt-auto-roles`):

`vault auth enable -path=jwt-auto-roles vault-plugin-auth-jwt-auto-roles`

For the plugin to be able to determine matching roles, it must be configured
with the host name and mount point of a builtin _jwt_ plugin mount, along with
all its roles and their bound claims:

`vault write auth/jwt-auto-roles/config @config.json`

```json5
// config.json
{
  "jwt_auth_host": "https://vault.org.com",
  "jwt_auth_path": "jwt",
  "user_claim": "user@example.com", // optional, generated based on policies and namespace otherwise
  "roles": {
    "role-a": {
      "project_path": ["foo/bar"]
    },
    "role-b": {
      "branch": ["main", "master"]
    }
  }
}
```

Then login as with the builtin _jwt_ auth, although without the role parameter:

`vault write auth/jwt-auto-roles/login jwt=$jwt`

## Future work

The plugin could be configured with policies directly, instead of role names,
and not rely on the builtin _jwt_ plugin for token verification. Although
simpler to configure, security guarantees will no longer be delegated to
HashiCorp.

There is a also [an upstream issue](https://github.com/hashicorp/vault/issues/23279)
to get this implemented in the official [_jwt_ plugin](https://github.com/hashicorp/vault-plugin-auth-jwt), obsoleting this plugin.
