package parser

import (
	"testing"
)

func Test_leadingCommentEnd(t *testing.T) {
	var list = []struct {
		text string
		body string
	}{
		{
			"/*test*/select 1;",
			"select 1;",
		},
		{
			"/*test*/ select 1",
			"select 1",
		},
		{
			"select databases()",
			"select databases()",
		},
	}

	for _, v := range list {
		end := leadingCommentEnd(v.text)
		if v.text[end:] != v.body {
			t.Log(end, v.text[end:])
			t.Fatalf("%v", v)
		}
	}
}

func Test_trailingCommentStart(t *testing.T) {
	var list = []struct {
		text string
		body string
	}{
		{
			"select 1/*test*/",
			"select 1",
		},
		{
			"select 1 /*test*/",
			"select 1",
		},
		{
			"select version()",
			"select version()",
		},
	}

	for _, v := range list {
		start := trailingCommentStart(v.text)
		if v.text[:start] != v.body {
			t.Log(start, v.text[:start])
			t.Fatalf("%v", v)
		}
	}
}
