/*
screenWidth 120
screenHeight 160
map 114
stone 6
*/
package main

import (
	"fmt"
	"image"
	"image/draw"
	"os"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/samples/flags"
	"github.com/nfnt/resize"
	_ "golang.org/x/image/bmp"
)

func appMain(driver gxui.Driver) {
	var theme = flags.CreateTheme(driver)

	////////board
	var fd_board, fd_err_board = os.Open("image/board.bmp")
	if fd_err_board != nil {
		fmt.Printf("Failed to open image : %v\n", fd_err_board)
		os.Exit(1)
	}

	var img_src_board, _, img_err_board = image.Decode(fd_board)
	if img_err_board != nil {
		fmt.Printf("Failed to read image : %v\n", img_err_board)
		os.Exit(1)
	}
	var img_src_board_resize = resize.Resize(456, 456, img_src_board, resize.Lanczos3)

	var img_board = theme.CreateImage()
	// Copy the image to a RGBA format before handing to a gxui.Texture
	var rgba_board = image.NewRGBA(img_src_board_resize.Bounds())
	draw.Draw(rgba_board, img_src_board_resize.Bounds(), img_src_board_resize, image.ZP, draw.Src)
	var texture_board = driver.CreateTexture(rgba_board, 1)
	img_board.SetTexture(texture_board)

	////////black
	var fd_black, fd_err_black = os.Open("image/black.bmp")
	if fd_err_black != nil {
		fmt.Printf("Failed to open image : %v\n", fd_err_black)
		os.Exit(1)
	}

	var img_src_black, _, img_err_black = image.Decode(fd_black)
	if img_err_black != nil {
		fmt.Printf("Failed to read image : %v\n", img_err_black)
		os.Exit(1)
	}
	var img_src_black_resize = resize.Resize(20, 20, img_src_black, resize.Lanczos3)

	var img_black = theme.CreateImage()
	// Copy the image to a RGBA format before handing to a gxui.Texture
	var rgba_black = image.NewRGBA(img_src_black_resize.Bounds())
	draw.Draw(rgba_black, img_src_black_resize.Bounds(), img_src_black_resize, image.ZP, draw.Src)
	var texture_black = driver.CreateTexture(rgba_black, 1)
	img_black.SetTexture(texture_black)

	///// window
	var window = theme.CreateWindow(img_src_board_resize.Bounds().Dx(), img_src_board_resize.Bounds().Dy(), "Sw Lee Developer")
	window.AddChild(img_board)
	window.AddChild(img_black)
	window.SetScale(flags.DefaultScaleFactor)
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}
