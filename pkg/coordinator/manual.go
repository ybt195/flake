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

package coordinator

import (
	"errors"

	"github.com/ybt195/flake/pkg/flake"
)

type manualCoordinator struct {
	buckets []uint64
}

// NewManualCoordinator returns a coordinator that pushes bucket id management to the user.
func NewManualCoordinator(buckets ...uint64) (Coordinator, error) {
	if len(buckets) == 0 {
		return nil, errors.New("must specify at least one bucket for the coordinator")
	}
	return &manualCoordinator{
		buckets: buckets,
	}, nil
}

func (c *manualCoordinator) Get() (flake.Generator, error) {
	generators := make([]flake.Generator, len(c.buckets))
	for _, bucket := range c.buckets {
		generator, err := flake.New(bucket)
		if err != nil {
			return nil, err
		}
		generators = append(generators, generator)
	}
	if len(generators) == 1 {
		return generators[0], nil
	}
	return flake.NewPool(generators...)
}
