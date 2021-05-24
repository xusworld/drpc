package protocol

type Response struct {
	// 返回码
	RetCode uint32

	// 返回消息
	RetMsg string

	// 透传数据
	Attachment map[string]string

	// 返回体
	Payload []byte
}
