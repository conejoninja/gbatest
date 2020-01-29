package main

// Draw a red square on the GameBoy Advance screen.

import (
	"image/color"
	"machine"
	"runtime/interrupt"
	"runtime/volatile"
	"unsafe"

	"tinygo.org/x/tinydraw"

	"github.com/conejoninja/gbatest/fonts"
	"tinygo.org/x/tinyfont"
)

var display = machine.Display

var regDISPSTAT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000004)))

func main() {
	display.Configure()
	regDISPSTAT.SetBits(1<<3 | 1<<4)

	interrupt.New(interrupt.IRQ_VBLANK, update).Enable()
	//display.Display()
	for {
	}
}

var (
	krgb = uint8(0)
)

func update(interrupt.Interrupt) {

	tinyfont.DrawChar(display, &fonts.Bold24pt7b, 36, 60, byte('T'), getRainbowRGB(krgb))
	tinyfont.DrawChar(display, &fonts.Bold24pt7b, 66, 60, byte('i'), getRainbowRGB(1+krgb))
	tinyfont.DrawChar(display, &fonts.Bold24pt7b, 79, 60, byte('n'), getRainbowRGB(2+krgb))
	tinyfont.DrawChar(display, &fonts.Bold24pt7b, 108, 60, byte('y'), getRainbowRGB(3+krgb))
	tinyfont.DrawChar(display, &fonts.Bold24pt7b, 134, 60, byte('G'), getRainbowRGB(4+krgb))
	tinyfont.DrawChar(display, &fonts.Bold24pt7b, 170, 60, byte('o'), getRainbowRGB(5+krgb))

	tinyfont.WriteLine(display, &tinyfont.TomThumb, 70, 90, []byte("Go compiler for small places"), color.RGBA{10, 30, 30, 255})
	tinydraw.Triangle(display, 60, 110, 60, 126, 174, 110, color.RGBA{20, 20, 0, 255})
	tinydraw.Triangle(display, 72, 136, 184, 136, 184, 120, color.RGBA{20, 0, 20, 255})
	tinyfont.WriteLine(display, &fonts.Bold9pt7b, 68, 130, []byte("FOSDEM '20"), color.RGBA{10, 30, 30, 255})

	krgb++
	if krgb >= 30 {
		krgb = 0
	}
	if krgb%2 == 0 {
		tinyfont.DrawChar(display, &fonts.Regular58pt, 14, 140, byte('N'), color.RGBA{0, 0, 0, 255})
		tinyfont.DrawChar(display, &fonts.Regular58pt, 200, 140, byte('G'), color.RGBA{0, 0, 0, 255})
		tinyfont.DrawChar(display, &fonts.Regular58pt, 200, 140, byte('N'), getRainbowRGB(krgb))
		tinyfont.DrawChar(display, &fonts.Regular58pt, 14, 140, byte('G'), getRainbowRGB(krgb))
	} else {
		tinyfont.DrawChar(display, &fonts.Regular58pt, 200, 140, byte('N'), color.RGBA{0, 0, 0, 255})
		tinyfont.DrawChar(display, &fonts.Regular58pt, 14, 140, byte('G'), color.RGBA{0, 0, 0, 255})
		tinyfont.DrawChar(display, &fonts.Regular58pt, 14, 140, byte('N'), getRainbowRGB(krgb))
		tinyfont.DrawChar(display, &fonts.Regular58pt, 200, 140, byte('G'), getRainbowRGB(krgb))
	}
}

func getRainbowRGB(i uint8) color.RGBA {
	if i < 10 {
		return color.RGBA{i * 3, 30 - i*3, 0, 255}
	} else if i < 20 {
		i -= 10
		return color.RGBA{30 - i*3, 0, i * 3, 255}
	}
	i -= 20
	return color.RGBA{0, i * 3, 30 - i*3, 255}
}
