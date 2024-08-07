package cmd

import (
	"os"

	"github.com/Slug-Boi/aion-cli/src/logger"
	"github.com/spf13/cobra"
)

// This is an example of how to setup the logger for any CMD command you can then use it when doing calls.
// A similar logger can be setup in any other file that requires it by importing CMD and calling the SetupLogger function
var Sugar = logger.SetupLogger()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aion",
	Short: "A scheduling tool that takes in wishes and outputs a schedule",
	Long: `                                                     
 _____ _                 _ _ 
|  _  |_|___ ___ ___ ___| |_|
|     | | . |   |___|  _| | |
|__|__|_|___|_|_|   |___|_|_|
                               
                        
This is a CLI tool scheduling tool that takes in timeslot wishes and outputs a schedule. 
The tool was designed around google forms format and should therefore work with any format that is similar to the google forms format.
                                              `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
