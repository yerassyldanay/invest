package tests

import (
	"github.com/yerassyldanay/invest/utils/helper"
	"testing"
)

func TestCheckSQLInjection(t *testing.T) {
	type testCaseStruct struct {
		Str					string				`json:"str"`
		IsErr				bool				`json:"is_err"`
		//Err					error				`json:"err"`
	}

	testCases := []testCaseStruct{
		{
			Str:   "abcABC",
			IsErr: false,
		},
		{
			Str:   "this is `injection",
			IsErr: true,
		},
	}

	for _, testCase := range testCases {
		err := helper.OnlyCheckSqlInjection(testCase.Str)
		if (err != nil) != testCase.IsErr {
			t.Error("REGEXP: sql injection | expected something else not ", )
		}
	}
}
