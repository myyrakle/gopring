package properties

import (
	"reflect"
	"testing"
)

func Test_ParseProperties(t *testing.T) {
	type testCase struct {
		inputText string
		want      Properties
	}

	testCases := []testCase{
		{
			inputText: `server.port = 8888
spring.profiles = dev
spring.application.name = demo
`,
			want: Properties{
				Property{
					Key:   "server.port",
					Value: "8888",
				},
				Property{
					Key:   "spring.profiles",
					Value: "dev",
				},
				Property{
					Key:   "spring.application.name",
					Value: "demo",
				},
			},
		},
	}

	for _, tc := range testCases {
		got := ParseProperties(tc.inputText)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("ParseProperties(%q) = %q; want %q", tc.inputText, got, tc.want)
		}
	}
}
