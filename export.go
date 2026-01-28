package connectip

type Http3Stream = http3Stream

func NewProxiedConn(str Http3Stream) *Conn {
	return newProxiedConn(str)
}
