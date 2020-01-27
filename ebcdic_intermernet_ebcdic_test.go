package ebcdic_test

import (
	"fmt"
	"runtime"
	"testing"

	intermernetebcdic "github.com/Intermernet/ebcdic"
	"golang.org/x/text/encoding"
)

func TestIntermernetEbcdicDecode(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range testCases {
		tc := tc
		switch tc.toASCII {
		case string(encoding.ASCIISub), "¬", "¦", "±":
			continue
		}
		t.Run(fmt.Sprintf("%#v %#v %d", tc.fromEBCDIC, tc.toASCII, tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			b := intermernetebcdic.Decode([]byte(tc.fromEBCDIC))
			if string(b) != tc.toASCII {
				t.Errorf("unexpected UTF-8 string %q, expected %q %s", string(b), tc.toASCII, linkToExample)
			}
		})
	}
}

func BenchmarkIntermernetEbcdicDecode(b *testing.B) {
	b.ReportAllocs()
	for _, tc := range testCases {
		if !tc.benchmark {
			continue
		}
		b.Run(fmt.Sprintf("%#v %#v %d", tc.fromEBCDIC, tc.toASCII, tc.line), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = intermernetebcdic.Decode([]byte(tc.fromEBCDIC))
			}
		})
	}
}

func TestIntermernetEbcdicEncode(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range testCases {
		tc := tc
		switch tc.fromASCII {
		case string(encoding.ASCIISub), "\x85", "\xad", "\xa0", "\xa2":
			continue
		}
		t.Run(fmt.Sprintf("%#v %#v %d", tc.fromASCII, tc.toEBCDIC, tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			b := intermernetebcdic.Encode([]byte(tc.fromASCII))
			if string(b) != tc.toEBCDIC {
				t.Errorf("unexpected EBCDIC string %#v, expected %#v %s", string(b), tc.toEBCDIC, linkToExample)
			}
		})
	}
}

func BenchmarkIntermernetEbcdicEncode(b *testing.B) {
	b.ReportAllocs()
	for _, tc := range testCases {
		if !tc.benchmark {
			continue
		}
		b.Run(fmt.Sprintf("%#v %#v %d", tc.fromASCII, tc.toEBCDIC, tc.line), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = intermernetebcdic.Encode([]byte(tc.toASCII))
			}
		})
	}
}
