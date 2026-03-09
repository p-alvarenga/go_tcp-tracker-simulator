package gt06

import "bytes"

func ExtractFrame(buf []byte) ([]byte, []byte, bool) {
	start := bytes.Index(buf, startBytes[:])

	if start == -1 {
		if len(buf) > 1 {
			buf = buf[len(buf)-1:]
		}
		return nil, buf[:len(buf):len(buf)], false
	}

	end := bytes.Index(buf[start+2:], stopBytes[:])
	if end == -1 {
		return nil, buf, false
	}

	end = start + end + 4
	frame := buf[start:end]
	next := buf[end:]

	return frame, next, true
}
