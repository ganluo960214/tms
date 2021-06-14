# TMS

generate map and slice by type

### download
```cmd
go get -u github.com/ganluo960214/tms
```

### flags
```
-type string
    type name; must be set
-map
    generate map (default true)
-slice
    generate slice (default true)
```


### usage
```cmd
//go:generate tms -type=Type
```

### example

```go
//go:generate tms -type=Type
type Type byte

const (
    TypeCA Type = iota + 1
    TypeCB
    TypeCC
    TypeCD
    TypeCE
)
```

```go
// Code generated by "tms -type=Type"; DO NOT EDIT.

var (
	TMS_Type_Map = map[Type]interface{}{
		TypeCA: nil,
		TypeCB: nil,
		TypeCC: nil,
		TypeCD: nil,
		TypeCE: nil,
	}

	TMS_Type_Slice = []Type{
		TypeCA,
		TypeCB,
		TypeCC,
		TypeCD,
		TypeCE,
	}
)
```
