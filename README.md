# IoT sensor simulator

Many industrial software solutions such as real-time analytics systems, SCADA, and Industrial IoT Alert Systems use data from IoT sensors. Development of such systems requires setting up testing environments that simulate real-life behavior of these sensors.

Deploying a software development team to a factory is in most cases not a viable option, since factories can be distant and working directly with sensors during the development process can interfere with or even disrupt factory operations.
## Use Case

This solution was developed to simulate smart sensors that send data in JSON format, particularly industrial pressure, temperature, and voltage sensors. It can be easily extended to other types of sensors by adding sensor parameters in `sensor/sensors`.
