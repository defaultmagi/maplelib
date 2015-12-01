/*
   Copyright 2014-2015 Franc[e]sco (lolisamurai@tfwno.gf)
   This file is part of maplelib-go.
   maplelib-go is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   maplelib-go is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with maplelib-go. If not, see <http://www.gnu.org/licenses/>.
*/

package wz

import (
	"image"
	_ "image/png" // must be loaded to decode png files
	"os"
	"path/filepath"
)

// 90% of this package is ported directly from OdinMS, so credits to them

// A FileStoredPngMapleCanvas is a wrapper around image.Image that holds info
// about a png image extracted from a wz file
type FileStoredPngMapleCanvas struct {
	filepath string
	width    int
	height   int
	img      *image.Image
}

// NewFileStoredPngMapleCanvas initializes a new FileStoredPngMapleCanvas object
// with the given file path and size
func NewFileStoredPngMapleCanvas(w, h int, path string,
) *FileStoredPngMapleCanvas {

	return &FileStoredPngMapleCanvas{
		filepath: path,
		width:    w,
		height:   h,
	}
}

func (f *FileStoredPngMapleCanvas) Height() int { return f.height }
func (f *FileStoredPngMapleCanvas) Width() int  { return f.width }
func (f *FileStoredPngMapleCanvas) Image() *image.Image {
	f.loadImageIfNecessary()
	return f.img
}

// setTestPathPrefix is internally used to append the absolute path to the
// image's path when running unit tests that are stored in temporary folders
func (f *FileStoredPngMapleCanvas) setTestPathPrefix(prefix string) {
	f.filepath = filepath.Join(prefix, f.filepath)
}

func (f *FileStoredPngMapleCanvas) loadImageIfNecessary() {
	if f.img != nil {
		return
	}

	file, err := os.Open(f.filepath)
	if err != nil {
		return
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return
	}

	f.img = &img
}
