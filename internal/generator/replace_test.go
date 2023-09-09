package generator

import (
	"reflect"
	"testing"
)

func Test_replaceImportPath(t *testing.T) {
	type testCase struct {
		input string
		want  string
	}

	testCases := []testCase{
		{
			input: `package controller

import "github.com/myyrakle/gopring/src/service"`,
			want: `package controller

import "github.com/myyrakle/gopring/dist/service"`,
		},
	}

	for _, testCase := range testCases {
		got := replaceImportPath(testCase.input)

		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("replaceImportPath(%s) = %#v, want %#v", testCase.input, got, testCase.want)
		}
	}
}
