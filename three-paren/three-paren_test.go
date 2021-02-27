package three_paren

import "testing"

func TestTheThreeParen(t *testing.T) {
	expects := make([]bool, 0)
	testCases := make([]string, 0)

	expects = append(expects, []bool{
		true,
		false,
		true,
		false,
		false,
		true,
		true,
		false,
	}...)
	testCases = append(testCases, []string{
		"(())",
		"(()))(()",
		"(()())()",
		"({])[]",
		"[({}])",
		"()[{}()]",
		"[{{()}{}}]",
		"[()[{}()[]{{]]()]",
	}...)

	for i, testCase := range testCases {
		if expects[i] != IsPerfectParen(testCase) {
			t.Errorf(`Error!! result should equal be %v, but it is %v`, expects[i], IsPerfectParen(testCase))
		}
	}
}
