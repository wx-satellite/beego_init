package requests

type User struct {
	Age  int    `form:"age" valid:"Required"`
	Name string `form:"name" valid:"Required"`
}
