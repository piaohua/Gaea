package parser

import (
	"testing"
)

func Test_Preview(t *testing.T) {
	var list = []struct {
		text     string
		stmtType int
	}{
		{
			"/*test*/select curdate();",
			StmtSelect,
		},
		{
			"/*test*/ select curtime()",
			StmtSelect,
		},
		{
			"select now()",
			StmtSelect,
		},
		{
			"/*test*/select/*test*/ name from test.users where id=1;",
			StmtSelect,
		},
		{
			`/* test */ 
			select * from 
			test.users where id=1`,
			StmtSelect,
		},
	}

	for _, v := range list {
		stmtType := Preview(v.text)
		if stmtType != v.stmtType {
			t.Log(StmtType(stmtType))
			t.Fatalf("%v", v)
		}
	}
}
