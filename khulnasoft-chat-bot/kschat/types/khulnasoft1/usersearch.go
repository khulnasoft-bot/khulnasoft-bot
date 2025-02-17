// Auto-generated to Go types using avdl-compiler v1.4.10 (https://github.com/khulnasoft/node-avdl-compiler)
//   Input file: ../client/protocol/avdl/khulnasoft1/usersearch.avdl

package khulnasoft1

type APIUserServiceID string

func (o APIUserServiceID) DeepCopy() APIUserServiceID {
	return o
}

type APIUserKhulnasoftResult struct {
	Username   string  `codec:"username" json:"username"`
	Uid        UID     `codec:"uid" json:"uid"`
	PictureUrl *string `codec:"pictureUrl,omitempty" json:"picture_url,omitempty"`
	FullName   *string `codec:"fullName,omitempty" json:"full_name,omitempty"`
	RawScore   float64 `codec:"rawScore" json:"raw_score"`
	Stellar    *string `codec:"stellar,omitempty" json:"stellar,omitempty"`
	IsFollowee bool    `codec:"isFollowee" json:"is_followee"`
}

func (o APIUserKhulnasoftResult) DeepCopy() APIUserKhulnasoftResult {
	return APIUserKhulnasoftResult{
		Username: o.Username,
		Uid:      o.Uid.DeepCopy(),
		PictureUrl: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.PictureUrl),
		FullName: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.FullName),
		RawScore: o.RawScore,
		Stellar: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.Stellar),
		IsFollowee: o.IsFollowee,
	}
}

type APIUserServiceResult struct {
	ServiceName APIUserServiceID `codec:"serviceName" json:"service_name"`
	Username    string           `codec:"username" json:"username"`
	PictureUrl  string           `codec:"pictureUrl" json:"picture_url"`
	Bio         string           `codec:"bio" json:"bio"`
	Location    string           `codec:"location" json:"location"`
	FullName    string           `codec:"fullName" json:"full_name"`
	Confirmed   *bool            `codec:"confirmed,omitempty" json:"confirmed,omitempty"`
}

