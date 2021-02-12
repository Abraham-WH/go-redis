package new_redis

var CommandTable = map[string]*CommandHandler{}

type CommandProc = func(client *Client)

// 可以定义成接口
type CommandHandler struct {
	commandName string
	proc        CommandProc
}

type SetHandler struct {
	commandName string
	proc CommandHandler
}

