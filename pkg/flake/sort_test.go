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
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDSlice_Len(t *testing.T) {
	tests := []struct {
		name string
		p    IDSlice
		want int
	}{
		{name: "zero-len", p: []ID{}, want: 0},
		{name: "one-len", p: []ID{1}, want: 1},
		{name: "two-len", p: []ID{1, 2}, want: 2},
		{name: "more-len", p: []ID{1, 2, 3}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIDSlice_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    IDSlice
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIDSlice_Sort(t *testing.T) {
	g, err := New(1)
	assert.Nil(t, err)

	largeList := make([]ID, 0, 5000)
	for i := 0; i < cap(largeList); i++ {
		largeList = append(largeList, g.Must())
	}

	tests := []struct {
		name string
		p    IDSlice
	}{
		{name: "empty", p: []ID{}},
		{name: "one-elem", p: []ID{g.Must()}},
		{name: "large-list", p: largeList},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, sort.IsSorted(tt.p))
			rand.Shuffle(len(tt.p), func(i, j int) {
				tt.p[i], tt.p[j] = tt.p[j], tt.p[i]
			})
			tt.p.Sort()
			assert.True(t, sort.IsSorted(tt.p))
		})
	}
}
