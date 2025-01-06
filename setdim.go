package setdim

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"path/filepath"
)

const Inch = 2.54
const Cm = 1.0

type Dims struct {
	Width float64
	Height float64
	XPixelsPerCm float64
	YPixelsPerCm float64
}

func GetDims(path string) (Dims, error) {
	var b strings.Builder
	cmd := exec.Command("identify", "-units", "PixelsPerCentimeter", "-format", "%w x %h %x x %y\\n", path)
	cmd.Stdout = &b
	cmd.Stderr = os.Stderr
	if e := cmd.Run(); e != nil {
		return Dims{}, e
	}
	var d Dims
	_, e := fmt.Sscanf(b.String(), "%f x %f %f x %f",
		&d.Width,
		&d.Height,
		&d.XPixelsPerCm,
		&d.YPixelsPerCm,
	)
	return d, e
}

func SetDensity(path, outpath string, densityCm float64) error {
	cmd := exec.Command("convert", "-units", "PixelsPerCentimeter", "-density", fmt.Sprint(densityCm), path, outpath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func SetDensityAndLabel(path, outpath string, densityCm float64, label string) error {
	cmd := exec.Command("convert",
		"-units", "PixelsPerCentimeter",
		"-density", fmt.Sprint(densityCm),
		"-bordercolor", "white",
		"-border", fmt.Sprint(densityCm / 2.0),
		"-annotate", "+0+0", label,
		"-pointsize", fmt.Sprint(densityCm * 12.0 / 72.0),
		"-gravity", "northwest",
		path,
		outpath,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func SetWidthRaster(path, outpath string, widthCm float64) error {
	dims, err := GetDims(path)
	if err != nil {
		return err
	}
	// newDensity := dims.XPixelsPerCm * (dims.Width / widthCm)
	newDensity := dims.Width / widthCm
	return SetDensity(path, outpath, newDensity)
}

func SetWidthRasterAndLabel(path, outpath string, widthCm float64, label string) error {
	dims, err := GetDims(path)
	if err != nil {
		return err
	}
	// newDensity := dims.XPixelsPerCm * (dims.Width / widthCm)
	newDensity := dims.Width / widthCm
	return SetDensityAndLabel(path, outpath, newDensity, label)
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
	initCm := dims.Width / dims.XPixelsPerCm
	firstDensity := widthCm * densityPerCm / initCm
	finalDensity := densityPerCm
	midpath := filepath.Join(tmpdir, "temp" + filepath.Ext(outpath))

	if err = SetDensity(path, midpath, firstDensity); err != nil {
		return err
	}
	return SetDensity(midpath, outpath, finalDensity)
}

func SetWidthVectorAndLabel(path, outpath string, widthCm, densityPerCm float64, label string) (err error) {
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
	initCm := dims.Width / dims.XPixelsPerCm
	firstDensity := widthCm * densityPerCm / initCm
	finalDensity := densityPerCm
	midpath := filepath.Join(tmpdir, "temp" + filepath.Ext(outpath))

	if err = SetDensity(path, midpath, firstDensity); err != nil {
		return err
	}
	return SetDensityAndLabel(midpath, outpath, finalDensity, label)
}

// identify -format "%w x %h %x x %y\n" "$@"
// convert -units PixelsPerInch figure1.png -density 600 figure1resized.png
