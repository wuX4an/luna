package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	lua "github.com/yuin/gopher-lua"
)

// http.request(method, url, opts) -> res
func Request(L *lua.LState) int {
	method := L.CheckString(1)
	rawurl := L.CheckString(2)
	opts := L.OptTable(3, nil)

	// Parámetros opcionales
	headers := map[string]string{}
	queryParams := map[string]string{}
	timeout := 5000 * time.Millisecond

	if opts != nil {
		// headers
		if tbl := opts.RawGetString("headers"); tbl.Type() == lua.LTTable {
			tblHeaders := tbl.(*lua.LTable)
			tblHeaders.ForEach(func(k, v lua.LValue) {
				headers[k.String()] = v.String()
			})
		}
		// query
		if tbl := opts.RawGetString("query"); tbl.Type() == lua.LTTable {
			tblQuery := tbl.(*lua.LTable)
			tblQuery.ForEach(func(k, v lua.LValue) {
				queryParams[k.String()] = v.String()
			})
		}
		// timeout
		if to := opts.RawGetString("timeout"); to.Type() == lua.LTNumber {
			ms := time.Duration(lua.LVAsNumber(to))
			timeout = ms * time.Millisecond
		}
	}

	// Parsear URL y añadir query params
	u, err := url.Parse(rawurl)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(fmt.Sprintf("invalid url: %v", err)))
		return 2
	}

	q := u.Query()
	for k, v := range queryParams {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	// Crear request
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, method, u.String(), nil)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(fmt.Sprintf("error creating request: %v", err)))
		return 2
	}

	// headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// Hacer request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(fmt.Sprintf("request error: %v", err)))
		return 2
	}
	defer resp.Body.Close()

	// Leer body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(fmt.Sprintf("error reading response: %v", err)))
		return 2
	}

	// Crear tabla resultado
	resTbl := L.NewTable()
	resTbl.RawSetString("status", lua.LNumber(resp.StatusCode))
	resTbl.RawSetString("body", lua.LString(string(body)))

	L.Push(resTbl)
	return 1
}
