package bitpackage

import (
	"fmt"
	"strings"
)

// Block with BitFields and binary data
type Block struct {
	Field  BitField
	Binary string
	Input  string
}

// FillString returns string block type
func (b *Block) FillString() string {
	if b.IsParity() {
		return " P "
	} else if b.IsUndef() {
		return fmt.Sprintf("#%d#", b.Field.ID)
	} else {
		return fmt.Sprintf("~%d~", b.Field.ID)
	}
}

// IsParity return true if field description is parity
func (b *Block) IsParity() bool {
	return strings.ToUpper(b.Field.Desc) == "PARITY"
}

// IsUndef return true if field description is undef
func (b *Block) IsUndef() bool {
	return strings.ToUpper(b.Field.Desc) == "UNDEF"
}

// CalculateParity against parity type parameter even/odd. Retruns the index of
// found PARITY block, the calculated parity and an error. If parity is not
// found the index is -1 and the parity value as string is empty
func CalculateParity(bs *[]Block, s string) (pidx int, pval string, err error) {
	pidx = -1
	pval = ""

	bits := ""

	for i, b := range *bs {
		if strings.ToUpper(b.Field.Desc) == "PARITY" {
			pidx = i
		} else {
			bits += b.Binary
		}
	}

	if pidx > -1 {
		pval = "0"
		if strings.ToLower(s) == "even" {
			if strings.Count(bits, "1")%2 == 1 {
				pval = "1"
			}
		} else if strings.ToLower(s) == "odd" {
			if strings.Count(bits, "1")%2 == 0 {
				pval = "1"
			}
		} else {
			pval = ""
			return pidx, pval, fmt.Errorf("Use undefined parity string: %s. Only 'even' or 'odd' are allowed", s)
		}
		((*bs)[pidx]).Binary = pval
	}

	return pidx, pval, nil
}
