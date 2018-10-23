/*
 * Copyright 2018 Jonathan Ben-tzur
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ybt195/flake/pkg/flake"
)

const (
	binary  string = "binary"
	octal   string = "octal"
	decimal string = "decimal"
	hex     string = "hex"
)

type generateOptions struct {
	bucket uint64
	count  int
	format string
}

func newGenerateCommand() *cobra.Command {
	var opts = generateOptions{}

	var cmd = &cobra.Command{
		Use:   "generate",
		Short: "Generates time-based 64-bit unsigned integers",
		RunE: func(cmd *cobra.Command, args []string) error {
			return generate(opts)
		},
	}

	cmd.Flags().Uint64VarP(&opts.bucket, "bucket", "b", 0, "Bucket id for all generated flake ids.")
	cmd.Flags().IntVarP(&opts.count, "count", "c", 1, "Number of flake ids to generate.")
	cmd.Flags().StringVarP(&opts.format, "format", "f", decimal, "Output format. Can be one of: binary, octal, int, or hex.")

	return cmd
}

func generate(opts generateOptions) error {
	g, err := flake.New(opts.bucket)
	if err != nil {
		return err
	}

	for i := 0; i < opts.count; i++ {
		id, err := formatID(g.Must(), opts.format)
		if err != nil {
			return err
		}
		fmt.Println(id)
	}

	return nil
}

func formatID(id flake.ID, format string) (string, error) {
	switch format {
	case binary:
		return id.Binary(), nil
	case octal:
		return id.Octal(), nil
	case decimal:
		return fmt.Sprint(id.Uint64()), nil
	case hex:
		return id.Hex(), nil
	default:
		return "", fmt.Errorf("unexpected format: %s", format)
	}
}
