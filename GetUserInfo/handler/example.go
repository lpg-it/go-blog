package handler

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	getuserinfo "ihome/GetUserInfo/proto/getuserinfo"
	"ihome/ihomeWeb/models"
	"ihome/ihomeWeb/utils"
	"time"
)

type Server struct{}

func (e *Server) GetUserInfo(ctx context.Context, req *getuserinfo.Request, rsp *getuserinfo.Response) error {
	/* 初始化返回值 */
	rsp.ErrNo = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)

	/* 获取数据 */
	// 获取 sessionId + "userId"，可得到 userId
	// 连接 redis 数据库
	redisConfig, _ := json.Marshal(map[string]string{
		"key":   utils.G_server_name,
		"conn":  utils.G_redis_addr + ":" + utils.G_redis_port,
		"dbNum": utils.G_redis_dbnum,
	})
	bm, err := cache.NewCache("redis", string(redisConfig))
	if err != nil {
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}
	userIdTemp := bm.Get(req.SessionId + "userId")
	userId := int(userIdTemp.([]uint8)[0])

	/* 处理数据 */
	// 从数据库中查询该用户
	o := orm.NewOrm()
	var user models.User
	user.Id = userId
	err = o.Read(&user)
	if err != nil {
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	// 更新 session
	_ = bm.Put(req.SessionId+"userId", user.Id, time.Second*3600)
	_ = bm.Put(req.SessionId+"name", user.Name, time.Second*3600)
	_ = bm.Put(req.SessionId+"mobile", user.Mobile, time.Second*3600)

	/* 返回数据 */
	rsp.UserId = int64(user.Id)
	rsp.UserName = user.Name
	rsp.Mobile = user.Mobile
	rsp.RealName = user.RealName
	rsp.IdCard = user.IdCard
	rsp.AvatarUrl = user.AvatarUrl
	return nil
}
