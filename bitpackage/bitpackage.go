package bitpackage

import (
	"errors"
	"strconv"
	"strings"
)

// BitPackage data structure
type BitPackage struct {
	Name      string     `json:"Name"`
	Len       int        `json:"len"`
	Parity    string     `json:"parity"`
	BitFields []BitField `json:"bitfields"`
}

// BitField data structure
type BitField struct {
	ID         int    `json:"id"`
	Pos        int    `json:"pos"`
	Len        int    `json:"len"`
	Desc       string `json:"desc"`
	Assignable bool   `json:"assignable"`
}

// EvaluateInputData check against definition file BitFields
// Returns false and error when evaluations fails otherwise true and nil
func (bpkg *BitPackage) EvaluateInputData(s string) (bool, error) {
	ba := 0
	for _, b := range bpkg.BitFields {
		if b.Assignable {
			ba++
		}
	}
	l := strings.Split(s, ",")
	if len(s) == 0 || len(l) != ba {
		return false, errors.New("Assignable input data must be the same as in definition file")
	}
	return true, nil
}

// ConvertDataBits convert all assignable input data into a bit string
func (bpkg *BitPackage) ConvertDataBits(s string) ([]Block, error) {
	pl := strings.Split(s, ",")
	pli := 0
	blocks := make([]Block, len(bpkg.BitFields))
	for i, f := range bpkg.BitFields {
		(blocks[i]).Field = f
		if f.Assignable {
			ds := pl[pli]
			(blocks[i]).Input = ds
			dv, err := strconv.ParseInt(ds, 10, 64)
			if err != nil {
				return []Block{}, err
			}
			bs := strings.Repeat("0", f.Len) + strconv.FormatInt(dv, 2)
			bs = bs[len(bs)-f.Len:]
			blocks[i].Binary = bs
			pli++
		} else {
			bs := strings.Repeat("0", f.Len)
			blocks[i].Binary = bs
			blocks[i].Input = "Non"
		}
	}

	return blocks, nil
}
