package network

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("hcloud_network", func(r *config.Resource) {
	r.ShortGroup = ""
    })

    p.AddResourceConfigurator("hcloud_network_subnet", func(r *config.Resource) {
        r.References["network"] = config.Reference{
            Type: "github.com/0xffea/provider-hcloud/apis/hcloud/v1alpha1.Network",
        }
    })

    p.AddResourceConfigurator("hcloud_network_route", func(r *config.Resource) {
        r.References["network"] = config.Reference{
            Type: "github.com/0xffea/provider-hcloud/apis/hcloud/v1alpha1.Network",
        }
    })
}
