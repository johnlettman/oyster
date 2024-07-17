package metadata

import (
	"github.com/johnlettman/oyster/types"
	"time"
)

// SensorInfo represents basic information about the sensor.
type SensorInfo struct {
	Status           types.OperatingMode `json:"status"`            // Current status of the sensor.
	InitializationID uint                `json:"initialization_id"` // Startup initialization number of the sensor.

	ProductLine  string `json:"prod_line"` // The market product line of the sensor.
	PartNumber   string `json:"prod_pn"`   // The part number of the sensor.
	SerialNumber string `json:"prod_sn"`   // The serial number of the sensor.

	BuildDate        time.Time `json:"build_date"`          // The build date of the sensor.
	BuildRevision    string    `json:"build_rev"`           // The build revision of the sensor.
	ImageRevision    string    `json:"image_rev"`           // The revision of the firmware on the sensor.
	ProtocolRevision string    `json:"proto_rev,omitempty"` // The revision of the communications protocol.
}
