package ui

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// static GtkScale *toGtkScale(void *p) { return (GTK_SCALE(p)); }
import "C"

import (
    "unsafe"

    "github.com/gotk3/gotk3/gtk"
)

type ScaleExt struct {
    *gtk.Scale
}

func gbool(b bool) C.gboolean {
    if b {
        return C.gboolean(1)
    } else {
        return C.gboolean(0)
    }
}

func (scale *ScaleExt) native() *C.GtkScale {
    if scale == nil || scale.GObject == nil {
        return nil
    }
    ptr := unsafe.Pointer(scale.GObject)
    return C.toGtkScale(ptr)
}

func (scale *ScaleExt) SetHasOrigin(hasOrigin bool) {
    C.gtk_scale_set_has_origin(scale.native(), gbool(hasOrigin))
}

func ScaleExtNew(orientaion gtk.Orientation, adjustment *gtk.Adjustment) (*ScaleExt, error) {
    scale, err := gtk.ScaleNew(orientaion, adjustment)
    return &ScaleExt{scale}, err
}
