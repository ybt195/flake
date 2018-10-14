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

func runFlake(_ *cobra.Command, _ []string) error {
	g, err := flake.New(opts.bucket)
	if err != nil {
		return err
	}

	for i := 0; i < opts.count; i++ {
		id, err := formatID(g.Must())
		if err != nil {
			return err
		}
		fmt.Println(id)
	}

	return nil
}

func formatID(id flake.ID) (string, error) {
	switch opts.format {
	case binary:
		return id.Binary(), nil
	case octal:
		return id.Octal(), nil
	case decimal:
		return fmt.Sprint(id.Uint64()), nil
	case hex:
		return id.Hex(), nil
	default:
		return "", fmt.Errorf("unexpected format: %s", opts.format)
	}
}
