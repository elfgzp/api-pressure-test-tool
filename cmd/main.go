package main

import (
	"flag"
	"github.com/elfgzp/api-pressure-test-tool/internal/pressuretesttool"
)

var (
	URL string
	Times int
	Concurrent int
	Percentile float64
)

func init() {
	flag.StringVar(&URL, "u", "", "压测 URL")
	flag.StringVar(&URL, "url", "", "压测 URL")

	flag.IntVar(&Times, "t", 100, "请求次数")
	flag.IntVar(&Times, "times", 100, "请求次数")

	flag.IntVar(&Concurrent, "c", 10, "并发次数")
	flag.IntVar(&Concurrent, "concurrent", 10, "并发次数")

	flag.Float64Var(&Percentile, "p", 0.95, "百分位耗时")
	flag.Float64Var(&Percentile, "percentile", 0.95, "百分位耗时")
}

func main() {
	flag.Parse()
	tool := pressuretesttool.New(URL, Times, Concurrent)
	tool.Run()
	tool.PrintResult(Percentile)
}