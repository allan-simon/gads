package gads

import (
	"encoding/xml"
	"fmt"
)

type CampaignExtensionSettingService struct {
	Auth
}

type CampaignExtensionSetting struct {
	CampaignID       int64             `xml:"campaignId,omitempty"`
	ExtensionType    string            `xml:"extensionType"`
	ExtensionSetting *ExtensionSetting `xml:"extensionSetting"`
}

type ExtensionSetting struct {
	Extensions            []ExtensionFeedItem `xml:"extensions"`
	PlateformRestrictions string              `xml:"platformRestrictions"`
}

type CampaignExtensionSettingOperations map[string][]CampaignExtensionSetting

func NewCampaignExtensionSettingService(auth *Auth) *CampaignExtensionSettingService {
	return &CampaignExtensionSettingService{Auth: *auth}
}

func (s *CampaignExtensionSettingService) Get(selector Selector) (
	extensionSettings []ExtensionSetting,
	totalCount int64,
	err error,
) {
	selector.XMLName = xml.Name{"", "selector"}
	respBody, err := s.Auth.request(
		campaignExtensionSettingServiceUrl,
		"get",
		struct {
			XMLName xml.Name
			Sel     Selector
		}{
			XMLName: xml.Name{
				Space: baseUrl,
				Local: "get",
			},
			Sel: selector,
		},
	)

	if err != nil {
		return extensionSettings, totalCount, err
	}
	getResp := struct {
		Size              int64              `xml:"rval>totalNumEntries"`
		ExtensionSettings []ExtensionSetting `xml:"rval>entries"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return extensionSettings, totalCount, err
	}
	return getResp.ExtensionSettings, getResp.Size, err

}

//
//
//
func (s *CampaignExtensionSettingService) Mutate(
	campaignExtensionSettingOperations CampaignExtensionSettingOperations,
) (campaignExtensionSettings []CampaignExtensionSetting, err error) {
	type operation struct {
		Action                   string                   `xml:"operator"`
		CampaignExtensionSetting CampaignExtensionSetting `xml:"operand"`
	}
	operations := []operation{}
	for action, campaignExtensionSettings := range campaignExtensionSettingOperations {
		for _, campaignExtensionSetting := range campaignExtensionSettings {
			operations = append(
				operations,
				operation{
					Action: action,
					CampaignExtensionSetting: campaignExtensionSetting,
				},
			)
		}
	}
	mutation := struct {
		XMLName xml.Name
		Ops     []operation `xml:"operations"`
	}{
		XMLName: xml.Name{
			Space: baseUrl,
			Local: "mutate",
		},
		Ops: operations,
	}
	respBody, err := s.Auth.request(
		campaignExtensionSettingServiceUrl,
		"mutate",
		mutation,
	)
	if err != nil {
		return campaignExtensionSettings, err
	}
	mutateResp := struct {
		CampaignExtensionSettings []CampaignExtensionSetting `xml:"rval>value"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &mutateResp)
	if err != nil {
		return campaignExtensionSettings, err
	}
	return mutateResp.CampaignExtensionSettings, err
}

// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201506/CampaignExtensionSettingService#query
//
func (s *CampaignExtensionSettingService) Query(query string) (campaignExtensionSettings []CampaignExtensionSetting, totalCount int64, err error) {

	respBody, err := s.Auth.request(
		adGroupServiceUrl,
		"query",
		AWQLQuery{
			XMLName: xml.Name{
				Space: baseUrl,
				Local: "query",
			},
			Query: query,
		},
	)

	if err != nil {
		return campaignExtensionSettings, totalCount, err
	}
	getResp := struct {
		Size                      int64                      `xml:"rval>totalNumEntries"`
		CampaignExtensionSettings []CampaignExtensionSetting `xml:"rval>entries"`
	}{}
	fmt.Printf("%s\n", respBody)
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return campaignExtensionSettings, totalCount, err
	}
	return getResp.CampaignExtensionSettings, getResp.Size, err

}
