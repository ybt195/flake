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

import "sort"

// IDSlice implements sort.Interface for an ID slice. IDs are sorted in ascending ordering based
// on time, bucket, and then sequence.
type IDSlice []ID

// The number of IDs in the collection
func (p IDSlice) Len() int { return len(p) }

// Less returns true if the ID at index i is before the ID at index j.
func (p IDSlice) Less(i, j int) bool { return p[i].Before(p[j]) }

// Swap swaps the IDs with indexes i and j.
func (p IDSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// Sort sorts the ID slice.
func (p IDSlice) Sort() { sort.Sort(p) }
