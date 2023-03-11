package test

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/braydonk/yaml"
)

func TestExplicitDocumentStart(t *testing.T) {
	reader := bytes.NewReader([]byte{})
	decoder := yaml.NewDecoder(reader)
	var n yaml.Node
	err := decoder.Decode(&n)
	if err != nil && !errors.Is(err, io.EOF) {
		t.Fatalf("expect EOF, got %v", err)
	}

	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	enc.SetExplicitDocumentStart(true)
	err = enc.Encode(&n)
	if err != nil {
		t.Fatalf("expected nil err, got %v", err)
	}
	if !strings.Contains(buf.String(), "---") {
		t.Fatalf("expected buffer to contain document start\n document:\n%s", buf.String())
	}
}

func TestIndentlessBlockSequence(t *testing.T) {
	yml := `a:
- 1
- 2
b:
    c:
    - 1
    - 2
`
	reader := bytes.NewReader([]byte(yml))
	decoder := yaml.NewDecoder(reader)
	var n yaml.Node
	err := decoder.Decode(&n)
	if err != nil && !errors.Is(err, io.EOF) {
		t.Fatalf("expect EOF, got:\n%v", err)
	}

	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	enc.SetIndentlessBlockSequence(true)
	err = enc.Encode(&n)
	if err != nil {
		t.Fatalf("expected nil err, got:\n%v", err)
	}
	if buf.String() != yml {
		t.Fatalf("expected:\n%s\nactual:\n%s", yml, buf.String())
	}
}
