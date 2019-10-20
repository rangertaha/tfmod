/* The MIT License (MIT)

Copyright © 2019 rangertaha@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE. */

package pkg

import (
	"github.com/kataras/iris"
	"github.com/spf13/viper"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"

	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/swaggerFiles"

	_ "github.com/rangertaha/tfmod/docs"
)

// Success: Return status is 200 on success with a body or 204 if there is no body data to return.
// Redirects: Moved or aliased endpoints redirect with a 301. Endpoints redirecting to the latest version of a module may redirect with 302 or 307 to indicate that they logically point to different resources over time.
// Client Errors: Invalid requests will receive the relevant 4xx status. Except where noted below, the request should not be retried.
// Rate Limiting: Clients placing excessive load on the service might be rate-limited and receive a 429 code. This should be interpreted as a sign to slow down, and wait some time before retrying the request.
// Service Errors: The usual 5xx errors will be returned for service failures. In all cases it is safe to retry the request after receiving a 5xx response.
// Load Shedding: A 503 response indicates that the service is under load and can't process your request immediately. As with other 5xx errors you may retry after some delay, although clients should consider being more lenient with retry schedule in this case.

type (
	// Request ...
	Request struct{}

	// Response ...
	Response struct {
		Meta    Meta
		Modules []Module
		Errors  []string `json:"errors"`
	}

	// Meta ...
	Meta struct {
		Limit         int    `json:"limit"`
		CurrentOffset int    `json:"current_offset"`
		NextOffset    int    `json:"next_offset"`
		NextURL       string `json:"next_url"`
	}

	// Module ...
	Module struct {
		ID          string `json:"id"`
		Owner       string `json:"owner"`
		Namespace   string `json:"namespace"`
		Name        string `json:"name"`
		Version     string `json:"version"`
		Provider    string `json:"provider"`
		Description string `json:"description"`
		Source      string `json:"source"`
		Published   string `json:"published_at"`
		Downloads   int    `json:"downloads"`
		Verified    bool   `json:"verified"`
		Root        Root

		Versions []Version `json:"versions"`
	}

	// Version ...
	Version struct {
		Version    string      `json:"version"`
		Submodules []Submodule `json:"submodules"`
		Root       struct {
			Dependencies []string   `json:"dependencies"`
			Providers    []Provider `json:"providers"`
		}
	}

	// Submodule ..
	Submodule struct {
		Path      string     `json:"path"`
		Providers []Provider `json:"providers"`
	}

	// Provider ...
	Provider struct {
		Name    string `json:"name"`
		Version string `json:"verstion"`
	}

	// Input ...
	Input struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Default     string `json:"default"`
	}

	// Output ...
	Output struct {
		Name        string `json:"name"`
		Description string `json:"descripiton"`
	}

	// Root ...
	Root struct {
		Path         string     `json:"path"`
		Readme       string     `json:"readme"`
		Empty        bool       `json:"empty"`
		Inputs       []Input    `json:"inputs"`
		Outputs      []Output   `json:"outputs"`
		Dependencies []string   `json:"dependencies"`
		Resources    []Resource `json:"resources"`
	}
	// Resource ...
	Resource struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}
)

// Server ...
func Server() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// Method:   GET
	// Resource: http://localhost:8080
	config := &swagger.Config{
		URL: "http://127.0.0.1:8080/doc.json",
	}
	// use swagger middleware to
	// app.Get("/", swagger.CustomWrapHandler(config, swaggerFiles.Handler))

	app.Get("/{any:path}", swagger.CustomWrapHandler(config, swaggerFiles.Handler))

	// // Method:   GET
	// // Resource: http://localhost:8080
	// app.Handle("GET", "/", func(ctx iris.Context) {
	// 	ctx.HTML("<h1>Welcome</h1>")
	// })

	// // same as app.Handle("GET", "/ping", [...])
	// // Method:   GET
	// // Resource: http://localhost:8080/ping
	// app.Get("/ping", func(ctx iris.Context) {
	// 	ctx.WriteString("pong")
	// })

	// // Method:   GET
	// // Resource: http://localhost:8080/hello
	// app.Get("/hello", func(ctx iris.Context) {
	// 	ctx.JSON(iris.Map{"message": "Hello Iris!"})
	// })

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello

	// var port = viper.GetString("http.port")
	var host = viper.GetString("http.host")

	app.Run(iris.Addr(host+":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
