package driver

import (
	"testing"
	"time"

	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

var d *RandomDriver

func init() {
	d = new(RandomDriver)
	d.lc = logger.NewClient("devicerandom", false, "", "DEBUG")
	d.randomDevices = make(map[string]*randomDevice)
}

func TestHandleReadCommands(t *testing.T) {
	deviceName := "testDevice"
	protocols := map[string]models.ProtocolProperties{
		"other": {
			"Address": "simple01",
			"Port":    "300",
		},
	}

	requests := []dsModels.CommandRequest{
		{
			DeviceResourceName: "RandomValue_Int8",
			Type: dsModels.Int8,
		},
		{
			DeviceResourceName: "RandomValue_Int16",
			Type: dsModels.Int16,
		},
		{
			DeviceResourceName: "RandomValue_Int32",
			Type: dsModels.Int32,
		},
	}

	res, err := d.HandleReadCommands(deviceName, protocols, requests)

	if err != nil {
		t.Fatalf("Failt to read command, %v", err)
	}
	if len(res) != len(requests) {
		t.Fatalf("Result amount '%v' should match '%v'", len(res), len(requests))
	}
	if res[0].DeviceResourceName != "RandomValue_Int8" || res[1].DeviceResourceName != "RandomValue_Int16" || res[2].DeviceResourceName != "RandomValue_Int32" {
		t.Fatalf("Unexpect test result. Wrong resource object.")
	}
	if res[0].Type != dsModels.Int8 || res[1].Type != dsModels.Int16 || res[2].Type != dsModels.Int32 {
		t.Fatalf("Unexpect test result. Wrong value type.")
	}
}

func TestHandleWriteCommands(t *testing.T) {
	deviceName := "testDevice"
	protocols := map[string]models.ProtocolProperties{
		"other": {
			"Address": "simple01",
			"Port":    "300",
		},
	}
	requests := []dsModels.CommandRequest{}

	now := time.Now().UnixNano() / int64(time.Millisecond)
	cv, err := dsModels.NewInt8Value("Max_Int8", now, int8(127))
	if err != nil {
		t.Fatalf("Failt to create command value, %v", err)
	}
	params := []*dsModels.CommandValue{cv}

	err = d.HandleWriteCommands(deviceName, protocols, requests, params)

	if err != nil {
		t.Fatalf("Failt to write command, %v", err)
	}
}
