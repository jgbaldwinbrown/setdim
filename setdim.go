package setdim

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"path/filepath"
)

const Inch = 2.54

type Dims struct {
	Width float64
	Height float64
	XPixelsPerCm float64
	YPixelsPerCm float64
}

func GetDims(path string) (Dims, error) {
	var b strings.Builder
	cmd := exec.Command("identify", "-format", "%w x %h %x x %y\\n", path)
	cmd.Stdout = &b
	cmd.Stderr = os.Stderr
	if e := cmd.Run(); e != nil {
		return Dims{}, e
	}
	var d Dims
	_, e := fmt.Sscanf(b.String(), "%f x %f %f x %f")
	return d, e
}

func SetDensity(path, outpath string, densityCm float64) error {
	cmd := exec.Command("convert", "-units", "PixelsPerCm", path, "-density", fmt.Sprint(densityCm), outpath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func SetWidthRaster(path, outpath string, widthCm float64) error {
	dims, err := GetDims(path)
	if err != nil {
		return err
	}
	newDensity := dims.XPixelsPerCm * (dims.Width / widthCm)
	return SetDensity(path, outpath, newDensity)
}

func SetWidthVector(path, outpath string, widthCm, densityPerCm float64) (err error) {
	tmpdir, err := os.MkdirTemp("", "setdim_SetWidthVector_*")
	if err != nil {
		return err
	}
	defer func() {
		e := os.RemoveAll(tmpdir)
		if err == nil {
			err = e
		}
	}()

	dims, err := GetDims(path)
	if err != nil {
		return err
	}
	firstDensity := widthCm * densityPerCm / dims.Width
	finalDensity := densityPerCm
	midpath := filepath.Join(tmpdir, "temp" + filepath.Ext(outpath))

	if err = SetDensity(path, midpath, firstDensity); err != nil {
		return err
	}
	return SetDensity(midpath, outpath, finalDensity)
}

// identify -format "%w x %h %x x %y\n" "$@"
// convert -units PixelsPerInch figure1.png -density 600 figure1resized.png
