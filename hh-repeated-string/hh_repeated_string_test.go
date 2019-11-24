package repeated_string

import "testing"

func TestAba2_1(t *testing.T) {
	var returned int64 = CountALetterCoccurrences("aba", 2)
	var expected int64 = 1
	if returned != expected {
		t.Errorf("expected: %v, returned: %v", expected, returned)
	}
}

func TestAba3_2(t *testing.T) {
	var returned int64 = CountALetterCoccurrences("aba", 3)
	var expected int64 = 2
	if returned != expected {
		t.Errorf("expected: %v, returned: %v", expected, returned)
	}
}

func TestAba4_3(t *testing.T) {
	var returned int64 = CountALetterCoccurrences("aba", 4)
	var expected int64 = 3
	if returned != expected {
		t.Errorf("expected: %v, returned: %v", expected, returned)
	}
}

func TestAba5_3(t *testing.T) {
	var returned int64 = CountALetterCoccurrences("aba", 5)
	var expected int64 = 3
	if returned != expected {
		t.Errorf("expected: %v, returned: %v", expected, returned)
	}
}

func TestAba6_4(t *testing.T) {
	var returned int64 = CountALetterCoccurrences("aba", 6)
	var expected int64 = 4
	if returned != expected {
		t.Errorf("expected: %v, returned: %v", expected, returned)
	}
}
