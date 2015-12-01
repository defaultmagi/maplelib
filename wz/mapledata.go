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
	"strconv"
)

// 90% of this package is ported directly from OdinMS, so credits to them

// A MapleData is a generic interface for a wz file entry
// with a retrievable value
type MapleData interface {
	Name() string
	Parent() MapleDataEntity
	Type() MapleDataType
	Children() []MapleData
	ChildByPath(path string) MapleData
	Get() interface{}
}

// GetString returns a pointer to the data's value as a string.
// returns nil if the value is not a valid string.
func GetString(d MapleData) *string {
	if d == nil {
		return nil
	}
	val, ok := d.Get().(string)
	if !ok {
		return nil
	}
	return &val
}

// GetStringD returns the data's value as a string.
// If the value can't be retrieved, defval will be returned.
func GetStringD(d MapleData, defval string) string {
	res := GetString(d)
	if res == nil {
		return defval
	}

	return *res
}

// GetDouble returns a pointer to the data's value as a float64.
// returns nil if the value is not a valid float64.
func GetDouble(d MapleData) *float64 {
	if d == nil {
		return nil
	}
	val, ok := d.Get().(float64)
	if !ok {
		return nil
	}
	return &val
}

// GetDoubleD returns the data's value as a float64.
// If the value can't be retrieved, defval will be returned.
func GetDoubleD(d MapleData, defval float64) float64 {
	res := GetDouble(d)
	if res == nil {
		return defval
	}

	return *res
}

// GetFloat returns a pointer to the data's value as a float32.
// returns nil if the value is not a valid float32.
func GetFloat(d MapleData) *float32 {
	if d == nil {
		return nil
	}
	val, ok := d.Get().(float32)
	if !ok {
		return nil
	}
	return &val
}

// GetFloatD returns the data's value as a float32.
// If the value can't be retrieved, defval will be returned.
func GetFloatD(d MapleData, defval float32) float32 {
	res := GetFloat(d)
	if res == nil {
		return defval
	}
	return *res
}

// GetInt returns a pointer to the data's value as an int32.
// returns nil if the value is not a valid int32.
func GetInt(d MapleData) *int32 {
	if d == nil {
		return nil
	}
	val, ok := d.Get().(int32)
	if !ok {
		return nil
	}
	return &val
}

// GetIntD returns the data's value as an int32.
// If the value can't be retrieved, defval will be returned.
func GetIntD(d MapleData, defval int32) int32 {
	res := GetInt(d)
	if res == nil {
		return defval
	}
	return *res
}

// GetIntConvert returns a pointer to the data's value as an int32.
// If the data's value is a string, it will convert the string to int32.
// returns nil if the value is not a valid int32.
func GetIntConvert(d MapleData) *int32 {
	if d == nil {
		return nil
	}

	if d.Type() == STRING {
		pstr := GetString(d)
		if pstr == nil {
			return nil
		}

		i, err := strconv.Atoi(*pstr)
		if err != nil {
			return nil
		}

		res := int32(i)
		return &res
	}

	return GetInt(d)
}

// GetIntConvertD returns the data's value as an int32.
// If the data's value is a string, it will convert the string to int32.
// If the value can't be retrieved, defval will be returned.
func GetIntConvertD(d MapleData, defval int32) int32 {
	res := GetIntConvert(d)
	if res == nil {
		return defval
	}
	return *res
}

// GetImage returns a pointer to the data's value as a pointer to image.Image.
// returns nil if the value is not a valid maple canvas.
func GetImage(d MapleData) *image.Image {
	if d == nil {
		return nil
	}
	val, ok := d.Get().(*FileStoredPngMapleCanvas)
	if !ok {
		return nil
	}
	return val.Image()
}

// GetImageD returns the data's value as a pointer to image.Image.
// If the value can't be retrieved, defval will be returned.
func GetImageD(d MapleData, defval *image.Image) *image.Image {
	res := GetImage(d)
	if res == nil {
		return defval
	}
	return res
}

// GetPoint returns a pointer to the data's value as an image.Point.
// returns nil if the value is not a valid image.Point.
func GetPoint(d MapleData) *image.Point {
	if d == nil {
		return nil
	}
	val, ok := d.Get().(image.Point)
	if !ok {
		return nil
	}
	return &val
}

// GetPointD returns the data's value as an image.Point.
// If the value can't be retrieved, defval will be returned.
func GetPointD(d MapleData, defval image.Point) image.Point {
	res := GetPoint(d)
	if res == nil {
		return defval
	}
	return *res
}

// GetFullDataPath returns the full, absolute path to the data
// by walking the mapledata backwards to the root node.
func GetFullDataPath(d MapleData) string {
	path := ""
	data := MapleDataEntity(d)
	for data != nil {
		path = data.Name() + "/" + path
		data = data.Parent()
	}
	return path[0 : len(path)-1]
}
