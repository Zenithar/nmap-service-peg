# NMAP Service Probe PEG Parser

**Status:** Proof of Concept (PoC) / Unmaintained

The **NMAP Service Probe PEG Parser** is a Proof of Concept (PoC) project 
designed to demonstrate the use of Parsing Expression Grammar (PEG) for parsing 
the NMAP service probe file format. This parser, implemented in Go, leverages 
the [pigeon](https://github.com/mna/pigeon) library to read and process the NMAP
service probe definitions and transform them into a structured Go representation.

## Features

- **PEG-based Parsing**: The parser uses a PEG grammar to efficiently parse NMAP
  service probe files, ensuring a clean and maintainable codebase.
- **Go Struct Output**: The parser generates a Go struct that accurately 
  represents the parsed data, making it easy to integrate with other Go 
  applications or tools.
- **Simple & Lightweight**: Designed for simplicity, the parser is a minimal 
  implementation intended to showcase the capabilities of PEG-based parsing.

## Overview

The NMAP service probe file defines patterns that NMAP uses to identify services
running on a remote host. These patterns are stored in a 
[plain-text file format](https://github.com/nmap/nmap/blob/master/nmap-service-probes),
where each probe is represented by a set of key-value fields. The parser 
transforms this textual representation into a Go struct that can be easily 
manipulated or processed further.

This parser is based on the NMAP service probe file format specification, which 
is available in the official [NMAP documentation](https://nmap.org/book/vscan-fileformat.html).

## Why PEG?

Parsing Expression Grammar (PEG) is preferred over regular expressions for 
parsing structured and complex data formats due to its:

- Support for hierarchical and recursive structures.
- Unambiguous, deterministic parsing approach.
- Flexibility and expressiveness in defining complex grammars.
- Better error handling and maintainability.

## Specification

The NMAP service probe file consists of multiple probes, with each probe 
containing a set of fields that describe how NMAP identifies a service. These 
fields include information about the probe's match patterns, conditions, and 
expected responses.

For more information on the file format, please refer to the official 
[NMAP service probe format documentation](https://nmap.org/book/vscan-fileformat.html).

## Limitations

- **Delimiter Restrictions**: The current implementation of the PEG grammar does 
  not support the flexible, sed-like matchers that allow any character to serve 
  as a delimiter. Instead, it is limited to specific delimiters: `/`, `|`, `=`, 
  `%`, and `@`.
- **Basic Functionality**: This project is still in the PoC stage, so some 
  advanced features and edge cases might not be fully handled yet.

## Usage

To use the NMAP Service Probe PEG Parser, follow these steps:

```bash
# To generate the parser code from the PEG grammar
$ go run github.com/mna/pigeon -o grammar.go grammar.peg
```

You can test the grammar by running the following command:

```bash
$ cd cmd/parser
$ go build
$ cat nmap-service-probes.txt | ./parser
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file 
for more details.
