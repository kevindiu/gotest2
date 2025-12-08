package models

// FunctionInfo represents a function or method to be tested.
type FunctionInfo struct {
	Name       string
	Receiver   *Receiver // nil if it's a standard function
	Parameters []*Field
	Results    []*Field
	TypeParams []*Field // For Generics: [T any, K comparable]
	IsExported bool
}

// Receiver represents the receiver of a method.
type Receiver struct {
	Name string // The variable name (e.g., "m")
	Type string // The type name (e.g., "*MyMap[K, V]")
}

// Field represents a parameter or return value.
type Field struct {
	Name       string // The variable name
	Type       string // The type representation (e.g., "string", "int", "[]T")
	Index      int
	IsVariadic bool
}

// Path represents a file path.
type Path string

// FileResult groups functions found in a single source file.
type FileResult struct {
	Path        string
	PackageName string
	Functions   []*FunctionInfo
}
