/*
Copyright © 2023 Vance Longwill <vance@evren.co.uk>
*/
package cmd

import (
	"dscli/dumper"
	"encoding/json"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/spf13/cobra"
)

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Get a dump of all entries for a given entity",
	Long:  `Runs a GetAll query for the given entity until all entries have been exhausted`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		dsClient, err := datastore.NewClient(ctx, project)
		if err != nil {
			return err
		}

		enc := json.NewEncoder(os.Stdout)
		entityName := "RESTTransaction"

		dumper := dumper.New(dumper.Config{
			Client:     dsClient,
			Encoder:    enc,
			BatchSize:  batchSize,
			EntityName: entityName,
		})

		return dumper.Dump(ctx)
	},
}

var (
	entity    string
	output    string
	batchSize int
)

func init() {

	dumpCmd.Flags().StringVarP(&entity, "entity", "e", "", "the datastore entity name")
	dumpCmd.Flags().IntVarP(&batchSize, "batch-size", "b", 1000, "the number of entities to fetch in a single page")
	dumpCmd.Flags().StringVarP(&output, "output", "o", "json", "the output format (defaults to JSON)")
	dumpCmd.MarkFlagRequired("entity")

	rootCmd.AddCommand(dumpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dumpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dumpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
