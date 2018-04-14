package marvel

// CharacterDataWrapper contains character wrapper information
type CharacterDataWrapper struct {
	DataWrapper
	Data CharacterDataContainer `json:"data,omitempty"`
}

// CharacterDataContainer contains character container information
type CharacterDataContainer struct {
	DataContainer
	Results []Character `json:"results,omitempty"`
}

// Character contains Marvel character information
type Character struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ResourceURI string `json:"resourceURI,omitempty"`
	URLs        []URL  `json:"urls,omitempty"`
	Thumbnail   Image `json:"thumbnail,omitempty"`
}
