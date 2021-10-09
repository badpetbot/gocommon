package net

import (
  "fmt"
)

type APIErrorCode int

type APIError struct {
  HTTPCode        int
  ErrorCode       APIErrorCode
  ResponseMessage string
  LogMessage      string
  CauseError      error
}

func (this *APIError) Error() string {
  return fmt.Sprintf("HTTP %d: %s", this.HTTPCode, this.LogMessage)
}

// 0-99     | Generic problems.
var APIErrorNone                APIErrorCode = 0
var APIErrorNotYetImplemented   APIErrorCode = 1
var APIErrorAreYouSure          APIErrorCode = 2
var APIErrorNotFound            APIErrorCode = 3
var APIErrorCantParseBody       APIErrorCode = 4
var APIErrorInvalidAction       APIErrorCode = 5
var APIErrorTooSoon             APIErrorCode = 6

// 100-199  | Validation problems.

// 200-299  | Account problems.

// 300-399  | Permissions problems.

// 400-499  | Input problems.

// 500-599  | Server problems.
var APIErrorDatabase            APIErrorCode = 500
var APIErrorExternalCall        APIErrorCode = 501


// 1000     | Unknown
var APIErrorUnknown             APIErrorCode = 1000