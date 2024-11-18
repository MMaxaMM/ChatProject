package models

import "io"

type Object struct {
	Payload     io.Reader
	PayloadSize int64
}
