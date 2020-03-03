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

// Replay status object
type ReplayStatus struct {
	// Name of the replay file currently running
	ReplayFileRunning string `json:"replayFileRunning,omitempty"`
	// Index of the last ran event
	Index int32 `json:"index,omitempty"`
	// Number of events in the replay file
	MaxIndex int32 `json:"maxIndex,omitempty"`
	// Loop mode enables
	LoopMode bool `json:"loopMode,omitempty"`
	// Total time remaining for the replay file after last event
	TimeRemaining int32 `json:"timeRemaining,omitempty"`
	// Time remaining until the next event for the replay file after last event
	TimeToNextEvent int32 `json:"timeToNextEvent,omitempty"`
}
