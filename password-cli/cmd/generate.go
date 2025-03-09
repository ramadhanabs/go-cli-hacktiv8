/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"math/rand"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a random password",
	Long:  "Generates a secure random password with customizable length and character options.",
	Run: func(cmd *cobra.Command, args []string) {
		length, _ := cmd.Flags().GetInt("length")
		includeDigits, _ := cmd.Flags().GetBool("digits")
		includeSpecial, _ := cmd.Flags().GetBool("special")
		isSave, _ := cmd.Flags().GetBool("save")

		password := generatePassword(length, includeDigits, includeSpecial, isSave)
		fmt.Println("Generated Password:", password)
	},
}

func generatePassword(length int, includeDigits, includeSpecial, isSave bool) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	special := "!@#$%^&*()_+-=[]{}|;:'\",.<>?/"

	var charset = letters
	if includeDigits {
		charset += digits
	}

	if includeSpecial {
		charset += special
	}

	rand.Seed(time.Now().UnixNano())
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	if isSave {
		os.WriteFile("last_password.txt", []byte(password), 0644)
	}
	return string(password)
}

func init() {
	generateCmd.Flags().IntP("length", "l", 6, "Length of password (max 12)")
	generateCmd.Flags().BoolP("digits", "d", false, "Include digits")
	generateCmd.Flags().BoolP("special", "c", false, "Include special characters")
	generateCmd.Flags().BoolP("save", "s", false, "Save to .txt format")
}
