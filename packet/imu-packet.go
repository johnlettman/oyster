package packet

import (
	"fmt"
	"github.com/johnlettman/buffergenerics"
	"github.com/johnlettman/oyster/packet/decode"
	"strings"
)

// IMUPacket represents the data packet from an IMU (Inertial Measurement Unit) sensor.
// It contains information about timestamps, linear acceleration, and angular velocity.
// For additional information, refer to [IMU Data Format].
//
// [IMU Data Format]: https://static.ouster.dev/sensor-docs/image_route1/image_route3/sensor_data/sensor-data.html#imu-data-format
type IMUPacket struct {
	DiagnosticSystemTime uint64 // Timestamp of the monotonic system time since boot in ns.
	AccelerometerTime    uint64 // Timestamp for the Accelerometer relative to types.TimeSource in ns.
	GyroscopeTime        uint64 // Timestamp for the Gyroscope relative to types.TimeSource in ns.

	LinearAccelerationX float32 // Measured linear acceleration in g for the X axis.
	LinearAccelerationY float32 // Measured linear acceleration in g for the Y axis.
	LinearAccelerationZ float32 // Measured linear acceleration in g for the Z axis.

	AngularVelocityX float32 // Measured angular velocity in °/sec for the X axis.
	AngularVelocityY float32 // Measured angular velocity in °/sec for the Y axis.
	AngularVelocityZ float32 // Measured angular velocity in °/sec for the Z axis.
}

// imuStringPrecision is the precision for formatting linear acceleration and angular velocity values.
const imuStringPrecision = "4"

func (p IMUPacket) String() string {
	return fmt.Sprintf("IMUPacket(t+%v, →a=%."+imuStringPrecision+"f, ω=%."+imuStringPrecision+"f)",
		p.DiagnosticSystemTime, p.LinearAcceleration(), p.AngularVelocity())
}

func (p IMUPacket) GoString() string {
	var lines []string

	if p.DiagnosticSystemTime > 0 {
		lines = append(lines, fmt.Sprintf("\tDiagnosticSystemTime: %v", p.DiagnosticSystemTime))
	}

	if p.AccelerometerTime > 0 {
		lines = append(lines, fmt.Sprintf("\tAccelerometerTime:    %v", p.AccelerometerTime))
	}

	if p.GyroscopeTime > 0 {
		lines = append(lines, fmt.Sprintf("\tGyroscopeTime:        %v", p.GyroscopeTime))
	}

	if p.LinearAccelerationX != 0 {
		lines = append(lines, fmt.Sprintf("\tLinearAccelerationX:  %v", p.LinearAccelerationX))
	}

	if p.LinearAccelerationY != 0 {
		lines = append(lines, fmt.Sprintf("\tLinearAccelerationY:  %v", p.LinearAccelerationY))
	}

	if p.LinearAccelerationZ != 0 {
		lines = append(lines, fmt.Sprintf("\tLinearAccelerationZ:  %v", p.LinearAccelerationZ))
	}

	if p.AngularVelocityX != 0 {
		lines = append(lines, fmt.Sprintf("\tAngularVelocityX:     %v", p.AngularVelocityX))
	}

	if p.AngularVelocityY != 0 {
		lines = append(lines, fmt.Sprintf("\tAngularVelocityY:     %v", p.AngularVelocityY))
	}

	if p.AngularVelocityZ != 0 {
		lines = append(lines, fmt.Sprintf("\tAngularVelocityZ:     %v", p.AngularVelocityZ))
	}

	if len(lines) == 0 {
		return "IMUPacket{}"
	} else {
		return fmt.Sprintf("IMUPacket{\n%s,\n}", strings.Join(lines, ",\n"))
	}
}

type PacketDecodeError struct {
	error
}

func (p IMUPacket) UnmarshalBinary(data []byte) error {
	var err error

	if p.DiagnosticSystemTime, err = buffergenerics.ReadOrderedT[uint64](data, 0, decode.ByteOrder); err != nil {
		return PacketDecodeError{err}
	}

	if p.AccelerometerTime, err = buffergenerics.ReadOrderedT[uint64](data, 8, decode.ByteOrder); err != nil {
		return PacketDecodeError{err}
	}

	if p.GyroscopeTime, err = buffergenerics.ReadOrderedT[uint64](data, 16, decode.ByteOrder); err != nil {
		return PacketDecodeError{err}
	}

	if p.LinearAccelerationX, err = buffergenerics.ReadOrderedT[float32](data, 24, decode.ByteOrder); err != nil {
		return PacketDecodeError{err}
	}

	if p.LinearAccelerationY, err = buffergenerics.ReadOrderedT[float32](data, 28, decode.ByteOrder); err != nil {
		return PacketDecodeError{err}
	}

	if p.LinearAccelerationZ, err = buffergenerics.ReadOrderedT[float32](data, 32, decode.ByteOrder); err != nil {
		return PacketDecodeError{err}
	}

	if p.AngularVelocityX, err = buffergenerics.ReadOrderedT[float32](data, 36, decode.ByteOrder); err != nil {
		return PacketDecodeError{err}
	}

	if p.AngularVelocityY, err = buffergenerics.ReadOrderedT[float32](data, 40, decode.ByteOrder); err != nil {
		return PacketDecodeError{err}
	}

	if p.AngularVelocityZ, err = buffergenerics.ReadOrderedT[float32](data, 44, decode.ByteOrder); err != nil {
		return PacketDecodeError{err}
	}

	return nil
}

// LinearAcceleration returns the linear acceleration values [x, y, z] of the IMUPacket.
func (p IMUPacket) LinearAcceleration() [3]float32 {
	return [3]float32{p.LinearAccelerationX, p.LinearAccelerationY, p.LinearAccelerationZ}
}

// AngularVelocity returns the angular velocity values [x, y, z] of the IMUPacket.
func (p IMUPacket) AngularVelocity() [3]float32 {
	return [3]float32{p.AngularVelocityX, p.AngularVelocityY, p.AngularVelocityZ}
}
