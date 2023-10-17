package environment

import (
	"reflect"
	"testing"
)

func TestGetSecretsEnvVars(t *testing.T) {
	t.Setenv("GETSECRETS_MYDB1_ID", "some-id")
	tests := []struct {
		name string
		want []string
	}{
		{name: "asdfdsf", want: []string{"GETSECRETS_MYDB1_ID=some-id"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSecretsEnvVars(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSecretsEnvVars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSecretsEnvVarsMap(t *testing.T) {
	type args struct {
		vars []string
	}
	tests := []struct {
		name       string
		args       args
		wantResult map[string]string
	}{
		{name: "sdf",
			args: args{vars: []string{
				"GETSECRETS_MYDB1_SECRETID=some=id",
			}},
			wantResult: map[string]string{"MYDB1_SECRETID": "some=id"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := GetSecretsEnvVarsMap(tt.args.vars); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("GetSecretsEnvVarsMap() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
