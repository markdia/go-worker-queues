package main

/*
Pulling in the binary files into memory

per BLambert's email
uint8_t     version;
uint8_t     flags;
uint64_t    timestamp;
uint16_t    gsr;
uint8_t     temperature;
float_t     accelerometerMagnitude;
uint16_t    heartRate;
float_t     rrIntervals[5];
*/

type WhoopStrapData struct {
	Version                uint8
	Flags                  uint8
	Timestamp              uint64
	Gsr                    uint16
	Temperature            uint8
	AccelerometerMagnitude float32
	HeartRate              uint16
	RRIntervals            [5]float32
}

type WhoopStrapDataRows []WhoopStrapData
