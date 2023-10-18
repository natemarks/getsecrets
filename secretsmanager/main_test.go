package secretsmanager

import "testing"

func TestGetSecretValue(t *testing.T) {
	type args struct {
		secretName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "sdf", args: args{secretName: "DNASecretsSandboxApplexones-zZcO0lSih1TE"}, want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSecretValue(tt.args.secretName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSecretValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetSecretValue() got = %v, want %v", got, tt.want)
			}
		})
	}
}
