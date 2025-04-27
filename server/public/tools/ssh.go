package tools

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type SSHClientConfig struct {
	UserName   string        `form:"username" json:"username" binding:"required"`
	Password   string        `form:"password" json:"password"`
	PublicIP   string        `form:"public_ip" json:"public_ip" binding:"required"`
	Port       int           `form:"port" json:"port" binding:"required"`
	Command    string        `form:"command" json:"command"`
	AuthModel  string        `form:"authmodel" json:"authmodel" binding:"required"`
	PrivateKey string        `form:"private_key" json:"private_key"`
	Timeout    time.Duration `form:"timeout" json:"timeout"` //超时时间
}

func SshCommand(conf *SSHClientConfig, command string) (string, error) {
	config := &ssh.ClientConfig{
		User:            conf.UserName,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 忽略know_hosts检查
		Timeout:         conf.Timeout,
	}
	switch conf.AuthModel {
	case "password":
		config.Auth = []ssh.AuthMethod{ssh.Password(conf.Password)}
	case "privateKey":
		signer, err := ssh.ParsePrivateKey([]byte(conf.PrivateKey))
		if err != nil {
			return "", err
		}
		config.Auth = []ssh.AuthMethod{ssh.PublicKeys(signer)}
	}
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", conf.PublicIP, conf.Port), config)
	if err != nil {
		return "失败", err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	output, err := session.CombinedOutput(command)
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// CreateFileOnRemoteServer 在远程服务器上创建文件
func CreateFileOnRemoteServer(sshConfig *SSHClientConfig, filename, typename, content string) (string, error) {
	absoluteFilePath := "/tmp/" + filename + "." + typename

	// Escape special characters and format the content for bash script
	escapedContent := strings.ReplaceAll(content, "'", `'\''`)

	// Construct the bash script content
	scriptContent := fmt.Sprintf("echo '%s' > %s", escapedContent, absoluteFilePath)

	// Execute the combined script as a single SSH command
	command := fmt.Sprintf("%s && %s %s", scriptContent, typename, absoluteFilePath)

	output, err := SshCommand(sshConfig, command)
	if err != nil {
		fmt.Println("SSH Error:", err)
		return "", fmt.Errorf("Failed to execute SSH command: %v", err)
	}

	fmt.Println("SSH Output:", output) // Optional: Print the SSH output

	return output, nil
}

// UploadFileToHost 用来把本地文件上传到远程主机
// conf: SSH 连接配置
// localPath: 本地文件完整路径，如 "/tmp/app.tar.gz"
// remoteDir: 远程目标目录，如 "/opt/app"（函数会自动 mkdir -p）
// 返回：上传后的远程文件全路径，或 error
func UploadFileToHost(conf *SSHClientConfig, localPath, remoteDir string) (string, error) {
	// --- 1. 准备 SSH clientConfig ---
	clientCfg := &ssh.ClientConfig{
		User:            conf.UserName,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         conf.Timeout,
	}
	switch strings.ToLower(conf.AuthModel) {
	case "password":
		clientCfg.Auth = []ssh.AuthMethod{ssh.Password(conf.Password)}
	case "privatekey", "private_key":
		signer, err := ssh.ParsePrivateKey([]byte(conf.PrivateKey))
		if err != nil {
			return "", fmt.Errorf("parse private key: %w", err)
		}
		clientCfg.Auth = []ssh.AuthMethod{ssh.PublicKeys(signer)}
	default:
		return "", fmt.Errorf("unsupported auth model: %s", conf.AuthModel)
	}

	// --- 2. 建立 SSH 连接 ---
	addr := fmt.Sprintf("%s:%d", conf.PublicIP, conf.Port)
	sshClient, err := ssh.Dial("tcp", addr, clientCfg)
	if err != nil {
		return "", fmt.Errorf("ssh dial %s failed: %w", addr, err)
	}
	defer sshClient.Close()

	// --- 3. 创建 SFTP 客户端 ---
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		return "", fmt.Errorf("new sftp client failed: %w", err)
	}
	defer sftpClient.Close()

	// --- 4. 打开本地文件 ---
	srcFile, err := os.Open(localPath)
	if err != nil {
		return "", fmt.Errorf("open local file %s: %w", localPath, err)
	}
	defer srcFile.Close()

	// --- 5. 确保远程目录存在 ---
	if err := sftpClient.MkdirAll(remoteDir); err != nil {
		return "", fmt.Errorf("mkdir remote dir %s: %w", remoteDir, err)
	}

	// --- 6. 在远程创建目标文件 ---
	remoteFileName := path.Base(localPath)
	remoteFullPath := path.Join(remoteDir, remoteFileName)
	dstFile, err := sftpClient.Create(remoteFullPath)
	if err != nil {
		return "", fmt.Errorf("create remote file %s: %w", remoteFullPath, err)
	}
	defer dstFile.Close()

	// --- 7. 拷贝文件内容 ---
	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return "", fmt.Errorf("copy data to remote file: %w", err)
	}

	// --- 8. （可选）设置权限，比如 0644 ---
	if err := sftpClient.Chmod(remoteFullPath, 0644); err != nil {
		return "", fmt.Errorf("chmod remote file: %w", err)
	}

	return remoteFullPath, nil
}
