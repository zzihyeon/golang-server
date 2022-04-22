package mongodb

// func InsertLocationData(input mongoModel.LoactionHistoryData) {
// 	client := GetMongoClient()
// 	input.Timestamp = time.Now()
// 	_, err := client.Database("cavea").Collection("location_log").InsertOne(context.Background(), input)
// 	if err != nil {
// 		log.Println("Insert Location Log Data Error: ", err.Error())
// 	}
// }

// func GetLastLocation(input mongoModel.LocationHistoryInput) (*mongoModel.LoactionHistoryData, error) {
// 	client := GetMongoClient()
// 	loc := &mongoModel.LoactionHistoryData{}
// 	findOptions := options.FindOne()
// 	findOptions.SetSort(bson.D{{Key: "timestamp", Value: -1}})
// 	data := client.Database("cavea").Collection("location_log").FindOne(context.Background(), bson.M{"metaData.user_id": input.UserID}, findOptions)
// 	if data == nil {
// 		log.Println("Get Location Error!!")
// 		return &mongoModel.LoactionHistoryData{}, nil
// 	}
// 	err := data.Decode(loc)
// 	if err != nil {
// 		log.Println("Get Location Error@@")
// 		return &mongoModel.LoactionHistoryData{}, err
// 	}
// 	return loc, nil
// }
