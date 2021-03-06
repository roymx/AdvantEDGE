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
 * AdvantEDGE Location Service REST API
 *
 * Location Service is AdvantEDGE's implementation of [ETSI MEC ISG MEC013 Location API](http://www.etsi.org/deliver/etsi_gs/MEC/001_099/013/01.01.01_60/gs_mec013v010101p.pdf) <p>The API is based on the Open Mobile Alliance's specification RESTful Network API for Zonal Presence <p>[Copyright (c) ETSI 2017](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-loc-serv](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-loc-serv) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about Users (UE) and Zone locations <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address:30000/api_ <p>**Default Port**<br>`30007`
 *
 * API version: 1.1.1
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

// A type containing user information.
type UserInfo struct {

	// Address of user (e.g. \"sip\" URI, \"tel\" URI, \"acr\" URI).
	Address string `json:"address"`

	// Identifier of access point, <E-CGI><Cell Portion ID> (reference ETSI TS 129 171). Where the E-CGI is made up of the PLMN and Cell Identity (28 bit string). Then the PLMN is made up of the 3 digit MCC & 2 or 3 digit MNC. The Cell Portion is an optional element
	AccessPointId string `json:"accessPointId"`

	// Identifier of zone
	ZoneId string `json:"zoneId"`

	// Self referring URL.
	ResourceURL string `json:"resourceURL"`

	LocationInfo *LocationInfo `json:"locationInfo,omitempty"`

	// Contextual information of a user location (e.g., aisle, floor, room number, etc.)
	ContextLocationInfo string `json:"contextLocationInfo,omitempty"`

	// Reserved for future use.
	AncillaryInfo string `json:"ancillaryInfo,omitempty"`
}
