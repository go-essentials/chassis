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
	"fmt"
	"io"
)

// Run executes app based on its configuration and writes output to w.
func (app Application) Run(w io.Writer) {
	app.Commands.sort()

	app.writeHeader(w)
	app.writeCommands(w)
}

// Write to w the data of app.
func (app Application) writeHeader(w io.Writer) {
	if app.Logo != "" {
		fmt.Fprintf(w, "%s\n", app.Logo)
		fmt.Fprintf(w, "\n")
	}

	if app.Description != "" {
		fmt.Fprintf(w, "  %s\n", app.Description)
	}

	if app.Author != "" {
		fmt.Fprintf(w, "  Author: %s\n", app.Author)
	}

	if app.Description != "" || app.Author != "" {
		fmt.Fprintf(w, "\n")
	}

	if app.Version != "" {
		fmt.Fprintf(w, "  Version: %s\n", app.Version)
		fmt.Fprintf(w, "\n")
	}
}

// Write to w the commands of app.
func (app Application) writeCommands(w io.Writer) {
	if len(app.Commands) == 0 {
		return
	}

	commandNameMaxLen := 0

	for _, cmd := range app.Commands {
		if len(cmd.Name) > commandNameMaxLen {
			commandNameMaxLen = len(cmd.Name)
		}
	}

	fmt.Fprintf(w, "Commands:\n")
	for _, cmd := range app.Commands {
		fmt.Fprintf(w, "  %-*s    %s\n", commandNameMaxLen, cmd.Name, cmd.Description)
	}

	fmt.Fprintf(w, "\n")
}
