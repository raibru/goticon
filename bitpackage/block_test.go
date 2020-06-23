package bitpackage

import (
	"reflect"
	"testing"

	"gotest.tools/assert"
)

func Test_BlockStructureField_Binary_Exists(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var b Block

	// When
	var f interface{} = b.Binary
	s := f.(string)

	// Then
	assert.Assert(t, reflect.TypeOf(b.Binary) == reflect.TypeOf(s))
}

func Test_BlockStructureField_Input_Exists(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var b Block

	// When
	var f interface{} = b.Input
	s := f.(string)

	// Then
	assert.Assert(t, reflect.TypeOf(b.Input) == reflect.TypeOf(s))
}

func Test_BlockStructureField_BitField_Exists(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var b Block

	// When
	var f interface{} = b.Field
	bf := f.(BitField)

	// Then
	assert.Assert(t, reflect.TypeOf(b.Field) == reflect.TypeOf(bf))
}

func Test_IsParity_Success(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	conditions := []struct {
		input  string
		expect bool
	}{
		{"PARITY", true},
		{"OTHER", false},
	}

	for _, c := range conditions {

		// When
		var b Block
		b.Field.Desc = c.input
		result := b.IsParity()

		// Then
		if c.expect != result {
			t.Errorf("\n\tFailure: for '%s' expect '%v' but get '%v'", c.input, c.expect, result)
		}
	}
}

func Test_IsUndef_Success(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	conditions := []struct {
		input  string
		expect bool
	}{
		{"UNDEF", true},
		{"OTHER", false},
	}

	for _, c := range conditions {

		// When
		var b Block
		b.Field.Desc = c.input
		result := b.IsUndef()

		// Then
		if c.expect != result {
			t.Errorf("\n\tFailure: for '%s' expect '%v' but get '%v'", c.input, c.expect, result)
		}
	}
}

func Test_FillString_Success(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	conditions := []struct {
		desc   string
		id     int
		expect string
	}{
		{"PARITY", 1, " P "},
		{"UNDEF", 2, "#2#"},
		{"OTHER", 3, "~3~"},
	}

	for _, c := range conditions {

		// When
		var b Block
		b.Field.Desc = c.desc
		b.Field.ID = c.id
		result := b.FillString()

		// Then
		if c.expect != result {
			t.Errorf("\n\tFailure: for '%s' expect '%s' but get '%s'", c.desc, c.expect, result)
		}
	}
}

func Test_CalculateParity_OneParityBlock_Success(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	terms := []struct {
		desc   string
		parity string
		desc1  string
		input1 string
		desc2  string
		input2 string
		desc3  string
		input3 string
		expect string
	}{
		{"1st Block even parity odd bits", "even", "PARITY", "0", "BLOCK2", "1", "BLOCK3", "0", "1"},
		{"1st Block even parity even bits", "even", "PARITY", "0", "BLOCK2", "1", "BLOCK3", "1", "0"},
		{"1st Block odd parity odd bits", "odd", "PARITY", "0", "BLOCK2", "1", "BLOCK3", "0", "0"},
		{"1st Block odd parity even bits", "odd", "PARITY", "0", "BLOCK2", "1", "BLOCK3", "1", "1"},
		{"2nd Block even parity odd bits", "even", "BLOCK1", "1", "PARITY", "0", "BLOCK3", "0", "1"},
		{"2nd Block even parity even bits", "even", "BLOCK1", "1", "PARITY", "0", "BLOCK3", "1", "0"},
		{"2nd Block odd parity odd bits", "odd", "BLOCK1", "1", "PARITY", "0", "BLOCK3", "0", "0"},
		{"2nd Block odd parity even bits", "odd", "BLOCK1", "1", "PARITY", "0", "BLOCK3", "1", "1"},
		{"3rd Block even parity odd bits", "even", "BLOCK1", "1", "BLOCK2", "0", "PARITY", "0", "1"},
		{"3rd Block even parity even bits", "even", "BLOCK1", "1", "BLOCK2", "1", "PARITY", "0", "0"},
		{"3rd Block odd parity odd bits", "odd", "BLOCK1", "1", "BLOCK2", "0", "PARITY", "0", "0"},
		{"3rd Block odd parity even bits", "odd", "BLOCK1", "1", "BLOCK2", "1", "PARITY", "0", "1"},
	}

	for _, el := range terms {

		var b1, b2, b3 Block
		blks := make([]Block, 3)

		b1.Binary = el.input1
		b1.Field.Desc = el.desc1
		b1.Field.Pos = 0
		b1.Field.Len = 1
		blks[0] = b1

		b2.Binary = el.input2
		b2.Field.Desc = el.desc2
		b2.Field.Pos = 1
		b2.Field.Len = 1
		blks[1] = b2

		b3.Binary = el.input3
		b3.Field.Desc = el.desc3
		b3.Field.Pos = 2
		b3.Field.Len = 1
		blks[2] = b3

		// When
		i, p, err := CalculateParity(&blks, el.parity)
		r := blks[i].Binary

		// Then
		assert.NilError(t, err)
		assert.Assert(t, i != -1)
		if el.expect != p {
			t.Errorf("\n\tFailure: for '%s' expect '%s' but get '%s'", el.desc, el.expect, p)
		}
		if el.expect != r {
			t.Errorf("\n\tFailure: for '%s' expect '%s' but get '%s'", el.desc, el.expect, r)
		}
	}
}

