package stack

import (
	"errors"
	gomock "go.uber.org/mock/gomock"
	mock_stack "stackpj/mocks"
	"testing"
)

func TestStackImpl_Push(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := mock_stack.NewMockStack(ctrl)
	s.EXPECT().Push(1).Return(nil)

	s.Push(1)
}

func TestStacImpl_Peek(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := mock_stack.NewMockStack(ctrl)
	s.EXPECT().Peek().Return(1, nil)

	val, e := s.Peek()
	if val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}

	if e != nil {
		t.Errorf(e.Error())
	}

	s.EXPECT().Peek().Return(nil, errors.New("Empty Stack"))

	val, e = s.Peek()
	if val != nil {
		t.Errorf("Expected nil, got %d", val)
	}

	if e == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestStackImpl_Pop(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := mock_stack.NewMockStack(ctrl)
	s.EXPECT().Pop().Return(1, nil)

	val, e := s.Pop()
	if val != 1 && e == nil {
		t.Errorf("Expected 1, got %d", val)
	}
}
