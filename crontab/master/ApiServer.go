package master

import (
	"crontab/common"
	"encoding/json"
	"net"
	"net/http"
	"strconv"
	"time"
)

// 任务的HTTP接口
type ApiServer struct {
	httpServer *http.Server
}

var (
	// 单例对象
	G_apiServer *ApiServer
)

//  保存任务接口
func handleJobSave(resp http.ResponseWriter, req *http.Request) {
	var (
		err error
		postJob string
		job common.Job
		oldJob *common.Job
		bytes []byte
	)

	// 1.解析POST表单
	if err = req.ParseForm(); err != nil {
		goto ERR
	}

	// 2.取表单字段
	postJob = req.PostForm.Get("job")

	// 3.反序列化
	if err = json.Unmarshal([]byte(postJob), &job); err != nil {
		goto ERR
	}

	// 4.保存到etcd
	if oldJob, err = G_jobMgr.SaveJob(&job); err != nil {
		goto ERR
	}

	if bytes, err = common.BuildResponse(0, "success", oldJob); err == nil {
		resp.Write(bytes)
	}

	return

ERR:
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err != nil {
		resp.Write(bytes)
	}

}

// 删除任务接口
func handleJobDelete(resp http.ResponseWriter, req *http.Request) {

	var (
		err error
		name string
		oldJob *common.Job
		bytes []byte
	)

	if err = req.ParseForm(); err != nil {
		goto ERR
	}

	// 删除任务名
	name = req.PostForm.Get("name")

	// 删除操作
	if oldJob, err = G_jobMgr.DeleteJob(name); err != nil {
		goto ERR
	}

	// 正常应答
	if bytes, err = common.BuildResponse(0, "success", oldJob); err == nil {
		resp.Write(bytes)
	}

	return

ERR:
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		resp.Write(bytes)
	}

}

// 列举所有crontab任务
func handleJobList(resp http.ResponseWriter, req *http.Request) {

	var (
		jobList []*common.Job
		err error
		bytes []byte
	)

	if jobList,err = G_jobMgr.ListJobs(); err != nil {
		goto ERR
	}

	if bytes, err = common.BuildResponse(0, "success", jobList); err == nil {
		resp.Write(bytes)
	}

	return

ERR:
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		resp.Write(bytes)
	}

}

//
func handleJobKill(resp http.ResponseWriter, req *http.Request) {
	var (
		err error
		name string
		bytes []byte
	)

	if err = req.ParseForm(); err != nil {
		goto ERR
	}

	name = req.PostForm.Get("name")

	if err = G_jobMgr.KillJob(name); err != nil {
		goto ERR
	}

	if bytes, err = common.BuildResponse(0, "success", nil); err == nil {
		resp.Write(bytes)
	}

	return

ERR:
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		resp.Write(bytes)
	}

}

// 初始化服务
func InitApiServer() (err error) {
	var (
		mux *http.ServeMux
		listener net.Listener
		httpServer *http.Server
		staticDir http.Dir
		staticHandler http.Handler
	)

	// 配置路由
	mux = http.NewServeMux()
	mux.HandleFunc("/job/save", handleJobSave)
	mux.HandleFunc("/job/delete", handleJobDelete)
	mux.HandleFunc("/job/list", handleJobList)
	mux.HandleFunc("/job/kill", handleJobKill)

	//

	// 静态文件目录
	staticDir = http.Dir(G_config.Webroot)
	staticHandler = http.FileServer(staticDir)
	mux.Handle("/", http.StripPrefix("/", staticHandler))

	// 启动TCP监听
	if listener, err  = net.Listen("tcp", ":" + strconv.Itoa(G_config.ApiPort)); err != nil {
		return
	}

	// 创建HTTP服务
	httpServer = &http.Server{
		ReadTimeout: time.Duration(G_config.ApiReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(G_config.ApiWriteTimeout) * time.Millisecond,
		Handler: mux,
	}

	G_apiServer = &ApiServer{
		httpServer: httpServer,
	}

	// 启动服务端
	go httpServer.Serve(listener)

	return
}