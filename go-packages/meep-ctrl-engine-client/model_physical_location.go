/*
 * Copyright (c) 2019  InterDigital Communications, Inc
 *
 * Licensed under the Apache License, Version 2.0 (the \"License\");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an \"AS IS\" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * AdvantEDGE Platform Controller REST API
 *
 * This API is the main platform API and mainly used by the AdvantEDGE frontend to interact with scenarios <p>**Micro-service**<br>[meep-ctrl-engine](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-ctrl-engine) <p>**Type & Usage**<br>Platform main interface used by controller software that want to interact with the AdvantEDGE platform <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address:30000/api_ <p>**Default Port**<br>`30000`
 *
 * API version: 1.0.0
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Physical location object
type PhysicalLocation struct {
	// Unique physical location ID
	Id string `json:"id,omitempty"`
	// Physical location name
	Name string `json:"name,omitempty"`
	// Physical location type
	Type_ string `json:"type,omitempty"`
	// true: Physical location is external to MEEP false: Physical location is internal to MEEP
	IsExternal              bool     `json:"isExternal,omitempty"`
	NetworkLocationsInRange []string `json:"networkLocationsInRange,omitempty"`
	// Key/Value Pair Map (string, string)
	Meta map[string]string `json:"meta,omitempty"`
	// Key/Value Pair Map (string, string)
	UserMeta  map[string]string `json:"userMeta,omitempty"`
	Processes []Process         `json:"processes,omitempty"`
	// Latency in ms between the physical location and the network (wired interface, air interface)
	LinkLatency int32 `json:"linkLatency,omitempty"`
	// Latency variation in ms between the physical location and the network (wired interface, air interface)
	LinkLatencyVariation int32 `json:"linkLatencyVariation,omitempty"`
	// The limit of the traffic supported between the physical location and the network (wired interface, air interface)
	LinkThroughput int32 `json:"linkThroughput,omitempty"`
	// Packet lost (in terms of percentage) between the physical location and the network (wired interface, air interface)
	LinkPacketLoss float64 `json:"linkPacketLoss,omitempty"`
}
