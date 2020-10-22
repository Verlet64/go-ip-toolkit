package ipv4

import (
	"errors"
	"strconv"
	"strings"
)

type IP struct {
	raw [4]byte
}

var (
	ipFormatError       = errors.New("invalid IPv4 address provided to constructor")
	ipSegmentUint8Error = errors.New("IPv4 segments must be unsigned 8-bit integer values (i.e between 0 and 255 inclusive)")
)

func NewIPFromString(addr string) (*IP, error) {
	segments := strings.Split(addr, ".")
	buffer := [4]byte{}

	if len(segments) != 4 {
		return nil, ipFormatError
	}

	for i, segment := range segments {
		val, err := strconv.Atoi(segment)
		if err != nil {
			return nil, err
		}

		// To prevent a panic, we'll verify that the
		// value is indeed a uint8 before trying to
		// cast it to a uint8
		if val < 0 || val > 255 {
			return nil, ipSegmentUint8Error
		}

		copy(buffer[i:i+1], []uint8{uint8(val)})
	}

	return &IP{
		raw: buffer,
	}, nil
}

func (i *IP) String() string {
	ip := ""

	for _, segment := range i.raw {
		if ip != "" {
			ip += "."
		}

		ip += strconv.Itoa(int(segment))
	}

	return ip
}
