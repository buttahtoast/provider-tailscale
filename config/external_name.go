package config

import "github.com/crossplane/upjet/v2/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"tailscale_acl":                     config.IdentifierFromProvider,
	"tailscale_aws_external_id":         config.IdentifierFromProvider,
	"tailscale_contacts":                config.IdentifierFromProvider,
	"tailscale_device_authorization":    config.IdentifierFromProvider,
	"tailscale_device_key":              config.IdentifierFromProvider,
	"tailscale_device_subnet_routes":    config.IdentifierFromProvider,
	"tailscale_device_tags":             config.IdentifierFromProvider,
	"tailscale_dns_configuration":       config.IdentifierFromProvider,
	"tailscale_dns_nameservers":         config.IdentifierFromProvider,
	"tailscale_dns_preferences":         config.IdentifierFromProvider,
	"tailscale_dns_search_paths":        config.IdentifierFromProvider,
	"tailscale_dns_split_nameservers":   config.IdentifierFromProvider,
	"tailscale_federated_identity":      config.IdentifierFromProvider,
	"tailscale_logstream_configuration": config.IdentifierFromProvider,
	"tailscale_oauth_client":            config.IdentifierFromProvider,
	"tailscale_posture_integration":     config.IdentifierFromProvider,
	"tailscale_service":                 config.IdentifierFromProvider,
	"tailscale_tailnet_key":             config.IdentifierFromProvider,
	"tailscale_tailnet_settings":        config.IdentifierFromProvider,
	"tailscale_webhook":                 config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
