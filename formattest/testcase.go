package formattest

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/braydonk/yaml"
)

type decoderConfigureFunc func(*yaml.Decoder)

func noopDecoder(dec *yaml.Decoder) {}

type encoderConfigureFunc func(*yaml.Encoder)

type formatTestCase struct {
	name             string
	folder           string
	configureDecoder decoderConfigureFunc
	configureEncoder encoderConfigureFunc
}

func (tc formatTestCase) Run(t *testing.T) {
	t.Run(tc.name, func(t *testing.T) {
		// Read test input
		input, err := tc.readTestdataFile("input.yaml")
		if err != nil {
			t.Fatal(err)
		}

		// Configure Decoder
		reader := bytes.NewReader(input)
		decoder := yaml.NewDecoder(reader)
		tc.configureDecoder(decoder)

		// Decode input document
		var n yaml.Node
		err = decoder.Decode(&n)
		if err != nil && !errors.Is(err, io.EOF) {
			t.Fatalf("expect EOF, got:\n%v", err)
		}

		// Configure Encoder
		var buf bytes.Buffer
		enc := yaml.NewEncoder(&buf)
		tc.configureEncoder(enc)

		// Encode the decoded input document
		err = enc.Encode(&n)
		if err != nil {
			t.Fatalf("expected nil err, got:\n%v", err)
		}

		// Read the expected output
		expected, err := tc.readTestdataFile("expected.yaml")
		if err != nil {
			t.Fatal(err)
		}
		if buf.String() != string(expected) {
			t.Fatalf("expected:\n%s\nactual:\n%s", string(expected), buf.String())
		}
	})
}

func (tc formatTestCase) readTestdataFile(path string) ([]byte, error) {
	fullPath := filepath.Join("testdata", tc.folder, path)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, fmt.Errorf("path %s not found", fullPath)
	}
	return content, nil
}
