package pkg

import "testing"

func TestExternelSort_Sort(t *testing.T) {
	tests := []struct {
		name    string
		es      *ExternelSort
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.es.Sort(); (err != nil) != tt.wantErr {
				t.Errorf("ExternelSort.Sort() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
