package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/raibru/pktfmt/bitpackage"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pktfmt",
	Short: "Package Structure Formatter",
	Long:  `Print package structure in nice format defined inside a yaml definition file`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := parseParam(cmd, args); err != nil {
			cmd.Help()
			fmt.Println("\nRoot command parsing error: ", err)
			os.Exit(1)
		}
	},
}

// Execute the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Execute root cmd has error", err)
		os.Exit(1)
	}
}

// parseParam parameter evaluation
func parseParam(cmd *cobra.Command, args []string) error {

	if prtVersion {
		PrintVersion(os.Stdout)
	} else if len(defFile) == 0 {
		return errors.New("Need a definition file")
	} else if len(defFile) >= 0 {
		bpkg := bitpackage.BitPackage{}
		if err := readDefinition(defFile, &bpkg); err != nil {
			return err
		}
		if _, err := bpkg.EvaluateInputData(dataParam); err != nil {
			return err
		}
		bls, err := bpkg.ConvertDataBits(dataParam)
		if err != nil {
			return err
		}
		if _, _, err := bitpackage.CalculateParity(&bls, bpkg.Parity); err != nil {
			return err
		}

		printData := bitpackage.BasePrintParam{}
		printData.Prefix = prefix
		printData.Postfix = postfix
		printData.IsUppercase = toUpper

		if prtHex {
			printData.Type = "hex"
		} else if prtDec {
			printData.Type = "dec"
		} else if prtOct {
			printData.Type = "oct"
		} else if prtBin {
			printData.Type = "bin"
		} else {
			printData.Type = "all"
		}

		bitpackage.PrintBasesValue(&bpkg, bls, &printData, os.Stdout)

		if rptResult {
			if err := writeReportFile(rptFile, &bpkg, bls, &printData); err != nil {
				return err
			}
		}

	} else {
		cmd.Help()
	}

	return nil
}

func readDefinition(f string, bpkg *bitpackage.BitPackage) error {
	content, errRead := ioutil.ReadFile(f)
	if errRead != nil {
		return errRead
	}

	errParse := json.Unmarshal([]byte(content), bpkg)
	if errParse != nil {
		return errParse
	}

	return nil
}

func writeReportFile(fn string, bpkg *bitpackage.BitPackage, bls []bitpackage.Block, bpp *bitpackage.BasePrintParam) error {
	var f *os.File
	if fn == "STDOUT" {
		f = os.Stdout
	} else if _, err := os.Stat(fn); os.IsNotExist(err) {
		fh, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		f = fh
	} else {
		fh, err := os.OpenFile(fn, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		f = fh
	}

	defer f.Close()
	bitpackage.PrintStructFormat(bpkg, bls, bpp, f)

	return nil
}
