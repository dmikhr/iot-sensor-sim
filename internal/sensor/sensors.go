package sensor

var data = []*Data{
	{SensorID: "P1", Parameter: ValueTypePressure, Unit: "Pa", Frequency: 10, Location: "input", Enabled: true},
	{SensorID: "P2", Parameter: ValueTypePressure, Unit: "Pa", Frequency: 10, Location: "emergency_valve", Enabled: true},
	{SensorID: "T1", Parameter: ValueTypeTemperature, Unit: "C", Frequency: 2, Location: "pump", Enabled: true},
	{SensorID: "V1", Parameter: ValueTypeVoltage, Unit: "V", Frequency: 5, Location: "pump", Enabled: true},
}

// Init - set up sensors for simulation
func Init() ([]*Sensor, []*Settings) {
	// define struct initial data for search sensor with settings, then split into 2 slices with filter by enabled, put into a separate file within package
	// map?
	// return filtered results
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
