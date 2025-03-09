/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "password",
	Short: "Password Generator CLI",
	Long: `
 ________  ________  ________   ________  ________  _______   ________      
|\   __  \|\   __  \|\   ____\ |\   ____\|\   ____\|\  ___ \ |\   ___  \    
\ \  \|\  \ \  \|\  \ \  \___|_\ \  \___|\ \  \___|\ \   __/|\ \  \\ \  \   
 \ \   ____\ \   __  \ \_____  \\ \_____  \ \  \  __\ \  \_|/_\ \  \\ \  \  
  \ \  \___|\ \  \ \  \|____|\  \\|____|\  \ \  \|\  \ \  \_|\ \ \  \\ \  \ 
   \ \__\    \ \__\ \__\____\_\  \ ____\_\  \ \_______\ \_______\ \__\\ \__\
    \|__|     \|__|\|__|\_________\\_________\|_______|\|_______|\|__| \|__|
                       \|_________\|_________|                              
                                                                            
                                                                        
A CLI tool to generate secure passwords with options for digits and special characters.
`,
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
	rootCmd.AddCommand(generateCmd)
}
