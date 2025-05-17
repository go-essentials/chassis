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

// App is the API for a "chassis" application.
type App interface {
	// Run the "chassis" application and write output to w.
	Run(io.Writer, []string)
}

// New returns a new App with the given logo, name, description, version, author and commands.
func New(logo, name, description, version, author string, commands CommandSet) App {
	return application{
		logo:              logo,
		name:              name,
		description:       description,
		version:           version,
		author:            author,
		commands:          commands.sort(),
		commandMap:        commands.asMap(),
		commandNameMaxLen: commands.getMaxNameLen(),
	}
}

// Run executes app based on its configuration and writes output to w.
func (app application) Run(w io.Writer, args []string) {
	for _, arg := range args {
		if handler, ok := app.commandMap[arg]; ok {
			handler(w)

			return
		}
	}

	app.writeHeader(w)
	app.writeCommands(w)
}

// Write to w the data of app.
func (app application) writeHeader(w io.Writer) {
	if app.logo != "" {
		fmt.Fprintf(w, "%s\n", app.logo)
		fmt.Fprintf(w, "\n")
	}

	if app.description != "" {
		fmt.Fprintf(w, "  %s\n", app.description)
	}

	if app.author != "" {
		fmt.Fprintf(w, "  Author: %s\n", app.author)
	}

	if app.description != "" || app.author != "" {
		fmt.Fprintf(w, "\n")
	}

	if app.version != "" {
		fmt.Fprintf(w, "  Version: %s\n", app.version)
		fmt.Fprintf(w, "\n")
	}
}

// Write to w the commands of app.
func (app application) writeCommands(w io.Writer) {
	if len(app.commands) == 0 {
		return
	}

	commandNameMaxLen := 0

	for _, cmd := range app.commands {
		if len(cmd.Name) > commandNameMaxLen {
			commandNameMaxLen = len(cmd.Name)
		}
	}

	fmt.Fprintf(w, "Commands:\n")
	for _, cmd := range app.commands {
		fmt.Fprintf(w, "  %-*s    %s\n", commandNameMaxLen, cmd.Name, cmd.Description)
	}

	fmt.Fprintf(w, "\n")
}
