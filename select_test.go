package mquery

import (
	"testing"
)

func Test_selectQueryBuidler_ToQuery(t *testing.T) {

	tests := []struct {
		name string
		sqb  SelectQueryBuilder
		want string
	}{
		{
			name: "select all",
			sqb:  NewSelectBuilder("user"),
			want: "SELECT * FROM user",
		},
		{
			name: "select with field",
			sqb:  NewSelectBuilder("user").Fields("id", "username", "password"),
			want: "SELECT id,username,password FROM user",
		},
		{
			name: "select where and",
			sqb:  NewSelectBuilder("user").Fields("username", "password").And("id", 5),
			want: "SELECT username,password FROM user WHERE id = 5",
		},
		{
			name: "select where or",
			sqb:  NewSelectBuilder("user").Fields("username", "password").Or("id", 5),
			want: "SELECT username,password FROM user WHERE id = 5",
		},
		{
			name: "select where and or mix",
			sqb:  NewSelectBuilder("user").Fields("username", "password").And("id", 5).Or("username", "hahaha").And("password", "haha"),
			want: "SELECT username,password FROM user WHERE id = 5 and password = \"haha\" or username = \"hahaha\"",
		},
		{
			name: "select order by ASC",
			sqb:  NewSelectBuilder("user").Fields("username", "password").OrderByASC("username"),
			want: "SELECT username,password FROM user ORDER BY username ASC",
		}, {
			name: "select order by ASC",
			sqb:  NewSelectBuilder("user").Fields("username", "password").OrderByASC("username"),
			want: "SELECT username,password FROM user ORDER BY username ASC",
		}, {
			name: "select JOIN",
			sqb:  NewSelectBuilder("user").Fields("username", "password").Join("account", "id_user", "id_user").And("balance", 5),
			want: "SELECT username,password FROM user JOIN account ON user.id_user = account.id_user WHERE balance = 5",
		}, {
			name: "select IN",
			sqb:  NewSelectBuilder("user").Fields("username", "password").In("id_user", 1, 2, 3, 4, 5),
			want: "SELECT username,password FROM user WHERE id_user IN (1,2,3,4,5)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qb := tt.sqb
			if got := qb.ToQuery(); got != tt.want {
				t.Errorf("selectQueryBuidler.ToQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}