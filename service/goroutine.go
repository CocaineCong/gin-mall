package service

import (
	"github.com/lunny/log"
	"runtime"
	"sync"
	"time"
)

//工厂模型
type Factory struct {
	Wg        *sync.WaitGroup 		//任务监控系统
	MaxWorker int             		//最大机器数
	MaxJobs   int             		//最大工作数量
	JobQueue  chan int        		//工作队列管道
	Quit      chan bool       		//是否关闭机器
}

//创建工厂模型
func NewFactory(maxWorker int, wg *sync.WaitGroup) Factory {
	return Factory{
		Wg:        wg,                        		//引用任务监控系统
		MaxWorker: maxWorker,                	 	//机器数量（数量多少，根据服务器性能而定）
		JobQueue:  make(chan int, maxWorker), 		//工作管道，数量大于等于机器数
		Quit:      make(chan bool),
	}
}

//设置最大订单数量
func (f *Factory) SetMaxJobs(taskNum int) {
	f.MaxJobs = taskNum
}

//开始上班
func (f *Factory) Start() {				//机器开机，MaxWorker
	for i := 0; i < f.MaxWorker; i++ {
		go func() {					//每一台机器开启后，去工作吧
			for {					//等待下发命令
				select {
				case i := <-f.JobQueue:
					f.doWork(i)				//接到工作，开工！
				case <-f.Quit:
					log.Println("机器关机")
					return
				}
			}
		}()
	}
}

//分配每个任务到管道中
func (f *Factory) AddTask(taskNum int) {
	f.Wg.Add(1)			//系统监控任务 +1
	f.JobQueue <- taskNum		//分配任务到管道中
}

//模拟耗时工作
func (f *Factory) doWork(taskNum int) {
	time.Sleep(200 * time.Millisecond)				//生产产品的工作
	f.Wg.Done()										//完成工作报告
	//log.Println("完工：", taskNum)
}

//创建工厂
func Begin() {
	gomaxprocs := runtime.GOMAXPROCS(runtime.NumCPU())		//配置工作核数
	log.Println("核数：", gomaxprocs)
	wg := new(sync.WaitGroup)						//配置监控系统
	factory := NewFactory(1000, wg)		//开工厂
	factory.SetMaxJobs(10000)	//订单量
	factory.Start()						//开始上班
	log.Println("开始生产")
	//讲所有的订单，添加到任务队列
	for i := 0; i < factory.MaxJobs; i++ {
		factory.AddTask(i)
	}
	factory.Wg.Wait()
	log.Println("所有订单任务生产完成")
}


