// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"blog-backend/app/dao/internal"
	"blog-backend/app/dao/utils"
	"blog-backend/app/dao/valid"
	"blog-backend/app/model"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"reflect"
)

// userDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type userDao struct {
	*internal.UserDao
}

const (
	redisKeyPrefix      = "user:"
	redisLoginKeyPrefix = "user:login:"
)

var (
	// User is globally public accessible object for table user operations.
	User       userDao
	userWorker *utils.TableWorker
	userLogger *glog.Logger
)

func init() {
	User = userDao{
		internal.NewUserDao(),
	}
	userWorker = utils.NewTableWorker(User.Table)
	userLogger = g.Log(`"user"表`)
}

// GetUserBuUID 通过uid来获取数据的信息，优先会从redis中去查询缓存数据，如果redis中查询不到，再从数据库中查找
// 事实上，这种设计方式是不安全的，因为一个user包含了所有关于user的信息数据，因此推荐采用加密的方式存储用户信息，
// @todo 将 model.User 数据加密之后存储到redis缓存库当中
func (d *userDao) GetUserBuUID(uid int64) (model.User, error) {
	if len(gconv.String(uid)) == 0 {
		return emptyUser(), gerror.New("Query field [uid] must have a positive length.")
	}

	user := emptyUser()
	val, err := g.Redis("user_cache").DoVar("GET", getRedisUserKey(gconv.String(uid)))
	// redis中找到
	if err == nil && !val.IsNil() && !val.IsEmpty() {
		err = val.Scan(&user)
		if err == nil {
			return user, nil
		}
	}
	// redis中没有，需要从mysql中查找
	err = g.Model(User.Table).Where(User.C.Uid, uid).Scan(&user)
	if err != nil {
		// 数据库获取失败，直接返回
		return emptyUser(), gerror.Newf("Cannot get [model.User] with uid:[%d] in mysql database", uid)
	}
	// 成功获取，需要写回到redis中
	//val, err = g.Redis("user_cache").DoVar("SET", getRedisUserKey(gconv.String(uid)), user)
	err = addUserToRedisWithKey(uid, user)
	if err != nil {
		// 遇上redis写失败，直接跳过
		userLogger.Infof("Cannot write [model.User] to redis, uid:[%d]", uid)
	}
	return user, nil
}

// GetUserByUIDs 批量获取 model.User
func (d *userDao) GetUserByUIDs(uids ...int64) ([]model.User, error) {
	if len(uids) == 0 {
		return emptyUserSlice(), nil
	}

	users := emptyUserSlice()
	if len(uids) == 1 {
		user, err := d.GetUserBuUID(uids[0])
		if err != nil {
			return users, err
		}
		users = append(users, user)
		return users, nil
	}

	sqlUid := make([]int64, 0)
	// 包含两个以上的 model.User 需要查询，优先查询redis
	for _, uid := range uids {
		val, err := g.Redis("user_cache").DoVar("GET", getRedisUserKey(gconv.String(uid)))
		if err == nil && !val.IsNil() && !val.IsEmpty() {
			user := emptyUser()
			err := val.Scan(&user)
			if err == nil {
				users = append(users, user)
				continue
			}
		}
		// 其余情况，从mysql中查找
		sqlUid = append(sqlUid, uid)
	}

	// 当前uids中存储的就是需要从mysql数据库中查找的
	sqlUsers := emptyUserSlice()
	err := g.Model(User.Table).WhereIn(User.C.Uid, sqlUid).Order(User.C.Uid).Scan(&sqlUsers)
	if err != nil {
		return users, err
	}
	// 写入到redis中
	for _, u := range sqlUsers {
		_ = addUserToRedisWithKey(u.Uid, u)
	}
	return append(users, sqlUsers...), nil
}

// AddUser 向mysql中添加 model.User，只有在实际成功插入数据库的时候才返回true
func (d *userDao) AddUser(user model.User) bool {
	if reflect.DeepEqual(user, emptyUser()) {
		return false
	}

	// 昵称和用户名不能重复
	all, e := g.Model(User.Table).Where(User.C.UserName, user.UserName).All()
	if e == nil && all != nil {
		userLogger.Info("用户已经存在")
		return false
	}

	all, e = g.Model(User.Table).Where(User.C.NickName, user.NickName).All()
	if e == nil && all != nil {
		userLogger.Info("用户已经存在")
		return false
	}

	err := g.Validator().Rules(valid.UserAddingRules).Messages(valid.UserAddingMsg).CheckStruct(user)
	if err != nil {
		// 数据校验失败
		return false
	}

	// 生成uid
	uid, e := userWorker.NextID()
	if e != nil {
		// 生成uid失败
		return false
	}
	user.Uid = uid
	_, e = g.Model(User.Table).OmitEmpty().Insert(user)
	if e != nil {
		return false
	}
	return true
}

