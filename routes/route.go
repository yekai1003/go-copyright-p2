package routes

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"net/http"
	"os"

	"go-copyright-p2/dbs"
	"go-copyright-p2/eths"
	"go-copyright-p2/utils"

	"go-copyright-p2/configs"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

func PingHandler(c echo.Context) error {

	return c.String(http.StatusOK, "pong")
}

//注册功能
func Register(c echo.Context) error {
	//1. 响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c, &resp)

	//2. 解析数据
	account := &dbs.Account{}
	if err := c.Bind(account); err != nil {
		fmt.Println(account)
		resp.Errno = utils.RECODE_PARAMERR
		return err
	}
	//3. 操作geth创建账户
	address, err := eths.NewAcc(account.IdentityID, configs.Config.Eth.Connstr)
	if err != nil {
		resp.Errno = utils.RECODE_IPCERR
		return err
	}
	//4. 操作mysql-插入数据
	//sql:insert into account(email,username,identity_id,address) values('1','2','3','4')
	sql := fmt.Sprintf("insert into account(email,username,identity_id,address) values('%s','%s','%s','%s')",
		account.Email,
		account.UserName,
		account.IdentityID,
		address,
	)
	fmt.Println(sql)
	_, err = dbs.Create(sql)
	if err != nil {
		fmt.Println("failed to insert account", err)
		resp.Errno = utils.RECODE_DBERR
		return err
	}
	//5. session处理
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["address"] = address
	sess.Save(c.Request(), c.Response())
	return nil
}

//登陆功能
func Login(c echo.Context) error {
	//1. 响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c, &resp)

	//2. 解析数据
	account := &dbs.Account{}
	if err := c.Bind(account); err != nil {
		fmt.Println(account)
		resp.Errno = utils.RECODE_PARAMERR
		return err
	}

	//3. 操作mysql-查询数据
	//sql:select * from account where username ='yekai' and identity_id='yekai';
	sql := fmt.Sprintf("select * from account where username ='%s' and identity_id='%s'",
		account.UserName,
		account.IdentityID,
	)
	fmt.Println(sql)
	m, n, err := dbs.DBQuery(sql)
	if err != nil || n <= 0 {
		fmt.Println("failed to query account", err)
		resp.Errno = utils.RECODE_DBERR
		return err
	}
	//m 是一个[]map[string]string
	rows := m[0]
	//4. session处理
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["address"] = rows["address"]
	sess.Save(c.Request(), c.Response())
	return nil
}

//上传图片功能
func Upload(c echo.Context) error {
	//1. 响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c, &resp)

	//2. 解析数据
	content := &dbs.Content{}

	h, err := c.FormFile("fileName")
	if err != nil {
		fmt.Println("failed to FormFile ", err)
		resp.Errno = utils.RECODE_PARAMERR
		return err
	}
	src, err := h.Open()
	defer src.Close()
	content.Content = "static/photo/" + h.Filename
	dst, err := os.Create(content.Content)
	if err != nil {
		fmt.Println("failed to create file ", err, content.Content)
		resp.Errno = utils.RECODE_IOERR
		return err
	}
	defer dst.Close()
	//计算hash
	cData := make([]byte, h.Size)
	n, err := src.Read(cData)
	if err != nil || h.Size != int64(n) {
		resp.Errno = utils.RECODE_IOERR
		return err
	}
	content.ContentHash = fmt.Sprintf("%x", sha256.Sum256(cData))
	dst.Write(cData)
	content.Title = h.Filename

	//写文件

	//3. 操作mysql-新增数据
	content.AddContent()

	//操作以太坊

	sess, _ := session.Get("session", c)
	fromAddr, ok := sess.Values["address"].(string)
	if fromAddr == "" || !ok {
		resp.Errno = utils.RECODE_SESSIONERR
		return errors.New("no session")
	}
	//UploadPic(from, pass, hash, data string)
	go eths.UploadPic(fromAddr, "yekai", content.ContentHash, content.Title)

	return nil
}

//session获取
func GetSession(c echo.Context) error {
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c, &resp)
	//处理session
	sess, err := session.Get("session", c)
	if err != nil {
		fmt.Println("failed to get session")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	address := sess.Values["address"]
	if address == nil {
		fmt.Println("failed to get session,address is nil")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	return nil
}
