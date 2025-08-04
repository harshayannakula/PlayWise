package resume

import (
	"testing"
)

func TestPauseAndResume(t *testing.T) {
	rm := NewResumeManager()

	rm.Pause("Rock", 3)
	rm.Pause("Rock", 5)

	index, err := rm.Resume("Rock")
	if err != nil || index != 5 {
		t.Errorf("Expected index 5, got %d", index)
	}

	index, err = rm.Resume("Rock")
	if err != nil || index != 3 {
		t.Errorf("Expected index 3, got %d", index)
	}
}

func TestPeek(t *testing.T) {
	rm := NewResumeManager()

	rm.Pause("Chill", 2)
	index, err := rm.Peek("Chill")
	if err != nil || index != 2 {
		t.Errorf("Peek failed: expected 2, got %d", index)
	}
}

func TestResumeEmpty(t *testing.T) {
	rm := NewResumeManager()
	_, err := rm.Resume("Unknown")
	if err == nil {
		t.Errorf("Expected error on empty resume stack")
	}
}

func TestClearStack(t *testing.T) {
	rm := NewResumeManager()
	rm.Pause("Sleep", 1)
	rm.ClearStack("Sleep")
	_, err := rm.Resume("Sleep")
	if err == nil {
		t.Errorf("Expected error after clearing stack")
	}
}
