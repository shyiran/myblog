package models

import "strings"

func HandleTagsListData(tags []string) map[string]int {
	var tagsMap = make(map[string]int)
	for _, tag := range tags {
		//普及&保险&基础知识
		// 普及、保险、基础知识
		// 客户端，get，post等方法
		tagList := strings.Split(tag, "&")
		for _, value := range tagList {
			tagsMap[value]++
		}
	}
	return tagsMap
}

