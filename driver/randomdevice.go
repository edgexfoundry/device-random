// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	defMinInt8, defMaxInt8   = -128, 127
	defMinInt16, defMaxInt16 = -32768, 32767
	defMinInt32, defMaxInt32 = -2147483648, 2147483647
)

type randomDevice struct {
	minInt8  int64
	maxInt8  int64
	minInt16 int64
	maxInt16 int64
	minInt32 int64
	maxInt32 int64
}

func (d *randomDevice) value(valueType string) (int64, error) {
	switch valueType {
	case "Int8":
		if d.maxInt8 <= d.minInt8 {
			return 0, fmt.Errorf("randomDevice.value: maximum: %d <= minimum : %d", d.maxInt8, d.minInt8)
		} else {
			return random(d.minInt8, d.maxInt8), nil
		}
	case "Int16":
		if d.maxInt16 <= d.minInt16 {
			return 0, fmt.Errorf("randomDevice.value: maximum: %d <= minimum : %d", d.maxInt16, d.minInt16)
		} else {
			return random(d.minInt16, d.maxInt16), nil
		}
	case "Int32":
		if d.maxInt32 <= d.minInt32 {
			return 0, fmt.Errorf("randomDevice.value: maximum: %d <= minimum : %d", d.maxInt32, d.minInt32)
		} else {
			return random(d.minInt32, d.maxInt32), nil
		}
	default:
		return 0, fmt.Errorf("randomDevice.value: wrong value type: %s", valueType)
	}
}

func newRandomDevice() *randomDevice {
	return &randomDevice{
		minInt8:  defMinInt8,
		maxInt8:  defMaxInt8,
		minInt16: defMinInt16,
		maxInt16: defMaxInt16,
		minInt32: defMinInt32,
		maxInt32: defMaxInt32,
	}
}

func random(min int64, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min) + min
}
