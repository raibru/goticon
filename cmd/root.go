package cmd

var defFile string
var rptResult bool
var rptFile string
var prefix string
var postfix string
var prtHex bool
var prtDec bool
var prtOct bool
var prtBin bool
var toUpper bool
var dataParam string
var prtVersion bool

func init() {
	rootCmd.PersistentFlags().StringVarP(&defFile, "define-file", "d", "", "definition file contains data structure (mandatory)")
	rootCmd.PersistentFlags().StringVarP(&rptFile, "report-file", "", "STDOUT", "report file contains data about calculated structure with data")
	rootCmd.PersistentFlags().BoolVarP(&rptResult, "report", "r", false, "Report the calculated result to STDOUT or a report file")
	rootCmd.PersistentFlags().StringVarP(&prefix, "prefix", "", "", "prefix string before the printed value. Only used in single hex, dec, oct and bin output")
	rootCmd.PersistentFlags().StringVarP(&postfix, "postfix", "", "", "postfix string after the printed value. Only used in single hex, dec, oct and bin output")
	rootCmd.PersistentFlags().BoolVarP(&prtHex, "hex", "", false, "print hexadecimal result only (default all)")
	rootCmd.PersistentFlags().BoolVarP(&prtDec, "dec", "", false, "print decimal result only (default all)")
	rootCmd.PersistentFlags().BoolVarP(&prtOct, "oct", "", false, "print octal result only (default all)")
	rootCmd.PersistentFlags().BoolVarP(&prtBin, "bin", "", false, "print binary result only (default all)")
	rootCmd.PersistentFlags().BoolVarP(&toUpper, "to-uppercase", "", false, "print base used letters in uppercase (default: lowercase)")
	rootCmd.PersistentFlags().BoolVarP(&prtVersion, "version", "v", false, "display bitpacks version")

	rootCmd.PersistentFlags().StringVarP(&dataParam, "input-data", "i", "", `use a list of data parameter input depends on blocks
	inside definition filei which are assignable (e.g. -i "1,2,..,n")
	(mandatory when assignable value in definition exists)`)
}
