// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	externalid "github.com/buttahtoast/provider-tailscale/internal/controller/namespaced/aws/externalid"
	authorization "github.com/buttahtoast/provider-tailscale/internal/controller/namespaced/device/authorization"
	key "github.com/buttahtoast/provider-tailscale/internal/controller/namespaced/device/key"
	subnetroutes "github.com/buttahtoast/provider-tailscale/internal/controller/namespaced/device/subnetroutes"
	tags "github.com/buttahtoast/provider-tailscale/internal/controller/namespaced/device/tags"
	configuration "github.com/buttahtoast/provider-tailscale/internal/controller/namespaced/dns/configuration"
	nameservers "github.com/buttahtoast/provider-tailscale/internal/controller/namespaced/dns/nameservers"
	preferences "github.com/buttahtoast/provider-tailscale/internal/controller/namespaced/dns/preferences"
	searchpaths "github.com/buttahtoast/provider-tailscale/internal/controller/namespaced/dns/searchpaths"
	splitnameservers "github.com/buttahtoast/provider-tailscale/internal/controller/namespaced/dns/splitnameservers"
	identity "github.com/buttahtoast/provider-tailscale/internal/controller/namespaced/federated/identity"
	configurationlogstream "github.com/buttahtoast/provider-tailscale/internal/controller/namespaced/logstream/configuration"
	client "github.com/buttahtoast/provider-tailscale/internal/controller/namespaced/oauth/client"
	integration "github.com/buttahtoast/provider-tailscale/internal/controller/namespaced/posture/integration"
	providerconfig "github.com/buttahtoast/provider-tailscale/internal/controller/namespaced/providerconfig"
	keytailnet "github.com/buttahtoast/provider-tailscale/internal/controller/namespaced/tailnet/key"
	settings "github.com/buttahtoast/provider-tailscale/internal/controller/namespaced/tailnet/settings"
	acl "github.com/buttahtoast/provider-tailscale/internal/controller/namespaced/tailscale/acl"
	contacts "github.com/buttahtoast/provider-tailscale/internal/controller/namespaced/tailscale/contacts"
	service "github.com/buttahtoast/provider-tailscale/internal/controller/namespaced/tailscale/service"
	webhook "github.com/buttahtoast/provider-tailscale/internal/controller/namespaced/tailscale/webhook"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		externalid.Setup,
		authorization.Setup,
		key.Setup,
		subnetroutes.Setup,
		tags.Setup,
		configuration.Setup,
		nameservers.Setup,
		preferences.Setup,
		searchpaths.Setup,
		splitnameservers.Setup,
		identity.Setup,
		configurationlogstream.Setup,
		client.Setup,
		integration.Setup,
		providerconfig.Setup,
		keytailnet.Setup,
		settings.Setup,
		acl.Setup,
		contacts.Setup,
		service.Setup,
		webhook.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		externalid.SetupGated,
		authorization.SetupGated,
		key.SetupGated,
		subnetroutes.SetupGated,
		tags.SetupGated,
		configuration.SetupGated,
		nameservers.SetupGated,
		preferences.SetupGated,
		searchpaths.SetupGated,
		splitnameservers.SetupGated,
		identity.SetupGated,
		configurationlogstream.SetupGated,
		client.SetupGated,
		integration.SetupGated,
		providerconfig.SetupGated,
		keytailnet.SetupGated,
		settings.SetupGated,
		acl.SetupGated,
		contacts.SetupGated,
		service.SetupGated,
		webhook.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
