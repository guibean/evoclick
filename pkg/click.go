package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/EricFrancis12/evoclick/prisma/db"
	"github.com/google/uuid"
)

func (s *Storer) GetClickById(ctx context.Context, id int) (Click, error) {
	model, err := s.Client.Click.FindUnique(
		db.Click.ID.Equals(id),
	).Exec(ctx)
	if err != nil {
		return Click{}, err
	}
	return formatClick(model), nil
}

func (s *Storer) GetClickByPublicId(ctx context.Context, publicId string) (Click, error) {
	model, err := s.Client.Click.FindUnique(
		db.Click.PublicID.Equals(publicId),
	).Exec(ctx)
	if err != nil {
		return Click{}, err
	}
	return formatClick(model), nil
}

func (s *Storer) CreateNewClick(ctx context.Context, creationReq ClickCreationReq) (Click, error) {
	clickTokensStr := marshallTokens(creationReq.Tokens)
	optParams := makeOptParams(creationReq)

	model, err := s.Client.Click.CreateOne(
		// Mandatory parameters:
		db.Click.PublicID.Set(creationReq.PublicId),
		db.Click.ExternalID.Set(creationReq.ExternalId),
		db.Click.Cost.Set(creationReq.Cost),
		db.Click.Revenue.Set(creationReq.Revenue),
		db.Click.ViewTime.Set(creationReq.ViewTime),
		db.Click.ViewOutputURL.Set(creationReq.ViewOutputURL),
		db.Click.Tokens.Set(clickTokensStr),
		db.Click.IP.Set(creationReq.IP),
		db.Click.UserAgent.Set(creationReq.UserAgent),
		db.Click.Language.Set(creationReq.Language),
		db.Click.DeviceType.Set(creationReq.DeviceType),
		db.Click.Device.Set(creationReq.Device),
		db.Click.ScreenResolution.Set(creationReq.ScreenResolution),
		db.Click.Os.Set(creationReq.Os),
		db.Click.OsVersion.Set(creationReq.OsVersion),
		db.Click.BrowserName.Set(creationReq.BrowserName),
		db.Click.BrowserVersion.Set(creationReq.BrowserVersion),
		db.Click.CampaignID.Set(creationReq.CampaignID),
		db.Click.TrafficSourceID.Set(creationReq.TrafficSourceID),
		// Optional parameters:
		optParams...,
	).Exec(ctx)
	if err != nil {
		return Click{}, err
	}

	return formatClick(model), nil
}

func (s *Storer) UpsertClickById(ctx context.Context, id int, click Click) (Click, error) {
	clickTokensStr := marshallTokens(click.Tokens)
	upsertParams := makeUpsertParams(click)

	model, err := s.Client.Click.UpsertOne(
		db.Click.ID.Equals(id),
	).Create(
		// Mandatory parameters:
		db.Click.PublicID.Set(click.PublicID),
		db.Click.ExternalID.Set(click.ExternalID),
		db.Click.Cost.Set(click.Cost),
		db.Click.Revenue.Set(click.Revenue),
		db.Click.ViewTime.Set(click.ViewTime),
		db.Click.ViewOutputURL.Set(click.ViewOutputURL),
		db.Click.Tokens.Set(clickTokensStr),
		db.Click.IP.Set(click.IP),
		db.Click.UserAgent.Set(click.UserAgent),
		db.Click.Language.Set(click.Language),
		db.Click.DeviceType.Set(click.DeviceType),
		db.Click.Device.Set(click.Device),
		db.Click.ScreenResolution.Set(click.ScreenResolution),
		db.Click.Os.Set(click.Os),
		db.Click.OsVersion.Set(click.OsVersion),
		db.Click.BrowserName.Set(click.BrowserName),
		db.Click.BrowserVersion.Set(click.BrowserVersion),
		db.Click.CampaignID.Set(click.CampaignID),
		db.Click.TrafficSourceID.Set(click.TrafficSourceID),
	).Update(
		upsertParams...,
	).Exec(ctx)
	if err != nil {
		return Click{}, err
	}
	return formatClick(model), nil
}

func formatClick(model *db.ClickModel) Click {
	clickTokens := parseClickTokens(model.Tokens)
	return Click{
		InnerClick: model.InnerClick,
		Tokens:     clickTokens,
	}
}

func makeOptParams(cp ClickCreationReq) []db.ClickSetParam {
	optParams := []db.ClickSetParam{}
	// Optional parameters that CANNOT accept default values, so they should be ommitted if they are 0
	optParams = appendIfTrue(optParams, db.Click.ClickTime.Set(cp.ClickTime), !cp.ClickTime.IsZero())
	optParams = appendIfTrue(optParams, db.Click.ConvTime.Set(cp.ConvTime), !cp.ConvTime.IsZero())
	optParams = appendIfTrue(optParams, db.Click.ClickOutputURL.Set(cp.ClickOutputURL), cp.ClickOutputURL != "")
	optParams = appendIfTrue(optParams, db.Click.Isp.Set(cp.Isp), cp.Isp != "")
	optParams = appendIfTrue(optParams, db.Click.Country.Set(cp.Country), cp.Country != "")
	optParams = appendIfTrue(optParams, db.Click.Region.Set(cp.Region), cp.Region != "")
	optParams = appendIfTrue(optParams, db.Click.City.Set(cp.City), cp.City != "")
	optParams = appendIfTrue(optParams, db.Click.AffiliateNetworkID.Set(cp.AffiliateNetworkID), cp.AffiliateNetworkID != 0)
	optParams = appendIfTrue(optParams, db.Click.LandingPageID.Set(cp.LandingPageID), cp.LandingPageID != 0)
	optParams = appendIfTrue(optParams, db.Click.OfferID.Set(cp.OfferID), cp.OfferID != 0)
	optParams = appendIfTrue(optParams, db.Click.SavedFlowID.Set(cp.SavedFlowID), cp.SavedFlowID != 0)
	return optParams
}

