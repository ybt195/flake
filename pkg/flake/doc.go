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

// Package flake generates 64-bit flake ids.
//
// Overview
//
// Flake ids are generated using a generator configured with a bucket ids. You can create a generator
// with:
//
//  g, err := New(0)
//
// With a generator you can generate unique ids with:
//
//  id, err := g.Next() // Generate an id, returning an error if one is unavailable.
//  id := g.Must() // Blocks until an id is available.
//
// When assigning a bucket id, ensure that no two generators use the same bucket ids.
package flake
