/*
 * ETSI GS MEC 028 - WLAN Access Information API
 *
 * WLAN Access Information Service is AdvantEDGE's implementation of [ETSI MEC ISG MEC028 WAI API](http://www.etsi.org/deliver/etsi_gs/MEC/001_099/028/02.01.01_60/gs_MEC028v020101p.pdf) <p>[Copyright (c) ETSI 2020](https://forge.etsi.org/etsi-forge-copyright-notice.txt) <p>**Micro-service**<br>[meep-wais](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-wais) <p>**Type & Usage**<br>Edge Service used by edge applications that want to get information about WLAN access information in the network <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address/api_
 *
 * API version: 2.1.1
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

type StaDataRateSubscription struct {
	Links *AssocStaSubscriptionLinks `json:"_links,omitempty"`
	// URI selected by the service consumer to receive notifications on the subscribed WLAN Access Information Service. This shall be included both in the request and in response.
	CallbackReference string `json:"callbackReference"`

	ExpiryDeadline *TimeStamp `json:"expiryDeadline,omitempty"`
	// Identifier(s) to uniquely specify the target client station(s) for the subscription
	StaId []StaIdentity `json:"staId"`
	// Shall be set to \"StaDataRateSubscription\".
	SubscriptionType string `json:"subscriptionType"`
}