func appendIfTrue(params []db.ClickSetParam, p db.ClickSetParam, condition bool) []db.ClickSetParam {
	if condition {
		params = append(params, p)
	}
	return params
}

func makeUpsertParams(click Click) []db.ClickSetParam {
	clickTokensStr := marshallTokens(click.Tokens)
	params := []db.ClickSetParam{
		// Mandatory parameters:
		db.Click.PublicID.Set(click.PublicID),
		db.Click.ExternalID.Set(click.ExternalID),
		db.Click.Cost.Set(click.Cost),
		db.Click.Revenue.Set(click.Revenue),
		db.Click.ViewTime.Set(click.ViewTime),
		db.Click.ViewOutputURL.Set(click.ViewOutputURL),
		db.Click.Tokens.Set(clickTokensStr),
		db.Click.IP.Set(click.IP),
		db.Click.UserAgent.Set(click.UserAgent),
		db.Click.Language.Set(click.Language),
		db.Click.DeviceType.Set(click.DeviceType),
		db.Click.Device.Set(click.Device),
		db.Click.ScreenResolution.Set(click.ScreenResolution),
		db.Click.Os.Set(click.Os),
		db.Click.OsVersion.Set(click.OsVersion),
		db.Click.BrowserName.Set(click.BrowserName),
		db.Click.BrowserVersion.Set(click.BrowserVersion),
		db.Click.CampaignID.Set(click.CampaignID),
		db.Click.TrafficSourceID.Set(click.TrafficSourceID),
	}

	// Optional parameters:
	optParams := makeOptParams(ClickCreationReq{
		ClickTime:          click.ClickTime,
		ConvTime:           click.ConvTime,
		ClickOutputURL:     click.ClickOutputURL,
		Isp:                click.Isp,
		Country:            click.Country,
		Region:             click.Region,
		City:               click.City,
		AffiliateNetworkID: click.AffiliateNetworkID,
		LandingPageID:      click.LandingPageID,
		OfferID:            click.OfferID,
	})
	params = append(params, optParams...)
	return params
}

func parseClickTokens(jsonStr string) []Token {
	clickTokens, err := ParseJSON[[]Token](jsonStr)
	if err != nil {
		return []Token{}
	}
	return clickTokens
}

func marshallTokens(tokens []Token) string {
	jsonData, err := json.Marshal(tokens)
	if err != nil {
		fmt.Printf("error marshalling to JSON: %s", err)
		return "[]"
	}
	return string(jsonData)
}

func NewPublicClickID() string {
	return uuid.New().String()
}

type PostbackMatcherMap = map[PostbackURLMatcher]string

func newPostbackMatcherMap(click Click) PostbackMatcherMap {
	return PostbackMatcherMap{
		PostbackURLMatcherID:                 strconv.Itoa(click.ID),
		PostbackURLMatcherPublicID:           click.PublicID,
		PostbackURLMatcherExternalID:         click.ExternalID,
		PostbackURLMatcherCost:               strconv.Itoa(click.Cost),
		PostbackURLMatcherRevenue:            strconv.Itoa(click.Revenue),
		PostbackURLMatcherViewTime:           timeString(click.ViewTime),
		PostbackURLMatcherClickTime:          timeString(click.ClickTime),
		PostbackURLMatcherConvTime:           timeString(click.ConvTime),
		PostbackURLMatcherViewOutputUrl:      click.ViewOutputURL,
		PostbackURLMatcherClickOutputUrl:     click.ClickOutputURL,
		PostbackURLMatcherIP:                 click.IP,
		PostbackURLMatcherIsp:                click.Isp,
		PostbackURLMatcherUserAgent:          click.UserAgent,
		PostbackURLMatcherLanguage:           click.Language,
		PostbackURLMatcherCountry:            click.Country,
		PostbackURLMatcherRegion:             click.Region,
		PostbackURLMatcherCity:               click.City,
		PostbackURLMatcherDeviceType:         click.DeviceType,
		PostbackURLMatcherDevice:             click.Device,
		PostbackURLMatcherScreenResolution:   click.ScreenResolution,
		PostbackURLMatcherOs:                 click.Os,
		PostbackURLMatcherOsVersion:          click.OsVersion,
		PostbackURLMatcherBrowserName:        click.BrowserName,
		PostbackURLMatcherBrowserVersion:     click.BrowserVersion,
		PostbackURLMatcherCreatedAt:          timeString(click.CreatedAt),
		PostbackURLMatcherUpdatedAt:          timeString(click.UpdatedAt),
		PostbackURLMatcherAffiliateNetworkID: strconv.Itoa(click.AffiliateNetworkID),
		PostbackURLMatcherCampaignID:         strconv.Itoa(click.CampaignID),
		PostbackURLMatcherFlowID:             strconv.Itoa(click.SavedFlowID),
		PostbackURLMatcherLandingPageID:      strconv.Itoa(click.LandingPageID),
		PostbackURLMatcherOfferID:            strconv.Itoa(click.OfferID),
		PostbackURLMatcherTrafficSourceID:    strconv.Itoa(click.TrafficSourceID),
	}
}
