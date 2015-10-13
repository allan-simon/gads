package gads

type CommonExtensionFeedItem struct {
	Type       string  `xml:"xsi:type,attr,omitempty"`
	FeedID     int64   `xml:"feedId,omitempty"`
	FeedItemID int64   `xml:"feedItemId,omitempty"`
	Status     string  `xml:"status,omitempty"`
	FeedType   string  `xml:"feedType,omitempty"`
	StartTime  string  `xml:"startTime,omitempty"`
	EndTime    *string `xml:"endTime"`
	// devicePreference
	// scheduling
	// policyData
	// end common
}

type ExtensionFeedItem interface {
	GetFeedID() int64
	GetFeedItemID() int64
}

func (e *CommonExtensionFeedItem) GetFeedID() int64 {
	return e.FeedID
}

func (e *CommonExtensionFeedItem) GetFeedItemID() int64 {
	return e.FeedItemID
}

type SitelinkFeedItem struct {
	*CommonExtensionFeedItem
	Text                string   `xml:"sitelinkText"`
	FinalUrls           *UrlList `xml:"sitelinkFinalUrls"`
	TrackingURLTemplate *string  `xml:"sitelinkTrackingUrlTemplate"`
}
