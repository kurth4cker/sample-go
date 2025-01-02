// SPDX-License-Identifier: GPL-3.0-or-later
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package subcmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "greet user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello world")
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)
}
