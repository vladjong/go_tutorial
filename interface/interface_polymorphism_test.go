package gotutorial

import (
	"go_tutorial/interface/mocks"
	"testing"
)

func TestUseCase_CheckUser(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		expect  func(m *mocks.Storage)
		wantErr bool
	}{
		{
			name: "ok",
			args: args{"1"},
			expect: func(m *mocks.Storage) {
				m.
					On("User", "1").
					Return("Vasy")
			},
		},
		{
			name: "empy name",
			args: args{"1"},
			expect: func(m *mocks.Storage) {
				m.
					On("User", "1").
					Return("")
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Иницилизация мока
			mockStorage := mocks.NewStorage(t)
			tt.expect(mockStorage)

			u := &UseCase{
				repo: mockStorage,
			}
			if err := u.CheckUser(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UseCase.CheckUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
