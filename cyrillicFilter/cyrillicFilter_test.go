package cyrillicFilter

import (
	"reflect"
	"testing"
)

type S struct {
	A int64
	B bool
	C string
}

type S2 struct {
	A S
	B string
	C int64
}

type S3 struct {
	A int
	B string
	C *string
	D float64
	E S2
}

type S4 struct {}

type S5 struct {
	A *string
}

type S6 struct {
	A int
}

// testing CyrFilter for different structs
func TestCyrFilter(t *testing.T) {
	strs := []string{
		"Прыфвлод! Alphabet! Алфавит!",
		"Apple!ЯблокоApple!",
	}
	expectedStrs := []string{
		"! Alphabet! !",
		"Apple!Apple!",
	}

	testTable := []struct {
		s          S
		s2         S2
		s3         S3
		s4         S4
		s5         S5
		s6         S6
		expectedS  S
		expectedS2 S2
		expectedS3 S3
		expectedS4 S4
		expectedS5 S5
		expectedS6 S6
	}{
		{ //test 0
			s3: S3{
				A: 32,
				B: "Hello! Привет!",
				C: &strs[0], // "Прыфвлод! Alphabet! Алфавит!"
				D: 78.0,
				E: S2{
					A: S{
						A: 0,
						B: true,
						C: "ваолдыЪGTX1080 Ti",
					},
					B: "AstanaАстанаAstana",
					C: 21,
				},
			},
			expectedS3: S3{
				A: 32,
				B: "Hello! !",
				C: &expectedStrs[0], //"! Alphabet! !",
				D: 78.0,
				E: S2{
					A: S{
						A: 0,
						B: true,
						C: "GTX1080 Ti",
					},
					B: "AstanaAstana",
					C: 21,
				},
			},
		},
		{ // test 1
			s: S{
				A: 100,
				B: false,
				C: "Camelot",
			},
			expectedS: S{
				A: 100,
				B: false,
				C: "Camelot",
			},
		},
		{ // test 2
			s: S{
				C: "ДЛявЩ",
			},
			expectedS: S{
				C: "",
			},
		},
		{ // test 3
			s4:         S4{},
			expectedS4: S4{},
		},
		{ // test 4
			s5: S5{
				A: &strs[1], // "Apple!ЯблокоApple!"
			},
			expectedS5: S5{
				A: &expectedStrs[1], // "Apple!Apple!"
			},
		},
		{ // test 5
			s6: S6{
				A: 10,
			},
			expectedS6: S6{
				A: 10,
			},
		},
	}

	for i, testCase := range testTable {
		CyrFilter(&testCase.s)
		CyrFilter(&testCase.s2)
		CyrFilter(&testCase.s3)
		CyrFilter(&testCase.s4)
		CyrFilter(&testCase.s5)
		CyrFilter(&testCase.s6)

		t.Logf("Running test %v", i)
		if !reflect.DeepEqual(testCase.s, testCase.expectedS) {
			t.Errorf("Incorrect result. Expect %#v, got %#v", testCase.expectedS, testCase.s)
		}
		if !reflect.DeepEqual(testCase.s2, testCase.expectedS2) {
			t.Errorf("Incorrect result. Expect %#v, got %#v", testCase.expectedS2, testCase.s2)
		}
		if !reflect.DeepEqual(testCase.s3, testCase.expectedS3) {
			t.Errorf("Incorrect result. Expect %#v, got %#v", testCase.expectedS3, testCase.s3)
		}
		if !reflect.DeepEqual(testCase.s4, testCase.expectedS4) {
			t.Errorf("Incorrect result. Expect %#v, got %#v", testCase.expectedS4, testCase.s4)
		}
		if !reflect.DeepEqual(testCase.s5, testCase.expectedS5) {
			t.Errorf("Incorrect result. Expect %#v, got %#v", testCase.expectedS5, testCase.s5)
		}
		if !reflect.DeepEqual(testCase.s6, testCase.expectedS6) {
			t.Errorf("Incorrect result. Expect %#v, got %#v", testCase.expectedS6, testCase.s6)
		}
	}
}