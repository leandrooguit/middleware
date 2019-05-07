package dist

type HeaderImpl struct {
	magic      string
	version    int
	order      bool
	headerType int
	size       int
}

func (h HeaderImpl) Magic() string {
	return h.magic
}

func (h HeaderImpl) Version() int {
	return h.version
}

func (h HeaderImpl) Order() bool {
	return h.order
}

func (h HeaderImpl) Type() int {
	return h.headerType
}

func (h HeaderImpl) Size() int {
	return h.size
}

type BodyImpl struct {
	requestHeader string
	requestBody   string
	replyHeader   string
	replyBody     interface{}
}

func (b BodyImpl) RequestHeader() string {
	return b.requestHeader
}

func (b BodyImpl) RequestBody() string {
	return b.requestBody
}

func (b BodyImpl) ReplyHeader() string {
	return b.replyHeader
}

func (b BodyImpl) ReplyBody() interface{} {
	return b.replyBody
}

type MessageImpl struct {
	header Header
	body   Body
}

func (m MessageImpl) Body() Body {
	return m.body
}

func (m MessageImpl) Header() Header {
	return m.header
}
