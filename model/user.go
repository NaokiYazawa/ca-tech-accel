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
	users := make([]User, 0, len(userMap))
	for _, v := range userMap {
		users = append(users, v)
	}
	return users, nil
}

//　Find ユーザ取得
func Find(id int) (User, error) {
	return userMap[id], nil
}

// Update ユーザ更新
// update には id と name が必要
// error のみを返す
func Update(id int, name string) (error) {
	user := User{ID: id, Name: name}
	userMap[id] = user
	return nil
}
