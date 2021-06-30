package mozillainstallhash

import "testing"

func TestMozillaInstallHash(t *testing.T) {
	type args struct {
		installPath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Default Windows Firefox Release Path",
			args: args{
				installPath: "C:\\Program Files\\Mozilla Firefox",
			},
			want: "308046B0AF4A39CB",
		},
		{
			name: "Default Windows Firefox Nightly Path",
			args: args{
				installPath: "C:\\Program Files\\Firefox Nightly",
			},
			want: "6F193CCC56814779",
		},
		{
			name: "Default macOS Firefox Release Path",
			args: args{
				installPath: "/Applications/Firefox.app/Contents/MacOS",
			},
			want: "2656FF1E876E9973",
		},
		{
			name: "Default macOS Firefox Nightly Path",
			args: args{
				installPath: "/Applications/Firefox Nightly.app/Contents/MacOS",
			},
			want: "31210A081F86E80E",
		},
		{
			name: "Default Arch Linux Firefox Release Path",
			args: args{
				installPath: "/usr/lib/firefox",
			},
			want: "4F96D1932A9F858E",
		},
		{
			name: "Default Arch Linux Firefox Nightly Path",
			args: args{
				installPath: "/opt/firefox-nightly",
			},
			want: "6BA5C87ECB35E12F",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := MozillaInstallHash(tt.args.installPath); err != nil || got != tt.want {
				t.Errorf("MozillaInstallHash() = %v, %v, want %v, <nil>", got, err, tt.want)
			}
		})
	}
}
