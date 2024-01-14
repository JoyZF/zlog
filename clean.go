package zlog

import (
	"fmt"
	"github.com/robfig/cron"
	defaultLog "log"
	"os"
	"path/filepath"
	"time"
)

type Cleaner interface {
	Division() // 日志分割
	Clean()    // 日志清理
}

type Clean struct {
	Interval time.Duration // 定时扫描
	Reserve  time.Duration // 保留时间
}

const DivisionSpec = "*/1 * * * * *"

func (c *Clean) Division() {
	// 开启一个定时任务, 每小时执行一次 将当前日志文件重命名为当前时间戳
	_ = cron.New().AddFunc(DivisionSpec, func() {
		fmt.Println("Division")
		log.mu.Lock()
		defer log.mu.Unlock()
		f := log.opt.serviceName + log.opt.fileName
		if !IsExist(f) {
			return
		}
		dateStr := time.Unix(time.Now().Unix()-60, 0).Format("2006010215")
		//dateStr := time.Unix(time.Now().Unix()-60, 0).Format("200601021504")
		os.Rename(f, f+"."+dateStr)
	})
	return
}

func (c *Clean) Clean() {
	go c.clean()
	return
}

func (c *Clean) clean() {
	defer func() {
		if err := recover(); err != nil {
			defaultLog.Println("log cleaner stop !!! error : ", err)
		}
	}()
	fmt.Println("start")
	ticker := time.NewTicker(c.Interval)
	go func() {
		for {
			<-ticker.C
			c.do()
		}
	}()
}

func (c *Clean) do() {
	path := log.opt.path
	// 遍历文件夹中的目录 最后修改时间在保留时间之前的文件删除
	_ = filepath.WalkDir(path, func(path string, d os.DirEntry, err error) error {
		if d.IsDir() {
			// 跳过 默认日志文件夹平铺
			return nil
		}
		fileInfo, _ := d.Info()
		if time.Now().Sub(fileInfo.ModTime()) > c.Reserve {
			// 释放文件
			_ = os.Remove(path)
		}

		return nil
	})
}

// IsExist  判断文件夹/文件是否存在  存在返回 true
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
