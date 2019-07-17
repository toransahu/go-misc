package jsons

// Message is an structure for message being passed as JSON
type Message struct {
	ID     uint16 `json:"id"`
	Name   string `json:"name"`
	secret string `json:"secret"` // if add json tag: vim-go lint error| struct field secret has json tag but is not exported
}

// for json alias/tags
// vim-go : visual select -> <leader> GoAddTags
// GoLand : tbd
