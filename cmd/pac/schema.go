package main

import (
	"github.com/spf13/cobra"
	"github.com/squillace/porter-pac/pkg/pac"
)

func buildSchemaCommand(m *pac.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schema",
		Short: "Print the json schema for the mixin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.PrintSchema()
		},
	}
	return cmd
}
