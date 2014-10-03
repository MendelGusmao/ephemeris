package stubs

type Result int

const (
	ResultSuccess   Result = 1
	ResultNoRows    Result = 2
	ResultFoundMany Result = 3
	ResultError     Result = 4
)
