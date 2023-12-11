package command

import (
	"github.com/spf13/cobra"

	"github.com/khulnasoft-lab/gobin/pkg/clio"
)

func Root(app clio.Application) *cobra.Command {
	return app.SetupRootCommand(&cobra.Command{})
}
