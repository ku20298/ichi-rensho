package main

import (
	"golang.org/x/image/colornames"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"golang.org/x/image/font"
	"github.com/ku20298/localstorage"
	
	"strconv"
	"math/rand"
	"log"
	"time"
)

const screenWidth = 360
const screenHeight = 640

var mainFont font.Face
var subFont font.Face
var rensho int
var win bool
var winCount int
var loseCount int
var saikou int

func init() {
	rand.Seed(time.Now().UnixNano())

	if localstorage.GetItem("saikou") == "null" {
		localstorage.SetItem("saikou", 0)
	}

	var err error
	saikou, err = strconv.Atoi(localstorage.GetItem("saikou"))
	if err != nil {
		log.Fatal(err)
	}

	mainFont = decodeFont(fontByte, 56)
	subFont = decodeFont(fontByte, 32)

	rensho = 0
}

func update(screen *ebiten.Image) error {
	screen.Fill(colornames.White)

	jsEvent()
	
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || len(inpututil.JustPressedTouchIDs()) > 0 {
		if rand.Intn(2) == 1 {
			rensho ++
			win = true
			winCount ++
		}else {
			if rensho >= saikou {
				saikou = rensho
				localstorage.SetItem("saikou", saikou)
			}
			rensho = 0
			win = false
			loseCount ++
		}
	}

	text.Draw(screen, "最高 " + strconv.Itoa(saikou), subFont, 134, 86, colornames.Black)
	
	text.Draw(screen, strconv.Itoa(rensho), mainFont, 100, 200, colornames.Black)
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
