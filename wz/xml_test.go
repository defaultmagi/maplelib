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
	"fmt"
	"image"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestXml(t *testing.T) {
	expected := []interface{}{
		int32(180), float32(100.0), "L3", image.Pt(29, 51)}

	path := filepath.Join(os.Getenv("GOPATH"), "src", "github.com",
		"Francesco149", "maplelib", "wz", "testfiles")

	x, err := NewMapleDataProvider(path)
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	// ----------------------------------------

	img, err := x.Get("TamingMob.wz/0003.img")
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	// INT
	val := img.ChildByPath("info/speed").Get()
	if val != expected[0] {
		v := reflect.ValueOf(val)
		e := reflect.ValueOf(expected[0])
		t.Errorf("TamingMob.wz/0003.img/info/speed = %v(%v), expected %v(%v)",
			v.Type(), val, e.Type(), expected[0])
	}
	conval := GetInt(img.ChildByPath("info/speed"))
	if conval == nil {
		t.Errorf("Failed to convert TamingMob.wz/0003.img/info/speed to int32")
	}
	if *conval != expected[0] {
		t.Errorf("converted TamingMob.wz/0003.img/info/speed = %v, expected %v",
			*conval, expected[0])
	}

	// FLOAT
	val = img.ChildByPath("info/swim").Get()
	v := reflect.ValueOf(val)
	e := reflect.ValueOf(expected[1])
	tmpstr := fmt.Sprintf("%v(%v)", v.Type(), val)
	tmpexpectedstr := fmt.Sprintf("%v(%v)", e.Type(), expected[1])
	if tmpstr != tmpexpectedstr {
		t.Errorf("TamingMob.wz/0003.img/info/speed = %v(%v), expected %v(%v)",
			v.Type(), val, e.Type(), expected[1])
	}
	conval1 := GetFloat(img.ChildByPath("info/swim"))
	if conval1 == nil {
		t.Errorf("Failed to convert TamingMob.wz/0003.img/info/swim to float32")
	}
	if *conval1 != expected[1] {
		t.Errorf("converted TamingMob.wz/0003.img/info/swim = %v, expected %v",
			*conval1, expected[1])
	}

	// ----------------------------------------

	img, err = x.Get("Mob.wz/0210100.img")
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	// STRING
	val = img.ChildByPath("info/elemAttr").Get()
	if val != expected[2] {
		v := reflect.ValueOf(val)
		e := reflect.ValueOf(expected[2])
		t.Errorf("Mob.wz/0210100.img/info/elemAttr = %v(%v), expected %v(%v)",
			v.Type(), val, e.Type(), expected[2])
	}
	conval2 := GetString(img.ChildByPath("info/elemAttr"))
	if conval2 == nil {
		t.Errorf(
			"Failed to convert TamingMob.wz/0003.img/info/elemAttr to string")
	}
	if *conval2 != expected[2] {
		t.Errorf(
			"converted TamingMob.wz/0003.img/info/elemAttr = %v, expected %v",
			*conval2, expected[2])
	}

	// CANVAS
	val = img.ChildByPath("move/0").Get()
	c, ok := val.(MapleCanvas)
	if !ok {
		t.Errorf("Mob.wz/0210100.img/move/0: is not a MapleCanvas")
	}

	c.setTestPathPrefix(path)
	canvas := c.Image()
	if canvas == nil {
		t.Errorf("Mob.wz/0210100.img/move/0: failed to load 0.png")
	}

	w := (*canvas).Bounds().Max.X
	h := (*canvas).Bounds().Max.Y
	if w != 237 {
		t.Errorf("Mob.wz/0210100.img/move/0/0.png: width=%v, expected 237")
	}
	if h != 248 {
		t.Errorf("Mob.wz/0210100.img/move/0/0.png: height=%v, expected 248")
	}

	// VECTOR
	val = img.ChildByPath("move/0/origin").Get()
	if val != expected[3] {
		v := reflect.ValueOf(val)
		e := reflect.ValueOf(expected[3])
		t.Errorf("Mob.wz/0210100.img/info/elemAttr = %v(%v), expected %v(%v)",
			v.Type(), val, e.Type(), expected[3])
	}
}
