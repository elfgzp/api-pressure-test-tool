package pressuretesttool

import (
	"fmt"
	"math"
	"net/http"
	"sort"
	"sync"
	"time"
)

func New(url string, times int, concurrent int) *PressureTestTool {
	return &PressureTestTool{
		url: url,
		times: times,
		concurrent: concurrent,
		client: &http.Client{},
		wg: sync.WaitGroup{},
		result: &PressureTestResult{
			timeConsumedArr: make([]time.Duration, 0),
			mux: sync.Mutex{},
		},
	}
}

type PressureTestResult struct {
	timeConsumedArr []time.Duration
	mux sync.Mutex
}

func (p *PressureTestResult) Add(timeConsumed time.Duration) {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.timeConsumedArr = append(p.timeConsumedArr, timeConsumed)
}

func (p *PressureTestResult) Avg() time.Duration {
	if len(p.timeConsumedArr) < 1 {
		return 0
	}

	var totalConsumed time.Duration
	for _, consumed := range p.timeConsumedArr {
		totalConsumed += consumed
	}

	avg := int(totalConsumed) / len(p.timeConsumedArr)

	return time.Duration(avg)
}

func (p PressureTestResult) Percentile(percentage float64) time.Duration {
	if len(p.timeConsumedArr) < 1 {
		return 0
	}

	sort.Slice(
		p.timeConsumedArr,
		func(i, j int) bool { return p.timeConsumedArr[i] < p.timeConsumedArr[j]},
	)

	i := int(math.Ceil(float64(len(p.timeConsumedArr)) * percentage)) - 1
	return p.timeConsumedArr[i]
}

type PressureTestTool struct {
	url string
	times int
	concurrent int
	client *http.Client
	wg sync.WaitGroup
	result *PressureTestResult
}


func (p *PressureTestTool) Run(){
	fmt.Printf("开始测试 URL %s 请求次数 %d 并发数 %d ... \n", p.url, p.times, p.concurrent)
	p.wg.Add(p.times)
	counter := &Counter{
		current: 0,
		mux: sync.Mutex{},
	}

	for i := 0; i < p.concurrent; i++ {
		num := i
		go func() {
			goname := fmt.Sprintf("并发序号 %d", num)
			j := 1
			for counter.LessThan(p.times) {
				start := time.Now()
				_, err := p.request()
				if err != nil {
					fmt.Printf("请求失败 error %s \n", err.Error())
				}
				consumed := time.Since(start)
				p.result.Add(consumed)
				fmt.Printf("%s 第 %d 次请求，耗时 %d ms \n", goname, j, consumed.Milliseconds())
				counter.PlusOne()
				j ++
				p.wg.Done()
			}
		}()
	}
	p.wg.Wait()
}

func (p *PressureTestTool) request() (*http.Response, error) {
	resp, err := p.client.Get(p.url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PressureTestTool) PrintResult(percentile float64) {
	fmt.Printf("平均请求耗时 %d ms \n", p.result.Avg().Milliseconds())
	fmt.Printf("百分位 %f 耗时 %d ms \n", percentile, p.result.Percentile(percentile).Milliseconds())
}

type Counter struct {
	current int
	mux sync.Mutex
}

func (c *Counter) PlusOne() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.current ++
}

func (c *Counter) LessThan(num int) bool {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.current < num
}