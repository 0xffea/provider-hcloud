// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	network "github.com/0xffea/provider-hcloud/internal/controller/hcloud/network"
	server "github.com/0xffea/provider-hcloud/internal/controller/hcloud/server"
	route "github.com/0xffea/provider-hcloud/internal/controller/network/route"
	subnet "github.com/0xffea/provider-hcloud/internal/controller/network/subnet"
	providerconfig "github.com/0xffea/provider-hcloud/internal/controller/providerconfig"
	networkserver "github.com/0xffea/provider-hcloud/internal/controller/server/network"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		network.Setup,
		server.Setup,
		route.Setup,
		subnet.Setup,
		providerconfig.Setup,
		networkserver.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
