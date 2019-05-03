package dist

type HeaderImpl struct {
}

func (HeaderImpl) Magic() string {
	panic("implement me")
}

func (HeaderImpl) Version() int {
	panic("implement me")
}

func (HeaderImpl) Order() bool {
	panic("implement me")
}

func (HeaderImpl) Type() int {
	panic("implement me")
}

func (HeaderImpl) Size() int {
	panic("implement me")
}

type BodyImpl struct {
}

func (BodyImpl) RequestHeader() string {
	panic("implement me")
}

func (BodyImpl) RequestBody() string {
	panic("implement me")
}

func (BodyImpl) ReplyHeader() string {
	panic("implement me")
}

func (BodyImpl) ReplyBody() string {
	panic("implement me")
}

type MessageImpl struct {
	header HeaderImpl
	body   BodyImpl
}

func (MessageImpl) Header() Header {
	panic("implement me")
}

func (MessageImpl) Body() Body {
	panic("implement me")
}
