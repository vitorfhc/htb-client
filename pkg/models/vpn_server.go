package models

import "strings"

type deepOption struct {
	Servers map[string]*VPNServer `json:"servers"`
}

type VPNServersData struct {
	Assigned *VPNServer                        `json:"assigned"`
	Options  map[string]map[string]*deepOption `json:"options"`
}

type VPNServersList []*VPNServer

type VPNServer struct {
	ID           uint   `json:"id"`
	FriendlyName string `json:"friendly_name"`
	Location     string `json:"location"`
}

func (s *VPNServer) SubscriptionType() HTBLabSubscription {
	var subscriptionType HTBLabSubscription
	friendlyName := strings.ToLower(s.FriendlyName)
	if strings.Contains(friendlyName, string(HTBLabSubscriptionFree)) {
		subscriptionType = HTBLabSubscriptionFree
	} else if strings.Contains(friendlyName, string(HTBLabSubscriptionVIPPlus)) {
		subscriptionType = HTBLabSubscriptionVIPPlus
	} else if strings.Contains(friendlyName, string(HTBLabSubscriptionVIP)) {
		subscriptionType = HTBLabSubscriptionVIP
	} else {
		panic("unknown subscription type")
	}

	return subscriptionType
}
