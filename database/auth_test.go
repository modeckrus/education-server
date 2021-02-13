package database

import (
	"context"
	"education/config"
	"education/model"
	"testing"
)

func TestDB_UpdateUser(t *testing.T) {
	type args struct {
		ctx         context.Context
		currUser    model.CheckedUser
		name        string
		displayName string
		iphotoID    *string
	}
	ctx := context.Background()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "UpdateTest",
			args: args{
				ctx,
				model.CheckedUser{
					ID: "6027c58a8ec27d4d1708bff7",
					Roles: []string{
						"user",
					},
				},
				"New Super Name",
				"How now a days",
				nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := config.NewServerConfig()
			db := Connect(conf)
			got, err := db.UpdateUser(tt.args.ctx, tt.args.currUser, tt.args.name, tt.args.displayName, tt.args.iphotoID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("got nil")
			}
			t.Log(got)
			want := model.User{
				Name:        tt.args.name,
				DisplayName: tt.args.displayName,
			}
			if got.Name != want.Name {
				t.Errorf("Name not same")
			}
			if got.DisplayName != want.DisplayName {
				t.Errorf("Display name not same")
			}

		})
	}
}
