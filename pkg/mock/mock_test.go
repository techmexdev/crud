package mock_test

import (
	"testing"

	"github.com/techmexdev/crud"
	"github.com/techmexdev/crud/pkg/mock"
)

func TestStorage(t *testing.T) {
	us := mock.NewUserStore()

	uu := []crud.User{
		{FirstName: "John", LastName: "Lennon", Username: "jlennon"},
		{FirstName: "Paul", LastName: "McCartney", Username: "pmccartney"},
		{FirstName: "George", LastName: "Harrison", Username: "gharrison"},
		{FirstName: "Ringo", LastName: "Star", Username: "rstarr"},
	}

	for _, u := range uu {
		err := us.Create(u)
		if err != nil {
			t.Fatalf("could not create %v: %s", u, err)
		}
	}

	ringo, err := us.Read("rstarr")
	if err != nil {
		t.Fatalf("could not read  ringo: %s", err)
	}

	ringo.LastName = "Starr"

	err = us.Update("rstarr", ringo)
	if err != nil {
		t.Fatalf("could not update ringo : %s", err)
	}

	newRingo, err := us.Read("rstarr")
	if err != nil {
		t.Fatalf("could not update read new ringo : %s", err)
	}

	if newRingo.LastName != "Starr" {
		t.Fatalf("did not update ringo")
	}

	err = us.Delete("pmccartney")
	if err != nil {
		t.Fatalf("could not delete paul : %s", err)
	}

	beatles, err := us.List()
	if err != nil {
		t.Fatalf("could not list beatles : %s", err)
	}

	for _, b := range beatles {
		if b.Username == "pmccartney" {
			t.Fatal("did not delete paul")
		}
	}
}
