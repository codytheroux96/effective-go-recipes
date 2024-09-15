package io

import (
	"bytes"
	"encoding/json"
	"io"
	"time"
)

type Ride struct {
	ID       int
	Time     time.Time
	Duration time.Duration
	Distance float64
	Price    float64
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{json.NewDecoder(r)}
}

func UnmarshalRide(data []byte, ride *Ride) error {
	r := bytes.NewReader(data)
	return NewDecoder(r).DecodeRide(ride)
}