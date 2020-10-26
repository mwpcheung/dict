package dict

import (
	"bytes"
	"unicode/utf8"
)

func isAnsiAsUtf8(buf []byte) bool {
	offset := 0
	for offset < len(buf) {
		r, size := utf8.DecodeRune(buf[offset:])
		if r == utf8.RuneError {
			return false
		}

		offset += size
	}

	return true
}

func ReadLines(filename String) []String {
	// file, err := os.Open(string(filename))
	// if err != nil {
	//     return nil, err
	// }

	// defer file.Close()

	// buf, err := ioutil.ReadAll(file)
	// if err != nil {
	//     return nil, err
	// }
	file := Open(filename.String())
	// file := Open(String())
	defer file.Close()

	buf := file.ReadAll()

	var text String

	switch {
	case bytes.Equal(buf[:3], BOM_UTF8):
		text = Decode(buf[3:], CP_UTF8)

	case bytes.Equal(buf[:2], BOM_UTF16_LE):
		text = Decode(buf[2:], CP_UTF16_LE)

	case bytes.Equal(buf[:2], BOM_UTF16_BE):
		text = Decode(buf[2:], CP_UTF16_BE)

	default:
		if isAnsiAsUtf8(buf) {
			text = Decode(buf, CP_UTF8)
		} else {
			text = Decode(buf, CP_GBK)
		}
	}

	return text.SplitLines()
}
