package banner_test

import (
	"fmt"
	"testing"

	"github.com/fnordic/banner"
)

func ExampleInline() {
	fmt.Println("start of banner")
	fmt.Println(banner.Inline("hey world."))
	fmt.Println("end of banner")
	// Output:
	// start of banner
	//  _                                    _     _
	// | |_   ___  _  _   __ __ __ ___  _ _ | | __| |
	// | ' \ / -_)| || |  \ V  V // _ \| '_|| |/ _` | _
	// |_||_|\___| \_, |   \_/\_/ \___/|_|  |_|\__,_|(_)
	//             |__/
	// end of banner
}

func ExampleInline_lowercase() {
	fmt.Println("start of banner")
	fmt.Println(banner.Inline("abcdefghij"))
	fmt.Println(banner.Inline("klmnopqrst"))
	fmt.Println(banner.Inline("uvwxyz"))
	fmt.Println("end of banner")
	// Output:
	// start of banner
	//        _            _        __        _     _    _
	//  __ _ | |__  __  __| | ___  / _| __ _ | |_  (_)  (_)
	// / _` || '_ \/ _|/ _` |/ -_)|  _|/ _` || ' \ | |  | |
	// \__,_||_.__/\__|\__,_|\___||_|  \__, ||_||_||_| _/ |
	//                                 |___/          |__/
	//  _    _                                         _
	// | |__| | _ __   _ _   ___  _ __  __ _  _ _  ___| |_
	// | / /| || '  \ | ' \ / _ \| '_ \/ _` || '_|(_-<|  _|
	// |_\_\|_||_|_|_||_||_|\___/| .__/\__, ||_|  /__/ \__|
	//                           |_|      |_|
	//  _  _ __ ____ __ ____ __ _  _  ___
	// | || |\ V /\ V  V /\ \ /| || ||_ /
	//  \_,_| \_/  \_/\_/ /_\_\ \_, |/__|
	//                          |__/
	// end of banner
}

func ExampleInline_uppercase() {
	fmt.Println("start of banner")
	fmt.Println(banner.Inline("ABCDEFGHIJ"))
	fmt.Println(banner.Inline("KLMNOPQRST"))
	fmt.Println(banner.Inline("UVWXYZ"))
	fmt.Println("end of banner")
	// Output:
	// start of banner
	//        _            _        __        _     _    _
	//  __ _ | |__  __  __| | ___  / _| __ _ | |_  (_)  (_)
	// / _` || '_ \/ _|/ _` |/ -_)|  _|/ _` || ' \ | |  | |
	// \__,_||_.__/\__|\__,_|\___||_|  \__, ||_||_||_| _/ |
	//                                 |___/          |__/
	//  _    _                                         _
	// | |__| | _ __   _ _   ___  _ __  __ _  _ _  ___| |_
	// | / /| || '  \ | ' \ / _ \| '_ \/ _` || '_|(_-<|  _|
	// |_\_\|_||_|_|_||_||_|\___/| .__/\__, ||_|  /__/ \__|
	//                           |_|      |_|
	//  _  _ __ ____ __ ____ __ _  _  ___
	// | || |\ V /\ V  V /\ \ /| || ||_ /
	//  \_,_| \_/  \_/\_/ /_\_\ \_, |/__|
	//                          |__/
	// end of banner
}

func TestInline(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"jjj", `
   _    _    _
  (_)  (_)  (_)
  | |  | |  | |
 _/ | _/ | _/ |
|__/ |__/ |__/
`},
		{"j j", `
   _      _
  (_)    (_)
  | |    | |
 _/ |   _/ |
|__/   |__/
`},
		{"j", `
   _
  (_)
  | |
 _/ |
|__/
`},
		{"@?!", `
 ___  ___  ___
|__ \|__ \|__ \
  /_/  /_/  /_/
 (_)  (_)  (_)
`},
		{"ccc", `
 __  __  __
/ _|/ _|/ _|
\__|\__|\__|
`},
		{" ", `

`},
		{"", `

`},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			output := banner.Inline(test.input)
			expected := test.expected[1 : len(test.expected)-1]
			if expected != output {
				t.Log("output: \n" + output)
				t.Log("expected: \n" + expected)
				t.Errorf("output differs")
			}
		})
	}
}
