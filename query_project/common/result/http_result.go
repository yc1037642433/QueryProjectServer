package result

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		//成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		logx.WithContext(r.Context()).Errorf("[API ERROR] : %v ", err)
		httpx.WriteJson(w, http.StatusBadRequest, Error(-1, err.Error()))
	}
}

func AuthHttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		//成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回
		logx.WithContext(r.Context()).Errorf("[GATEWAY ERROR] : %+v ", err)
		httpx.WriteJson(w, http.StatusUnauthorized, Error(-2, err.Error()))
	}
}

func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	logx.Error(err)
	httpx.WriteJson(w, http.StatusBadRequest, Error(-3, err.Error()))
}
