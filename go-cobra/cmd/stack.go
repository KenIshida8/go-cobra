package cmd

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

)

func init() {
	RootCmd.AddCommand(newStackCmd())
}

func newStackCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stack",
		Short: "Manage Stack resources",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(
		newStackShowCmd(),
	)

	return cmd
}

func newStackShowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show <StackID>",
		Short: "Show Stack",
		RunE:  runStackShowCmd,
	}

	return cmd
}

func runStackShowCmd(cmd *cobra.Command, args []string) error {
	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	if len(args) == 0 {
		return errors.New("StackID is required")
	}

	appStackID, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.Wrapf(err, "failed to parse StackID: %s", args[0])
	}

	req := AppStackShowRequest{
		ID: appStackID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.StackShow(ctx, req)
	if err != nil {
		return errors.Wrapf(err, "StackShow was failed: req = %+v, res = %+v", req, res)
	}

	appStack := res.AppStack
	fmt.Printf(
		"id: %d, name: %s, inserted_at: %v, updated_at: %v\n",
		appStack.ID, appStack.Name, appStack.InsertedAt, appStack.UpdatedAt,
	)

	return nil
}
