package merge

import (
	"fmt"

	measurements "github.com/slugbus/backend-measurements"
	"github.com/slugbus/taps"
)

// This function merges a new ping response with the current
// response
func mergeWithState(newPing taps.BusMap, time float64, currentBusMap taps.UpdatedBusMap) taps.UpdatedBusMap {
	// Prepare the new state.
	newUpdatedBusMap := taps.UpdatedBusMap{}

	for key, pingBus := range newPing {
		// Prepare a new UpdatedBus struct for each bus in the new ping
		newUpdatedBus := taps.UpdatedBus{
			Bus: pingBus,
		}
		// Check to see if the bus we're looking
		// at was in our previous state, if it was
		// update it with speed and angle value
		//isBusStillRunning = currentBusMap[key] // is the key in the existing UpdatedBusMap
		if _, exists := currentBusMap[key]; exists {
			//testing
			//newUpdatedBus.Speed = 100.0
			//newUpdatedBus.Angle = 3000.0
			//distance := measusurements.getDistance(currentBusMap[key].Lat, currentBusMap[key].Lon, newUpdatedBus.Lat, newUpdatedBus.Lon)
			//newUpdatedBus.Speed = measusurements.Speed(distance, time)
			fmt.Printf("CurrentBusMap bus Lat: %f   Lon: %f\n", currentBusMap[key].Lat, currentBusMap[key].Lon)
			fmt.Printf("Ping bus Lat: %f   Lon: %f\n", pingBus.Lat, pingBus.Lon)
			newUpdatedBus.Angle = measurements.Angle(currentBusMap[key].Lat, currentBusMap[key].Lon, pingBus.Lat, pingBus.Lon)
		} else {
			newUpdatedBus.Angle = 30.0
		}

		// Always push the new bus into the new updated bus map
		newUpdatedBusMap[key] = newUpdatedBus
	}

	return newUpdatedBusMap
}
