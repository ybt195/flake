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
	"database/sql/driver"
	"fmt"
)

// Value returns a form of the id that database drivers must be able to handle.
func (f ID) Value() (driver.Value, error) {
	return int64(f.Uint64()), nil
}

// Scan assigns a value from a database driver to the id.
func (f *ID) Scan(src interface{}) (err error) {
	switch src := src.(type) {
	case int64:
		*f = ID(uint64(src))
	case []byte:
		*f, err = FromBytes(src)
	case nil:
		return nil
	default:
		return fmt.Errorf("incompatible type for flake id: %T", src)
	}
	return err
}
