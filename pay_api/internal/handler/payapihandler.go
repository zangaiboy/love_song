package handler

import (
	"net/http"

	"love_song/pay_api/internal/logic"
	"love_song/pay_api/internal/svc"
	"love_song/pay_api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func Pay_apiHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewPay_apiLogic(r.Context(), ctx)
		resp, err := l.Pay_api(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
