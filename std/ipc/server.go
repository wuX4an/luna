package ipc

import (
	"fmt"
	"net"
	"os"
	"sync"

	lua "github.com/yuin/gopher-lua"
)

// Server userdata
type LServer struct {
	path     string
	listener net.Listener
	closed   bool
	mu       sync.Mutex
}

// Constructor Lua: ipc.server{ path = "/tmp/misocket.sock" }
func Server(L *lua.LState) int {
	tbl := L.CheckTable(1)
	path := tbl.RawGetString("path")
	if path.Type() != lua.LTString {
		L.ArgError(1, "'path' string expected")
		return 0
	}

	s := &LServer{
		path: path.String(),
	}

	ud := L.NewUserData()
	ud.Value = s

	mt := L.NewTypeMetatable("ipc_server")
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), serverMethods))

	L.SetMetatable(ud, mt)
	L.Push(ud)
	return 1
}

var serverMethods = map[string]lua.LGFunction{
	"start": serverStart,
	"close": serverClose,
}

func (s *LServer) acceptLoop(L *lua.LState, handler lua.LValue) {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			// Listener closed or error
			break
		}
		// Cada conexión en goroutine para no bloquear
		go func(c net.Conn) {
			L2, _ := L.NewThread()
			connUD := newConnection(L2, c)

			// Ejecutar handler(conn) en nueva coroutine
			L2.Push(handler)
			L2.Push(connUD)
			if err := L2.PCall(1, lua.MultRet, nil); err != nil {
				fmt.Println("Error en handler:", err)
			}
		}(conn)
	}
}

func newConnection(L *lua.LState, conn net.Conn) *lua.LUserData {
	ud := L.NewUserData()
	ud.Value = conn

	mt := L.NewTypeMetatable("ipc_connection")
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), connectionMethods))
	L.SetMetatable(ud, mt)

	return ud
}

var connectionMethods = map[string]lua.LGFunction{
	"send":  connectionSend,
	"recv":  connectionRecv,
	"close": connectionClose,
}

func checkConn(L *lua.LState) net.Conn {
	ud := L.CheckUserData(1)
	conn, ok := ud.Value.(net.Conn)
	if !ok {
		L.ArgError(1, "ipc_connection expected")
	}
	return conn
}

func connectionSend(L *lua.LState) int {
	conn := checkConn(L)
	data := L.CheckString(2)
	_, err := conn.Write([]byte(data))
	if err != nil {
		L.RaiseError("failed to send: %v", err)
	}
	return 0
}

func connectionRecv(L *lua.LState) int {
	conn := checkConn(L)
	maxlen := 1024
	if L.GetTop() >= 2 {
		maxlen = L.CheckInt(2)
	}
	buf := make([]byte, maxlen)
	n, err := conn.Read(buf)
	if err != nil {
		L.Push(lua.LNil)
		return 1
	}
	L.Push(lua.LString(buf[:n]))
	return 1
}

func connectionClose(L *lua.LState) int {
	conn := checkConn(L)
	conn.Close()
	return 0
}

func serverStart(L *lua.LState) int {
	ud := L.CheckUserData(1)
	s, ok := ud.Value.(*LServer)
	if !ok {
		L.ArgError(1, "ipc_server expected")
		return 0
	}

	handler := L.CheckAny(2)
	if handler.Type() != lua.LTFunction {
		L.ArgError(2, "function expected")
		return 0
	}

	// Eliminar socket si existe
	if _, err := os.Stat(s.path); err == nil {
		os.Remove(s.path)
	}

	l, err := net.Listen("unix", s.path)
	if err != nil {
		L.RaiseError("error listening on unix socket: %v", err)
		return 0
	}

	s.listener = l

	go s.acceptLoop(L, handler)
	select {}
}

func serverClose(L *lua.LState) int {
	ud := L.CheckUserData(1)
	s, ok := ud.Value.(*LServer)
	if !ok {
		L.ArgError(1, "ipc_server expected")
		return 0
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.closed {
		s.listener.Close()
		os.Remove(s.path)
		s.closed = true
	}

	return 0
}