func (o APIUserServiceResult) DeepCopy() APIUserServiceResult {
	return APIUserServiceResult{
		ServiceName: o.ServiceName.DeepCopy(),
		Username:    o.Username,
		PictureUrl:  o.PictureUrl,
		Bio:         o.Bio,
		Location:    o.Location,
		FullName:    o.FullName,
		Confirmed: (func(x *bool) *bool {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.Confirmed),
	}
}

type APIUserServiceSummary struct {
	ServiceName APIUserServiceID `codec:"serviceName" json:"service_name"`
	Username    string           `codec:"username" json:"username"`
}

func (o APIUserServiceSummary) DeepCopy() APIUserServiceSummary {
	return APIUserServiceSummary{
		ServiceName: o.ServiceName.DeepCopy(),
		Username:    o.Username,
	}
}

type ImpTofuSearchResult struct {
	Assertion       string `codec:"assertion" json:"assertion"`
	AssertionValue  string `codec:"assertionValue" json:"assertionValue"`
	AssertionKey    string `codec:"assertionKey" json:"assertionKey"`
	Label           string `codec:"label" json:"label"`
	PrettyName      string `codec:"prettyName" json:"prettyName"`
	KhulnasoftUsername string `codec:"khulnasoftUsername" json:"khulnasoftUsername"`
}

func (o ImpTofuSearchResult) DeepCopy() ImpTofuSearchResult {
	return ImpTofuSearchResult{
		Assertion:       o.Assertion,
		AssertionValue:  o.AssertionValue,
		AssertionKey:    o.AssertionKey,
		Label:           o.Label,
		PrettyName:      o.PrettyName,
		KhulnasoftUsername: o.KhulnasoftUsername,
	}
}

type APIUserSearchResult struct {
	Score           float64                                    `codec:"score" json:"score"`
	Khulnasoft         *APIUserKhulnasoftResult                      `codec:"khulnasoft,omitempty" json:"khulnasoft,omitempty"`
	Service         *APIUserServiceResult                      `codec:"service,omitempty" json:"service,omitempty"`
	Contact         *ProcessedContact                          `codec:"contact,omitempty" json:"contact,omitempty"`
	Imptofu         *ImpTofuSearchResult                       `codec:"imptofu,omitempty" json:"imptofu,omitempty"`
	ServicesSummary map[APIUserServiceID]APIUserServiceSummary `codec:"servicesSummary" json:"services_summary"`
	RawScore        float64                                    `codec:"rawScore" json:"rawScore"`
}

func (o APIUserSearchResult) DeepCopy() APIUserSearchResult {
	return APIUserSearchResult{
		Score: o.Score,
		Khulnasoft: (func(x *APIUserKhulnasoftResult) *APIUserKhulnasoftResult {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Khulnasoft),
		Service: (func(x *APIUserServiceResult) *APIUserServiceResult {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Service),
		Contact: (func(x *ProcessedContact) *ProcessedContact {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Contact),
		Imptofu: (func(x *ImpTofuSearchResult) *ImpTofuSearchResult {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Imptofu),
		ServicesSummary: (func(x map[APIUserServiceID]APIUserServiceSummary) map[APIUserServiceID]APIUserServiceSummary {
			if x == nil {
				return nil
			}
			ret := make(map[APIUserServiceID]APIUserServiceSummary, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.ServicesSummary),
		RawScore: o.RawScore,
	}
}

type NonUserDetails struct {
	IsNonUser            bool                  `codec:"isNonUser" json:"isNonUser"`
	AssertionValue       string                `codec:"assertionValue" json:"assertionValue"`
	AssertionKey         string                `codec:"assertionKey" json:"assertionKey"`
	Description          string                `codec:"description" json:"description"`
	Contact              *ProcessedContact     `codec:"contact,omitempty" json:"contact,omitempty"`
	Service              *APIUserServiceResult `codec:"service,omitempty" json:"service,omitempty"`
	SiteIcon             []SizedImage          `codec:"siteIcon" json:"siteIcon"`
	SiteIconDarkmode     []SizedImage          `codec:"siteIconDarkmode" json:"siteIconDarkmode"`
	SiteIconFull         []SizedImage          `codec:"siteIconFull" json:"siteIconFull"`
	SiteIconFullDarkmode []SizedImage          `codec:"siteIconFullDarkmode" json:"siteIconFullDarkmode"`
}

func (o NonUserDetails) DeepCopy() NonUserDetails {
	return NonUserDetails{
		IsNonUser:      o.IsNonUser,
		AssertionValue: o.AssertionValue,
		AssertionKey:   o.AssertionKey,
		Description:    o.Description,
		Contact: (func(x *ProcessedContact) *ProcessedContact {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Contact),
		Service: (func(x *APIUserServiceResult) *APIUserServiceResult {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Service),
		SiteIcon: (func(x []SizedImage) []SizedImage {
			if x == nil {
				return nil
			}
			ret := make([]SizedImage, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.SiteIcon),
		SiteIconDarkmode: (func(x []SizedImage) []SizedImage {
			if x == nil {
				return nil
			}
			ret := make([]SizedImage, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.SiteIconDarkmode),
		SiteIconFull: (func(x []SizedImage) []SizedImage {
			if x == nil {
				return nil
			}
			ret := make([]SizedImage, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.SiteIconFull),
		SiteIconFullDarkmode: (func(x []SizedImage) []SizedImage {
			if x == nil {
				return nil
			}
			ret := make([]SizedImage, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.SiteIconFullDarkmode),
	}
}

type EmailOrPhoneNumberSearchResult struct {
	Input          string `codec:"input" json:"input"`
	Assertion      string `codec:"assertion" json:"assertion"`
	AssertionValue string `codec:"assertionValue" json:"assertionValue"`
	AssertionKey   string `codec:"assertionKey" json:"assertionKey"`
	FoundUser      bool   `codec:"foundUser" json:"foundUser"`
	Username       string `codec:"username" json:"username"`
	FullName       string `codec:"fullName" json:"fullName"`
}

func (o EmailOrPhoneNumberSearchResult) DeepCopy() EmailOrPhoneNumberSearchResult {
	return EmailOrPhoneNumberSearchResult{
		Input:          o.Input,
		Assertion:      o.Assertion,
		AssertionValue: o.AssertionValue,
		AssertionKey:   o.AssertionKey,
		FoundUser:      o.FoundUser,
		Username:       o.Username,
		FullName:       o.FullName,
	}
}
