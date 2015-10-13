package gads

import (
	"testing"
	//	"encoding/xml"
)

func testAdGroupCriterionService(t *testing.T) (service *AdGroupCriterionService) {
	return &AdGroupCriterionService{Auth: testAuthSetup(t)}
}

func TestAdGroupCriterion(t *testing.T) {
	adGroup, cleanupAdGroup := testAdGroup(nil)
	defer cleanupAdGroup()

	agcs := testAdGroupCriterionService(t)
	adGroupCriterions, err := agcs.Mutate(
		AdGroupCriterionOperations{
			"ADD": {
				/*
									NegativeAdGroupCriterion{
					          AdGroupId: adGroup.Id,
					          Criterion: AgeRangeCriterion{AgeRangeType:"AGE_RANGE_25_34"},
					        },
									NegativeAdGroupCriterion{
					          AdGroupId: adGroup.Id,
					          Criterion: GenderCriterion{},
					        },
									NewBiddableAdGroupCriterion{
					          AdGroupId: adGroup.Id,
					          Criterion: MobileAppCategoryCriterion{
					            60000,"My Google Play Android Apps"
					          }
					        },
				*/
				BiddableAdGroupCriterion{
					Type:       "BiddableAdGroupCriterion",
					AdGroupId:  adGroup.Id,
					Criterion:  KeywordCriterion{Type: "Keyword", Text: "test1", MatchType: "EXACT"},
					UserStatus: "PAUSED",
				},
				BiddableAdGroupCriterion{
					Type:       "BiddableAdGroupCriterion",
					AdGroupId:  adGroup.Id,
					Criterion:  KeywordCriterion{Type: "Keyword", Text: "test2", MatchType: "PHRASE"},
					UserStatus: "PAUSED",
				},
				BiddableAdGroupCriterion{
					Type:       "BiddableAdGroupCriterion",
					AdGroupId:  adGroup.Id,
					Criterion:  KeywordCriterion{Type: "Keyword", Text: "test3", MatchType: "BROAD"},
					UserStatus: "PAUSED",
				},
				NegativeAdGroupCriterion{
					AdGroupId: adGroup.Id,
					Criterion: KeywordCriterion{Type: "Keyword", Text: "test4", MatchType: "BROAD"},
				},
				BiddableAdGroupCriterion{
					Type:       "BiddableAdGroupCriterion",
					AdGroupId:  adGroup.Id,
					Criterion:  PlacementCriterion{Type: "Placement", Url: "https://classdo.com"},
					UserStatus: "PAUSED",
				},
				// NewBiddableAdGroupCriterion(adGroup.Id, NewUserInterestCriterion()),
				// NewBiddableAdGroupCriterion(adGroup.Id, NewUserListCriterion()),
				// NewBiddableAdGroupCriterion(adGroup.Id, NewVerticalCriterion(0, 0, []string{"Pets & Anamals","Pets","Dogs"})),
				BiddableAdGroupCriterion{
					Type:      "BiddableAdGroupCriterion",
					AdGroupId: adGroup.Id,
					Criterion: WebpageCriterion{
						Type: "Webpage",
						Parameter: WebpageParameter{
							CriterionName: "test criterion",
							Conditions: []WebpageCondition{
								WebpageCondition{
									Operand:  "URL",
									Argument: "example.com",
								},
							},
						},
					},
					UserStatus: "PAUSED",
				},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", adGroupCriterions)

	defer func() {
		_, err = agcs.Mutate(AdGroupCriterionOperations{"REMOVE": adGroupCriterions})
		if err != nil {
			t.Error(err)
		}
	}()
	/*
	   reqBody, err := xml.MarshalIndent(adGroupCriterions,"  ", "  ")
	   t.Fatalf("%s\n",reqBody)
	*/
}
