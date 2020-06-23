package bitpackage

import (
	"reflect"
	"testing"

	"gotest.tools/assert"
)

func Test_BitPackageStructureField_Name_Exists(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bp BitPackage

	// When
	var f interface{} = bp.Name
	v := f.(string)

	// Then
	assert.Assert(t, reflect.TypeOf(bp.Name) == reflect.TypeOf(v))
}

func Test_BitPackageStructureField_Len_Exists(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bp BitPackage

	// When
	var f interface{} = bp.Len
	v := f.(int)

	// Then
	assert.Assert(t, reflect.TypeOf(bp.Len) == reflect.TypeOf(v))
}

func Test_BitPackageStructureField_Parity_Exists(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bp BitPackage

	// When
	var f interface{} = bp.Parity
	v := f.(string)

	// Then
	assert.Assert(t, reflect.TypeOf(bp.Parity) == reflect.TypeOf(v))
}

func Test_BitPackageStructureField_BitFields_Exists(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bp BitPackage

	// When
	var f interface{} = bp.BitFields
	v := f.([]BitField)

	// Then
	assert.Assert(t, reflect.TypeOf(bp.BitFields) == reflect.TypeOf(v))
}

func Test_BitFieldStructureField_ID_Exists(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bf BitField

	// When
	var f interface{} = bf.ID
	v := f.(int)

	// Then
	assert.Assert(t, reflect.TypeOf(bf.ID) == reflect.TypeOf(v))
}

func Test_BitFieldStructureField_Pos_Exists(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bf BitField

	// When
	var f interface{} = bf.Pos
	v := f.(int)

	// Then
	assert.Assert(t, reflect.TypeOf(bf.Pos) == reflect.TypeOf(v))
}

func Test_BitFieldStructureField_Len_Exists(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bf BitField

	// When
	var f interface{} = bf.Len
	v := f.(int)

	// Then
	assert.Assert(t, reflect.TypeOf(bf.Len) == reflect.TypeOf(v))
}

func Test_BitFieldStructureField_Desc_Exists(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bf BitField

	// When
	var f interface{} = bf.Desc
	v := f.(string)

	// Then
	assert.Assert(t, reflect.TypeOf(bf.Desc) == reflect.TypeOf(v))
}

func Test_BitFieldStructureField_Assignable_Exists(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bf BitField

	// When
	var f interface{} = bf.Assignable
	v := f.(bool)

	// Then
	assert.Assert(t, reflect.TypeOf(bf.Assignable) == reflect.TypeOf(v))
}

func Test_EvaluateInputData_AssignableParamterEqual_Success(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	terms := []struct {
		context   string
		desc1     string
		assigne1  bool
		desc2     string
		assigne2  bool
		desc3     string
		assigne3  bool
		parameter string
	}{
		{"Assignable 1st field", "BLOCK1", true, "BLOCK2", false, "BLOCK3", false, "p1"},
		{"Assignable 2nd field", "BLOCK1", false, "BLOCK2", true, "BLOCK3", false, "p2"},
		{"Assignable 3rd field", "BLOCK1", false, "BLOCK2", false, "BLOCK3", true, "p3"},
		{"Assignable 1st, 2nd field", "BLOCK1", true, "BLOCK2", true, "BLOCK3", false, "p1,p2"},
		{"Assignable 1st, 3rd field", "BLOCK1", true, "BLOCK2", false, "BLOCK3", true, "p1,p3"},
		{"Assignable 2nd, 3rd field", "BLOCK1", false, "BLOCK2", true, "BLOCK3", true, "p2,p3"},
		{"Assignable 1st, 2nd, 3rd field", "BLOCK1", true, "BLOCK2", true, "BLOCK3", true, "p1,p2,p3"},
	}

	for _, el := range terms {

		var f1, f2, f3 BitField
		bfs := make([]BitField, 3)
		bp := BitPackage{}

		f1.Desc = el.desc1
		f1.Assignable = el.assigne1
		bfs[0] = f1

		f2.Desc = el.desc2
		f2.Assignable = el.assigne2
		bfs[1] = f2

		f3.Desc = el.desc3
		f3.Assignable = el.assigne3
		bfs[2] = f3

		bp.BitFields = bfs

		// When
		pass, err := bp.EvaluateInputData(el.parameter)

		// Then
		if !pass {
			t.Errorf("\n\tFailure: for '%s' expect passing but get false for '%s'", el.context, el.parameter)
		}
		assert.NilError(t, err)
	}
}

