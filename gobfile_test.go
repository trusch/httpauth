package httpauth

import (
	"os"
	"testing"
)

func TestInitGobFileAuthBackend(t *testing.T) {
	os.Remove(file)
	b, err := NewGobFileAuthBackend(file)
	if err != ErrMissingBackend {
		t.Fatal(err.Error())
	}

	_, err = os.Create(file)
	if err != nil {
		t.Fatal(err.Error())
	}
	b, err = NewGobFileAuthBackend(file)
	if err != nil {
		t.Fatal(err.Error())
	}
	if b.filepath != file {
		t.Fatal("File path not saved.")
	}
	if len(b.users) != 0 {
		t.Fatal("Users initialized with items.")
	}

	testBackend(t, b)
}

func TestGobReopen(t *testing.T) {
	b, err := NewGobFileAuthBackend(file)
	if err != nil {
		t.Fatal(err.Error())
	}
	b.Close()
	b, err = NewGobFileAuthBackend(file)
	if err != nil {
		t.Fatal(err.Error())
	}

	testBackend2(t, b)
}
