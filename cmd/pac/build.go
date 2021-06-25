package main

import (
	"github.com/spf13/cobra"
	"github.com/squillace/porter-pac/pkg/pac"
)

func buildBuildCommand(m *pac.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Generate Dockerfile lines for the bundle invocation image",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Build()
		},
	}
	return cmd
}
