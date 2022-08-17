package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type FileData struct {
	LastEditedTime time.Time
	Path           string
}

type MatchingPair struct {
	OriginalFile string
	CopiedFile   string
}

func findDuplicates(rootDir string) ([]*MatchingPair, error) {
	seen := make(map[string]*FileData)
	dirStack := &Stack{
		items: []string{},
	}
	dirStack.Push(rootDir)
	duplicates := []*MatchingPair{}
	for len(dirStack.items) > 0 {
		currentPath := dirStack.Pop()
		file, err := os.Open(currentPath)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		fileInfo, err := file.Stat()
		if err != nil {
			return nil, err
		}
		if fileInfo.IsDir() {
			files, err := ioutil.ReadDir(currentPath)
			if err != nil {
				return nil, err
			}
			for _, path := range files {
				dirStack.Push(filepath.Join(currentPath, path.Name()))
			}
		} else {
			hash, err := sha256sum(currentPath)
			if err != nil {
				return nil, err
			}
			currentLastEditTime := fileInfo.ModTime()
			if _, ok := seen[hash]; ok {
				existingFileData := seen[hash]
				diff := existingFileData.LastEditedTime.Sub(currentLastEditTime)
				if diff > 0 {
					duplicates = append(duplicates, &MatchingPair{
						OriginalFile: currentPath,
						CopiedFile:   existingFileData.Path,
					})
					seen[hash] = &FileData{
						LastEditedTime: currentLastEditTime,
						Path:           currentPath,
					}
				} else {
					duplicates = append(duplicates, &MatchingPair{
						OriginalFile: existingFileData.Path,
						CopiedFile:   currentPath,
					})
				}
			} else {
				seen[hash] = &FileData{
					LastEditedTime: currentLastEditTime,
					Path:           currentPath,
				}
			}
		}
	}
	return duplicates, nil
}

func sha256sum(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

type Stack struct {
	items []string
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() string {
	if len(s.items) < 1 {
		return ""
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Peek() string {
	if len(s.items) < 1 {
		return ""
	}
	return s.items[len(s.items)-1]
}
