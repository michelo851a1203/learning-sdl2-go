package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Renderer struct {
	r *sdl.Renderer
}

func (r *Renderer) Begin(
	colorHex string,
) error {
	color, _ := HexToRGB(colorHex)
	err := r.r.SetDrawColor(color.R, color.G, color.B, color.A)
	if err != nil {
		return err
	}
	err = r.r.Clear()
	if err != nil {
		return err
	}
	return nil
}

func (r *Renderer) End() error {
	r.r.Present()
	return nil
}

func (r *Renderer) DrawRect(
	pos Vector,
	w int32,
	h int32,
	colorHex string,
) error {
	color, _ := HexToRGB(colorHex)
	r.r.SetDrawColor(color.R, color.G, color.B, color.A)
	rect := sdl.Rect{
		X: pos.X,
		Y: pos.Y,
		W: w,
		H: h,
	}
	r.r.FillRect(&rect)

	return nil
}

type Vector struct {
	X int32
	Y int32
}

type Color struct {
	R, G, B, A uint8
}

type Entity interface {
	Render(chart *Chart) error
	Update() error
}

type Panel interface {
	Entity
}

type Chart struct {
	Width     int32
	Height    int32
	Window    *sdl.Window
	Renderer  *Renderer
	IsRunning bool

	Panels []Panel
}
type PriceBar struct {
	Width    int32
	Height   int32
	Position Vector
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func NewPriceBar() (*PriceBar, error) {
	return &PriceBar{
		Width:  60,
		Height: 0,
		Position: Vector{
			X: 10,
			Y: 10,
		},
	}, nil
}

func (p *PriceBar) Render(chart *Chart) error {
	position := Vector{
		X: chart.Width - p.Width,
		Y: 0,
	}
	chart.Renderer.DrawRect(
		position,
		p.Width,
		chart.Height,
		"#ffffff",
	)

	return nil
}

func (p *PriceBar) Update() error {
	return nil
}

func NewChart(
	width,
	height int32,
	title string,
) (*Chart, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	window, err := sdl.CreateWindow(
		title,
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		int32(width),
		int32(height),
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		return nil, err
	}
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	chart := &Chart{
		Window:    window,
		Renderer:  &Renderer{r: renderer},
		Width:     width,
		Height:    height,
		IsRunning: true,
		Panels:    []Panel{},
	}

	return chart, nil
}

func (c *Chart) AddPanel(p Panel) error {
	c.Panels = append(c.Panels, p)
	return nil
}

func (c *Chart) MainLoop() error {
	for c.IsRunning {
		c.ProcessEvent()
		c.Update()
		c.Render()
		sdl.Delay(16)
	}
	return nil
}

func (c *Chart) Render() error {
	c.Renderer.Begin("#000000")

	for _, panel := range c.Panels {
		panel.Render(c)
	}
	c.Renderer.End()
	return nil
}

func (c *Chart) Update() error {
	for _, panel := range c.Panels {
		panel.Update()
	}
	return nil
}

func (c *Chart) ProcessEvent() error {
LOOP:
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch ev := event.(type) {
		case *sdl.KeyboardEvent:
			switch ev.Keysym.Sym {
			case sdl.K_ESCAPE:
				c.IsRunning = false
				break LOOP
			}
		case *sdl.QuitEvent:
			fmt.Println("Quit")
			c.IsRunning = false
			break LOOP
		}
	}
	return nil
}
func main() {
	chart, err := NewChart(800, 600, "testing chart")
	CheckError(err)
	priceBar, err := NewPriceBar()
	CheckError(err)
	chart.AddPanel(priceBar)
	chart.MainLoop()
}
