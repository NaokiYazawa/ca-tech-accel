package model

import "sync"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TODO: 排他制御しよう
// TODO: mysqlを導入して永続化しよう
var userMap = map[int]User{1: {ID: 1, Name: "Naoki Yazawa"}}
// model package の中で参照できるように小文字の変数
var rwlock sync.RWMutex

// Create 新規ユーザ登録
func Create(name string) (User, error) {
	rwlock.RLock()
	userID := len(userMap) + 1
	rwlock.RUnlock()
	user := User{ID: userID, Name: name}
	rwlock.Lock()
	userMap[userID] = user
	rwlock.Unlock()
	return user, nil
}

// List ユーザ一覧取得
func List() ([]User, error) {
	rwlock.RLock()
	users := make([]User, 0, len(userMap))
	for _, v := range userMap {
		users = append(users, v)
	}
	rwlock.RUnlock()
	return users, nil
}

//　Find ユーザ取得
func Find(id int) (User, error) {
	rwlock.RLock()
	defer rwlock.RUnlock()
	return userMap[id], nil
}

// Update ユーザ更新
// update には id と name が必要
// error のみを返す
func Update(id int, name string) error {
	user := User{ID: id, Name: name}
	rwlock.Lock()
	userMap[id] = user
	rwlock.Unlock()
	return nil
}
