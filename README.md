# ZPL Label Generator

The `zpl-gen` package is a Go utility for generating Zebra Programming Language (ZPL) label files. It enables you to use struct-based data to dynamically fill in placeholder values in ZPL templates. This is especially useful for automating label printing in manufacturing, logistics, or retail environments.

---

## Features

- Replace placeholders in ZPL templates with struct field values.
- Generate label content from either files or raw strings.
- Simple and extensible design.
- Supports printable ASCII characters sanitization.

---

## Installation

```bash
go get github.com/android-lewis/zpl-gen
````

---

## Usage

### Define a Template

Use placeholders in your ZPL template like this:

```
^XA
^FO50,50^ADN,36,20^FD<<ProductName>>^FS
^FO50,100^ADN,36,20^FD<<ProductCode>>^FS
^XZ
```

### Define a Struct

```go
type Product struct {
    ProductName string
    ProductCode string
}
```

### Generate the Label

```go
package main

import (
    "fmt"
    "github.com/android-lewis/zpl-gen"
)

func main() {
    product := Product{
        ProductName: "Widget A",
        ProductCode: "W12345",
    }

    placeholders := zpl.GenerateDetailMap(product)

    output, err := zpl.GenerateLabelFile("template.zpl", placeholders)
    if err != nil {
        panic(err)
    }

    fmt.Println(output)
}
```

Or for string-based templates:

```go
zplTemplate := `^XA
^FO50,50^FD<<ProductName>>^FS
^FO50,100^FD<<ProductCode>>^FS
^XZ`

label, _ := zpl.GenerateLabelString(zplTemplate, placeholders)
fmt.Println(label)
```

---

## API

### `GenerateDetailMap(details interface{}) map[string]string`

Creates a map of placeholders from a struct's exported fields. Each key is in the format `<<FieldName>>`.

---

### `GenerateLabelFile(filename string, detailsMap map[string]string) (string, error)`

Loads a file and replaces placeholders with values from the `detailsMap`.

---

### `GenerateLabelString(file string, detailsMap map[string]string) (string, error)`

Processes a raw ZPL string and replaces placeholders with values from the `detailsMap`.

---

## Placeholder Format

All placeholders must follow the format:

```
<<FieldName>>
```

Field names are **case-sensitive** and must match the exported fields of your struct.

---

## License

MIT © Lewis Baston

```
