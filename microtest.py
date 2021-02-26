import gi
import numpy as np
gi.require_version("Gtk", "3.0")
from gi.repository import Gtk, Gdk, GdkPixbuf

def on_activate(app):
    win = Gtk.ApplicationWindow(application=app)
    win.set_default_size(800, 600)
    scrollbox = Gtk.ScrolledWindow(hadjustment=None, vadjustment=None)
    flowBox = Gtk.FlowBox()
    flowBox.set_homogeneous(True)
    flowBox.set_max_children_per_line(10)
    for x in range(1000):
        # the following line creates a 44x44 green image
        img = np.full((44,44,4), [0,0,0xff,0xff], dtype=np.uint8).tobytes()
        pixbuf = GdkPixbuf.Pixbuf.new_from_data(img, GdkPixbuf.Colorspace.RGB, True, 8, 44, 44 ,44 * 4, None, None)
        gtkImage = Gtk.Image.new_from_pixbuf(pixbuf)
        flowBox.add(gtkImage)
    scrollbox.add(flowBox)
    win.add(scrollbox)
    win.show_all()


app = Gtk.Application(application_id='org.microtest.PixbufBugs')
app.connect('activate', on_activate)
app.run(None)
