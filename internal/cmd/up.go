package cmd

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/google/shlex"
	"github.com/spf13/cobra"

	"github.com/twelvelabs/depctl/internal/core"
)

func NewUpCmd(app *core.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "up [group]",
		Short: "Ensure all dependencies are installed",
		Args:  cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: configure context cancellation on SIGINT.
			ctx := cmd.Context()

			if len(args) == 0 {
				args = append(args, "default")
			}

			for _, key := range args {
				deps, ok := app.Config.Dependencies[key]
				if !ok {
					return fmt.Errorf("unknown group: %s", key)
				}

				for _, dep := range deps {
					path, err := exec.LookPath(dep.Name)
					if err != nil && !errors.Is(err, exec.ErrNotFound) {
						return err
					}

					if path != "" {
						app.UI.Out(app.UI.SuccessIcon()+" Found dependency: %s \n", dep.Name)
						continue
					}
					app.UI.Out(app.UI.InfoIcon()+" Installing dependency: %s \n", dep.Name)

					// Parse the install string into args.
					args, err := shlex.Split(dep.Install)
					if err != nil {
						return err
					}
					if len(args) == 0 {
						return fmt.Errorf("install command missing for %s", dep.Name)
					}
					// Create the install command.
					command := exec.CommandContext(ctx, args[0], args[1:]...) //nolint: gosec
					command.Stderr = app.IO.Err
					command.Stdout = app.IO.Out
					if err := command.Run(); err != nil {
						return err
					}
				}
			}

			return nil
		},
	}

	return cmd
}
