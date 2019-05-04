package dist

type Header interface {
	Magic() string
	Version() int
	Order() bool
	Type() int
	Size() int
}

type Body interface {
	RequestHeader() string
	RequestBody() string
	ReplyHeader() string
	ReplyBody() interface{}
}

type Message interface {
	Header() Header
	Body() Body
}
