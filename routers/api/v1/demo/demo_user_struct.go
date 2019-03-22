package demo

type DemoUser struct {
	User string `form:"user" json:"user" xml:"user"  binding:"required"`
	Pass string `form:"pass" json:"pass" xml:"pass"  binding:"required"`
}