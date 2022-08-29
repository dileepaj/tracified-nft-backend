package requestDtos

type SVGHashUpdateRequst struct {
	SvgID string `json:"svgid" bson:"_id,omitempty"`
	Hash  string `json:"hash" bson:"hash,omitempty"`
}
type GenerateSVGReqeust struct {
	BatchID string `json:"batchid" bson:"batchid,omitempty"`
	Email   string `json:"email" bson:"email,omitempty"`
}
