package util

import "github.com/bwmarrin/snowflake"

var NodeID int64 = 1

func SnowflakeID() (string, error) {
	snowflake.NodeBits = 8                 //工作机器id
	snowflake.StepBits = 14                //序列号id
	node, err := snowflake.NewNode(NodeID) //传入工作节点nodeID  不同机器上要维护不同的nodeId
	if err != nil {
		return "", err
	}
	Id := node.Generate() //产生一个64bit大小的整数为一个int64型 (转换成字符后长度19位)

	s := Id.String()
	return s, nil
}
