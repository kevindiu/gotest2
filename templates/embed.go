package templates

import "embed"

// FS holds the embedded templates.
// Updated to ensuring recompilation after template fixes. 4
//
//go:embed *.tmpl
//go:embed *.tmpl
var FS embed.FS
