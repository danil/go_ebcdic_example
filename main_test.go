package ebcdic_test

import (
	"runtime"
)

func testLine() int {
	_, _, line, _ := runtime.Caller(1)
	return line
}

var testCases = []struct {
	fromEBCDIC string
	toEBCDIC   string
	fromASCII  string
	toASCII    string
	line       int
	benchmark  bool
}{
	{"\x00", "\x00", "\x00", "\x00", testLine(), false},
	{"\x01", "\x01", "\x01", "\x01", testLine(), false},
	{"\x02", "\x02", "\x02", "\x02", testLine(), false},
	{"\x03", "\x03", "\x03", "\x03", testLine(), false},
	{"\x04", "\x04", "\x1a", "\x1a", testLine(), false},
	{"\x05", "\x05", "\t", "\t", testLine(), false},
	{"\x06", "\x06", "\x1a", "\x1a", testLine(), false},
	{"\x07", "\x07", "\x7f", "\x7f", testLine(), false},
	{"\x08", "\x08", "\x1a", "\x1a", testLine(), false},
	{"\x09", "\x09", "\x1a", "\x1a", testLine(), false},
	{"\x0a", "\x0a", "\x1a", "\x1a", testLine(), false},
	{"\x0b", "\x0b", "\v", "\v", testLine(), false},
	{"\x0c", "\x0c", "\f", "\f", testLine(), false},
	{"\x0d", "\x0d", "\r", "\r", testLine(), false},
	{"\x0e", "\x0e", "\x0e", "\x0e", testLine(), false},
	{"\x0f", "\x0f", "\x0f", "\x0f", testLine(), false},
	{"\x10", "\x10", "\x10", "\x10", testLine(), false},
	{"\x11", "\x11", "\x11", "\x11", testLine(), false},
	{"\x12", "\x12", "\x12", "\x12", testLine(), false},
	{"\x13", "\x13", "\x13", "\x13", testLine(), false},
	{"\x14", "\x14", "\x1a", "\x1a", testLine(), false},
	{"\x15", "\x15", "\x85", "\x85", testLine(), false},
	{"\x16", "\x16", "\b", "\b", testLine(), false},
	{"\x17", "\x17", "\x1a", "\x1a", testLine(), false},
	{"\x18", "\x18", "\x18", "\x18", testLine(), false},
	{"\x19", "\x19", "\x19", "\x19", testLine(), false},
	{"\x1a", "\x1a", "\x1a", "\x1a", testLine(), false},
	{"\x1b", "\x1b", "\x1a", "\x1a", testLine(), false},
	{"\x1c", "\x1c", "\x1c", "\x1c", testLine(), false},
	{"\x1d", "\x1d", "\x1d", "\x1d", testLine(), false},
	{"\x1e", "\x1e", "\x1e", "\x1e", testLine(), false},
	{"\x1f", "\x1f", "\x1f", "\x1f", testLine(), false},
	{"\x20", "\x20", "\x1a", "\x1a", testLine(), false},
	{"\x21", "\x21", "\x1a", "\x1a", testLine(), false},
	{"\x22", "\x22", "\x1a", "\x1a", testLine(), false},
	{"\x23", "\x23", "\x1a", "\x1a", testLine(), false},
	{"\x24", "\x24", "\x1a", "\x1a", testLine(), false},
	{"\x25", "\x25", "\n", "\n", testLine(), false},
	{"\x26", "\x26", "\x17", "\x17", testLine(), false},
	{"\x27", "\x27", "\x1b", "\x1b", testLine(), false},
	{"\x28", "\x28", "\x1a", "\x1a", testLine(), false},
	{"\x29", "\x29", "\x1a", "\x1a", testLine(), false},
	{"\x2a", "\x2a", "\x1a", "\x1a", testLine(), false},
	{"\x2b", "\x2b", "\x1a", "\x1a", testLine(), false},
	{"\x2c", "\x2c", "\x1a", "\x1a", testLine(), false},
	{"\x2d", "\x2d", "\x05", "\x05", testLine(), false},
	{"\x2e", "\x2e", "\x06", "\x06", testLine(), false},
	{"\x2f", "\x2f", "\a", "\a", testLine(), false},
	{"\x30", "\x30", "\x1a", "\x1a", testLine(), false},
	{"\x31", "\x31", "\x1a", "\x1a", testLine(), false},
	{"\x32", "\x32", "\x16", "\x16", testLine(), false},
	{"\x33", "\x33", "\x1a", "\x1a", testLine(), false},
	{"\x34", "\x34", "\x1a", "\x1a", testLine(), false},
	{"\x35", "\x35", "\x1a", "\x1a", testLine(), false},
	{"\x36", "\x36", "\x1a", "\x1a", testLine(), false},
	{"\x37", "\x37", "\x04", "\x04", testLine(), false},
	{"\x38", "\x38", "\x1a", "\x1a", testLine(), false},
	{"\x39", "\x39", "\x1a", "\x1a", testLine(), false},
	{"\x3a", "\x3a", "\x1a", "\x1a", testLine(), false},
	{"\x3b", "\x3b", "\x1a", "\x1a", testLine(), false},
	{"\x3c", "\x3c", "\x14", "\x14", testLine(), false},
	{"\x3d", "\x3d", "\x15", "\x15", testLine(), false},
	{"\x3e", "\x3e", "\x1a", "\x1a", testLine(), false},
	{"\x3f", "\x3f", "\x1a", "\x1a", testLine(), false},
	{"\x40", "\x40", " ", " ", testLine(), false},
	{"\x41", "\x41", "\x1a", "\x1a", testLine(), false},
	{"\x42", "\x42", "\x1a", "\x1a", testLine(), false},
	{"\x43", "\x43", "\x1a", "\x1a", testLine(), false},
	{"\x44", "\x44", "\x1a", "\x1a", testLine(), false},
	{"\x45", "\x45", "\x1a", "\x1a", testLine(), false},
	{"\x46", "\x46", "\x1a", "\x1a", testLine(), false},
	{"\x47", "\x47", "\x1a", "\x1a", testLine(), false},
	{"\x48", "\x48", "\x1a", "\x1a", testLine(), false},
	{"\x49", "\x49", "\x1a", "\x1a", testLine(), false},
	{"\x4a", "\x4a", "\xa2", "\xa2", testLine(), false}, // ¢
	{"\x4b", "\x4b", ".", ".", testLine(), false},
	{"\x4c", "\x4c", "<", "<", testLine(), false},
	{"\x4d", "\x4d", "(", "(", testLine(), false},
	{"\x4e", "\x4e", "+", "+", testLine(), false},
	{"\x4f", "\x4f", "|", "|", testLine(), false},
	{"\x50", "\x50", "&", "&", testLine(), false},
	{"\x51", "\x51", "\x1a", "\x1a", testLine(), false},
	{"\x52", "\x52", "\x1a", "\x1a", testLine(), false},
	{"\x53", "\x53", "\x1a", "\x1a", testLine(), false},
	{"\x54", "\x54", "\x1a", "\x1a", testLine(), false},
	{"\x55", "\x55", "\x1a", "\x1a", testLine(), false},
	{"\x56", "\x56", "\x1a", "\x1a", testLine(), false},
	{"\x57", "\x57", "\x1a", "\x1a", testLine(), false},
	{"\x58", "\x58", "\x1a", "\x1a", testLine(), false},
	{"\x59", "\x59", "\x1a", "\x1a", testLine(), false},
	{"\x5a", "\x5a", "!", "!", testLine(), false},
	{"\x5b", "\x5b", "$", "$", testLine(), false},
	{"\x5c", "\x5c", "*", "*", testLine(), false},
	{"\x5d", "\x5d", ")", ")", testLine(), false},
	{"\x5e", "\x5e", ";", ";", testLine(), false},
	{"\x5f", "\x5f", "¬", "¬", testLine(), false},
	{"\x60", "\x60", "-", "-", testLine(), false},
	{"\x61", "\x61", "/", "/", testLine(), false},
	{"\x62", "\x62", "\x1a", "\x1a", testLine(), false},
	{"\x63", "\x63", "\x1a", "\x1a", testLine(), false},
	{"\x64", "\x64", "\x1a", "\x1a", testLine(), false},
	{"\x65", "\x65", "\x1a", "\x1a", testLine(), false},
	{"\x66", "\x66", "\x1a", "\x1a", testLine(), false},
	{"\x67", "\x67", "\x1a", "\x1a", testLine(), false},
	{"\x68", "\x68", "\x1a", "\x1a", testLine(), false},
	{"\x69", "\x69", "\x1a", "\x1a", testLine(), false},
	{"\x6a", "\x6a", "¦", "¦", testLine(), false},
	{"\x6b", "\x6b", ",", ",", testLine(), false},
	{"\x6c", "\x6c", "%", "%", testLine(), false},
	{"\x6d", "\x6d", "_", "_", testLine(), false},
	{"\x6e", "\x6e", ">", ">", testLine(), false},
	{"\x6f", "\x6f", "?", "?", testLine(), false},
	{"\x70", "\x70", "\x1a", "\x1a", testLine(), false},
	{"\x71", "\x71", "\x1a", "\x1a", testLine(), false},
	{"\x72", "\x72", "\x1a", "\x1a", testLine(), false},
	{"\x73", "\x73", "\x1a", "\x1a", testLine(), false},
	{"\x74", "\x74", "\x1a", "\x1a", testLine(), false},
	{"\x75", "\x75", "\x1a", "\x1a", testLine(), false},
	{"\x76", "\x76", "\x1a", "\x1a", testLine(), false},
	{"\x77", "\x77", "\x1a", "\x1a", testLine(), false},
	{"\x78", "\x78", "\x1a", "\x1a", testLine(), false},
	{"\x79", "\x79", "`", "`", testLine(), false},
	{"\x7a", "\x7a", ":", ":", testLine(), false},
	{"\x7b", "\x7b", "#", "#", testLine(), false},
	{"\x7c", "\x7c", "@", "@", testLine(), false},
	{"\x7d", "\x7d", "'", "'", testLine(), false},
	{"\x7e", "\x7e", "=", "=", testLine(), false},
	{"\x7f", "\x7f", `"`, `"`, testLine(), false},
	{"\x80", "\x80", "\x1a", "\x1a", testLine(), false},
	{"\x81", "\x81", "a", "a", testLine(), false},
	{"\x82", "\x82", "b", "b", testLine(), false},
	{"\x83", "\x83", "c", "c", testLine(), false},
	{"\x84", "\x84", "d", "d", testLine(), false},
	{"\x85", "\x85", "e", "e", testLine(), false},
	{"\x86", "\x86", "f", "f", testLine(), false},
	{"\x87", "\x87", "g", "g", testLine(), false},
	{"\x88", "\x88", "h", "h", testLine(), false},
	{"\x89", "\x89", "i", "i", testLine(), false},
	{"\x8a", "\x8a", "\x1a", "\x1a", testLine(), false},
	{"\x8b", "\x8b", "\x1a", "\x1a", testLine(), false},
	{"\x8c", "\x8c", "\x1a", "\x1a", testLine(), false},
	{"\x8d", "\x8d", "\x1a", "\x1a", testLine(), false},
	{"\x8e", "\x8e", "\x1a", "\x1a", testLine(), false},
	{"\x8f", "\x8f", "±", "±", testLine(), false},
	{"\x90", "\x90", "\x1a", "\x1a", testLine(), false},
	{"\x91", "\x91", "j", "j", testLine(), false},
	{"\x92", "\x92", "k", "k", testLine(), false},
	{"\x93", "\x93", "l", "l", testLine(), false},
	{"\x94", "\x94", "m", "m", testLine(), false},
	{"\x95", "\x95", "n", "n", testLine(), false},
	{"\x96", "\x96", "o", "o", testLine(), false},
	{"\x97", "\x97", "p", "p", testLine(), false},
	{"\x98", "\x98", "q", "q", testLine(), false},
	{"\x99", "\x99", "r", "r", testLine(), false},
	{"\x9a", "\x9a", "\x1a", "\x1a", testLine(), false},
	{"\x9b", "\x9b", "\x1a", "\x1a", testLine(), false},
	{"\x9c", "\x9c", "\x1a", "\x1a", testLine(), false},
	{"\x9d", "\x9d", "\x1a", "\x1a", testLine(), false},
	{"\x9e", "\x9e", "\x1a", "\x1a", testLine(), false},
	{"\x9f", "\x9f", "\x1a", "\x1a", testLine(), false},
	{"\xa0", "\xa0", "\x1a", "\x1a", testLine(), false},
	{"\xa1", "\xa1", "~", "~", testLine(), false},
	{"\xa2", "\xa2", "s", "s", testLine(), false},
	{"\xa3", "\xa3", "t", "t", testLine(), false},
	{"\xa4", "\xa4", "u", "u", testLine(), false},
	{"\xa5", "\xa5", "v", "v", testLine(), false},
	{"\xa6", "\xa6", "w", "w", testLine(), false},
	{"\xa7", "\xa7", "x", "x", testLine(), false},
	{"\xa8", "\xa8", "y", "y", testLine(), false},
	{"\xa9", "\xa9", "z", "z", testLine(), false},
	{"\xaa", "\xaa", "\x1a", "\x1a", testLine(), false},
	{"\xab", "\xab", "\x1a", "\x1a", testLine(), false},
	{"\xac", "\xac", "\x1a", "\x1a", testLine(), false},
	{"\xad", "\xad", "\x1a", "\x1a", testLine(), false},
	{"\xae", "\xae", "\x1a", "\x1a", testLine(), false},
	{"\xaf", "\xaf", "\x1a", "\x1a", testLine(), false},
	{"\xb0", "\xb0", "^", "^", testLine(), false},
	{"\xb1", "\xb1", "\x1a", "\x1a", testLine(), false},
	{"\xb2", "\xb2", "\x1a", "\x1a", testLine(), false},
	{"\xb3", "\xb3", "\x1a", "\x1a", testLine(), false},
	{"\xb4", "\xb4", "\x1a", "\x1a", testLine(), false},
	{"\xb5", "\xb5", "\x1a", "\x1a", testLine(), false},
	{"\xb6", "\xb6", "\x1a", "\x1a", testLine(), false},
	{"\xb7", "\xb7", "\x1a", "\x1a", testLine(), false},
	{"\xb8", "\xb8", "\x1a", "\x1a", testLine(), false},
	{"\xb9", "\xb9", "\x1a", "\x1a", testLine(), false},
	{"\xba", "\xba", "[", "[", testLine(), false},
	{"\xbb", "\xbb", "]", "]", testLine(), false},
	{"\xbc", "\xbc", "\x1a", "\x1a", testLine(), false},
	{"\xbd", "\xbd", "\x1a", "\x1a", testLine(), false},
	{"\xbe", "\xbe", "\x1a", "\x1a", testLine(), false},
	{"\xbf", "\xbf", "\x1a", "\x1a", testLine(), false},
	{"\xc0", "\xc0", "{", "{", testLine(), false},
	{"\xc1", "\xc1", "A", "A", testLine(), false},
	{"\xc2", "\xc2", "B", "B", testLine(), false},
	{"\xc3", "\xc3", "C", "C", testLine(), false},
	{"\xc4", "\xc4", "D", "D", testLine(), false},
	{"\xc5", "\xc5", "E", "E", testLine(), false},
	{"\xc6", "\xc6", "F", "F", testLine(), false},
	{"\xc7", "\xc7", "G", "G", testLine(), false},
	{"\xc8", "\xc8", "H", "H", testLine(), false},
	{"\xc9", "\xc9", "I", "I", testLine(), false},
	{"\xca", "\xca", "\x1a", "\x1a", testLine(), false},
	{"\xcb", "\xcb", "\x1a", "\x1a", testLine(), false},
	{"\xcc", "\xcc", "\x1a", "\x1a", testLine(), false},
	{"\xcd", "\xcd", "\x1a", "\x1a", testLine(), false},
	{"\xce", "\xce", "\x1a", "\x1a", testLine(), false},
	{"\xcf", "\xcf", "\x1a", "\x1a", testLine(), false},
	{"\xd0", "\xd0", "}", "}", testLine(), false},
	{"\xd1", "\xd1", "J", "J", testLine(), false},
	{"\xd2", "\xd2", "K", "K", testLine(), false},
	{"\xd3", "\xd3", "L", "L", testLine(), false},
	{"\xd4", "\xd4", "M", "M", testLine(), false},
	{"\xd5", "\xd5", "N", "N", testLine(), false},
	{"\xd6", "\xd6", "O", "O", testLine(), false},
	{"\xd7", "\xd7", "P", "P", testLine(), false},
	{"\xd8", "\xd8", "Q", "Q", testLine(), false},
	{"\xd9", "\xd9", "R", "R", testLine(), false},
	{"\xda", "\xda", "\x1a", "\x1a", testLine(), false},
	{"\xdb", "\xdb", "\x1a", "\x1a", testLine(), false},
	{"\xdc", "\xdc", "\x1a", "\x1a", testLine(), false},
	{"\xdd", "\xdd", "\x1a", "\x1a", testLine(), false},
	{"\xde", "\xde", "\x1a", "\x1a", testLine(), false},
	{"\xdf", "\xdf", "\x1a", "\x1a", testLine(), false},
	{"\xe0", "\xe0", "\\", "\\", testLine(), false},
	{"\xe1", "\xe1", "\x1a", "\x1a", testLine(), false},
	{"\xe2", "\xe2", "S", "S", testLine(), false},
	{"\xe3", "\xe3", "T", "T", testLine(), false},
	{"\xe4", "\xe4", "U", "U", testLine(), false},
	{"\xe5", "\xe5", "V", "V", testLine(), false},
	{"\xe6", "\xe6", "W", "W", testLine(), false},
	{"\xe7", "\xe7", "X", "X", testLine(), false},
	{"\xe8", "\xe8", "Y", "Y", testLine(), false},
	{"\xe9", "\xe9", "Z", "Z", testLine(), false},
	{"\xea", "\xea", "\x1a", "\x1a", testLine(), false},
	{"\xeb", "\xeb", "\x1a", "\x1a", testLine(), false},
	{"\xec", "\xec", "\x1a", "\x1a", testLine(), false},
	{"\xed", "\xed", "\x1a", "\x1a", testLine(), false},
	{"\xee", "\xee", "\x1a", "\x1a", testLine(), false},
	{"\xef", "\xef", "\x1a", "\x1a", testLine(), false},
	{"\xf0", "\xf0", "0", "0", testLine(), false},
	{"\xf1", "\xf1", "1", "1", testLine(), false},
	{"\xf2", "\xf2", "2", "2", testLine(), false},
	{"\xf3", "\xf3", "3", "3", testLine(), false},
	{"\xf4", "\xf4", "4", "4", testLine(), false},
	{"\xf5", "\xf5", "5", "5", testLine(), false},
	{"\xf6", "\xf6", "6", "6", testLine(), false},
	{"\xf7", "\xf7", "7", "7", testLine(), false},
	{"\xf8", "\xf8", "8", "8", testLine(), false},
	{"\xf9", "\xf9", "9", "9", testLine(), false},
	{"\xfa", "\xfa", "\x1a", "\x1a", testLine(), false},
	{"\xfb", "\xfb", "\x1a", "\x1a", testLine(), false},
	{"\xfc", "\xfc", "\x1a", "\x1a", testLine(), false},
	{"\xfd", "\xfd", "\x1a", "\x1a", testLine(), false},
	{"\xfe", "\xfe", "\x1a", "\x1a", testLine(), false},
	{"\xff", "\xff", "\x1a", "\x1a", testLine(), false},
	{"\xc8\x85\x93\x93\x96\x6b\x40\xe6\x96\x99\x93\x84\x5a", "\xc8\x85\x93\x93\x96\x6b\x40\xe6\x96\x99\x93\x84\x5a", "Hello, World!", "Hello, World!", testLine(), true},
}