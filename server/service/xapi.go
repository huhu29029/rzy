package service

import (
	"context"
	"online-exam/global"
	"online-exam/model"
)

func NewXAPIData(data model.XAPIData) {
	client := global.MongoDB

	coll := client.Database(global.CONFIG.Mongo.Dbname).Collection("xAPI")

	_, err := coll.InsertOne(context.TODO(), data)
	if err != nil {
		panic(err)
	}

	// var result bson.M
	// err = coll.FindOne(context.TODO(), data).Decode(&result)
	// if err == mongo.ErrNoDocuments {
	// 	fmt.Print("No document was found")
	// 	return
	// }
	// if err != nil {
	// 	panic(err)
	// }

	// jsonData, err := json.MarshalIndent(result, "", "    ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", jsonData)
}
