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
	"testing"

	"github.com/go-essentials/assert"
	"github.com/go-essentials/chassis"
)

// QA: Verify that `Application.Run` is implemented correctly.
func TestRun(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	for tcName, tc := range map[string]struct {
		app  chassis.Application
		want string
	}{
		"Nothing is printed when the `Application` is empty.": {
			app:  chassis.Application{},
			want: "",
		},
		"The logo is printed when it's defined.": {
			app: chassis.Application{
				Logo: "Chassis",
			},
			want: "Chassis\n\n",
		},
		"The logo and the description are printed when they are defined.": {
			app: chassis.Application{
				Logo:        "Chassis",
				Description: "A framework for building CLI applications.",
			},
			want: "Chassis\n\n  A framework for building CLI applications.\n\n",
		},
		"The logo, the description and the author are printed when they are defined.": {
			app: chassis.Application{
				Logo:        "Chassis",
				Description: "A framework for building CLI applications.",
				Author:      "Kevin De Coninck <kevin.dconinck@gmail.com>",
			},
			want: "Chassis\n\n  A framework for building CLI applications.\n  Author: Kevin De Coninck <kevin.dconinck@gmail.com>\n\n",
		},
		"The logo, the description, the author and the version are printed when they are defined.": {
			app: chassis.Application{
				Logo:        "Chassis",
				Description: "A framework for building CLI applications.",
				Version:     "1.0.0",
				Author:      "Kevin De Coninck <kevin.dconinck@gmail.com>",
			},
			want: "Chassis\n\n  A framework for building CLI applications.\n  Author: Kevin De Coninck <kevin.dconinck@gmail.com>\n\n  Version: 1.0.0\n\n",
		},
		"The logo and the version are printed when they are defined.": {
			app: chassis.Application{
				Logo:    "Chassis",
				Version: "1.0.0",
			},
			want: "Chassis\n\n  Version: 1.0.0\n\n",
		},
	} {
		tc := tc // Rebind 'tc'. Note: This is required to support "parallel" execution.

		// EXECUTION.
		t.Run(tcName, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// ARRANGE.
			var buf bytes.Buffer

			// ACT.
			tc.app.Run(&buf)

			// ASSERT.
			assert.Equal(t, buf.String(), tc.want, "", "\n\n"+
				"UT Name:  %s\n"+
				"\033[32mExpected: %s\033[0m\n"+
				"\033[31mActual:   %s\033[0m\n\n", tcName, tc.want, buf.String())
		})
	}
}
