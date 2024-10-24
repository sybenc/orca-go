package code

func init() {
	register(Success, 200, "请求成功")
	register(ErrInternalServer, 500, "服务器内部错误")
	register(ErrBadRequest, 400, "请求存在错误")
	register(ErrNotFound, 404, "资源未找到")
	register(ErrValidate, 400, "字段验证错误")
	register(ErrBind, 400, "参数绑定错误")
	register(ErrMenuAlreadyExist, 409, "菜单已存在")
	register(ErrMenuNotFound, 404, "菜单未找到")
}
