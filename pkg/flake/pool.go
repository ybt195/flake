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

type generatorPool struct {
	pool chan Generator
}

// NewPool returns a new flake id generator using the set of bucket ids.
func NewPool(bucketIDs ...uint64) (Generator, error) {
	gPool := &generatorPool{
		pool: make(chan Generator, len(bucketIDs)),
	}
	for _, bucketID := range bucketIDs {
		g, err := New(bucketID)
		if err != nil {
			return nil, err
		}
		gPool.put(g)
	}
	return gPool, nil
}

func (g *generatorPool) Next() (ID, error) {
	sub := <-g.pool
	defer g.put(sub)
	return sub.Next()
}

func (g *generatorPool) Must() ID {
	sub := <-g.pool
	defer g.put(sub)
	return sub.Must()
}

func (g *generatorPool) put(sub Generator) {
	g.pool <- sub
}
