package merge

import (
	measusurements "github.com/slugbus/backend-measurements"
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
			newUpdatedBus.Angle = measusurements.Direction(currentBusMap[key].Lat, currentBusMap[key].Lon, pingBus.Lat, pingBus.Lon)
		}

		// Always push the new bus into the new updated bus map
		newUpdatedBusMap[key] = newUpdatedBus
	}

	return newUpdatedBusMap
}

/* // Merge update takes in two regular responses from the server
// and t (that is in milliseconds) and combines them to get speed
// and angle data.
func mergeUpdate(p, q Bus, t float64) UpdatedBusMap {
	// Make of map of strings
	// to buses
	mb := map[string]taps.Bus{}
	// Loop through first
	// ping
	for _, bus := range p {
		// Map the bus ID to the
		// bus datastructure
		mb[bus.ID] = bus
	}
	// Prepare a result
	result := UpdatedBus{}
	// Loop through the second ping
	for _, pingTwoBus := range q {
		// Make a bus with angles and speed
		d := BusDataPlusPlus{}
		// Add the buses' data to the bus++?
		d.Bus = pingTwoBus
		// Check if the current bus exists in ping one
		if pingOneBus, contains := mb[d.ID]; contains {
			// If it does, calculate its distance, speed , and angle
			distance := geo.Dist(pingOneBus.Lat, pingOneBus.Lon, pingTwoBus.Lat, pingTwoBus.Lon)
			d.Speed = geo.Speed(distance, t)
			d.Angle = geo.Dir(pingOneBus.Lat, pingOneBus.Lon, pingTwoBus.Lat, pingTwoBus.Lon)
		}
		// push the bus to the result
		result = append(result, d)
	}
	return result
}
*/
