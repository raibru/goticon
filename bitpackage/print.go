package bitpackage

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

// BasePrintParam hold the base print parameter
type BasePrintParam struct {
	Type        string
	Prefix      string
	Postfix     string
	IsUppercase bool
}

// PrintBasesValue print
func PrintBasesValue(p *BitPackage, bs []Block, bpp *BasePrintParam, w io.Writer) error {
	s := ""
	for _, b := range bs {
		s += b.Binary
	}
	dec, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return err
	}

	hl := p.Len / 4
	hlz := strings.Repeat("0", hl) + strconv.FormatInt(dec, 16)
	hlz = hlz[len(hlz)-hl:]

	if bpp.IsUppercase {
		hlz = strings.ToUpper(hlz)
	}

	switch bpp.Type {
	case "hex":
		fmt.Fprintf(w, "%s%s%s\n", bpp.Prefix, hlz, bpp.Postfix)
		break
	case "dec":
		fmt.Fprintf(w, "%s%s%s\n", bpp.Prefix, strconv.FormatInt(dec, 10), bpp.Postfix)
		break
	case "oct":
		fmt.Fprintf(w, "%s%s%s\n", bpp.Prefix, strconv.FormatInt(dec, 8), bpp.Postfix)
		break
	case "bin":
		fmt.Fprintf(w, "%s%s%s\n", bpp.Prefix, strconv.FormatInt(dec, 2), bpp.Postfix)
		break
	default:
		fmt.Fprintf(w, "%s %s %s %s\n", hlz, strconv.FormatInt(dec, 10), strconv.FormatInt(dec, 8), s)
		break
	}

	return nil
}

// PrintStructFormat prints BitPackage in formated output
func PrintStructFormat(p *BitPackage, bs []Block, bpp *BasePrintParam, w io.Writer) error {
	sb := ""
	for _, b := range bs {
		sb += b.Binary
	}

	dec, err := strconv.ParseInt(sb, 2, 64)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "Definition : %s\n", p.Name)
	fmt.Fprintf(w, "Bit Size   : %d\n", p.Len)
	fmt.Fprintf(w, "Parity     : %s\n", p.Parity)
	fmt.Fprintln(w, "Bit field structure with input values:")

	for _, b := range bs {
		fmt.Fprintf(w, "  %d. %s: %s\n", b.Field.ID, b.Field.Desc, b.Input)
	}
	fmt.Fprintln(w, "")

	s1 := "+"
	s2 := "|"
	s3 := "+"
	s4 := "|"
	s6 := "+"

	blkIdx := 0

	for i := 0; i < p.Len; i++ {

		curBlk := bs[blkIdx]

		isFieldEnd := (curBlk.Field.Pos + curBlk.Field.Len) == i
		isEnd := i == p.Len-1

		if isFieldEnd && !isEnd {
			blkIdx++
			curBlk = bs[blkIdx]
		}

		s1 += "---"
		s2 += curBlk.FillString()
		s3 += "---"
		s4 += fmt.Sprintf(" %s ", string(sb[i]))
		s6 += "---"

		if (i+1)%8 == 0 {
			s1 += "+"
			s2 += "|"
			s3 += "+"
			s4 += "|"
			s6 += "+"
		} else {
			s1 += "-"
			s2 += " "
			s3 += "-"
			s4 += "|"
			s6 += "-"
		}
	}

	fmt.Fprintln(w, s1)
	fmt.Fprintln(w, s2)
	fmt.Fprintln(w, s3)
	fmt.Fprintln(w, s4)
	fmt.Fprintln(w, s6)

	hl := p.Len / 4
	hlz := strings.Repeat("0", hl) + strconv.FormatInt(dec, 16)
	hlz = hlz[len(hlz)-hl:]

	if bpp.IsUppercase {
		hlz = strings.ToUpper(hlz)
	}

	fmt.Fprintf(w, "\nBases:\nHEX:%s DEC:%s OCT:%s BIN:%s\n\n", hlz, strconv.FormatInt(dec, 10), strconv.FormatInt(dec, 8), sb)
	fmt.Fprintln(w, strings.Repeat("=", p.Len*4))
	fmt.Fprintln(w, "")

	return nil
}
