package audioutil

import (
	"math"
)

type Stream struct {
	position  int64
	remaining []byte
}

const (
	sampleRate   = 48000
	frequency    = 440
)


func (s *Stream) Read(buf []byte) (int, error) {
	// var origBuf []byte
	// if len(buf)%4 > 0 {
	// 	origBuf = buf
	// 	buf = make([]byte, len(origBuf)+4-len(origBuf)%4)
	// }

	const length = int64(sampleRate / frequency)
	p := s.position / 4
	for i := 0; i < len(buf)/4; i++ {
		const max = 32767
		b := int16(math.Sin(2*math.Pi*float64(p)/float64(length)) * max)
		buf[4*i] = byte(b)
		buf[4*i+1] = byte(b >> 8)
		buf[4*i+2] = byte(b)
		buf[4*i+3] = byte(b >> 8)
		p++
	}

	// s.position += int64(len(buf))
	// s.position %= length * 4

	// if origBuf != nil {
	// 	n := copy(origBuf, buf)
	// 	s.remaining = buf[n:]
	// 	return n, nil
	// }



	return len(buf), nil
}