// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"os"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/samples/flags"
)

var gcn_window gxui.Window
var gcn_ce gxui.CodeEditor
var gcn_shortcut1 gxui.KeyboardEvent

var gce_Keydown_fn = func(evt gxui.KeyboardEvent) {
	if evt.Key == 100 {
		gcn_shortcut1.Key = 100
	}
	if gcn_shortcut1.Key == 100 && evt.Key == 24 {
		fmt.Println(evt.Key)
	}

}
var gce_Keyup_fn = func(evt gxui.KeyboardEvent) {
	if evt.Key == 100 {
		gcn_shortcut1.Key = 0
	}
}

func appFileProc(filename string) {
	// open input file
	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	//
	gcn_window.SetTitle(gcn_window.Title() + " >> " + fi.Name())

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		//fmt.Println(string(buf[:n]))

		gcn_ce.SetText(gcn_ce.Text() + string(buf[:n]))
		// write a chunk
		//		if _, err := fo.Write(buf[:n]); err != nil {
		//			panic(err)
		//		}
	}
}

func appMain(driver gxui.Driver) {
	var theme = flags.CreateTheme(driver)

	gcn_window = theme.CreateWindow(800, 600, "sldEditer")
	var layout = theme.CreateTableLayout()
	layout.SetGrid(1, 1)

	gcn_ce = theme.CreateCodeEditor()
	gcn_ce.SetTextColor(gxui.Green90)
	gcn_ce.SetMultiline(true)
	layout.SetChildAt(0, 0, 1, 1, gcn_ce)

	appFileProc("D:\\Dev\\go1.6\\src\\fmt\\print.go")
	gcn_window.OnKeyDown(func(evt gxui.KeyboardEvent) { gce_Keydown_fn(evt) })
	gcn_window.OnKeyUp(func(evt gxui.KeyboardEvent) { gce_Keyup_fn(evt) })
	gcn_window.AddChild(layout)

	gcn_window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}
