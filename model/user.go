package model

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TODO: 排他制御しよう
// TODO: mysqlを導入して永続化しよう
var userMap = map[int]User{1: {ID: 1, Name: "Naoki Yazawa"}}

// Create 新規ユーザ登録
func Create(name string) (User, error) {
	userID := len(userMap) + 1
	user := User{ID: userID, Name: name}
	userMap[userID] = user
	return user, nil
}

// List ユーザ一覧取得
func List() ([]User, error) {
	beasts := make([]User, 0, len(userMap))
	for _, v := range userMap {
		beasts = append(beasts, v)
	}
	return beasts, nil
}

//　Find ユーザ取得
func Find(id int) (User, error) {
	return userMap[id], nil
}
