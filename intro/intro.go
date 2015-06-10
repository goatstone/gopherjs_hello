package intro

import "honnef.co/go/js/dom"

type Intro struct {
	Name string
	Greet string
	ID int
	El dom.HTMLElement
} 
func (p *Intro) Init() bool {
	p.El.SetTextContent(p.Name)
	style :=  p.El.Style();
	style.Set("margin", "0 auto")
	style.Set("text-align", "center")
	style.Set("padding", "30px")
	style.Set("font-family", "sans-serif")
	style.Set("border-radius", "20px")
	style.Set("overflow", "hidden")
	style.Set("color", "#ccc")
	style.Set("font-size", "3em")
	style.Set("display", "block")
	style.Set("background-color", "green")
	style.Set("width", "100px")
	style.Set("height", "300px")
	style.Set("-webkit-transition", "width 4s ease-in")	
	style.Set("transition", "width 4s ease-in")	
	style.Set("width", "30px")
    return true
}
func (p *Intro) Hello() bool {
	p.El.SetTextContent(p.Name + ", "+ p.Greet)
	p.El.Style().Set("width", "90%")
    return true
}