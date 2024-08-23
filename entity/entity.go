package entity

type User struct{
	Name string `json:"name"`
	UserName string `json:"username"`
	Posts[] string `json:"posts"`
		Followings := make(map[string]string) `json:"followings"`
	Followers := make(map[string]string) `json:"followers"`

}
type FollowUser struct{
	CurrentuserName string `json:"current"`
	FollowingUserName string `json:"following"`
}
