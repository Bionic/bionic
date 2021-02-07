package health

import "encoding/xml"

func (p *health) parseActivitySummary(data *Data, decoder *xml.Decoder, start *xml.StartElement) error {
	var activitySummary ActivitySummary

	if err := decoder.DecodeElement(&activitySummary, start); err != nil {
		return err
	}

	data.ActivitySummaries = append(data.ActivitySummaries, &activitySummary)

	return p.DB().
		FirstOrCreate(&activitySummary, activitySummary.Constraints()).
		Error
}