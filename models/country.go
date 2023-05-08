package models

type Countries []struct {
	Name           string       `json:"name"`
	TopLevelDomain []string     `json:"topLevelDomain"`
	Alpha2Code     string       `json:"alpha2Code"`
	Alpha3Code     string       `json:"alpha3Code"`
	CallingCodes   []string     `json:"callingCodes"`
	Capital        string       `json:"capital"`
	AltSpellings   []string     `json:"altSpellings"`
	Subregion      string       `json:"subregion"`
	Region         string       `json:"region"`
	Population     int          `json:"population"`
	Latlng         []int        `json:"latlng"`
	Demonym        string       `json:"demonym"`
	Area           float64      `json:"area"`
	Gini           float64      `json:"gini"`
	Timezones      []string     `json:"timezones"`
	Borders        []string     `json:"borders"`
	NativeName     string       `json:"nativeName"`
	NumericCode    string       `json:"numericCode"`
	Flags          Flags        `json:"flags"`
	Currencies     []Currencies `json:"currencies"`
	Languages      []Languages  `json:"languages"`
	Translations   Translations `json:"translations"`
	Flag           string       `json:"flag"`
	Cioc           string       `json:"cioc"`
	Independent    bool         `json:"independent"`
}
type Flags struct {
	Svg string `json:"svg"`
	Png string `json:"png"`
}
type Currencies struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
type Languages struct {
	Iso6391    string `json:"iso639_1"`
	Iso6392    string `json:"iso639_2"`
	Name       string `json:"name"`
	NativeName string `json:"nativeName"`
}
type Translations struct {
	Br string `json:"br"`
	Pt string `json:"pt"`
	Nl string `json:"nl"`
	Hr string `json:"hr"`
	Fa string `json:"fa"`
	De string `json:"de"`
	Es string `json:"es"`
	Fr string `json:"fr"`
	Ja string `json:"ja"`
	It string `json:"it"`
	Hu string `json:"hu"`
}


// 1.Coutry name
// 2.Calling_code
// 3.Capital
// 4.Population
// 5.Area
// 6.Numeric Code
// 7.Currency_code
// 8.Language -->name
// 9.Flag url