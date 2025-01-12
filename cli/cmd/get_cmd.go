package cmd

import (
	"context"
	"fmt"

	"github.com/kubeshop/tracetest/cli/analytics"
	"github.com/kubeshop/tracetest/cli/formatters"
	"github.com/spf13/cobra"
)

var resourceID string

var getCmd = &cobra.Command{
	GroupID: cmdGroupResources.ID,
	Use:     "get [resource type]",
	Long:    "Get a resource from your Tracetest server",
	Short:   "Get resource",
	PreRun:  setupCommand(),
	Args:    cobra.MinimumNArgs(1),
	Run: WithResultHandler(func(cmd *cobra.Command, args []string) (string, error) {
		if resourceID == "" {
			return "", fmt.Errorf("id of the resource to get must be specified")
		}

		resourceType := args[0]
		ctx := context.Background()

		analytics.Track("Resource Get", "cmd", map[string]string{
			resourceType: resourceType,
		})

		resourceActions, err := resourceRegistry.Get(resourceType)
		if err != nil {
			return "", err
		}

		resource, err := resourceActions.Get(ctx, resourceID)
		if err != nil {
			return "", err
		}

		resourceFormatter := resourceActions.Formatter()
		formatter := formatters.BuildFormatter(output, formatters.YAML, resourceFormatter)

		result, err := formatter.Format(resource)
		if err != nil {
			return "", err
		}

		return result, nil
	}),
	PostRun: teardownCommand,
}

func init() {
	getCmd.Flags().StringVar(&resourceID, "id", "", "id of the resource to get")
	rootCmd.AddCommand(getCmd)
}
