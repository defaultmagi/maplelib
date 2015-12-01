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

// 90% of this package is ported directly from OdinMS, so credits to them

import (
	"errors"
	"image"
	"os"
	"path/filepath"
	"strings"
)

import "github.com/jteeuwen/go-pkg-xmlx"

// A XMLDomMapleData is a wrapper around xmlx.Node and
// holds the data for a parsed wz xml file's node and provides
// access to all of its children
type XMLDomMapleData struct {
	node         *xmlx.Node
	imageDataDir string
}

// NewXMLDomMapleData parses the given xml file into a tree and returns the
// first node
func NewXMLDomMapleData(file *os.File, path string) (
	res *XMLDomMapleData, err error) {

	doc := xmlx.New()
	err = doc.LoadStream(file, nil)

	res = &XMLDomMapleData{
		node:         doc.Root.Children[1],
		imageDataDir: path,
	}
	return
}

// fromNode is internally used to wrap child nodes as XMLDomMapleData objects
func fromNode(node *xmlx.Node) *XMLDomMapleData {
	return &XMLDomMapleData{
		node: node,
	}
}

// ChildByPath finds and returns a value by xmlpath relative to the wz xml file
func (x *XMLDomMapleData) ChildByPath(path string) MapleData {
	segments := strings.Split(path, "/")
	if segments[0] == ".." {
		par := x.Parent()
		res, ok := par.(MapleData)
		if !ok {
			panic(errors.New(
				"XMLDomMapleData.ChildByPath: par failed type assertion"))
			return nil
		}
		return res.ChildByPath(path[strings.Index(path, "/")+1:])
	}

	mynode := x.node
	newdatadir := x.imageDataDir
	// we will append the full node xmlpath to newdatadir

	// iterate the path subfolder by subfolder
	for i := 0; i < len(segments); i++ {
		foundChild := false

		// look for the node in the children nodes
		for j := 0; j < len(mynode.Children); j++ {
			if mynode.Children[j].Type != xmlx.NT_ELEMENT {
				continue
			}

			// found the node we're lookgin for
			if mynode.Children[j].As("", "name") == segments[i] {
				mynode = mynode.Children[j]
				newdatadir = filepath.Join(newdatadir, mynode.As("", "name"))
				foundChild = true
				break
			}
		}

		// node not found, no data is returned
		if !foundChild {
			return nil
		}
	}

	// return the desired node
	res := fromNode(mynode)
	res.imageDataDir = newdatadir
	// imageDataDir now holds the correct path for the png file

	return res
}

// Children returns the children entries of this node
func (x *XMLDomMapleData) Children() []MapleData {
	res := make([]MapleData, 0)
	childNodes := x.node.Children
	for j := 0; j < len(childNodes); j++ {
		childNode := childNodes[j]
		if childNode.Type != xmlx.NT_ELEMENT {
			continue
		}

		child := fromNode(childNode)
		child.imageDataDir = filepath.Join(x.imageDataDir, x.Name())
		res = append(res, child)
	}

	return res
}

// Get returns the value of this node as an interface.
// If the value is invalid or absent, the return value is nil.
// All the possible types returned by Get are float64, float32, int32, int16,
// string, image.Point and FileStoredPngMapleCanvas.
func (x *XMLDomMapleData) Get() interface{} {
	datatype := x.Type()

	switch datatype {
	case DOUBLE:
		return x.node.Af64("", "value")
	case FLOAT:
		return x.node.Af32("", "value")
	case INT:
		return x.node.Ai32("", "value")
	case SHORT:
		return x.node.Ai16("", "value")
	case STRING, UOL:
		return x.node.As("", "value")

	case VECTOR:
		vx := x.node.Ai("", "x")
		vy := x.node.Ai("", "y")
		return image.Pt(vx, vy)

	case CANVAS:
		w := x.node.Ai("", "width")
		h := x.node.Ai("", "height")
		return NewFileStoredPngMapleCanvas(w, h,
			x.imageDataDir+".png")
	}

	return nil
}

// Type returns the maple data type of this node.
// See MapleDataType for more information.
func (x *XMLDomMapleData) Type() MapleDataType {
	nodeName := x.node.Name.Local

	switch nodeName {
	case "imgdir":
		return PROPERTY
	case "canvas":
		return CANVAS
	case "convex":
		return CONVEX
	case "sound":
		return SOUND
	case "uol":
		return UOL
	case "double":
		return DOUBLE
	case "float":
		return FLOAT
	case "int":
		return INT
	case "short":
		return SHORT
	case "string":
		return STRING
	case "vector":
		return VECTOR
	case "null":
		return IMG_0x00
	default:
		return INVALID
	}

	return INVALID
}

// Parent returns the parent node of this wz xml entry
func (x *XMLDomMapleData) Parent() MapleDataEntity {
	parentNode := x.node.Parent
	if parentNode.Type == xmlx.NT_ROOT {
		return nil
	}

	parentData := fromNode(parentNode)
	parentData.imageDataDir = x.imageDataDir[0:strings.LastIndex(x.imageDataDir, "/")]
	return parentData
}

// Name returns the name of the wz xml entry
func (x *XMLDomMapleData) Name() string {
	return x.node.As("", "name")
}
