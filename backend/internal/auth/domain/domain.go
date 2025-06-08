package domain

type User struct {
	Id     string
	Name   string
	Avatar string
}

// 有効なユーザーかどうかを確認するメソッド
// Entity (Domain) 層はかなり抽象的
func (u *User) IsValid() bool {
	return u.Id != "" && u.Name != "" && u.Avatar != ""
}
