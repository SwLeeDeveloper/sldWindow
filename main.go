/***********************************************
** sldWindow 단순한 라벨이랑 text 버튼입니다
** By SwLee
***********************************************/

package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/math"
	"github.com/google/gxui/samples/flags"
)

func appMain(driver gxui.Driver) {

	var theme = flags.CreateTheme(driver)

	//임시.
	var pen1 = gxui.CreatePen(1, gxui.Red90)
	var pen2 = gxui.CreatePen(1, gxui.Blue60)

	//layout := theme.CreateLinearLayout()
	//var layout gxui.LinearLayout
	var layout = theme.CreateTableLayout()
	layout.SetGrid(10, 10)

	///// row 1
	var label1_1 = theme.CreateLabel()
	label1_1.SetText("IP : ")
	label1_1.SetColor(gxui.White)
	label1_1.SetMargin(math.Spacing{L: 10, R: 10, T: 10, B: 10})
	label1_1.SetHorizontalAlignment(gxui.AlignRight)
	// column, row, horizontal span, vertical span
	layout.SetChildAt(0, 0, 1, 1, label1_1)

	var txt1_1 = theme.CreateTextBox()
	txt1_1.SetTextColor(gxui.Red60)
	txt1_1.SetText("127.0.0.1")
	layout.SetChildAt(1, 0, 1, 1, txt1_1)

	var label1_2 = theme.CreateLabel()
	label1_2.SetText("PORT : ")
	label1_2.SetColor(gxui.White)
	label1_2.SetHorizontalAlignment(gxui.AlignRight)
	// column, row, horizontal span, vertical span
	layout.SetChildAt(2, 0, 1, 1, label1_2)

	var txt1_2 = theme.CreateTextBox()
	txt1_2.SetTextColor(gxui.Red60)
	txt1_2.SetText("99999")
	layout.SetChildAt(3, 0, 1, 1, txt1_2)

	///// row 2
	var label2_1 = theme.CreateLabel()
	label2_1.SetText("Data : ")
	label2_1.SetColor(gxui.White)
	label2_1.SetHorizontalAlignment(gxui.AlignRight)
	// column, row, horizontal span, vertical span
	layout.SetChildAt(0, 1, 1, 1, label2_1)

	var txt2_1 = theme.CreateTextBox()
	txt2_1.SetTextColor(gxui.Red60)
	txt2_1.SetText("Txt2")
	layout.SetChildAt(1, 1, 6, 1, txt2_1)

	var label2_2 = theme.CreateLabel()
	label2_2.SetText("LEN : ")
	label2_2.SetColor(gxui.White)
	label2_2.SetHorizontalAlignment(gxui.AlignRight)
	// column, row, horizontal span, vertical span
	layout.SetChildAt(7, 1, 1, 1, label2_2)

	var txt2_2 = theme.CreateTextBox()
	txt2_2.SetTextColor(gxui.Red60)
	txt2_2.SetText("Txt2")
	layout.SetChildAt(8, 1, 1, 1, txt2_2)

	var btn2 = theme.CreateButton()
	btn2.SetBorderPen(pen2)
	btn2.SetText("Send2")
	layout.SetChildAt(9, 1, 1, 1, btn2)

	///// window
	var window = theme.CreateWindow(700, 300, "Sw Lee Developer")
	window.SetScale(flags.DefaultScaleFactor)

	window.SetBorderPen(pen1)
	window.AddChild(layout)

	window.OnClose(driver.Terminate)

}

func main() {
	gl.StartDriver(appMain)
}