// UpdateUser 更新用户
func (d *userDao) UpdateUser(user model.User) bool {
	if user.Uid == 0 && user.UserName == "" && user.NickName == "" {
		return false
	}

	e := g.Validator().Rules(valid.UserUpdateRules).CheckStruct(user)
	if e != nil {
		return false
	}

	if user.Uid != 0 {
		_, err := g.Model(User.Table).Data(user).OmitEmpty().Where(User.C.Uid, user.Uid).Update()
		if err == nil {
			deleteUserFromRedis(user.Uid)
			return true
		}
	} else if user.UserName != "" {
		_, err := g.Model(User.Table).Data(user).OmitEmpty().Where(User.C.UserName, user.UserName).Update()
		m, _ := g.Model(User.Table).Where(User.C.UserName, user.UserName).One()
		if err == nil {
			deleteUserFromRedis(m["uid"].Uint64())
			return true
		}
	} else if user.NickName != "" {
		_, err := g.Model(User.Table).Data(user).OmitEmpty().Where(User.C.NickName, user.NickName).Update()
		m, _ := g.Model(User.Table).Where(User.C.NickName, user.NickName).One()
		if err == nil {
			deleteUserFromRedis(m["uid"].Uint64())
			return true
		}
	}
	return false
}

// DeleteUserByUID 根据uid删除
func (d *userDao) DeleteUserByUID(uid uint64) bool {
	if uid == 0 {
		return false
	}

	_, err := g.Model(User.Table).Where(User.C.Uid, uid).Delete()
	if err != nil {
		return false
	}
	deleteUserFromRedis(uid)
	return true
}

func (d *userDao) DeleteUserByUIDs(uids ...uint64) bool {
	if len(uids) == 0 {
		return true
	}

	_, err := g.Model(User.Table).WhereIn(User.C.Uid, uids).Delete()
	if err != nil {
		return false
	}
	for _, uid := range uids {
		deleteUserFromRedis(uid)
	}
	return true
}

// CheckIfUserLogin 通过用户名，昵称和Uid检查用户是否已经登录，已经登录的会存储到redis中，其中redis前缀为user:login:
func (d *userDao) CheckIfUserLogin(uid interface{}, userName interface{}, nickName interface{}) bool {
	if uid != nil {
		uid = uid.(uint64)
		if uid == 0 {
			return false
		}
		val, err := g.Redis("user_cache").DoVar("GET", getRedisLoginKey(gconv.String(uid)))
		if err == nil && !val.IsNil() && !val.IsEmpty() {
			return true
		}
	} else if userName != "" {
		user := emptyUser()
		err := g.Model(User.Table).Where(User.C.UserName, userName).Scan(&user)
		if err == nil {
			val, err := g.Redis("user_cache").DoVar("GET", getRedisLoginKey(gconv.String(user.Uid)))
			if err == nil && !val.IsNil() && !val.IsEmpty() {
				return true
			}
		}
	} else if nickName != "" {
		user := emptyUser()
		err := g.Model(User.Table).Where(User.C.NickName, nickName).Scan(&user)
		if err == nil {
			val, err := g.Redis("user_cache").DoVar("GET", getRedisLoginKey(gconv.String(user.Uid)))
			if err == nil && !val.IsNil() && !val.IsEmpty() {
				return true
			}
		}
	}
	return false
}

func deleteUserFromRedis(uid uint64) {
	if uid == 0 {
		return
	}

	_, _ = g.Redis("user_cache").DoVar("DEL", getRedisUserKey(gconv.String(uid)))
}

func addUserToRedisWithKey(key interface{}, val interface{}) error {
	key = getRedisUserKey(gconv.String(key))

	user := val.(model.User)
	user.Password = utils.EncryptPassword(user.Password)

	_, err := g.Redis("user_cache").DoVar("SET", key, user)
	return err
}

// getUserByUID 直接从mysql获取用户，不写入redis
func getUserByUID(uid uint64) (model.User, bool) {
	if uid == 0 {
		return emptyUser(), false
	}
	user := emptyUser()
	err := g.Model(User.Table).Where(User.C.Uid, uid).Scan(&user)
	if err != nil {
		return emptyUser(), false
	}
	return user, true
}

// emptyUser 返回一个空的 model.User 对象，当出现错误，或者需要进一步填充值的时候需要使用
func emptyUser() model.User {
	return model.User{}
}

func emptyUserSlice() []model.User {
	return make([]model.User, 0)
}

func getRedisUserKey(str string) string {
	return redisKeyPrefix + str
}

func getRedisLoginKey(str string) string {
	return redisLoginKeyPrefix + str
}

// CheckUserByUid 检查用户是否存在
func CheckUserByUid(uid uint64) bool {
	if uid == 0 {
		return false
	}

	_, b := getUserByUID(uid)
	return b
}

func checkUserByName(name string) bool {
	if name == "" {
		return false
	}

	c, err := g.Model(User.Table).Where(User.C.UserName, name).Count()
	if err != nil || c == 0 {
		return false
	}
	return true
}

func checkUserByNickName(name string) bool {
	if name == "" {
		return false
	}

	c, err := g.Model(User.Table).Where(User.C.NickName, name).Count()
	if err != nil || c == 0 {
		return false
	}
	return true
}
