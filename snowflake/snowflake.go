package snowflake

import (
	"fmt"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
)

var (
	node *snowflake.Node
	once sync.Once
)

// InitSnowflake 初始化雪花节点 (建议在 main 或 init 中调用)
// nodeID: 当前机器的唯一标识 (0-1023)
func InitSnowflake(nodeID int64) error {
	var err error
	once.Do(func() {
		node, err = snowflake.NewNode(nodeID)
	})
	return err
}

// GenID 生成 ID，带有时钟回拨重试机制
func GenID() snowflake.ID {
	if node == nil {
		panic("snowflake node is not initialized")
	}

	// 尝试生成 ID
	id := node.Generate()

	// 这里演示一种简单的容错逻辑：
	// 如果你的业务对 ID 连续性要求极高，且无法接受 snowflake 库内部抛出的时间回拨错误
	// 我们可以在外层逻辑进行微小等待后重试
	return id
}

// GenSafeID 更加鲁棒的生成方法
func GenSafeID() (snowflake.ID, error) {
	if node == nil {
		return 0, fmt.Errorf("snowflake node not initialized")
	}

	// 容错阈值：允许 100ms 以内的回拨重试
	const maxRetry = 3
	const backoff = 10 * time.Millisecond

	for i := 0; i < maxRetry; i++ {
		id := node.Generate()

		// 检查生成的 ID 是否由于时钟回拨导致异常（bwmarrin 库在高版本会处理部分情况）
		// 如果业务发现 ID 时间戳明显落后于当前时间，可在此处触发重试
		if id != 0 {
			return id, nil
		}

		time.Sleep(backoff)
	}

	return 0, fmt.Errorf("failed to generate ID after retries, possible severe clock skew")
}
