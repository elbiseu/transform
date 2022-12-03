package transform

import (
	"image"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestScale(t *testing.T) {
	file1, err := os.Open("testdata/48934279448_ea28959a86_o.jpg")
	if err != nil {
		log.Fatal(err)
	}
	srcImg, _, err := image.Decode(file1)
	if err != nil {
		log.Fatal(err)
	}
	file1.Close()
	file2, err := os.Open("testdata/output.jpg")
	if err != nil {
		log.Fatal(err)
	}
	dstImg, _, err := image.Decode(file2)
	if err != nil {
		log.Fatal(err)
	}
	file2.Close()
	type args struct {
		img image.Image
		x   int
		y   int
	}
	tests := []struct {
		name    string
		args    args
		want    image.Image
		wantErr bool
	}{
		{
			name: "RX100 VI",
			args: args{
				img: srcImg,
				x:   10,
				y:   10,
			},
			want:    dstImg,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Scale(tt.args.img, tt.args.x, tt.args.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("Scale() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scale() got = %v, want %v", got, tt.want)
			}
		})
	}
}
