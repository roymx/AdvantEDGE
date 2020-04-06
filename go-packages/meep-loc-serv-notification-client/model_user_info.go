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
 * AdvantEDGE Location Service Subscription Notification REST API
 *
 * This API enables the Location Service to post location events to edge applications <p>**Micro-service**<br>None <p>**Type & Usage**<br>User's Edge Applications subscribing to Location events must implement this API <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address/api_ <p>**Note**<br>This API is not exposed by default on the AdvantEDGE platform
 *
 * API version: 1.0.0
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type UserInfo struct {
	// Unique identifier of a user or client node
	UserId string `json:"userId,omitempty"`
	// Address of user (e.g. \"sip\" URI, \"tel\" URI, \"acr\" URI).
	Address string `json:"address"`
	// Identifier of access point, <E-CGI><Cell Portion ID> (reference ETSI TS 129 171). Where the E-CGI is made up of the PLMN and Cell Identity (28 bit string). Then the PLMN is made up of the 3 digit MCC & 2 or 3 digit MNC. The Cell Portion is an optional element
	AccessPointId string `json:"accessPointId"`
	// Identifier of zone
	ZoneId string `json:"zoneId"`
	// URL where the UE may be contacted for some service (application specific)
	ResourceURL string `json:"resourceURL"`
}
