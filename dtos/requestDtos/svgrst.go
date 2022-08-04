package requestDtos

type SVGHashUpdateRequst struct {
	SvgID string `json:"svgid" bson:"_id,omitempty"`
	Hash  string `json:"hash" bson:"hash,omitempty"`
}