func Test_CalculateParity_NoneParityBlock_Success(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	terms := []struct {
		desc   string
		parity string
		desc1  string
		input1 string
		desc2  string
		input2 string
		desc3  string
		input3 string
		expect string
	}{
		{"No Block parity", "none", "BLOCK1", "1", "BLOCK2", "1", "BLOCK3", "1", ""},
	}

	for _, el := range terms {

		var b1, b2, b3 Block
		blks := make([]Block, 3)

		b1.Binary = el.input1
		b1.Field.Desc = el.desc1
		b1.Field.Pos = 0
		b1.Field.Len = 1
		blks[0] = b1

		b2.Binary = el.input2
		b2.Field.Desc = el.desc2
		b2.Field.Pos = 1
		b2.Field.Len = 1
		blks[1] = b2

		b3.Binary = el.input3
		b3.Field.Desc = el.desc3
		b3.Field.Pos = 2
		b3.Field.Len = 1
		blks[2] = b3

		// When
		i, p, err := CalculateParity(&blks, el.parity)

		// Then
		assert.NilError(t, err)
		assert.Assert(t, i == -1)
		if el.expect != p {
			t.Errorf("\n\tFailure: for '%s' expect '%s' (empty string) but get '%s'", el.desc, el.expect, p)
		}
	}
}

func Test_CalculateParity_WrongParityType_ErrorMessage(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	terms := []struct {
		desc   string
		parity string
		desc1  string
		input1 string
		desc2  string
		input2 string
		desc3  string
		input3 string
		expect string
	}{
		{"1st Block parity with none", "none", "PARITY", "0", "BLOCK2", "1", "BLOCK3", "1", ""},
		{"1st Block parity with empty string", "", "PARITY", "0", "BLOCK2", "1", "BLOCK3", "1", ""},
		{"2nd Block parity with none", "none", "Block1", "1", "PARITY", "0", "BLOCK3", "1", ""},
		{"2nd Block parity with empty string", "", "Block1", "1", "PARITY", "0", "BLOCK3", "1", ""},
		{"3rd Block parity with none", "none", "Block1", "1", "BLOCK2", "1", "PARITY", "0", ""},
		{"3rd Block parity with empty string", "", "Block1", "1", "BLOCK2", "1", "PARITY", "0", ""},
	}

	for _, el := range terms {

		var b1, b2, b3 Block
		blks := make([]Block, 3)

		b1.Binary = el.input1
		b1.Field.Desc = el.desc1
		b1.Field.Pos = 0
		b1.Field.Len = 1
		blks[0] = b1

		b2.Binary = el.input2
		b2.Field.Desc = el.desc2
		b2.Field.Pos = 1
		b2.Field.Len = 1
		blks[1] = b2

		b3.Binary = el.input3
		b3.Field.Desc = el.desc3
		b3.Field.Pos = 2
		b3.Field.Len = 1
		blks[2] = b3

		// When
		i, p, err := CalculateParity(&blks, el.parity)

		// Then
		assert.ErrorContains(t, err, "Only 'even' or 'odd' are allowed")
		assert.Assert(t, i > -1)
		if el.expect != p {
			t.Errorf("\n\tFailure: for '%s' expect '%s' (empty string) but get '%s'", el.desc, el.expect, p)
		}
	}
}
