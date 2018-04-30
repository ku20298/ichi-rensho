package main

import (
	"strconv"
	"golang.org/x/image/colornames"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"golang.org/x/image/font"
	
	"math/rand"
	"log"
	"time"
)

const screenWidth = 360
const screenHeight = 640

var mainFont font.Face
var rensyo int
var renpai int
var win bool
var lose bool
var winCount float64
var loseCount float64
var syoritsu string

func init() {
	rand.Seed(time.Now().UnixNano())

	mainFont = decodeFont(fontByte, 56)

	rensyo = 0
	renpai = 0
	win = false
	lose = false
}

func update(screen *ebiten.Image) error {
	jsEvent()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || len(inpututil.JustPressedTouches()) > 0 {
		if rand.Intn(2) == 1 {
			rensyo ++
			lose = false			
			win = true
			winCount ++
		}else {
			rensyo = 0
			win = false
			lose = true
			loseCount ++
		}
	}

	screen.Fill(colornames.White)

	text.Draw(screen, strconv.Itoa(rensyo), mainFont, 100, 200, colornames.Black)
	text.Draw(screen, "連勝", mainFont, 160, 200, colornames.Black)
	if win {
		for i := 0; i < rensyo; i++ {
			text.Draw(screen, "勝ち", mainFont, 124, 340, colornames.Red)
		}
	}
	if lose {
		text.Draw(screen, "負け", mainFont, 124, 340, colornames.Blue)
	}

	if winCount + loseCount > 0 {
		syoritsu = "勝率:" + strconv.Itoa(int(winCount / (winCount + loseCount) * 100)) + "%"
	} 
	text.Draw(screen, syoritsu, mainFont, 70, 500, colornames.Black)
	
	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "一連勝"); err != nil {
		log.Fatal(err)
	}
}
