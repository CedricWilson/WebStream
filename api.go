package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func DirRequest(c *gin.Context) {
	dirPath := c.DefaultQuery("dir", ".")
	var folderList []ListingInfo
	var fileList []ListingInfo

	fmt.Println(dirPath)
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}

	for _, e := range entries {

		if e.IsDir() {
			folderList = append(folderList, ListingInfo{FileName: e.Name(), Type: "folder"})
		} else {
			filepath := dirPath[1:] + e.Name()
			fileList = append(fileList, ListingInfo{FileName: e.Name(), Type: "file", FilePath: filepath})
		}

	}

	c.JSON(200, append(folderList, fileList...))
}

type ListingInfo struct {
	Type     string `json:"type"`
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
}
