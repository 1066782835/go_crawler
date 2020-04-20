package models

import "github.com/astaxie/goredis"
const (
	URL_QUEUE = "url_queue"
	URL_VISIT_SET = "url_visit_set"
)

var (
	client goredis.Client
)

//链接redis
func ConnectRedis(addr string){
	client.Addr = addr
}

func ConnectRedisConf(addr string, password string, db int){
	client.Addr = addr
	client.Password = password
	client.Db = db
}
// 存入到queue
func PutinQueue(url string){
	client.Lpush(URL_QUEUE, []byte(url))
}

//从queue中获取
func PopfromQueue() string{
	res,err := client.Rpop(URL_QUEUE)
	if err != nil{
		panic(err)
	}

	return string(res)
}

//获取队列的长度
func IsQueueEmpty() bool{
	length,err := client.Llen(URL_QUEUE)
	if err != nil || length == 0{
		return true
	}
	return false
}

//添加到已经访问过的set
func AddToSet(url string){
	client.Sadd(URL_VISIT_SET, []byte(url))
}

//是否已经被访问过
func IsVisit(url string) bool{
	bIsVisit, err := client.Sismember(URL_VISIT_SET, []byte(url))
	if err != nil{
		return false
	}

	return bIsVisit
}

