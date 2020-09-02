package model

import "strings"

var FilterCache map[string]int64

func SuggestFilter(_input string) []string {
    result := make([]string, 0)
    for k, _ := range FilterCache {
        if strings.Contains(strings.ToLower(k), strings.ToLower(_input)) {
            result = append(result, k)
        }
    }
    return result
}

// 建立过滤器缓存
func cacheFilter() error {
    //TODO use memcache or redis

    FilterCache = make(map[string]int64)

    dao := NewCollectionDAO(nil)
    count, err := dao.Count()
    if nil != err {
        return err
    }

    tags, err := dao.List(0, count)
    if nil != err {
        return err
    }

    for _, tag := range tags {
        // 添加Code
        if _, ok := FilterCache[tag.Code]; !ok {
            FilterCache[tag.Code] = 0
        }
        FilterCache[tag.Code] = FilterCache[tag.Code] + 1
        // 添加Name
        if _, ok := FilterCache[tag.Name]; !ok {
            FilterCache[tag.Name] = 0
        }
        FilterCache[tag.Name] = FilterCache[tag.Name] + 1
        // 添加Keyword
        for _, kw := range tag.Keyword {
            if _, ok := FilterCache[kw]; !ok {
                FilterCache[kw] = 0
            }
            FilterCache[kw] = FilterCache[kw] + 1
        }
        // 添加Alias
        for _, kw := range tag.Alias {
            if _, ok := FilterCache[kw]; !ok {
                FilterCache[kw] = 0
            }
            FilterCache[kw] = FilterCache[kw] + 1
        }
    }
    return nil
}

func increaseFilter(_tags []*Tag) {
    for _, tag := range _tags {
        FilterCache[tag.Code] = FilterCache[tag.Code] + 1
        FilterCache[tag.Name] = FilterCache[tag.Name] + 1
        for _, kw := range tag.Keyword {
            FilterCache[kw] = FilterCache[kw] + 1
        }
        for _, kw := range tag.Alias {
            FilterCache[kw] = FilterCache[kw] + 1
        }
    }
}

func reduceFilter(_tags []*Tag) {
    for _, tag := range _tags {
        FilterCache[tag.Code] = FilterCache[tag.Code] - 1
        FilterCache[tag.Name] = FilterCache[tag.Name] - 1
        for _, kw := range tag.Keyword {
            FilterCache[kw] = FilterCache[kw] - 1
        }
        for _, kw := range tag.Alias {
            FilterCache[kw] = FilterCache[kw] - 1
        }
    }
}
