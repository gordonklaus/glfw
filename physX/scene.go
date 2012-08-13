package physX

// #cgo LDFLAGS: CphysX.so
// #include "scene.h"
import "C"

import "unsafe"

type Scene struct { s unsafe.Pointer }

func NewScene() Scene { return Scene{C.newScene()} }
func (s Scene) NewDynamicActor(pos Vector) DynamicActor { return DynamicActor{C.Scene_newDynamicActor(s.s, pos.floatptr())} }
func (s Scene) Simulate(dt float32) { C.Scene_simulate(s.s, C.float(dt)) }

