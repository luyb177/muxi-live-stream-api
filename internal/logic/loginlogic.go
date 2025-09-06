package logic

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"

	"muxi-live-stream-api/internal/svc"
	"muxi-live-stream-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.Response, err error) {
	// 判断账号密码
	if req.UserName == "" || req.PassWord == "" {
		return &types.Response{
			Code:    400,
			Message: "用户名或密码不能为空",
		}, nil
	}

	loginUrl := "https://account.ccnu.edu.cn/cas/login?service=http://kjyy.ccnu.edu.cn/loginall.aspx?page="

	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}

	res, err := client.Get(loginUrl)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	//使用正则表达式找到It和execution
	ret := regexp.MustCompile(` <input type="hidden" name="lt" value="(.*?)" />`)
	lts := ret.FindAllStringSubmatch(string(body), -1) //查找所有的lt
	lt := lts[0][1]

	//获取execution
	ret1 := regexp.MustCompile(` <input type="hidden" name="execution" value="(.*?)" />`)
	executions := ret1.FindAllStringSubmatch(string(body), -1)
	execution := executions[0][1]

	// 构造表单
	formData := url.Values{
		"username":  {req.UserName},
		"password":  {req.PassWord},
		"lt":        {lt},
		"execution": {execution},
		"_eventId":  {"submit"},
		"submit":    {"登录"},
	}

	request, err := http.NewRequest("POST", loginUrl, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Println(err)
		return &types.Response{
			Code:    400,
			Message: "出现未知错误",
		}, err
	}

	// 添加请求头
	request.Header.Set("content-type", "application/x-www-form-urlencoded")
	request.Header.Set("origin", "https://account.ccnu.edu.cn")
	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")

	// 发送请求
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return &types.Response{
			Code:    400,
			Message: "出现未知错误",
		}, err
	}
	defer response.Body.Close()

	cookies := response.Cookies()
	cookieString := make([]string, len(cookies))
	for i, c := range cookies {
		cookieString[i] = c.String()
	}

	return &types.Response{
		Code:    200,
		Message: "登录成功",
		Data:    types.LoginResponse{Cookie: cookieString},
	}, nil
}
