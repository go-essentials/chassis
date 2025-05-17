// =====================================================================================================================
// = LICENSE:       Copyright (c) 2025 Kevin De Coninck
// =
// =                Permission is hereby granted, free of charge, to any person
// =                obtaining a copy of this software and associated documentation
// =                files (the "Software"), to deal in the Software without
// =                restriction, including without limitation the rights to use,
// =                copy, modify, merge, publish, distribute, sublicense, and/or sell
// =                copies of the Software, and to permit persons to whom the
// =                Software is furnished to do so, subject to the following
// =                conditions:
// =
// =                The above copyright notice and this permission notice shall be
// =                included in all copies or substantial portions of the Software.
// =
// =                THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// =                EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// =                OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// =                NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// =                HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// =                WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// =                FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// =                OTHER DEALINGS IN THE SOFTWARE.
// =====================================================================================================================

// Package chassis implements a framework for building CLI applications.
package chassis

import (
	"io"
	"slices"
)

// Sorts the elements in set by their name.
func (set CommandSet) sort() CommandSet {
	slices.SortFunc(set, func(a, b Command) int {
		if a.Name < b.Name {
			return -1
		}

		return 1
	})

	return set
}

// Returns the length of the largest name of all elements in set.
func (set CommandSet) getMaxNameLen() int {
	commandNameMaxLen := 0

	for _, cmd := range set {
		if len(cmd.Name) > commandNameMaxLen {
			commandNameMaxLen = len(cmd.Name)
		}
	}

	return commandNameMaxLen
}

// Converts set into a map where the key matches the name of the command.
// This increases the lookup speed.
func (set CommandSet) asMap() map[string]func(io.Writer) {
	cmdMap := make(map[string]func(io.Writer), len(set))

	for _, cmd := range set {
		cmdMap[cmd.Name] = cmd.Handler
	}

	return cmdMap
}
