package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/plugins/session"
	user "github.com/xiaobudongzhang/micro-user-srv/proto/user"

	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/v2/client"
	auth "github.com/xiaobudongzhang/micro-auth/proto/auth"
)

var (
	serviceClient user.UserService
	authClient    auth.Service
)

type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	serviceClient = user.NewUserService("mu.micro.book.service.user", client.DefaultClient)
	authClient = auth.NewService("mu.micro.book.service.auth", client.DefaultClient)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	r.ParseForm()

	rsp, err := serviceClient.QueryUserByName(context.TODO(), &user.UserRequest{
		UserName: r.Form.Get("userName"),
	})

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}

	if rsp.User.Pwd == r.Form.Get("pwd") {
		response["sucees"] = true

		rsp.User.Pwd = ""
		response["data"] = rsp.User

		log.Logf("生成token")

		resp2, err := authClient.MakeAccessToken(context.TODO(), &auth.Request{
			UserId:   rsp.User.Id,
			UserName: rsp.User.Name,
		})

		if err != nil {
			log.Logf("create token error %s", err)
			http.Error(w, err.Error(), 500)
			return
		}

		log.Logf("token %s", resp2.Token)
		response["token"] = resp2.Token

		w.Header().Add("set-cookie", "application/json; charset=utf-8")

		expire := time.Now().Add(300 * time.Minute)

		cookie := http.Cookie{Name: "remeber-me-token", Value: resp2.Token, Path: "/", Expires: expire, MaxAge: 9000, HttpOnly: false, Secure: false}

		http.SetCookie(w, &cookie)

		cookie2 := http.Cookie{Name: "test", Value: resp2.Token, Path: "/", Expires: expire, MaxAge: 9000, HttpOnly: false, Secure: false}

		http.SetCookie(w, &cookie2)
		//同步到session
		sess := session.GetSession(w, r)
		sess.Values["userId"] = rsp.User.Id
		sess.Values["userName"] = rsp.User.Name
		_ = sess.Save(r, w)
	} else {
		response["success"] = false
		response["error"] = &Error{
			Detail: "密码错误",
		}
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Logf("err param")
		http.Error(w, "err param", 400)
		return
	}

	tokenCookie, err := r.Cookie("remember-me-token")

	if err != nil {
		log.Logf("token get fail")
		http.Error(w, "error fail", 400)
		return
	}

	_, err = authClient.DelUserAccessToken(context.TODO(), &auth.Request{
		Token: tokenCookie.Value,
	})

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	cookie := http.Cookie{Name: "remember-me-token", Value: "", Path: "/", Expires: time.Now().Add(0 * time.Second), MaxAge: 0}
	http.SetCookie(w, &cookie)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	response := map[string]interface{}{
		"ref":     time.Now().UnixNano(),
		"success": true,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

func TestSession(w http.ResponseWriter, r *http.Request) {
	sess := session.GetSession(w, r)

	if v, ok := sess.Values["path"]; !ok {
		sess.Values["path"] = r.URL.Query().Get("path")
		log.Logf("path:" + r.URL.Query().Get("path"))
	} else {
		log.Logf(v.(string))
	}

	log.Logf(sess.ID)
	log.Logf(sess.Name())

	w.Write([]byte("OK"))
}
