package main

import (
	"github.com/jgbaldwinbrown/setdim/pkg"
)

//#include "cdefs.h"
import "C"

//export getDims
func getDims(path *C.char, err *C.int) C.struct_dims {
	*err = 0
	gpath := C.GoString(path)
	gdims, e := setdim.GetDims(gpath)
	if e != nil {
		*err = 1
		return C.struct_dims{}
	}
	dims := C.struct_dims{
		width: C.double(gdims.Width),
		height: C.double(gdims.Height),
		x_pixels_per_cm: C.double(gdims.XPixelsPerCm),
		y_pixels_per_cm: C.double(gdims.YPixelsPerCm),
	}
	return dims
}

//export setDensity
func setDensity(path, outpath *C.char, densityCm C.double) C.int {
	gpath := C.GoString(path)
	goutpath := C.GoString(outpath)
	if err := setdim.SetDensity(gpath, goutpath, float64(densityCm)); err != nil {
		return 1
	}
	return 0
}

//export setDensityAndLabel
func setDensityAndLabel(path, outpath *C.char, densityCm C.double, label *C.char) C.int {
	gpath := C.GoString(path)
	goutpath := C.GoString(outpath)
	glabel := C.GoString(label)
	if err := setdim.SetDensityAndLabel(gpath, goutpath, float64(densityCm), glabel); err != nil {
		return 1
	}
	return 0
}

//export setWidthRaster
func setWidthRaster(path, outpath *C.char, widthCm C.double) C.int {
	gpath := C.GoString(path)
	goutpath := C.GoString(outpath)
	if err := setdim.SetWidthRaster(gpath, goutpath, float64(widthCm)); err != nil {
		return 1
	}
	return 0
}

//export setWidthRasterAndLabel
func setWidthRasterAndLabel(path, outpath *C.char, widthCm C.double, label *C.char) C.int {
	gpath := C.GoString(path)
	goutpath := C.GoString(outpath)
	glabel := C.GoString(label)
	if err := setdim.SetWidthRasterAndLabel(gpath, goutpath, float64(widthCm), glabel); err != nil {
		return 1
	}
	return 0
}


//export setWidthVector
func setWidthVector(path, outpath *C.char, widthCm, densityPerCm C.double) C.int {
	gpath := C.GoString(path)
	goutpath := C.GoString(outpath)
	if err := setdim.SetWidthVector(gpath, goutpath, float64(widthCm), float64(densityPerCm)); err != nil {
		return 1
	}
	return 0
}

//export setWidthVectorAndLabel
func setWidthVectorAndLabel(path, outpath *C.char, widthCm, densityPerCm C.double, label *C.char) C.int {
	gpath := C.GoString(path)
	goutpath := C.GoString(outpath)
	glabel := C.GoString(label)
	if err := setdim.SetWidthVectorAndLabel(gpath, goutpath, float64(widthCm), float64(densityPerCm), glabel); err != nil {
		return 1
	}
	return 0
}

func main() {
}
