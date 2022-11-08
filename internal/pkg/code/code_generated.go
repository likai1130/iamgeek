// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Code generated by "codegen -type=int /home/colin/workspace/golang/src/github.com/marmotedu/iam/internal/pkg/code"; DO NOT EDIT.

package code

// init register error codes defines in this source code to `github.com/marmotedu/errors`
func init() {
	register(ErrUserNotFound, 404, "User not found","用户不存在")
	register(ErrUserAlreadyExist, 400, "User already exist","用户已存在")
	register(ErrSuccess, 200, "OK","成功")
	register(ErrUnknown, 500, "Internal server error","内部服务器错误")
	register(ErrBind, 400, "Error occurred while binding the request body to the struct","将请求正文绑定到结构时出错")
	register(ErrValidation, 400, "Validation failed","验证失败")
	register(ErrTokenInvalid, 401, "Token invalid","token失效")
	register(ErrPageNotFound, 404, "Page not found","未找到页面")
	register(ErrDatabase, 500, "Database error","数据库错误")
	register(ErrEncrypt, 401, "Error occurred while encrypting the user password","加密用户密码时出错")
	register(ErrSignatureInvalid, 401, "Signature is invalid","签名无效")
	register(ErrExpired, 401, "Token expired","token过期")
	register(ErrInvalidAuthHeader, 401, "Invalid authorization header","授权标头无效")
	register(ErrMissingHeader, 401, "The `Authorization` header was empty","Authorization标头为空")
	register(ErrPasswordIncorrect, 401, "Password was incorrect","密码不正确")
	register(ErrPermissionDenied, 403, "Permission denied","权限被拒绝")
	register(ErrEncodingFailed, 500, "Encoding failed due to an error with the data","由于数据错误，编码失败")
	register(ErrDecodingFailed, 500, "Decoding failed due to an error with the data","由于数据错误，解码失败")
	register(ErrInvalidJSON, 500, "Data is not valid JSON","数据不是有效JSON")
	register(ErrEncodingJSON, 500, "JSON data could not be encoded","JSON数据被不能被编码")
	register(ErrDecodingJSON, 500, "JSON data could not be decoded","JSON数据被不能被解码")
	register(ErrInvalidYaml, 500, "Data is not valid Yaml","yaml验证失败")
	register(ErrEncodingYaml, 500, "Yaml data could not be encoded","yaml不能被编码")
	register(ErrDecodingYaml, 500, "Yaml data could not be decoded","Yaml不能被解码")

	register(ErrBlockChainNotFound,404,"Blockchain not found","此区块链未找到")
	register(ErrBlockChainIdIllegal,400,"Blockchain chainId must be between[1,100000]","链Id必须在[1-100000]之间")




}
