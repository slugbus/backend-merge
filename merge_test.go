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
	testBus2Before := taps.Bus{ID: "Bus two", Lat: 80.0, Lon: 80.0, Type: "type"}
	testBus3Before := taps.Bus{ID: "Bus three", Lat: 90.0, Lon: 90.0, Type: "type"}

	testBus1After := taps.Bus{ID: "Bus One", Lat: 120.0, Lon: 120.0, Type: "type"}
	testBus2After := taps.Bus{ID: "Bus two", Lat: 160.0, Lon: 160.0, Type: "type"}
	testBus3After := taps.Bus{ID: "Bus three", Lat: 180.0, Lon: 180.0, Type: "type"}
	// UpdatedBus structs
	updatedTestBus1 := taps.UpdatedBus{Bus: testBus1Before, Speed: 30.0, Angle: 30.0}
	updatedTestBus2 := taps.UpdatedBus{Bus: testBus2Before, Speed: 10.0, Angle: 60.0}
	updatedTestBus3 := taps.UpdatedBus{Bus: testBus3Before, Speed: 15.0, Angle: 15.0}
	// Map of Bus structs
	busMap := taps.BusMap{}
	// Hashing bus structs into map
	busMap[testBus1After.ID] = testBus1After
	busMap[testBus2After.ID] = testBus2After
	busMap[testBus3After.ID] = testBus3After
	// Map of updatedBus structs
	updatedBusMap := taps.UpdatedBusMap{}
	// Hashing updatedBus Structs into map
	updatedBusMap[updatedTestBus1.ID] = updatedTestBus1
	updatedBusMap[updatedTestBus2.ID] = updatedTestBus2
	updatedBusMap[updatedTestBus3.ID] = updatedTestBus3

	// Call mergeWithState function
	mergedUpdatedBusMap := mergeWithState(busMap, 30.0, updatedBusMap)
	// print results to console
	fmt.Printf("ID: %s\nLat: %f Expected: 120\nLon: %f Expected: 120\nSpeed: %f Expected: 100\nAngle: %f Expected: 3000\n", mergedUpdatedBusMap["Bus One"].ID, mergedUpdatedBusMap["Bus One"].Lat,
		mergedUpdatedBusMap["Bus One"].Lon, mergedUpdatedBusMap["Bus One"].Speed, mergedUpdatedBusMap["Bus One"].Angle)
	fmt.Printf("ID: %s\nLat: %f Expected: 160\nLon: %f Expected: 160\nSpeed: %f Expected: 100\nAngle: %f Expected: 3000\n", mergedUpdatedBusMap["Bus two"].ID, mergedUpdatedBusMap["Bus two"].Lat,
		mergedUpdatedBusMap["Bus two"].Lon, mergedUpdatedBusMap["Bus two"].Speed, mergedUpdatedBusMap["Bus two"].Angle)
	fmt.Printf("ID: %s\nLat: %f Expected: 180\nLon: %f Expected: 180\nSpeed: %f Expected: 100\nAngle: %f Expected: 3000\n", mergedUpdatedBusMap["Bus three"].ID, mergedUpdatedBusMap["Bus three"].Lat,
		mergedUpdatedBusMap["Bus three"].Lon, mergedUpdatedBusMap["Bus three"].Speed, mergedUpdatedBusMap["Bus three"].Angle)

}
