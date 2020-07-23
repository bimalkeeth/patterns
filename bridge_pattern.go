package main

import "fmt"

type VectorRenderer struct {

}

type Renderer interface {
	RenderCircle(radius float32)
}

func(v *VectorRenderer)RenderCircle(radius float32){

	fmt.Println("Drawing a circle of radius",radius)
}

type RasterRenderer struct {
	Dpi int
}

func(v *RasterRenderer)RenderCircle(radius float32){

	fmt.Println("Drawing a pixels for radius of radius",radius)
}

type Circle struct {
	renderer Renderer
	radius float32
}


func NewCircle(renderer Renderer,radius float32) *Circle{
	return &Circle{renderer: renderer,radius: radius}
}

func(c *Circle)Draw(){
	c.renderer.RenderCircle(c.radius)
}

func(c *Circle)Resize(factor float32){
	c.radius *=factor
}


func main() {

	raster :=RasterRenderer{}
	//vector :=VectorRenderer{}
	circle:=NewCircle(&raster,5)
	circle.Draw()
	circle.Resize(10)
	
}
