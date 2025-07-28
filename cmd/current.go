package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use:  "current",
	Short: "Display the current weather information",
	Long: `This command retrieves and displays the current weather information for a specified location.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		location := strings.Join(args, "%20")
		var response Response
		var err error
		response, err = CurrentWeather(location)
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(response)
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)
}