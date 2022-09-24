package model

type Model interface {
	TableName() string
	ModelName() string
}
