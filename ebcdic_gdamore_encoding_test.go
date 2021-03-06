package ebcdic_test

import (
	"fmt"
	"runtime"
	"testing"

	gdamoreencoding "github.com/gdamore/encoding"
	"golang.org/x/text/encoding"
)

func TestGdamoreEncodingDecode(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range testCases {
		tc := tc
		switch tc.toASCII {
		case string(encoding.ASCIISub), "\x85", "\xad", "\xa0", "\xa2":
			continue
		}
		t.Run(fmt.Sprintf("%#v %#v %d", tc.fromEBCDIC, tc.toASCII, tc.line), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)
			decoder := gdamoreencoding.EBCDIC.NewDecoder()
			dst := make([]byte, len(tc.toASCII))
			nDst, nSrc, err := decoder.Transform(dst, []byte(tc.fromEBCDIC), true)
			if err != nil {
				t.Fatalf("unexpected decoder transform error: %v - %s", err, linkToExample)
			}
			if nSrc != len(tc.fromEBCDIC) {
				t.Errorf("unexpected source bytes number %d, expected %d %s",
					nSrc, len(tc.fromEBCDIC), linkToExample)
			}
			if nDst != len(tc.toASCII) {
				t.Errorf("unexpected destination bytes number %d, expected %d %s",
					nDst, len(tc.toASCII), linkToExample)
			}
			if string(dst[:nDst]) != tc.toASCII {
				t.Errorf("unexpected UTF-8 string %q, expected %q %s", string(dst[:nDst]), tc.toASCII, linkToExample)
			}
		})
	}
}

func BenchmarkGdamoreEncodingDecode(b *testing.B) {
	b.ReportAllocs()
	for _, tc := range testCases {
		if !tc.benchmark {
			continue
		}
		b.Run(fmt.Sprintf("%#v %#v %d", tc.fromEBCDIC, tc.toASCII, tc.line), func(b *testing.B) {
			decoder := gdamoreencoding.EBCDIC.NewDecoder()
			for i := 0; i < b.N; i++ {
				dst := make([]byte, len(tc.toASCII))
				_, _, _ = decoder.Transform(dst, []byte(tc.fromEBCDIC), true)
			}
		})
	}
}

func TestGdamoreEncodingEncode(t *testing.T) {
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
			encoder := gdamoreencoding.EBCDIC.NewEncoder()
			dst := make([]byte, len(tc.toEBCDIC))
			nDst, nSrc, err := encoder.Transform(dst, []byte(tc.fromASCII), true)
			if err != nil {
				t.Fatalf("unexpected encoder transform error: %v - %s", err, linkToExample)
			}
			if nSrc != len(tc.fromASCII) {
				t.Errorf("unexpected source bytes number %d, expected %d %s",
					nSrc, len(tc.fromASCII), linkToExample)
			}
			if nDst != len(tc.toEBCDIC) {
				t.Errorf("unexpected destination bytes number %d, expected %d %s",
					nDst, len(tc.toEBCDIC), linkToExample)
			}
			if string(dst[:nDst]) != tc.toEBCDIC {
				t.Errorf("unexpected EBCDIC string %#v, expected %#v %s", string(dst[:nDst]), tc.toEBCDIC, linkToExample)
			}
		})
	}
}

func BenchmarkGdamoreEncodingEncode(b *testing.B) {
	b.ReportAllocs()
	for _, tc := range testCases {
		if !tc.benchmark {
			continue
		}
		b.Run(fmt.Sprintf("%#v %#v %d", tc.fromASCII, tc.toEBCDIC, tc.line), func(b *testing.B) {
			encoder := gdamoreencoding.EBCDIC.NewEncoder()
			for i := 0; i < b.N; i++ {
				dst := make([]byte, len(tc.toEBCDIC))
				_, _, _ = encoder.Transform(dst, []byte(tc.fromASCII), true)
			}
		})
	}
}
