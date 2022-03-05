package main

import "testing"

func Test_md5Hash(t *testing.T) {
	tests := []struct {
		name string
		in   []byte
		want string
	}{
		{
			name: "empty string",
			in:   []byte(""),
			want: "d41d8cd98f00b204e9800998ecf8427e",
		},
		{
			name: "Answer to the Ultimate Question of Life, the Universe, and Everything",
			in:   []byte("42"),
			want: "a1d0c6e83f027327d8461063f4ac58a6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := md5Hash(tt.in); got != tt.want {
				t.Errorf("md5Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_worker(t *testing.T) {
	tests := []struct {
		name      string
		in        chan string
		outResult [][2]string
	}{
		{
			name: "simple",
			in: func() chan string {
				ch := make(chan string, 1)
				ch <- ""
				close(ch)
				return ch
			}(),
			outResult: [][2]string{{"", ""}},
		},
		{
			name: "simple",
			in: func() chan string {
				ch := make(chan string, 1)
				ch <- "google"
				close(ch)
				return ch
			}(),
			outResult: [][2]string{{"google", "google"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := make(chan [2]string, len(tt.in))
			worker(tt.in, out, func(addr string) string { return addr })
			for _, expected := range tt.outResult {
				if current := <-out; current != expected {
					t.Errorf("expected %s got %s", expected, current)
				}
			}
		})
	}
}
