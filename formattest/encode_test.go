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

func TestPadLineComments(t *testing.T) {
	formatTestCase{
		name:             "pad line comments",
		folder:           "pad_line_comments",
		configureDecoder: noopDecoder,
		configureEncoder: func(enc *yaml.Encoder) {
			enc.SetPadLineComments(2)
		},
	}.Run(t)
}

func TestAltArrayIndent(t *testing.T) {
	formatTestCase{
		name:             "alternate array indent",
		folder:           "alt_array_indent",
		configureDecoder: noopDecoder,
		configureEncoder: func(enc *yaml.Encoder) {
			enc.SetIndent(4)
			enc.SetArrayIndent(1)
		},
	}.Run(t)
}

func TestAltArrayIndentRoot(t *testing.T) {
	formatTestCase{
		name:             "alternate array indent (root)",
		folder:           "alt_array_indent_root",
		configureDecoder: noopDecoder,
		configureEncoder: func(enc *yaml.Encoder) {
			enc.SetIndent(4)
			enc.SetArrayIndent(1)
		},
	}.Run(t)
}
