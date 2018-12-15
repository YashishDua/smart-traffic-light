package endpoints

import (
  //"time"
	//"github.com/go-chi/chi"
	//"fmt"
  "net/http"
	"encoding/json"
  "smart-signal/util"
)

type TrafficData struct {
    StartNode     					string  `json:"start_node"`
    EndNode       		 			string  `json:"end_node"`
  	WithTrafficTime    			int64   `json:"with_traffic_time"`
    WithoutTrafficTime 			int64  	`json:"without_traffic_time"`
  	DefaultRedSignalTime  	int64  	`json:"default_red_signal_time"`
    DefaultGreenSignalTime  int64  	`json:"default_green_signal_time"`
		Direction					 			string  `json:"direction"`
}

type CongestionData struct {
		Node 							string `json:"node"`
		DeltaTime 				int64 `json:"delta_time"`
		RedLightTime 			int64 `json:"red_light_time"`
		GreenLightTime 			int64 `json:"green_light_time"`
		CongestionFactor 	float64 `json:"congestion_factor"`
}

// seconds
var MINIMUM_RED_LIGHT_TIME int64 = 30
var MAXIMUM_RED_LIGHT_TIME int64 = 150

func GetTrafficData(r *http.Request) (interface{}, *util.HTTPError) {
  decoder := json.NewDecoder(r.Body)
  var trafficDataArray []TrafficData
  err := decoder.Decode(&trafficDataArray)
  if err != nil {
      return nil, util.BadRequest("You missed something!")
  }
	var congestionData []CongestionData
	for _, element := range trafficDataArray {
	  delta := element.WithTrafficTime - element.WithoutTrafficTime
		congestionFactor := (float64(delta)/float64(element.WithoutTrafficTime))
		var c CongestionData
		c.CongestionFactor = congestionFactor
		c.Node = element.EndNode
		c.RedLightTime = element.DefaultRedSignalTime
		c.GreenLightTime = element.DefaultGreenSignalTime
		congestionData = append(congestionData, c)
    // index is the index where we are
    // element is the element from someSlice for where we are
	}

	for index, element := range congestionData {
		if element.CongestionFactor > 0.5 {
			var time int64
			for _, x := range congestionData {
				if x.Node == element.Node {
					continue
				}
				if element.CongestionFactor > x.CongestionFactor {
					if (x.CongestionFactor <= 0.4) {
						time -= 3
					}
					if (x.CongestionFactor > 0.4 && x.CongestionFactor <= 0.6) {
						time -= 2
					}
					if (x.CongestionFactor > 0.6) {
						time -= 1
					}
				}
				congestionData[index].DeltaTime = time
				var temp = congestionData[index].RedLightTime + time
				if temp > MINIMUM_RED_LIGHT_TIME {
					congestionData[index].RedLightTime += time
				}
				if temp <= MINIMUM_RED_LIGHT_TIME {
					congestionData[index].RedLightTime = MINIMUM_RED_LIGHT_TIME
				}
				var greenTemp = congestionData[index].GreenLightTime - time
				if greenTemp >= MAXIMUM_RED_LIGHT_TIME {
					congestionData[index].GreenLightTime = MAXIMUM_RED_LIGHT_TIME
				}
				if greenTemp < MAXIMUM_RED_LIGHT_TIME {
					congestionData[index].GreenLightTime -= time
				}
			}
		}
		if element.CongestionFactor <= 0.5 {
			congestionData[index].DeltaTime = 0
		}
	}

  return congestionData, nil
}
