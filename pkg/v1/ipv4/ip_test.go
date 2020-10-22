package ipv4_test

import (
	"errors"
	"go-ip-toolkit/pkg/v1/ipv4"
	"reflect"
	"testing"
)

func TestNewIPFromString(t *testing.T) {
	input := "192.168.0.1"
	_, err := ipv4.NewIPFromString(input)
	if err != nil {
		t.Fatalf("Expected IP struct to be constructed without error for valid IPv4 address: %v, encountered error: %v", input, err)
	}
}

func TestNewIPFromString_InvalidIPV4AddressFormat(t *testing.T) {
	input := "192.168,0.1"
	_, err := ipv4.NewIPFromString(input)
	expected := errors.New("invalid IPv4 address provided to constructor")
	if !reflect.DeepEqual(expected, err) {
		t.Fatalf("Expected error to be %v, received %v", expected, err)
	}
}

func TestNewIPFromString_OutOfBoundsValue(t *testing.T) {
	input := "257.168.0.1"
	_, err := ipv4.NewIPFromString(input)
	expected := errors.New("IPv4 segments must be unsigned 8-bit integer values (i.e between 0 and 255 inclusive)")
	if !reflect.DeepEqual(expected, err) {
		t.Fatalf("Expected error to be %v, received %v", expected, err)
	}
}

func TestString(t *testing.T) {
	input := "192.168.0.1"
	ip, _ := ipv4.NewIPFromString(input)
	expected := input
	got := ip.String()

	if expected != got {
		t.Fatalf("Expected IP address to be %v, found %v", expected, got)
	}
}
