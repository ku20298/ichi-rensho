package main

import (
	"golang.org/x/image/colornames"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"golang.org/x/image/font"
	
	"strconv"
	"math/rand"
	"log"
	"time"
)

const screenWidth = 360
const screenHeight = 640

var mainFont font.Face
var subFont font.Face
var rensyo int
var win bool
var winCount int
var loseCount int
var saikou int

func init() {
	rand.Seed(time.Now().UnixNano())

	mainFont = decodeFont(fontByte, 56)
	subFont = decodeFont(fontByte, 32)

	rensyo = 0
}

func update(screen *ebiten.Image) error {
	screen.Fill(colornames.White)

	jsEvent()
	
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || len(inpututil.JustPressedTouches()) > 0 {
		if rand.Intn(2) == 1 {
			rensyo ++
			win = true
			winCount ++
		}else {
			if rensyo >= saikou {
				saikou = rensyo
			}
			rensyo = 0
			win = false
			loseCount ++
		}
	}

	text.Draw(screen, "最高 " + strconv.Itoa(saikou), subFont, 134, 86, colornames.Black)
	
	text.Draw(screen, strconv.Itoa(rensyo), mainFont, 100, 200, colornames.Black)
	text.Draw(screen, "連勝", mainFont, 160, 200, colornames.Black)

	if winCount + loseCount == 0 {
		text.Draw(screen, "タップで勝負!", subFont, 82, 340, colornames.Black)
	}else {
		if win {
			text.Draw(screen, "勝ち", mainFont, 124, 340, colornames.Red)
		}else {
			text.Draw(screen, "負け", mainFont, 124, 340, colornames.Blue)
		}
	}
	
	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "1連勝"); err != nil {
		log.Fatal(err)
	}
}
