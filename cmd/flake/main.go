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
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

type flakeOptions struct {
	bucket uint64
	count  int
	format string
}

var opts = flakeOptions{}

func newFlakeCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:     "flake",
		Version: fmt.Sprintf("%s (build %s on %s)", version, commit, date),
		Short:   "Flake generates time-based 64-bit unsigned integers",
		RunE:    runFlake,
	}

	rootCmd.Flags().Uint64VarP(&opts.bucket, "bucket", "b", 0, "Bucket id for all generated flake ids.")
	rootCmd.Flags().IntVarP(&opts.count, "count", "c", 1, "Number of flake ids to generate.")
	rootCmd.Flags().StringVarP(&opts.format, "format", "f", decimal, "Output format. Can be one of: binary, octal, int, or hex.")

	return rootCmd
}

func main() {
	if err := newFlakeCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
