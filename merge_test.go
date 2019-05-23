package merge

import (
	"fmt"
	"testing"

	"github.com/slugbus/taps"
)

func TestMergeWithState(*testing.T) {
	// dummy data used to test merging a BusMap with an UpdatedBus Map

	// Bus structs: Before will be in the current UpdatedBus map and after will be the new bus ping
	testBus1Before := taps.Bus{ID: "Bus One", Lat: 60.0, Lon: 60.0, Type: "type"}
	testBus2Before := taps.Bus{ID: "Bus two", Lat: 122.0, Lon: 122.0, Type: "type"}
	testBus3Before := taps.Bus{ID: "Bus three", Lat: 200.0, Lon: 200.0, Type: "type"}

	testBus2After := taps.Bus{ID: "Bus two", Lat: 71.0, Lon: 71.0, Type: "type"}
	testBus3After := taps.Bus{ID: "Bus three", Lat: 125.0, Lon: 132.0, Type: "type"}
	testBus4After := taps.Bus{ID: "Bus four", Lat: 200.0, Lon: 200.0, Type: "type"}

	// UpdatedBus structs
	updatedTestBus1 := taps.UpdatedBus{Bus: testBus1Before, Speed: 30.0, Angle: 30.0}
	updatedTestBus2 := taps.UpdatedBus{Bus: testBus2Before, Speed: 10.0, Angle: 60.0}
	updatedTestBus3 := taps.UpdatedBus{Bus: testBus3Before, Speed: 15.0, Angle: 15.0}
	// Map of Bus structs
	busMap := taps.BusMap{}
	// Hashing bus structs into map
	busMap[testBus2After.ID] = testBus2After
	busMap[testBus3After.ID] = testBus3After
	busMap[testBus4After.ID] = testBus4After
	// Map of updatedBus structs
	updatedBusMap := taps.UpdatedBusMap{}
	// Hashing updatedBus Structs into map
	updatedBusMap[updatedTestBus1.ID] = updatedTestBus1
	updatedBusMap[updatedTestBus2.ID] = updatedTestBus2
	updatedBusMap[updatedTestBus3.ID] = updatedTestBus3

	// Call mergeWithState function
	mergedUpdatedBusMap := mergeWithState(busMap, 30.0, updatedBusMap)
	// print results to console
	for k := range mergedUpdatedBusMap {
		fmt.Printf("ID: %s\nLat: %f\nLon: %f\nSpeed: %f\nAngle: %f\n", mergedUpdatedBusMap[k].ID, mergedUpdatedBusMap[k].Lat,
			mergedUpdatedBusMap[k].Lon, mergedUpdatedBusMap[k].Speed, mergedUpdatedBusMap[k].Angle)
	}

}
