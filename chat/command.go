package chat

type commandId int

const (
	CMD_NICK  commandId = 0
	CMD_JOIN  commandId = 1
	CMD_ROOMS commandId = 2
	CMD_MSG   commandId = 3
	CMD_QUIT  commandId = 4
)

type command struct {
	id     commandId
	client *client
	args   []string
}
