package formattest

import (
	"testing"

	"github.com/braydonk/yaml"
)

func TestExplicitDocumentStart(t *testing.T) {
	formatTestCase{
		name:             "explicit document start",
		folder:           "document_start",
		configureDecoder: noopDecoder,
		configureEncoder: func(enc *yaml.Encoder) {
			enc.SetExplicitDocumentStart(true)
		},
	}.Run(t)
}

func TestIndentless(t *testing.T) {
	formatTestCase{
		name:             "indentless array",
		folder:           "indentless",
		configureDecoder: noopDecoder,
		configureEncoder: func(enc *yaml.Encoder) {
			enc.SetIndentlessBlockSequence(true)
		},
	}.Run(t)
}

func TestIndentedToIndentless(t *testing.T) {
	formatTestCase{
		name:             "indented to indentless array",
		folder:           "indented_to_indentless",
		configureDecoder: noopDecoder,
		configureEncoder: func(enc *yaml.Encoder) {
			enc.SetIndentlessBlockSequence(true)
		},
	}.Run(t)
}

func TestBlockScalar(t *testing.T) {
	formatTestCase{
		name:   "block scalar decoding and encoding",
		folder: "block_scalar",
		configureDecoder: func(dec *yaml.Decoder) {
			dec.SetScanBlockScalarAsLiteral(true)
		},
		configureEncoder: func(enc *yaml.Encoder) {
			enc.SetAssumeBlockAsLiteral(true)
		},
	}.Run(t)
}

func TestDropMergeTag(t *testing.T) {
	formatTestCase{
		name:             "drop merge tag",
		folder:           "drop_merge_tag",
		configureDecoder: noopDecoder,
		configureEncoder: func(enc *yaml.Encoder) {
			enc.SetDropMergeTag(true)
		},
	}.Run(t)
}
