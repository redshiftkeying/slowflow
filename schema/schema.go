package schema

import "bytes"

// GetRootSchema get file schema.graphql
func GetRootSchema() string {
	buf := bytes.Buffer{}
	for _, name := range AssetNames() {
		b := MustAsset(name)
		buf.Write(b)

		if len(b) > 0 && b[len(b)-1] != '\n' {
			buf.WriteByte('\n')
		}
	}
	return buf.String()
}
// Resolver for schema
type Resolver struct{}

// Hello function
func (*Resolver) Hello() string { return "Hello, world!" }
