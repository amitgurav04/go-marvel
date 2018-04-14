package marvel

// DataWrapper contains the common wrapper attributes
type DataWrapper struct {
	Code            int    `json:"code,omitempty"`
	Status          string `json:"status,omitempty"`
	Copyright       string `json:"copyright,omitempty"`
	AttributionText string `json:"attributionText,omitempty"`
	AttributionHTML string `json:"attributionHTML,omitempty"`
	ETag            string `json:"etag,omitempty"`
}

// DataContainer contains the common container attributes
type DataContainer struct {
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
	Total  int `json:"total,omitempty"`
	Count  int `json:"count,omitempty"`
}

// URL: A set of public web site URLs for the resource
type URL struct {
	Type string `json:"type,omitempty"`
	URL  string `json:"url,omitempty"`
}

// Image: The representative image for the resource
type Image struct {
	Path      string `json:"path,omitempty"`
	Extension string `json:"extension,omitempty"`
}
