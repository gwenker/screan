package cmd

import (
	"github.com/gwenker/screan/server"
	"github.com/spf13/cobra"
)

var port int

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start server",
	Long: `Start command to launch server`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Start(port)
	},
}

func init() {
	RootCmd.AddCommand(startCmd)

	// flags
	startCmd.Flags().IntVarP(&port, "port", "p", 1323, "port for server")

}
