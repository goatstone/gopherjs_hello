package main

import "github.com/gopherjs/gopherjs/js"
import "honnef.co/go/js/dom"
import "github.com/goatstone/gopherjs_hello/intro"

func main() {
	el := dom.GetWindow().Document().QuerySelector("div")
	htmlEl := el.(dom.HTMLElement)
  	p := intro.Intro{  "GopherJS", "Hello!!!!", 2, htmlEl}
	p.Init()
	js.Global.Get("window").Call("addEventListener", "load", func() {
	  go func() {
	  	p.Hello()
	  }()
	})
}
