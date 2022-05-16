package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"online-exam/global"
	"online-exam/model"
	"os"
	"sort"
	"strconv"
)

// 获取算法推荐内容
// input:
//   [[q1, q2, ...], [a1, a2, ...]]
//   eg: [[10, 4, 3], [1, 0, 1]]
// output:
//   [q1, q2, q3]
func GetKTRecommendations(exercises [][]int) ([]int, error) {
	WINDOW_SIZE := 10 // 算法推荐题目数量
	ALG_LEN := 50     // 算法所需的用户最大练习长度

	if len(exercises) > ALG_LEN {
		exercises = exercises[:ALG_LEN]
	}

	type SeqRequest struct {
		Seq [][]int `json:"seq"`
	}
	type TagsMap struct {
		key string
	}
	var tags map[string]int
	tagsFile, _ := os.Open("kt/tags.json")
	tagsData, _ := ioutil.ReadAll(tagsFile)
	json.Unmarshal(tagsData, &tags)
	fmt.Println("exercises", exercises)
	for i, _ := range exercises[0] {
		prevTagId := exercises[0][i]
		exercises[0][i] = tags[strconv.Itoa(prevTagId)]
	}
	defer tagsFile.Close()
	// exercises = [][]int{{1}, {1}}
	seq := SeqRequest{
		Seq: exercises,
	}
	jsonExer, _ := json.Marshal(seq)
	fmt.Println(string(jsonExer))
	// res, err := http.Post("http://127.0.0.1:8080/v1/models/dkt3:explain", "application/json", bytes.NewBuffer(jsonExer))
	res, err := http.Post("https://exam-kt.dev.feel.ac.cn:8443/v1/models/dkt3:explain", "application/json", bytes.NewBuffer(jsonExer))
	if err != nil || res.StatusCode != 200 {
		fmt.Println(err)
		return []int{}, errors.New("获取算法推荐题目出错")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return []int{}, errors.New("获取算法推荐题目出错")
	}
	var results []float64
	json.Unmarshal(body, &results)
	fmt.Println(results, len(results))
	// 推荐策略：预测答对率最低的10个知识点，每个知识点选择一道题目
	type RecommendElement struct {
		RightPossiblity float64
		TagID           int
	}
	topResults := make([]RecommendElement, len(results))
	// 初始化
	for i, _ := range topResults {
		topResults[i].RightPossiblity = results[i]
		topResults[i].TagID = i
	}
	sort.Slice(topResults, func(i, j int) bool {
		return topResults[i].RightPossiblity < topResults[j].RightPossiblity
	})
	topResults = topResults[:WINDOW_SIZE]
	fmt.Println(topResults)

	// 根据算法推荐检索题目详情
	tagsFile, _ = os.Open("kt/tags_r.json")
	tagsData, _ = ioutil.ReadAll(tagsFile)
	json.Unmarshal(tagsData, &tags)

	questionIDs := make([]int, WINDOW_SIZE)
	for i, _ := range topResults {
		var questionTag model.QuestionsTags
		fmt.Println(topResults[i].TagID)
		global.DB.Where("tags_id = ?", tags[strconv.Itoa(topResults[i].TagID)]).Select("*").Order("rand()").First(&questionTag)
		fmt.Println(questionTag)
		questionIDs[i] = int(questionTag.TagsId)
	}
	fmt.Println(questionIDs)
	return questionIDs, nil
}
