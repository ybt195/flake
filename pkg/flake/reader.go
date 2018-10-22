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

package flake

import (
	"fmt"
	"io"
)

type rCloserGenerator struct {
	reader io.ReadCloser
}

// NewReader returns a new flake id generator using the underlying reader as a source. The source
// is processed in 8-byte intervals.
func NewReader(reader io.ReadCloser) (Generator, error) {
	return &rCloserGenerator{
		reader: reader,
	}, nil
}

func (g *rCloserGenerator) Next() (ID, error) {
	buffer := make([]byte, Size)
	n, err := g.reader.Read(buffer)
	if err != nil {
		return Nil, err
	}
	if n < Size {
		return Nil, fmt.Errorf("expected to read %d bytes but only %d provided", Size, n)
	}

	id, err := FromBytes(buffer)
	return id, err
}

func (g *rCloserGenerator) Must() ID {
	id, err := g.Next()
	if err != nil {
		panic(err)
	}
	return id
}

func (g *rCloserGenerator) Close() error {
	return g.reader.Close()
}
