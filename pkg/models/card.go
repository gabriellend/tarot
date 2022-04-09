package cards

type Card struct {
	Name             string            `json:"name" bson:"name"`
	ZodiacSign       string            `json:"zodiacSign,omitempty" bson:"zodiacSign,omitempty"`
	Element          string            `json:"element,omitempty" bson:"element,omitempty"`
	UprightMeaning   string            `json:"uprightMeaning" bson:"uprightMeaning"`
	ReversedMeaning  string            `json:"reversedMeaning" bson:"reversedMeaning"`
	IsUpright        bool              `json:"isUpright" bson:"isUpright"`
	UprightKeywords  []string          `json:"uprightKeywords" bson:"uprightKeywords"`
	ReversedKeywords []string          `json:"reversedKeywords" bson:"reversedKeywords"`
	Number           int               `json:"number" bson:"number"`
	IsMajor          bool              `json:"isMajor" bson:"isMajor"`
	IsCourt          bool              `json:"isCourt" bson:"isCourt"`
	Planet           string            `json:"planet" bson:"planet"`
	Symbols          map[string]string `json:"symbols" bson:"symbols"`
}
