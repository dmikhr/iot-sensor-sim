// Package sensor provides functions for working with sensor.
package sensor

var data = []*Data{
	{SensorID: "P1", Parameter: ValueTypePressure, Unit: "Pa", Frequency: 10, Location: "input", Enabled: true},
	{SensorID: "P2", Parameter: ValueTypePressure, Unit: "Pa", Frequency: 10, Location: "emergency_valve", Enabled: true},
	{SensorID: "T1", Parameter: ValueTypeTemperature, Unit: "C", Frequency: 2, Location: "pump", Enabled: true},
	{SensorID: "V1", Parameter: ValueTypeVoltage, Unit: "V", Frequency: 5, Location: "pump", Enabled: true},
}

// Init sets up sensors for simulation.
func Init() ([]*Sensor, []*Settings) {
	var sensors []*Sensor
	var sensorsSettings []*Settings
	for _, sensor := range data {
		if sensor.Enabled {
			sensors = append(sensors, New(sensor.SensorID, sensor.Parameter, sensor.Unit))
			sensorsSettings = append(sensorsSettings,
				NewSettings(sensor.SensorID, sensor.Frequency, sensor.Location))
		}

	}
	return sensors, sensorsSettings
}
