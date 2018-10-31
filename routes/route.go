package routes

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"go-copyright-p2/dbs"
	"go-copyright-p2/eths"
	"go-copyright-p2/utils"

	"go-copyright-p2/configs"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

const PAGE_MAX_PIC = 5

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

//select content_hash,title,token_id from content a,account_content b where a.content_hash=b.content_hash and address=''
//查看用户所有图片
func GetContents(c echo.Context) error {
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
	//查询数据库
	sql := fmt.Sprintf("select a.content_hash,title,token_id from content a,account_content b where a.content_hash=b.content_hash and address='%s'", address)
	fmt.Println(sql)
	contents, num, err := dbs.DBQuery(sql)
	if err != nil || num <= 0 {
		fmt.Println("failed to query contents", err)
		resp.Errno = utils.RECODE_DBERR
		return err
	}
	total_page := int(num)/PAGE_MAX_PIC + 1
	current_page := 1
	mapResp := make(map[string]interface{})
	mapResp["total_page"] = total_page
	mapResp["current_page"] = current_page
	mapResp["contents"] = contents

	resp.Data = mapResp

	return nil
}

//查看单个图片
func GetContent(c echo.Context) error {
	content := &dbs.Content{}
	content.Title = c.Param("title")
	fmt.Println("get title ", content.Title)
	//最好查数据库获得文件路径
	content.Content = "static/photo/" + content.Title
	http.ServeFile(c.Response(), c.Request(), content.Content)
	return nil
}

//登陆功能
func Auction(c echo.Context) error {
	//1. 响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c, &resp)

	//2. 解析数据
	auction := &dbs.Auction{}
	if err := c.Bind(auction); err != nil {
		fmt.Println(auction)
		resp.Errno = utils.RECODE_PARAMERR
		return err
	}

	sess, err := session.Get("session", c)
	if err != nil {
		fmt.Println("failed to get session")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	address, ok := sess.Values["address"].(string)
	if address == "" || !ok {
		fmt.Println("failed to get session,address is nil")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	auction.Address = address
	//3. 操作mysql-新增
	//sql:insert into auction(content_hash,address,token_id,percent,price,status) values();
	sql := fmt.Sprintf("insert into auction(content_hash,address,token_id,percent,price,status) values('%s','%s',%d,%d,%d,0)",
		auction.ContentHash,
		auction.Address,
		auction.TokenID,
		auction.Percent,
		auction.Price,
	)
	fmt.Println(sql)
	_, err = dbs.Create(sql)
	if err != nil {
		fmt.Println("failed to insert  auction", err)
		resp.Errno = utils.RECODE_DBERR
		return err
	}
	return nil
}

//查看拍卖
func GetAuctions(c echo.Context) error {
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
	//查询数据库
	sql := fmt.Sprintf("select a.content_hash,title,price,token_id from content a,auction b where a.content_hash=b.content_hash and address <> '%s' and status <> 1", address)
	fmt.Println(sql)
	auctions, _, err := dbs.DBQuery(sql)
	if err != nil {
		fmt.Println("failed to query auctions", err)
		resp.Errno = utils.RECODE_DBERR
		return err
	}

	resp.Data = auctions

	return nil
}

//结束拍卖
func BidAuction(c echo.Context) error {
	//1. 组织响应数据
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c, &resp)
	//2. 获取参数 price+tokenid
	price := c.QueryParam("price")
	tokenID := c.QueryParam("tokenid")
	if price == "" || tokenID == "" {
		fmt.Println("failed to get params ", price, tokenID)
		resp.Errno = utils.RECODE_PARAMERR
		return errors.New("params not found")
	}
	//3. session?
	sess, err := session.Get("session", c)
	if err != nil {
		fmt.Println("failed to get session")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	address, ok := sess.Values["address"].(string)
	if address == "" || !ok {
		fmt.Println("failed to get session,address is nil")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	//4. 数据库操作 - 修改使拍卖结束
	//更新操作 update auction set price=111,status=1 where token_id = 0
	sql := fmt.Sprintf("update auction set price=%s,status=1 where token_id = %s", price, tokenID)
	_, err = dbs.Create(sql)
	if err != nil {
		fmt.Println("failed to update auction", err)
		resp.Errno = utils.RECODE_DBERR
		return err
	}
	//查看weight
	sql = fmt.Sprintf("select percent,address from auction where token_id = %s", tokenID)
	mData, num, err := dbs.DBQuery(sql)
	if num <= 0 || err != nil {
		fmt.Println("failed to query auction", err)
		resp.Errno = utils.RECODE_DBERR
		return err
	}
	weight := mData[0]["percent"]
	seller := mData[0]["address"]
	//5. 操作以太坊- 分割资产，erc20转账
	//EthSplitAsset(foundation, pass, buyer string, tokenID, weight int64)
	//EthTransfer20(from, pass, seller string, price int64)
	//启动协程
	go func() {
		//转换字符串为int64
		_tokenid, _ := strconv.ParseInt(tokenID, 10, 32)
		_price, _ := strconv.ParseInt(price, 10, 32)
		_weight, _ := strconv.ParseInt(weight, 10, 32)
		//资产分割
		err := eths.EthSplitAsset(configs.Config.Eth.Foundation, "found", address, _tokenid, _weight)
		if err != nil {
			fmt.Println("failed to call EthSplitAsset", err)
			return
		}
		//转账
		err = eths.EthTransfer20(address, "yekai", seller, _price)
		if err != nil {
			fmt.Println("failed to call EthTransfer20", err)
			return
		}
	}()
	return nil

}
