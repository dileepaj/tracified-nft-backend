package requestDtos

type RequestWidget struct {
	WidgetId string `json:"widgetId" bson:"widgetid"  validate:"required"`
	Query    string `json:"Query" bson:"query" validate:"required"`
}
