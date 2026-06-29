// Copyright (c) 2022, R.I. Pienaar and the Choria Project contributors
//
// SPDX-License-Identifier: Apache-2.0

package appbuilder

import (
	"github.com/choria-io/appbuilder/builder"
	"github.com/choria-io/go-choria/client/discovery"
)

func ProcessStdDiscoveryOptions(b *builder.AppBuilder, f *discovery.StandardOptions, arguments map[string]any, flags map[string]any) error {
	var err error

	if f.DiscoveryMethod != "" {
		f.DiscoveryMethod, err = b.RenderTemplate(f.DiscoveryMethod, arguments, flags)
		if err != nil {
			return err
		}
	}

	for k, v := range f.DiscoveryOptions {
		f.DiscoveryOptions[k], err = b.RenderTemplate(v, arguments, flags)
		if err != nil {
			return err
		}
	}

	if f.Collective != "" {
		f.Collective, err = b.RenderTemplate(f.Collective, arguments, flags)
		if err != nil {
			return err
		}
	}

	if f.NodesFile != "" {
		f.NodesFile, err = b.RenderTemplate(f.NodesFile, arguments, flags)
		if err != nil {
			return err
		}
	}

	if f.CompoundFilter != "" {
		f.CompoundFilter, err = b.RenderTemplate(f.CompoundFilter, arguments, flags)
		if err != nil {
			return err
		}
	}

	for i, item := range f.CombinedFilter {
		f.CombinedFilter[i], err = b.RenderTemplate(item, arguments, flags)
		if err != nil {
			return err
		}
	}

	for i, item := range f.IdentityFilter {
		f.IdentityFilter[i], err = b.RenderTemplate(item, arguments, flags)
		if err != nil {
			return err
		}
	}

	for i, item := range f.AgentFilter {
		f.AgentFilter[i], err = b.RenderTemplate(item, arguments, flags)
		if err != nil {
			return err
		}
	}

	for i, item := range f.ClassFilter {
		f.ClassFilter[i], err = b.RenderTemplate(item, arguments, flags)
		if err != nil {
			return err
		}
	}

	for i, item := range f.FactFilter {
		f.FactFilter[i], err = b.RenderTemplate(item, arguments, flags)
		if err != nil {
			return err
		}
	}

	return nil
}
