package main

import "fmt"

type bar struct {
	percent int64  //完成百分比,即cur/total
	cur     int64  //目前完成数目
	total   int64  //总共完成数目
	rate    string //进度条
	graph   string //进度条的显示符号
}

// 返回一个新的Bar
func NewBar(cur int64, total int64) *bar {
	b := &bar{
		cur:   cur,
		total: total,
		graph: "#",
	}
	b.getPercent()
	for i := int64(0); i < b.percent; i++ {
		b.rate += b.graph
	}

	return b
}

// 返回一个新的Bar,自定义graph
func NewBarWithGraph(cur int64, total int64, graph string) *bar {
	b := NewBar(cur, total)
	b.graph = graph
	return b
}

func (b *bar) getPercent() int64 {
	return int64(float32(b.cur) / float32(b.total) * 100)
}

// 主要函数，展示百分比进度
func (b *bar) Display(cur int64) {
	b.cur = cur
	lastPercent := b.percent
	percent := b.getPercent()
	if percent != lastPercent {
		b.percent = percent
		for i := lastPercent; i < b.percent; i++ {
			b.rate += b.graph
		}
	}
	fmt.Printf("\r[%-100s]%3d%%  %8d/%d", b.rate, b.percent, b.cur, b.total)
}
