package merge

import (
	"fmt"
)

func main() {
	testBus1 := Bus{"Bus One", 60.0, 60.0, "type"}
	updatedTestBus := UpdatedBus{testBus1, 30.0, 30.0}

	busOneMap := BusMap{}
	busOneMap[testBus1.ID] = testBus1
	updatedBusOneMap := UpdatedBusMap{}
	updatedBusOneMap[updatedTestBus.ID] = updatedTestBus

	merge.mergedUpdatedBusMap := mergeWithState(busOneMap, 30.0, updatedBusOneMap)

	fmt.Printf("ID: %s\n, Lat: %s\n, Lon: %s\n, Speed: %f\n, Angle: %f\n")

}
