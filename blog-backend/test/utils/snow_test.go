package utils

import (
	"blog-backend/app/dao/utils"
	"fmt"
	"testing"
	"time"
)

// @Author: OxCAFFEE
// @Github: https://github.com/OxCaffee
// @Email: wwh2021@mail.ustc.edu.cn
// @Date: 2021/8/11-16:53

func TestConstTimestamp(t *testing.T) {
	fmt.Println(time.Now().UnixNano() / 1e6)
}

func TestTableSnowUIDGeneration(t *testing.T) {
	tableWorker1 := utils.TableWorker{
		TableId:       1,
		SequenceId:    1,
		LastTimeStamp: uint64(time.Now().UnixNano()/1e6 - 1),
	}

	tableWorker2 := utils.TableWorker{
		TableId:       2,
		SequenceId:    1,
		LastTimeStamp: uint64(time.Now().UnixNano()/1e6 - 1000),
	}

	go func() {
		for {
			id, err := tableWorker1.NextID()
			<-time.After(5 * time.Second / 10)
			if err != nil {
				fmt.Println(err.Error())
				break
			} else {
				fmt.Println("worker1:", id)
			}
		}
	}()

	go func() {
		go func() {
			for {
				id, err := tableWorker2.NextID()
				<-time.After(5 * time.Second / 10)
				if err != nil {
					fmt.Println(err.Error())
					break
				} else {
					fmt.Println("worker2:", id)
				}
			}
		}()
	}()

	<-time.After(100 * time.Second)
}
