package config

import (
	"reflect"
	"testing"
)

func TestGetConfig(t *testing.T) {
	type args struct {
		varMap map[string]string
	}
	tests := []struct {
		name        string
		args        args
		wantConfig  Config
		wantSecrets map[string]Secret
	}{
		{
			name: "sdf",
			args: args{
				varMap: map[string]string{
					"DEBUG":          "true",
					"MYDB1_FILENAME": "/etc/mydb1",
					"MYDB1_SECRETID": "some-id",
					"MYDB1_DATABASE": "mydb1",
					"MYDB2_FILENAME": "/etc/mydb2",
					"MYDB2_SECRETID": "some-id2",
					"MYDB2_DATABASE": "mydb2",
				},
			},
			wantConfig: Config{
				Debug: true,
			},
			wantSecrets: map[string]Secret{
				"MYDB1": Secret{
					Name:     "MYDB1",
					Filename: "/etc/mydb1",
					SecretId: "some-id",
					Database: "mydb1",
				},
				"MYDB2": Secret{
					Name:     "MYDB2",
					Filename: "/etc/mydb2",
					SecretId: "some-id2",
					Database: "mydb2",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotConfig, gotSecrets := GetConfig(tt.args.varMap)
			if !reflect.DeepEqual(gotConfig, tt.wantConfig) {
				t.Errorf("GetConfig() gotConfig = %v, want %v", gotConfig, tt.wantConfig)
			}
			if !reflect.DeepEqual(gotSecrets, tt.wantSecrets) {
				t.Errorf("GetConfig() gotSecrets = %v, want %v", gotSecrets, tt.wantSecrets)
			}
		})
	}
}
