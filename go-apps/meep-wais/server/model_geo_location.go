/*
 * Copyright (c) 2020  InterDigital Communications, Inc
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
 * AdvantEDGE WLAN Access Information Service REST API
 *
 * WLAN Access Information Service is AdvantEDGE's implementation of [ETSI MEC ISG MEC028 WAI API](http://www.etsi.org/deliver/etsi_gs/MEC/001_099/028/02.01.01_60/gs_MEC028v020101p.pdf) <p>[Copyright (c) ETSI 2020](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-wais](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-wais) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about WLAN access information in the network <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address/api_
 *
 * API version: 2.1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type GeoLocation struct {

	// Uncertainty for Latitude information.
	LatUncertainty int32 `json:"latUncertainty"`

	// Latitude value of the location.
	Lat int32 `json:"lat"`

	// Uncertainty for Longitude information.
	LongUncertainty int32 `json:"longUncertainty"`

	// Longitude value of the location.
	Long int32 `json:"long"`

	// Type description for altitude information.
	AltitudeType int32 `json:"altitudeType,omitempty"`

	// Uncertainty for altitude information.
	AltitudeUncertainty int32 `json:"altitudeUncertainty,omitempty"`

	// Altitude value of the location.
	Altitude int32 `json:"altitude,omitempty"`

	// Datum value to express how coordinates are organized.
	Datum int32 `json:"datum"`
}