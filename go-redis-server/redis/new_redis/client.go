package new_redis

import "strings"

type Client struct {
	requestBuf []byte

	argv        []string
	argc        int
	responseBuf []byte

	handler *CommandHandler
}

func NewClient(req []byte) *Client{
	return &Client{requestBuf: req}
}

func (c *Client) ParseRequest() *Client {
	inputStr := string(c.requestBuf)
	c.argv = strings.Split(strings.Trim(inputStr, "\r\n"), " ")
	c.argc = len(c.argv)
	c.handler = CommandTable[c.argv[0]]
	return c
}

func (c *Client) Exec() *Client {
	proc := c.handler.proc
	proc(c)
	return c
}

func (c *Client) WriteResponse()  {

}