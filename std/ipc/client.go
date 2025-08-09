package ipc

import (
	"net"
	"time"

	lua "github.com/yuin/gopher-lua"
)

type LClient struct {
	path string
	conn net.Conn
}

func Client(L *lua.LState) int {
	tbl := L.CheckTable(1)
	path := tbl.RawGetString("path")
	if path.Type() != lua.LTString {
		L.ArgError(1, "'path' string expected")
		return 0
	}

	c := &LClient{
		path: path.String(),
	}

	ud := L.NewUserData()
	ud.Value = c

	mt := L.NewTypeMetatable("ipc_client")
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), clientMethods))

	L.SetMetatable(ud, mt)
	L.Push(ud)
	return 1
}

var clientMethods = map[string]lua.LGFunction{
	"connect": clientConnect,
	"send":    clientSend,
	"recv":    clientRecv,
	"close":   clientClose,
}

func (c *LClient) checkConn(L *lua.LState) {
	if c.conn == nil {
		L.RaiseError("not connected")
	}
}

func clientConnect(L *lua.LState) int {
	ud := L.CheckUserData(1)
	c, ok := ud.Value.(*LClient)
	if !ok {
		L.ArgError(1, "ipc_client expected")
		return 0
	}

	conn, err := net.DialTimeout("unix", c.path, 5*time.Second)
	if err != nil {
		L.RaiseError("failed to connect: %v", err)
		return 0
	}

	c.conn = conn
	return 0
}

func clientSend(L *lua.LState) int {
	ud := L.CheckUserData(1)
	c, ok := ud.Value.(*LClient)
	if !ok {
		L.ArgError(1, "ipc_client expected")
		return 0
	}
	c.checkConn(L)

	data := L.CheckString(2)
	_, err := c.conn.Write([]byte(data))
	if err != nil {
		L.RaiseError("failed to send: %v", err)
		return 0
	}
	return 0
}

func clientRecv(L *lua.LState) int {
	ud := L.CheckUserData(1)
	c, ok := ud.Value.(*LClient)
	if !ok {
		L.ArgError(1, "ipc_client expected")
		return 0
	}
	c.checkConn(L)

	maxlen := 1024
	if L.GetTop() >= 2 {
		maxlen = L.CheckInt(2)
	}

	buf := make([]byte, maxlen)
	n, err := c.conn.Read(buf)
	if err != nil {
		L.RaiseError("failed to receive: %v", err)
		return 0
	}

	L.Push(lua.LString(buf[:n]))
	return 1
}

func clientClose(L *lua.LState) int {
	ud := L.CheckUserData(1)
	c, ok := ud.Value.(*LClient)
	if !ok {
		L.ArgError(1, "ipc_client expected")
		return 0
	}
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}
	return 0
}
