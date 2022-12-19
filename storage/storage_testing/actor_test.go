package storage_test

import (
	"context"
	"testing"

	"github.com/mokh1rbek/film_CRUD/models"

	faker "github.com/bxcodec/faker/v3"
	"github.com/google/go-cmp/cmp"
)

func TestActorCreate(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.CreateActor
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.CreateActor{
				First_name: faker.Name(),
				Last_name: faker.Name(),
			},
			WantErr: false,
		},
		{
			Name: "case 2",
			Input: &models.CreateActor{
				First_name: faker.Name(),
				Last_name: faker.Name(),
			},
			WantErr: true,
		},
		{
			Name: "case 3",
			Input: &models.CreateActor{
				First_name: faker.Name(),
				Last_name: faker.Name(),
			},
			WantErr: false,
		},
	}

	for _, tc := range tests {

		t.Run(tc.Name, func(t *testing.T) {

			got, err := actorRepo.Create(
				context.Background(),
				tc.Input,
			)

			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			if got == "" {
				t.Errorf("%s: got: %v", tc.Name, err)
				return
			}
		})
	}

}

func TestActorGetById(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.ActorPrimarKey
		Output  *models.Actor
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.ActorPrimarKey{
				Id: "",
			},
			Output: &models.Actor{
				Id: "",
				First_name: "",
				Last_name: "",
				CreatedAt: "",
				UpdatedAt: "",
				},
			WantErr: false,
		},
	}

	for _, tc := range tests {

		t.Run(tc.Name, func(t *testing.T) {

			got, err := actorRepo.GetByPKey(
				context.Background(),
				tc.Input,
			)

			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			comparer := cmp.Comparer(func(x, y models.Category) bool {
				return x.Name == y.Name
			})

			if diff := cmp.Diff(tc.Output, got, comparer); diff != "" {
				t.Errorf(diff)
				return
			}
		})
	}

}

func TestActorUpdate(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.UpdateActor
		Output  *models.Actor
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.UpdateActor{
				First_name: "",
				Last_name: "",
			},
			Output: &models.Actor{
				First_name: "",
				Last_name: "",
			},
			WantErr: false,
		},
	}

	for _, tc := range tests {

		t.Run(tc.Name, func(t *testing.T) {

			_, err := actorRepo.Update(
				context.Background(),
				"",
				tc.Input,
			)

			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			res, err := actorRepo.GetByPKey(
				context.Background(),
				&models.ActorPrimarKey{
					Id: tc.Input.Id,
				},
			)

			comparer := cmp.Comparer(func(x, y models.Film) bool {
				return x.Title == y.Title
			})

			if diff := cmp.Diff(tc.Output, res, comparer); diff != "" {
				t.Errorf(diff)
				return
			}
		})
	}

}

func TestActorDelete(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.ActorPrimarKey
		Want    string
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.ActorPrimarKey{
				Id: "",
			},
			Want:    "no rows in result set",
			WantErr: false,
		},
	}

	for _, tc := range tests {

		t.Run(tc.Name, func(t *testing.T) {

			err := actorRepo.Delete(
				context.Background(),
				tc.Input,
			)

			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			_, err = actorRepo.GetByPKey(
				context.Background(),
				&models.ActorPrimarKey{
					Id: tc.Input.Id,
				},
			)

			comparer := cmp.Comparer(func(x, y string) bool {
				return x == y
			})

			if diff := cmp.Diff(tc.Want, err.Error(), comparer); diff != "" {
				t.Errorf(diff)
				return
			}
		})
	}

}
