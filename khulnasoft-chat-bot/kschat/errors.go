package kschat

import (
	"errors"
	"fmt"
)

type ErrorCode int

var errAPIDisconnected = errors.New("chat API disconnected")

const (
	RevisionErrorCode          ErrorCode = 2760
	DeleteNonExistentErrorCode ErrorCode = 2762
)

// Error is for unmarshaling CLI json responses
type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("received error response from khulnasoft api: %s", e.Message)
}

type APIError struct {
	err error
}

func (e APIError) Error() string {
	return fmt.Sprintf("failed to call khulnasoft api: %v", e.err)
}

type UnmarshalError struct {
	err error
}

func (e UnmarshalError) Error() string {
	return fmt.Sprintf("failed to parse output from khulnasoft api: %v", e.err)
}
