package chat

type commandId int

const (
	CMD_NICK commandId = 0
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
)

type command struct {
	id     commandId
	client *client
}
