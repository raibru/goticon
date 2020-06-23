# Packet Formatter (pktfmt)

## Bit Package Structure (bitpacks)

This tool take a definition file in JSON format which describe a data package and print a report with filled data into the package structure to stdout with base output for hexadecimal, decimal, octal and binarx. Additional report of the defined data structure can be included. This will print a graphical output of the data structure.

## Command Line Help

```sh
$ bitpacks -h
Create a package structure of defined word length with internal contents
Usage:
  bitpacks [flags]

Flags:
      --bin                  print binary result only (default all)
      --dec                  print decimal result only (default all)
  -d, --define-file string   definition file contains data structure (mandatory)
  -h, --help                 help for bitpacks
      --hex                  print hexadecimal result only (default all)
  -i, --input-data string    use a list of data parameter input depends on blocks
                                inside definition filei which are assignable (e.g. -i "1,2,..,n")
                                (mandatory when assignable value in definition exists)
      --oct                  print octal result only (default all)
      --postfix string       postfix string after the printed value. Only used in single hex, dec, oct and bin output
      --prefix string        prefix string before the printed value. Only used in single hex, dec, oct and bin output
  -r, --report               Report the calculated result to STDOUT or a report file
      --report-file string   report file contains data about calculated structure with data (default "STDOUT")
  -v, --version              display bitpacks version
```

## Example output

```text
$ C:\Projects\HomeWork-go\bitpackstruct> go run bitpacks.go -d .\definitions\struct-example.json -i "9" -r 
0009 9 11 0000000000001001
Definition : Example
Bit Size   : 16
Parity     : even
Bit field structure with input values:
  1. PARITY: Non
  2. UNDEF: Non
  3. IDValue: 9

+-------------------------------+-------------------------------+
| P  #2# #2# #2# #2# #2# #2# #2#|#2# #2# #2# #2# ~3~ ~3~ ~3~ ~3~|
+-------------------------------+-------------------------------+
| 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 1 | 0 | 0 | 1 |
+-------------------------------+-------------------------------+

Bases:
HEX:0009 DEC:9 OCT:11 BIN:0000000000001001

================================================================

```

```text
C:\Projects\HomeWork-go\bitpackstruct> go run bitpacks.go -d .\definitions\struct-example.json -i "8" -r
8008 32776 100010 1000000000001000
Definition : Example
Bit Size   : 16
Parity     : even
Bit field structure with input values:
  1. PARITY: Non
  2. UNDEF: Non
  3. IDValue: 8

+-------------------------------+-------------------------------+
| P  #2# #2# #2# #2# #2# #2# #2#|#2# #2# #2# #2# ~3~ ~3~ ~3~ ~3~|
+-------------------------------+-------------------------------+
| 1 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 1 | 0 | 0 | 0 |
+-------------------------------+-------------------------------+

Bases:
HEX:8008 DEC:32776 OCT:100010 BIN:1000000000001000

================================================================

```

## Definition of Bit Data Structure

The definition file use json structure for the output data. The definition file describe the package structureof one package.

***Package Definition:***

| Item | Description |
| ---- | ----------- |
| name | unique name of the package |
| len | the bit lengths of the package |
| parity | describe the kind of parity which is used for the package. Possible values are: even, odd, none |
| bitfields | holds an array of bit fields definitions |

***Bit-Field Definition:***

| Item | Description |
| ---- | ----------- |
| id | running identifcation number |
| pos | bit start position of the field inside the package |
| len | bit length of the field |
| desc | holds the reported field name for the output |
| assignable | descibe if this field get his value from the command parameter. The parameter count and order depends on the field definitions. Possible values are '***true***' for one parameter otherwise '***false***' not use a parameter for this field |

*Example JSON defintion*

```json
{
  "name": "Example",
  "len": 16,
  "parity": "even",
  "bitfields":
    [
      {
        "id": 1,
        "pos": 0,
        "len": 1,
        "desc": "PARITY",
        "assignable": false
      },
      {
        "id": 2,
        "pos": 1,
        "len": 11,
        "desc": "UNDEF",
        "assignable": false
      },
      {
        "id": 3,
        "pos": 12,
        "len": 4,
        "desc": "IDValue",
        "assignable": true
      }
    ]
}

```