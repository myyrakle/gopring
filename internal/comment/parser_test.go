package comment

import (
	"reflect"
	"testing"
)

func Test_ParseCommentBlocks(t *testing.T) {
	type testCase struct {
		input string
		want  []string
	}

	testCases := []testCase{
		{
			input: `// @PathVariable("id")
			/* @PathVariable("id2")*/
			/*
					@PathVariable("id3")
			*/`,
			want: []string{
				`@PathVariable("id")`,
				`@PathVariable("id2")`,
				`@PathVariable("id3")`,
			},
		},
	}

	for _, tc := range testCases {
		got := ParseCommentBlocks(tc.input)

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("ParseParameters(%s) = %#v, want %#v", tc.input, got, tc.want)
		}
	}
}
