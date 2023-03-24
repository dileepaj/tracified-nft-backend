package requestDtos

type SVGHashUpdateRequst struct {
	SvgID string `json:"svgid" bson:"_id,omitempty"`
	Hash  string `json:"hash" bson:"hash,omitempty"`
}
type GenerateSVGReqeust struct {
	BatchID       string `json:"batchID" bson:"batchID,omitempty"`
	ProductID     string `json:"productID" bson:"productID,omitempty"`
	Email         string `json:"email" bson:"email,omitempty"`
	ReciverName   string `json:"recivername" bson:"recivername,omitempty"`
	CustomMessage string `json:"custommsg" bson:"custommsg,omitempty"`
}
