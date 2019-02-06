package roman_test

import (
	. "github.com/Ethik-69/roman_tdd"
	"reflect"
	"testing"
)

func TestI(t *testing.T) {
	v := RomanConverter(1)
	if v != "I" {
		t.Errorf("Returned %s instead of I", v)
	}
}

func TestII(t *testing.T) {
	v := RomanConverter(2)
	if v != "II" {
		t.Errorf("Returned %s instead of II", v)
	}
}

func TestV(t *testing.T) {
	v := RomanConverter(5)
	if v != "V" {
		t.Errorf("Returned %s instead of V", v)
	}
}

func TestX(t *testing.T) {
	v := RomanConverter(10)
	if v != "X" {
		t.Errorf("Returned %s instead of X", v)
	}
}

func TestVI(t *testing.T) {
	v := RomanConverter(6)
	if v != "VI" {
		t.Errorf("Returned %s instead of VI", v)
	}
}

func TestIV(t *testing.T) {
	v := RomanConverter(4)
	if v != "IV" {
		t.Errorf("Returned %s instead of IV", v)
	}
}

func TestMLXVI(t *testing.T) {
	v := RomanConverter(1066)
	if v != "MLXVI" {
		t.Errorf("Returned %s instead of MLXVI", v)
	}
}

func TestMDCCCXCIII(t *testing.T) {
	v := RomanConverter(1893)
	if v == "MDCCCLXXXXIII" {
		t.Errorf("Returned %s, it's a trap !", v)
	}
	if v != "MDCCCXCIII" {
		t.Errorf("Returned %s instead of MDCCCXCIII", v)
	}
}

func TestAlmostLast(t *testing.T) {
	a := [7]int{56, 854, 11, 986, 325, 222, 3652}
	// 11 56  222    325    854     986      3652
	// XI LVI CCXXII CCCXXV DCCCLIV CMLXXXVI MMMDCLII
	v := RomanConverterArray(a)
	if typeOfV := reflect.TypeOf(v); typeOfV.Kind() != reflect.Array {
		t.Errorf("returned type should be an array, got %v", typeOfV.Kind())
	}

	b := [7]string{"LVI", "DCCCLIV", "XI", "CMLXXXVI", "CCCXXV", "CCXXII", "MMMDCLII"}
	if v != b {
		t.Errorf("Returned %s instead of %s", v, b)
	}
}

func TestLast(t *testing.T) {
	a := [7]int{56, 854, 11, 986, 325, 222, 3652}
	v := RomanConverterArray(a)
	if typeOfV := reflect.TypeOf(v); typeOfV.Kind() != reflect.Array {
		t.Errorf("returned type should be an array, got %v", typeOfV.Kind())
	}

	b := [7]string{"XI", "LVI", "CCXXII", "CCCXXV", "DCCCLIV", "CMLXXXVI", "MMMDCLII"}
	if v != b {
		t.Errorf("Returned %s instead of %s", v, b)
	}
}
