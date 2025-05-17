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

import "io"

// Represents a "chassis" application.
type application struct {
	logo        string     // ASCII Logo or visual identifier.
	name        string     // Application's name.
	description string     // Application's description.
	version     string     // Application's version.
	author      string     // Author name or contact information.
	commands    CommandSet // Set of commands available for the application.

	// Section: "Internal" fields.
	commandMap        map[string]func(io.Writer)
	commandNameMaxLen int // The largest length of the name of ALL commands.
}

// Command represents a single Application command.
type Command struct {
	Name        string          // Name (identifier) used to specify this command on the CLI.
	Description string          // Short description about this command.
	Handler     func(io.Writer) // Function to execute when the command is specified on the CLI.
}
