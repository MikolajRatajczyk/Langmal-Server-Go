package models

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStorabeStringArray_ScanWhenValidData(t *testing.T) {
	original := StorableStringArray{"d", "e", "f"}
	validData, err := json.Marshal(original)
	if err != nil {
		t.Error("Can't convert to []byte and continue the test")
	}
	sut := StorableStringArray{}

	err = sut.Scan(validData)

	if err != nil {
		t.Error("Should not fail")
	}

	if !cmp.Equal(sut, original) {
		t.Error("Original and scanned should be equal")
	}
}

func TestStorabeStringArray_ScanWhenEmptyBytes(t *testing.T) {
	emptyBytes := []byte{}
	sut := StorableStringArray{}

	err := sut.Scan(emptyBytes)

	if err == nil {
		t.Error("Should fail for empty bytes")
	}
}

func TestStorabeStringArray_ValueWhenFilled(t *testing.T) {
	testStorabeStringArray_Value(t, StorableStringArray{"a", "b", "c"})
}

func TestStorabeStringArray_ValueWhenEmpty(t *testing.T) {
	testStorabeStringArray_Value(t, StorableStringArray{})
}

func testStorabeStringArray_Value(t *testing.T, sut StorableStringArray) {
	value, err := sut.Value()

	if err != nil {
		t.Error("Should return no error, got: " + err.Error())
	}

	if value == nil {
		t.Error("Should return not nil")
	}
}
