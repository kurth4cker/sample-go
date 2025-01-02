// SPDX-License-Identifier: GPL-3.0-or-later
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package subcmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	User  bool
	World bool
)

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "greet user",
	Run:   Greet,
}

func Greet(cmd *cobra.Command, args []string) {
	if User {
		fmt.Printf("hello %v\n", os.Getenv("USER"))
	}
	if !User || World {
		fmt.Printf("hello world\n")
	}
}

func init() {
	rootCmd.AddCommand(greetCmd)
	greetCmd.Flags().BoolVarP(&User, "user", "u", false, "greet current user")
	greetCmd.Flags().BoolVarP(&World, "world", "w", false, "greet current user")
}
