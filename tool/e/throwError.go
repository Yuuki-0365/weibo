package e

import "weibo/tool"

func ThrowError(code int) tool.Response {
	return tool.Response{
		Status: code,
		Msg:    GetMsg(code),
	}
}

func ThrowSuccess() tool.Response {
	return tool.Response{
		Status: Success,
		Msg:    GetMsg(Success),
	}
}
