package UUID

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

/*
假设某互联网公司将唯一 ID生成器立为项目，该公司目前的条件如下。
采用3个机房（北京、上海、深圳）的多活架构，均匀承接用户请求。
按照当前的日活跃用户数量做乐观估计，将来每秒有近10亿次获取唯一 ID的需求。
当前的硬件和服务器框架可支持单个服务实例每秒最多处理100万个请求。
期望唯一 ID生成器可以工作30年。
根据如上条件，采用Snowflake算法设计的唯一 ID按位从高到低分段如下。
- 1位：依然是符号位，固定值为0,以保证生成正整数。
- 40位：系统运行的总毫秒数，30年约为9461亿毫秒，可以用40位二进制数表示。
- 2位：用于区分3个机房，并满足将来增加一个新机房的需求。
- 接下来的10位：单个服务实例每秒处理100万个请求，即每毫秒处理1000个请求。使用10位二进制数来区分同一毫秒内的并发请求。
- 最后的11位：可全部用于区分单个机房内唯一 ID生成器服务的不同实例，即可支持部署最多2048个服务实例。
最终的唯一ID生成器服务在全部机房每秒可承接的用户请求量为3 X 2048 X 100万 ≈ 61亿个，远超公司预期的请求量，并可实际运行34年有余。这个服务的核心代码实现
也很简单，需要额外注意的细节是，如果在某一毫秒内处理的用户请求量超过1024个， 那么服务将报错返回或者使请求阻塞等待到下一毫秒：
*/

// IdGeneratorService ID生成器服务结构体
type IdGeneratorService struct {
	lock         sync.Mutex
	dataCenterId int64     // 机房ID,人为指定
	workerId     int64     // 服务实例ID,人为指定
	startTime    time.Time // 系统初始时间，人为指定
	millisPassed int64     // 上一次请求的处理时间距离初始时间有多少毫秒
	concurrency  int64     // 记录在同一毫秒内生成的ID数
}

// GenID 生成唯一 ID
func (svr *IdGeneratorService) GenID() (id int64, err error) {
	// 计算当前时间距离初始时间有多少毫秒
	millis := time.Now().Sub(svr.startTime).Milliseconds()
	svr.lock.Lock()
	defer svr.lock.Unlock()
	// 与上一次生成ID的请求处于同一毫秒内
	var concurrencyValue int64
	if millis == svr.millisPassed {
		// 本毫秒内生成的ID数已超过最大并发数范围，报错返回
		if svr.concurrency >= 1<<10 {
			return 0, errors.New("concurrency limit")
		} else {
			concurrencyValue = svr.concurrency // 当前并发数
			svr.concurrency += 1               // 更新并发数
		}
	} else {
		concurrencyValue = 0      // 此请求是本毫秒内的第一个请求，标记为0
		svr.concurrency = 1       // 更新并发数
		svr.millisPassed = millis // 更新最新毫秒总数
	}
	// 按1-40-2-11-10的占位排布各变量
	fmt.Printf("millis: %064b\n", millis)
	fmt.Printf("dataCenterId: %064b\n", svr.dataCenterId)
	fmt.Printf("workerId: %064b\n", svr.workerId)
	fmt.Printf("concurrencyValue: %064b\n", concurrencyValue)
	id = (millis << 23) | (svr.dataCenterId << 21) | (svr.workerId << 10) | concurrencyValue
	fmt.Printf("id: %064b\n", id)
	return
}
