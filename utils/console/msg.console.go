package console

type ConsoleMsgStruct struct {
	Msg    string
	MsgErr string
}

var (
	ConsoleMsg = ConsoleMsgStruct{
		Msg:    "packages -> Product",
		MsgErr: "packages -> Product -> error",
	}
	ConsoleRedis = ConsoleMsgStruct{
		Msg:    "cache -> redis",
		MsgErr: "cache -> redis -> error",
	}
)
