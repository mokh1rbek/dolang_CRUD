package order_service

import (
	"fmt"
	"net/http"
	"sync"
	"testing"

	"github.com/mokh1rbek/film_CRUD/models"

	"github.com/bxcodec/faker/v3"
	"github.com/test-go/testify/assert"
)

var k int64

func TestActor(t *testing.T) {
	k = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			id := createActor(t)
			deleteActor(t, id)
		}()

	}

	wg.Wait()

	fmt.Println("k: ", k)
}

func createActor(t *testing.T) string {
	response := &models.Actor{}

	request := &models.CreateActor{
		First_name: faker.Name(),
		Last_name:  faker.Name(),
	}

	resp, err := PerformRequest(http.MethodPost, "/actor", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	fmt.Println(response)

	return response.Id
}

func UpdateActor(t *testing.T, id string) string {
	response := &models.Actor{}
	request := &models.UpdateActor{
		Id:         id,
		First_name: faker.Name(),
		Last_name:  faker.Name(),
	}

	resp, err := PerformRequest(http.MethodPut, "/actor/"+id, request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 200)
	}

	fmt.Println(resp)

	return response.Id
}

func deleteActor(t *testing.T, id string) string {

	resp, _ := PerformRequest(
		http.MethodDelete,
		fmt.Sprintf("/actor/%k", id),
		nil,
		nil,
	)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return ""
}
