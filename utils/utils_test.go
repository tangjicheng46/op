package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestRemoveSpecific(t *testing.T) {
	// 创建临时目录作为测试数据
	root, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			fmt.Println(err)
			return
		}
	}(root)

	// 在临时目录中创建要删除的文件
	fileToDelete := filepath.Join(root, "delete.me")
	if err := os.WriteFile(fileToDelete, []byte("test data"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// 在临时目录中创建一个子目录并在其中创建要删除的文件
	subDir := filepath.Join(root, "subdir")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("Failed to create subdir: %v", err)
	}

	fileToDeleteInSubdir := filepath.Join(subDir, "delete.me")
	if err := os.WriteFile(fileToDeleteInSubdir, []byte("test data"), 0644); err != nil {
		t.Fatalf("Failed to create test file in subdir: %v", err)
	}

	// 执行 RemoveSpecific
	if err := RemoveSpecific(root, "delete.me"); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// 检查文件是否被正确删除
	if _, err := os.Stat(fileToDelete); !os.IsNotExist(err) {
		t.Errorf("Expected file to be deleted, but it still exists")
	}
	if _, err := os.Stat(fileToDeleteInSubdir); !os.IsNotExist(err) {
		t.Errorf("Expected file in subdir to be deleted, but it still exists")
	}

	// 测试无法遍历的目录
	invalidDir := filepath.Join(root, "nonexistent")
	if err := RemoveSpecific(invalidDir, "delete.me"); err == nil {
		t.Errorf("Expected error for non-existent directory, got nil")
	}

	// 测试删除时遇到的错误（例如权限问题）- 这是一个稍微复杂的测试，可能需要更多的设置
}
