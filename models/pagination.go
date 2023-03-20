package models

type PaginationTemplate struct {
	TotalElements int32 `json:"totalelements" bson:"totalelements" validate:"required"`
	TotalPages    int32 `json:"totalPages" bson:"totalPages" validate:"required"`
	PageSize      int32 `json:"pagesize" bson:"pagesize" validate:"required"`
	Previouspage  int32 `json:"previouspage" bson:"previouspage" validate:"required"`
	Currentpage   int32 `json:"currentpg" bson:"currentpg" validate:"required"`
	NextPage      int32 `json:"nextpage" bson:"nextpage" validate:"required"`
}
