package responseDtos

type SVGforNFTResponse struct {
	SvgID     string `json:"svgid" bson:"_id,omitempty"`
	Thumbnail string `json:"thumbnail" bson:"thumbnail,omitempty"`
	SVG       string `json:"svg" bson:"svg,omitempty"`
}
