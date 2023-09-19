package utils

import (
	"blog/internal/models"
	"context"
	"fmt"
)

var ctx = context.Background()

// 进入队列
func Pub(data []byte) error {
	_, err := models.Rdb.Publish(ctx, "comment", data).Result()
	if err != nil {
		return err
	}
	return nil
}

// 读取
func Sub() {
	pubsub := models.Rdb.Subscribe(ctx, "comment")
	// 使用完毕，记得关闭
	defer pubsub.Close()
	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Println(msg.Channel, msg.Payload)
	}

}
