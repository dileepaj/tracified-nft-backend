package responseDtos

type SVGforNFTResponse struct {
	SvgID string `json:"svgid" bson:"_id,omitempty"`
	SVG   string `json:"svg" bson:"svg,omitempty"`
}
