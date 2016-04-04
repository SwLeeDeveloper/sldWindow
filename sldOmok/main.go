/*
Middle Center : image.Point{-263,-263}
Map Size : 526, 526
Stone : 28, 28
Cell Size : 28, 28
Array : 19,19
*/
package main

import (
	"fmt"
	"image"
	"image/draw"
	"os"

	"github.com/google/gxui"
	_ "github.com/google/gxui/math"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/samples/flags"	
	"github.com/nfnt/resize"
	_ "golang.org/x/image/bmp"
)


//임시 
//Opponents : 적
//Place to attack
/* board_place '0' 없는상태 , '1' 흑돌, '2' 힌돌 그외 미존재. */
type GameObj struct {	
	my_nm           string
	opp_nm          string
	attck_stone     int8
	board_place     [19][19]int8
}



func getFileImage(fullfilename string) image.Image {	

	var lv_fd, lv_fd_err = os.Open(fullfilename)
	if lv_fd_err != nil {
		fmt.Printf("Failed to open image : %v\n", lv_fd_err)
		os.Exit(1)
	}
	
	var lv_img, _, lv_img_err = image.Decode(lv_fd)
	if lv_img_err != nil {
		fmt.Printf("Failed to read image : %v\n", lv_img_err)
		os.Exit(1)
	}
	
	return lv_img
}

func gameing() {
	var obj = GameObj{}
	obj.my_nm = "SwLee"
	
	obj.board_place = [19][19]int8{{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
						{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
						{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
						{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
						{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
						{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
						{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
						{ 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
						{ 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
						{ 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
						{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
						{ 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
						{ 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
						{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
						{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
						{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
						{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
						{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
						{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 }	}	
}

func appMain(driver gxui.Driver) {	
	var theme = flags.CreateTheme(driver)
	
	//board : 이대로 두어야 함.
	var img_src_board_resize = resize.Resize(526, 526, getFileImage("image/board.bmp"), resize.MitchellNetravali)
	var img_board = theme.CreateImage()
	// Copy the image to a RGBA format before handing to a gxui.Texture
	var rgba_board = image.NewRGBA(img_src_board_resize.Bounds())
	draw.Draw(rgba_board, img_src_board_resize.Bounds(), img_src_board_resize, image.ZP, draw.Src)
	var texture_board = driver.CreateTexture(rgba_board, 1)
	img_board.SetTexture(texture_board)

	////////black
	var img_src_black_resize = resize.Resize(28, 28, getFileImage("image/black.png"), resize.MitchellNetravali)

	draw.Draw(rgba_board, img_src_board_resize.Bounds(), img_src_black_resize, image.Point{-235,-235}, draw.Src)
	var texture_black = driver.CreateTexture(rgba_board, 1)
	img_board.SetTexture(texture_black)
	
	////////black
	//var img_src_black_resize = resize.Resize(28, 28, getFileImage("image/black.png"), resize.MitchellNetravali)
	draw.Draw(rgba_board, img_src_board_resize.Bounds(), img_src_black_resize, image.Point{-263,-263}, draw.Src)
	var texture_black1 = driver.CreateTexture(rgba_board, 1)
	img_board.SetTexture(texture_black1)
	
	gameing()
	
	
	///// window
	var window = theme.CreateWindow(img_src_board_resize.Bounds().Dx(), img_src_board_resize.Bounds().Dy(), "Sw Lee Developer")
	window.AddChild(img_board)
	//window.AddChild(img_black)
	window.SetScale(flags.DefaultScaleFactor)
	window.OnClose(driver.Terminate)
	
}

func main() {
	gl.StartDriver(appMain)
}
