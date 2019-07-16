package engine

type ReadyNotifier interface {
	WorkerReady(w chan Request)
}

type Scheduler interface {
	ReadyNotifier
	Submit(request Request)
	WorkerChan() chan Request
	//ConfigureMasterWorkerChan(chan Request)	//设置chan
	Run()
}

type Processor func(Request) (ParseResult, error)

// 有点类似java中写的各种manager
type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int //worker的数量 就是配置worker协程的数量
	ItemChan         chan Item
	RequestProcessor Processor
}

func (c *ConcurrentEngine) Run(seeds ...Request) {
	//worker公用一个in，out
	//in := make(chan Request)
	out := make(chan ParseResult)

	// 设置chan
	//c.Scheduler.ConfigureMasterWorkerChan(in)
	c.Scheduler.Run()

	for i := 0; i < c.WorkerCount; i++ {
		// 创建worker,此处是创建多个协程 , 进行爬取和解析
		//createWorker(in, out)
		//createWorker(out, c.Scheduler)
		//createWorker(c.Scheduler.WorkerChan(), out, c.Scheduler)
		//createWorker(c.Scheduler.WorkerChan(), out, c.Scheduler)
		c.createWorker(c.Scheduler.WorkerChan(), out, c.Scheduler)
	}

	//向worker中提交种子
	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}
	//itemCount := 0
	//从out中获取result 对于request，就继续分配
	for {
		result := <-out

		//   此处只是用于打印,之后可以存储数据库
		for _, item := range result.Items {
			//log.Printf("Got %d  item %v", itemCount, item)
			//itemCount++
			go func(it Item) {
				c.ItemChan <- it
			}(item)

		}
		for _, r := range result.Requests {
			c.Scheduler.Submit(r) //提交余下的种子
		}
	}
}

//func createWorker(in <-chan Request, out chan<- ParseResult) {
//func createWorker(out chan<- ParseResult, s Scheduler) {
func createWorker(in chan Request, out chan<- ParseResult, ready ReadyNotifier) {
	//in := make(chan Request)
	go func() {
		for {
			//r := <-in
			//result, err :=Worker(r)
			//if err != nil {
			//	continue
			//}
			//out <- parseResult

			//需要让scheduler知道已经就绪了
			ready.WorkerReady(in)
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

// 创建远程的Worker
func (c *ConcurrentEngine) createWorker(in chan Request, out chan<- ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := c.RequestProcessor(request) //rpc调用
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
