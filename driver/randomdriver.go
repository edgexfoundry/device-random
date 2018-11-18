// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

// This package provides a implementation of a ProtocolDriver interface.
//
package driver

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
	"github.com/edgexfoundry/edgex-go/pkg/clients/logging"
	"github.com/edgexfoundry/edgex-go/pkg/models"
)

type RandomDriver struct {
	lc      logger.LoggingClient
	asyncCh chan<- *dsModels.AsyncValues
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func (d *RandomDriver) DisconnectDevice(address *models.Addressable) error {
	d.lc.Info(fmt.Sprintf("RandomDriver.DisconnectDevice: device-random driver is disconnecting to %v", address))
	return nil
}

func (d *RandomDriver) Initialize(lc logger.LoggingClient, asyncCh chan<- *dsModels.AsyncValues) error {
	d.lc = lc
	d.asyncCh = asyncCh
	return nil
}

func (d *RandomDriver) HandleReadCommands(addr *models.Addressable, reqs []dsModels.CommandRequest) (res []*dsModels.CommandValue, err error) {
	res = make([]*dsModels.CommandValue, len(reqs))
	now := time.Now().UnixNano() / int64(time.Millisecond)
	rand.Seed(time.Now().UnixNano())

	for i, req := range reqs {
		t := req.DeviceObject.Properties.Value.Type
		defMin, defMax := 0, 0

		switch t {
		case "Int8":
			defMin, defMax = -128, 127
		case "Int16":
			defMin, defMax = -32768, 32767
		case "Int32":
			defMin, defMax = -2147483648, 2147483647
		}

		min, e := strconv.Atoi(req.DeviceObject.Properties.Value.Minimum)
		if e != nil || min < defMin || min > defMax {
			err = fmt.Errorf("RandomDriver.HandleReadCommands: minimum value %d of %T must be int between %d ~ %d", min, min, defMin, defMax)
			return
		}

		max, e := strconv.Atoi(req.DeviceObject.Properties.Value.Maximum)
		if e != nil || max < defMin || max > defMax || max < min {
			err = fmt.Errorf("RandomDriver.HandleReadCommands: maximum value %d of %T must be int between %d ~ %d and greater than min", max, max, defMin, defMax)
			return
		}

		var cv *dsModels.CommandValue
		switch t {
		case "Int8":
			cv, _ = dsModels.NewInt8Value(&req.RO, now, int8(random(min, max)))
		case "Int16":
			cv, _ = dsModels.NewInt16Value(&req.RO, now, int16(random(min, max)))
		case "Int32":
			cv, _ = dsModels.NewInt32Value(&req.RO, now, int32(random(min, max)))
		}
		res[i] = cv
	}

	return
}

func (d *RandomDriver) HandleWriteCommands(addr *models.Addressable, reqs []dsModels.CommandRequest,
	params []*dsModels.CommandValue) error {
	return fmt.Errorf("RandomDriver.HandleWriteCommands: device-random doesn't support write operation")
}

func (d *RandomDriver) Stop(force bool) error {
	d.lc.Info("RandomDriver.Stop: device-random driver is stopping...")
	return nil
}
