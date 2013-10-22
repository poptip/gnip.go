package gnip

type Rules struct {
	Rules []Rule `json:"rules"`
}

type Rule struct {
	Tag   string `json:"tag"`
	Value string `json:"value"`
}
