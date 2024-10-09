package typeMapper

import "unsafe"

func typelinks2() (sections []unsafe.Pointer, offset [][]int32)

func resolveTypeOff(rtype unsafe.Pointer, off int32) unsafe.Pointer

type emptyInterface struct {
	typ  unsafe.Pointer
	data unsafe.Pointer
}
