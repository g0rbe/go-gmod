package checksum_test

import (
	"bytes"
	"testing"

	"github.com/g0rbe/gmod/cryptography/checksum"
)

func TestData256(t *testing.T) {

	d := []byte{0}
	r := []byte{0x6e, 0x34, 0xb, 0x9c, 0xff, 0xb3, 0x7a, 0x98, 0x9c, 0xa5, 0x44, 0xe6, 0xbb, 0x78, 0xa, 0x2c, 0x78, 0x90, 0x1d, 0x3f, 0xb3, 0x37, 0x38, 0x76, 0x85, 0x11, 0xa3, 0x6, 0x17, 0xaf, 0xa0, 0x1d}

	s := checksum.Data256(d)
	if !bytes.Equal(s, r) {
		t.Fatalf("Checksum mismatch\n")
	}

}
