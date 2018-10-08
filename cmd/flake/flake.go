package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ybt195/flake/pkg/flake"
)

const (
	BINARY  string = "binary"
	OCTAL   string = "octal"
	DECIMAL string = "decimal"
	HEX     string = "hex"
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

func formatID(id flake.FlakeID) (string, error) {
	switch opts.format {
	case BINARY:
		return id.Binary(), nil
	case OCTAL:
		return id.Octal(), nil
	case DECIMAL:
		return fmt.Sprint(id.Uint64()), nil
	case HEX:
		return id.Hex(), nil
	default:
		return "", fmt.Errorf("unexpected format: %s", opts.format)
	}
}
