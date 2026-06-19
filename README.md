# provider-tailscale

`provider-tailscale` is a [Crossplane](https://crossplane.io/) provider built with [Upjet v2](https://github.com/crossplane/upjet) that exposes XRM-conformant managed resources for the [Tailscale Terraform provider](https://registry.terraform.io/providers/tailscale/tailscale/latest/docs).

## Package

Install the provider from the Upbound registry:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-tailscale
spec:
  package: xpkg.upbound.io/buttah-cloud/provider-tailscale:v0.1.0
```

## Supported resources

This provider generates managed resources for all Tailscale Terraform resources in v0.29.2:

- ACL, contacts, services, webhooks
- Device authorization, keys, subnet routes, tags
- DNS configuration, nameservers, preferences, search paths, split nameservers
- AWS external ID, federated identities, OAuth clients
- Logstream configuration, posture integrations
- Tailnet keys and settings

Both cluster-scoped (`*.tailscale.buttah-cloud.io`) and namespaced (`*.tailscale.m.buttah-cloud.io`) APIs are available.

## Authentication

Create a Kubernetes secret with Tailscale credentials and reference it from a `ProviderConfig` or `ClusterProviderConfig`:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: tailscale-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "api_key": "tskey-api-...",
      "tailnet": "example.com"
    }
```

OAuth client credentials are also supported:

```json
{
  "oauth_client_id": "k...",
  "oauth_client_secret": "...",
  "tailnet": "example.com"
}
```

See the [Tailscale provider authentication docs](https://registry.terraform.io/providers/tailscale/tailscale/latest/docs#authentication) for all supported credential fields.

## Developing

Generate APIs, controllers, and CRDs:

```console
make submodules
make generate
```

Run locally against a cluster:

```console
make run
```

Build binaries (without container images):

```console
make go.build
```

## Report a Bug

File an issue in this repository.