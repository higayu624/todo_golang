package controllers

type Context interface {
	Param(string) string
	ShouldBind(interface{}) error
	Status(int)
	JSON(int, interface{})
}