func Test_EvaluateInputData_AssignableParamterNotEqual_FailureError(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	terms := []struct {
		context   string
		desc1     string
		assigne1  bool
		desc2     string
		assigne2  bool
		desc3     string
		assigne3  bool
		parameter string
	}{
		{"Assignable 1st field 2 paramter", "BLOCK1", true, "BLOCK2", false, "BLOCK3", false, "p1,p2"},
		{"Assignable 1st field no paramter", "BLOCK1", true, "BLOCK2", false, "BLOCK3", false, ""},
		{"Assignable 2nd field 2 paramter", "BLOCK1", false, "BLOCK2", true, "BLOCK3", false, "p1,p2"},
		{"Assignable 2nd field no paramter", "BLOCK1", false, "BLOCK2", true, "BLOCK3", false, ""},
		{"Assignable 3rd field 2 paramter", "BLOCK1", false, "BLOCK2", false, "BLOCK3", true, "p1,p2"},
		{"Assignable 3rd field no paramter", "BLOCK1", false, "BLOCK2", false, "BLOCK3", true, ""},
		{"Assignable 1st, 2nd field 3 paramter", "BLOCK1", true, "BLOCK2", true, "BLOCK3", false, "p1,p2,p3"},
		{"Assignable 1st, 2nd field 1 paramter", "BLOCK1", true, "BLOCK2", true, "BLOCK3", false, "p1"},
		{"Assignable 1st, 2nd field no paramter", "BLOCK1", true, "BLOCK2", true, "BLOCK3", false, ""},
		{"Assignable 1st, 3rd field 3 paramter", "BLOCK1", true, "BLOCK2", false, "BLOCK3", true, "p1,p2,p3"},
		{"Assignable 1st, 3rd field 1 paramter", "BLOCK1", true, "BLOCK2", false, "BLOCK3", true, "p1"},
		{"Assignable 1st, 3rd field no paramter", "BLOCK1", true, "BLOCK2", false, "BLOCK3", true, ""},
		{"Assignable 2nd, 3rd field 3 paramter", "BLOCK1", false, "BLOCK2", true, "BLOCK3", true, "p1,p2,p3"},
		{"Assignable 2nd, 3rd field 1 paramter", "BLOCK1", false, "BLOCK2", true, "BLOCK3", true, "p1"},
		{"Assignable 2nd, 3rd field no paramter", "BLOCK1", false, "BLOCK2", true, "BLOCK3", true, ""},
		{"Assignable 1st, 2nd, 3rd field 4 paramter", "BLOCK1", true, "BLOCK2", true, "BLOCK3", true, "p1,p2,p3,p4"},
		{"Assignable 1st, 2nd, 3rd field 2 paramter", "BLOCK1", true, "BLOCK2", true, "BLOCK3", true, "p1,p2"},
		{"Assignable 1st, 2nd, 3rd field 1 paramter", "BLOCK1", true, "BLOCK2", true, "BLOCK3", true, "p1"},
		{"Assignable 1st, 2nd, 3rd field no paramter", "BLOCK1", true, "BLOCK2", true, "BLOCK3", true, ""},
	}

	for _, el := range terms {

		var f1, f2, f3 BitField
		bfs := make([]BitField, 3)
		bp := BitPackage{}

		f1.Desc = el.desc1
		f1.Assignable = el.assigne1
		bfs[0] = f1

		f2.Desc = el.desc2
		f2.Assignable = el.assigne2
		bfs[1] = f2

		f3.Desc = el.desc3
		f3.Assignable = el.assigne3
		bfs[2] = f3

		bp.BitFields = bfs

		// When
		pass, err := bp.EvaluateInputData(el.parameter)

		// Then
		if pass {
			t.Errorf("\n\tFailure: for '%s' expect NOT passing but get true for '%s'", el.context, el.parameter)
		}
		assert.ErrorContains(t, err, "Assignable input data must be the same as in definition file")
	}
}
