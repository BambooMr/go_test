package master

import (
	"context"
	"crontab/common"
	"encoding/json"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"time"
)

type jobMgr struct {
	client *clientv3.Client
	kv clientv3.KV
	lease clientv3.Lease
}

var (
	G_jobMgr *jobMgr
)

func InitJobMgr() (err error){

	var (
		config clientv3.Config
		client *clientv3.Client
	)

	config = clientv3.Config {
		Endpoints: G_config.EtcdEndpoints,	//集群地址
		DialTimeout: time.Duration(G_config.EtcdDialTimeout) * time.Millisecond,
	}

	if client, err = clientv3.New(config); err != nil {
		return
	}

	// 得到KV和lease的API子集
	kv := clientv3.NewKV(client)
	lease := clientv3.NewLease(client)

	G_jobMgr = &jobMgr {
		client: client,
		kv: kv,
		lease: lease,
	}

	return

}

// 保存任务
func (jobMgr *jobMgr) SaveJob(job *common.Job) (oldJob *common.Job, err error) {

	var (
		jobKey string
		jobValue []byte
		putResp *clientv3.PutResponse
		oldJobObj common.Job
	)

	jobKey = common.JOB_SAVE_DIR + job.Name
	if jobValue, err = json.Marshal(job); err != nil {
		return
	}

	if putResp, err = jobMgr.kv.Put(context.TODO(), jobKey, string(jobValue), clientv3.WithPrevKV()); err != nil {
		return
	}

	if putResp.PrevKv != nil {
		if err = json.Unmarshal(putResp.PrevKv.Value, &oldJobObj); err != nil {
			err = nil
			return
		}

		oldJob = &oldJobObj
	}

	return
}

func (jobMgr *jobMgr) DeleteJob (name string) (oldJob *common.Job, err error) {

	var (
		jobKey string
		delResp *clientv3.DeleteResponse
		oldJobObj common.Job
	)

	jobKey = common.JOB_SAVE_DIR + name

	if delResp, err = jobMgr.kv.Delete(context.TODO(), jobKey, clientv3.WithPrevKV()); err != nil {
		return
	}

	if len(delResp.PrevKvs) != 0 {
		if err = json.Unmarshal(delResp.PrevKvs[0].Value, &oldJobObj); err != nil {
			err = nil
			return
		}
		oldJob = &oldJobObj
	}

	return
}

func (jobMgr *jobMgr) ListJobs()(jobList []*common.Job, err error) {

	var (
		dirKey string
		getResp *clientv3.GetResponse
		kvPair *mvccpb.KeyValue
		job *common.Job
	)

	dirKey = common.JOB_SAVE_DIR

	if getResp, err = jobMgr.kv.Get(context.TODO(), dirKey, clientv3.WithPrefix()); err != nil {
		return
	}

	// 初始化数组空间
	jobList = make([]*common.Job, 0)

	for _, kvPair = range getResp.Kvs {
		job = &common.Job{}
		if err = json.Unmarshal(kvPair.Value, job); err != nil {
			err = nil
			continue
		}

		jobList = append(jobList, job)
	}

	return
}

func (jobMgr *jobMgr) KillJob(name string)(err error) {
	var (
		killKey string
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseld clientv3.LeaseID
	)

	killKey = common.JOB_KILLER_DIR + name

	if leaseGrantResp, err = jobMgr.lease.Grant(context.TODO(), 1); err != nil {
		return
	}

	leaseld = leaseGrantResp.ID
	if _, err = jobMgr.kv.Put(context.TODO(), killKey, "", clientv3.WithLease(leaseld)); err != nil {
		return
	}

	return
}

