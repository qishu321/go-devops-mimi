package tools

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"go-devops-mimi/server/config"
	"net"
)

func CheckSafeKey(key string) bool {
	if key == config.Conf.AgentConfig.Token {
		return true
	}

	return false
}

// IPNameToNum 生成唯一的 uint64 ID
func IPNameToNum(name string, ip string) uint64 {
	ipBytes := net.ParseIP(ip).To4()
	if ipBytes == nil {
		fmt.Println("Invalid IP:", ip)
		return 0
	}

	// IP 转 uint32
	ipNum := binary.BigEndian.Uint32(ipBytes)

	// 计算 name 哈希并取前 4 字节
	nameHash := md5.Sum([]byte(name))
	nameNum := binary.BigEndian.Uint32(nameHash[:4])

	// 组合两个数并取模 10^10
	return (uint64(ipNum)<<32 | uint64(nameNum)) % 10000000000
}

// func IPNameToNum(name string, ip string) uint64 {
// 	ret := big.NewInt(0)
// 	ipBytes := net.ParseIP(ip).To4()
// 	if ipBytes == nil {
// 		return 0
// 	}

// 	// 计算 name 的哈希值，确保唯一性
// 	nameHash := md5.Sum([]byte(name))

// 	// 取前 4 个字节（IP）+ 取 nameHash 的前 4 个字节 组合成 8 字节
// 	fullBytes := append(ipBytes, nameHash[:4]...) // 总共 8 字节
// 	ret.SetBytes(fullBytes)

// 	return ret.Uint64()
// }
