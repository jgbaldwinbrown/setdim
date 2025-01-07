package main

import (
	"github.com/jgbaldwinbrown/setdim/pkg"
)

//#include "cdefs.h"
import "C"

func toNum(err error) C.int {
	if err != nil {
		return 1
	}
	return 0
}

//export getDims
func getDims(path *C.char, err *C.int) C.struct_dims {
	gdims, e := setdim.GetDims(C.GoString(path))
	*err = toNum(e)
	return C.struct_dims{
		width: C.double(gdims.Width),
		height: C.double(gdims.Height),
		x_pixels_per_cm: C.double(gdims.XPixelsPerCm),
		y_pixels_per_cm: C.double(gdims.YPixelsPerCm),
	}
}

//export setDensity
func setDensity(path, outpath *C.char, densityCm C.double) C.int {
	return toNum(setdim.SetDensity(C.GoString(path), C.GoString(outpath), float64(densityCm)))
}

//export setDensityAndLabel
func setDensityAndLabel(path, outpath *C.char, densityCm C.double, label *C.char) C.int {
	return toNum(setdim.SetDensityAndLabel(C.GoString(path), C.GoString(outpath), float64(densityCm), C.GoString(label)))
}

//export setWidthRaster
func setWidthRaster(path, outpath *C.char, widthCm C.double) C.int {
	return toNum(setdim.SetWidthRaster(C.GoString(path), C.GoString(outpath), float64(widthCm)))
}

//export setWidthRasterAndLabel
func setWidthRasterAndLabel(path, outpath *C.char, widthCm C.double, label *C.char) C.int {
	return toNum(setdim.SetWidthRasterAndLabel(C.GoString(path), C.GoString(outpath), float64(widthCm), C.GoString(label)))
}


//export setWidthVector
func setWidthVector(path, outpath *C.char, widthCm, densityPerCm C.double) C.int {
	return toNum(setdim.SetWidthVector(C.GoString(path), C.GoString(outpath), float64(widthCm), float64(densityPerCm)))
}

//export setWidthVectorAndLabel
func setWidthVectorAndLabel(path, outpath *C.char, widthCm, densityPerCm C.double, label *C.char) C.int {
	return toNum(setdim.SetWidthVectorAndLabel(C.GoString(path), C.GoString(outpath), float64(widthCm), float64(densityPerCm), C.GoString(label)))
}

func main() {
}
