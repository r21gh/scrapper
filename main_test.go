package main

import (
	"reflect"
	"testing"
)

func Test_dumpCSV(t *testing.T) {
	type args struct {
		path    string
		headers []string
		records <-chan []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := dumpCSV(tt.args.path, tt.args.headers, tt.args.records); (err != nil) != tt.wantErr {
				t.Errorf("dumpCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fetch(t *testing.T) {
	type args struct {
		url     string
		id      int
		queries []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fetch(tt.args.url, tt.args.id, tt.args.queries)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fetch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
