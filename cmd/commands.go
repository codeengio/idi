package cmd

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

var allowedArgs = []string{"new"}

func NewRootCmd(run func(cmd *cobra.Command, args []string) error) *cobra.Command {
	return &cobra.Command{
		Use:   "new",
		Short: "Creates a new app",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires at least one arg")
			}

			if !slices.Contains(allowedArgs, args[0]) {
				return fmt.Errorf("the arg must be one of: %s", strings.Join(allowedArgs, ", "))
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			run(cmd, args)
		},
	}
}
