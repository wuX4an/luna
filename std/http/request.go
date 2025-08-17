package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	lua "github.com/yuin/gopher-lua"
)

// http.request(method, url, opts) -> res
func Request(L *lua.LState) int {
	method := L.CheckString(1)
	rawurl := L.CheckString(2)
	opts := L.OptTable(3, nil)

	// Optional parameters
	headers := map[string]string{}
	queryParams := map[string]string{}
	timeout := 5000 * time.Millisecond
	var bodyReader io.Reader = nil

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

		// body
		if b := opts.RawGetString("body"); b.Type() == lua.LTString {
			bodyReader = strings.NewReader(b.String())
		}
	}

	// Parse URL and add query parameters
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

	// Create request
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, method, u.String(), bodyReader)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(fmt.Sprintf("error creating request: %v", err)))
		return 2
	}

	// Set headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(fmt.Sprintf("request error: %v", err)))
		return 2
	}
	defer resp.Body.Close()

	// Read body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(fmt.Sprintf("error reading response: %v", err)))
		return 2
	}

	// Create result table
	resTbl := L.NewTable()
	resTbl.RawSetString("status", lua.LNumber(resp.StatusCode))
	resTbl.RawSetString("body", lua.LString(string(body)))
	L.Push(resTbl)
	return 1
}
