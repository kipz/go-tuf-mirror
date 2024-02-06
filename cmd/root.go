package cmd

import (
	_ "embed"
	"fmt"

	"github.com/spf13/cobra"
)

type rootOptions struct {
	source      string
	destination string
}

func defaultRootOptions() *rootOptions {
	return &rootOptions{}
}

func newRootCmd(version string) *cobra.Command {
	o := defaultRootOptions()
	cmd := &cobra.Command{
		Use:   "go-tuf-mirror",
		Short: "Mirror TUF metadata to and between OCI registries, filesystems etc",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.Flags().StringVar(&o.source, "source", "s", "Source location (http://|oci://|file://)")
	cmd.Flags().StringVar(&o.destination, "destination", "d", "Destination location (oci://|file://)")

	cmd.MarkFlagRequired("source")
	cmd.MarkFlagRequired("destination")

	cmd.AddCommand(newVersionCmd(version)) // version subcommand

	return cmd
}

// Execute invokes the command.
func Execute(version string) error {
	if err := newRootCmd(version).Execute(); err != nil {
		return fmt.Errorf("error executing root command: %w", err)
	}

	return nil
}
