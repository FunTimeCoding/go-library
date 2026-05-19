package goquery

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/client"
	"github.com/spf13/cobra"
)

func search(c *client.Client) *cobra.Command {
	var limit int
	var collection string
	var keywordOnly bool
	var full bool
	var sourceType string
	result := &cobra.Command{
		Use:   "search [query]",
		Short: "Search indexed documents",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			mode := client.Hybrid

			if keywordOnly {
				mode = client.Keyword
			}

			params := &client.GetSearchParams{
				Query: arguments[0],
				Limit: &limit,
				Mode:  &mode,
				Full:  &full,
			}

			if collection != "" {
				params.Collection = &collection
			}

			if sourceType != "" {
				params.SourceType = &sourceType
			}

			r, e := c.GetSearch(context.Background(), params)
			errors.PanicOnError(e)
			defer errors.PanicClose(r.Body)
			var results []client.SearchResult
			errors.PanicOnError(json.NewDecoder(r.Body).Decode(&results))

			for _, v := range results {
				fmt.Printf(
					"%.4f  %-7s  %s  %s\n",
					v.Score,
					v.Source,
					v.VirtualPath,
					v.Title,
				)
			}
		},
	}
	result.Flags().IntVar(&limit, "limit", 10, "Maximum number of results")
	result.Flags().StringVar(
		&collection,
		"collection",
		"",
		"Restrict search to a collection",
	)
	result.Flags().BoolVar(
		&keywordOnly,
		"keyword",
		false,
		"Use keyword search only",
	)
	result.Flags().BoolVar(&full, "full", false, "Include full document body")
	result.Flags().StringVar(
		&sourceType,
		"source-type",
		"",
		"Filter by source type",
	)

	return result
}
