package main

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func onActivate(app *gtk.Application) {
	win, err := gtk.ApplicationWindowNew(app)
	if err != nil {
		log.Fatal(err)
	}
	win.SetDefaultSize(800, 600)
	scrollbox, err := gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	flowBox, err := gtk.FlowBoxNew()
	if err != nil {
		log.Fatal(err)
	}
	flowBox.SetHomogeneous(true)
	flowBox.SetMinChildrenPerLine(10)
	for x := 0; x < 1000; x++ {
		img := make([]byte, 44*44*4)
		for i := 0; i < 44*44; i++ {
			img[i*4+0] = 0   //r
			img[i*4+1] = 0   //g
			img[i*4+2] = 255 //b
			img[i*4+3] = 255 //a
		}
		pixbuf, err := gdk.PixbufNewFromData(
			img,
			gdk.COLORSPACE_RGB,
			true,
			8,
			44,
			44,
			44*4,
		)
		if err != nil {
			log.Fatal(err)
		}
		gtkImage, err := gtk.ImageNewFromPixbuf(pixbuf)
		if err != nil {
			log.Fatal(err)
		}
		flowBox.Add(gtkImage)

	}
	scrollbox.Add(flowBox)

	win.Add(scrollbox)
	win.ShowAll()
}

func main() {
	app, err := gtk.ApplicationNew("org.microtest.PixbufBugs", glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		log.Fatal(err)
	}
	app.Connect("activate", onActivate)
	os.Exit(app.Run(os.Args))
}
