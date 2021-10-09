package net

// APIResponseMetadata defines the metadata of a web response.
type APIResponseMetadata struct {
  Error   bool                `json:"error"`
  Code    APIErrorCode        `json:"code"`
  Message string              `json:"message"`
  Display bool                `json:"display"`
}

// APIResponse defines a web response structure.
type APIResponse struct {
  Success  bool                `json:"success"`
  Payload  interface{}         `json:"payload"`
  Metadata APIResponseMetadata `json:"metadata"`
}
