package test

import (
	"bytes"
	"testing"

	"github.com/braydonk/yaml"
)

func TestBlockScalar(t *testing.T) {
	yml := `commands: >
    [ -f "/usr/local/bin/foo" ] &&
    echo "skip install" ||
    go install github.com/foo/foo@latest
`
	dec := yaml.NewDecoder(bytes.NewReader([]byte(yml)))
	dec.SetScanBlockScalarAsLiteral(true)
	var n yaml.Node
	dec.Decode(&n)
	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	enc.SetAssumeBlockAsLiteral(true)
	err := enc.Encode(&n)
	if err != nil {
		t.Fatal(err)
	}
	resultStr := buf.String()
	if resultStr != yml {
		t.Fatalf("expected result string to match input:\nexpected:\n%s\nresult:\n%s\n", yml, resultStr)
	}
}
