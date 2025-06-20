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

// QA: Verify the public API of the `chassis` package.
package chassis_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/go-essentials/assert"
	"github.com/go-essentials/chassis"
)

// QA: Verify that `Application.Run` is implemented correctly.
func TestRun(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// UTILITY FUNCTIONS.
	fmtOutput := func(output string) string {
		return strings.Replace(output, "\n", "\n          ", -1)
	}

	for tcName, tc := range map[string]struct {
		logo, name, description, version, author string
		commands                                 []chassis.Command
		args                                     []string
		want                                     string
	}{
		"Nothing is printed when the `Application` is empty.": {
			want: "",
		},
		"The logo is printed when it's defined.": {
			logo: "Chassis",
			want: "Chassis\n\n",
		},
		"The logo and the description are printed when they are defined.": {
			logo:        "Chassis",
			description: "A framework for building CLI applications.",
			want:        "Chassis\n\n  A framework for building CLI applications.\n\n",
		},
		"The logo, the description and the author are printed when they are defined.": {
			logo:        "Chassis",
			description: "A framework for building CLI applications.",
			author:      "Kevin De Coninck <kevin.dconinck@gmail.com>",
			want:        "Chassis\n\n  A framework for building CLI applications.\n  Author: Kevin De Coninck <kevin.dconinck@gmail.com>\n\n",
		},
		"The logo, the description, the author and the version are printed when they are defined.": {
			logo:        "Chassis",
			description: "A framework for building CLI applications.",
			version:     "1.0.0",
			author:      "Kevin De Coninck <kevin.dconinck@gmail.com>",
			want:        "Chassis\n\n  A framework for building CLI applications.\n  Author: Kevin De Coninck <kevin.dconinck@gmail.com>\n\n  Version: 1.0.0\n\n",
		},
		"The logo and the version are printed when they are defined.": {
			logo:    "Chassis",
			version: "1.0.0",
			want:    "Chassis\n\n  Version: 1.0.0\n\n",
		},
		"The commands are printed (and sorted) when they are defined.": {
			commands: []chassis.Command{
				{
					Name:        "version",
					Description: "Prints version information and quits.",
				},
				{
					Name:        "bulk",
					Description: "Perform a bulk operation.",
				},
				{
					Name:        "clean-temp",
					Description: "Clean the temporary folder.",
				},
			},
			want: "Commands:\n  bulk          Perform a bulk operation.\n  clean-temp    Clean the temporary folder.\n  version       Prints version information and quits.\n\n",
		},
		"The handler is executed when the arguments maps to a valid command.": {
			args: []string{"analyze"},
			commands: []chassis.Command{
				{
					Name:        "analyze",
					Description: "Perform an analysis.",
					Handler: func(w io.Writer) {
						fmt.Fprintf(w, "Analysis performed.")
					},
				},
				{
					Name:        "clean-temp",
					Description: "Clean the temporary folder.",
					Handler: func(w io.Writer) {
						fmt.Fprintf(w, "Cleaned temporary folder.")
					},
				},
			},
			want: "Analysis performed.",
		},
		"The last handler is executed when the arguments maps to multiple valid commands.": {
			args: []string{"analyze"},
			commands: []chassis.Command{
				{
					Name:        "analyze",
					Description: "Perform an analysis.",
					Handler: func(w io.Writer) {
						fmt.Fprintf(w, "Analysis performed.")
					},
				},
				{
					Name:        "analyze",
					Description: "Perform an analysis.",
					Handler: func(w io.Writer) {
						fmt.Fprintf(w, "ALT. Analysis performed.")
					},
				},
			},
			want: "ALT. Analysis performed.",
		},
	} {
		tc := tc // Rebind 'tc'. Note: This is required to support "parallel" execution.

		// EXECUTION.
		t.Run(tcName, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// ARRANGE.
			var buf bytes.Buffer
			app := chassis.New(tc.logo, tc.name, tc.description, tc.version, tc.author, tc.commands)

			// ACT.
			app.Run(&buf, tc.args)

			// ASSERT.
			assert.Equalf(t, buf.String(), tc.want, "\n\n"+
				"UT Name:  %s\n"+
				"\033[32mExpected: %s\033[0m\n"+
				"\033[31mActual:   %s\033[0m\n\n", tcName, fmtOutput(tc.want), fmtOutput(buf.String()))
		})
	}
}

// BENCHMARK: Measure the performance for looking up a command.
func benchmarkLookup(b *testing.B, size int) {
	// UTILITY FUNCTION: Populate some test data.
	populateData := func(amount int) (chassis.CommandSet, []string) {
		commandSet := make(chassis.CommandSet, amount)
		argsSet := make([]string, amount)

		for idx := range amount {
			commandSet[idx] = chassis.Command{
				Name: fmt.Sprint(idx),
				Handler: func(w io.Writer) {
					// NOTE: Intentionally left blank.
				},
			}

			argsSet[idx] = fmt.Sprint(amount - idx - 1)
		}

		return commandSet, argsSet
	}

	// ARRANGE.
	var buf bytes.Buffer

	commandSet, argsSet := populateData(size)
	app := chassis.New("", "", "", "", "", commandSet)

	b.ResetTimer()

	// RUN.
	for b.Loop() {
		app.Run(&buf, argsSet)
	}
}

// BENCHMARK COLLECTION: Measure the performance for looking up a command.
func BenchmarkLookup_Small(b *testing.B)  { benchmarkLookup(b, 10) }
func BenchmarkLookup_Medium(b *testing.B) { benchmarkLookup(b, 1_000) }
func BenchmarkLookup_Large(b *testing.B)  { benchmarkLookup(b, 10_000) }
func BenchmarkLookup_XLarge(b *testing.B) { benchmarkLookup(b, 1_000_000) }
