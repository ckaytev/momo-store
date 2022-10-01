package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.praktikum-services.ru/Stasyan/momo-store/cmd/api/dependencies"
)

func TestFakeAppIntegrational(t *testing.T) {
	store, err := dependencies.NewFakeDumplingsStore()
	assert.NoError(t, err)
	app, err := NewInstance(store)
	assert.NoError(t, err)

	t.Run("create_order", func(t *testing.T) {
		for i := 1; i <= 10; i++ {
			t.Run("id"+strconv.Itoa(i), func(t *testing.T) {
				r := httptest.NewRequest("POST", "/orders", nil)
				w := httptest.NewRecorder()
				app.CreateOrderController(w, r)

				assert.Equal(t, http.StatusOK, w.Code)
				assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
				fmt.Fprintln(os.Stdout, "_____")
				fmt.Fprintln(os.Stdout, w.Body.String())
				fmt.Fprintln(os.Stdout, "_____")

				expectedJSON, err := json.Marshal(map[string]interface{}{"id": i})
				assert.NoError(t, err)
				assert.JSONEq(t, string(expectedJSON), w.Body.String())
			})
		}
	})

	t.Run("list_dumplings", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/packs", nil)
		w := httptest.NewRecorder()
		app.ListDumplingsController(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

		fmt.Fprintln(os.Stdout, "_____")
		fmt.Fprintln(os.Stdout, w.Body.String())
		fmt.Fprintln(os.Stdout, "_____")

		expectedJSON := "{\"results\":[{\"id\":1,\"name\":\"Пельмени\",\"price\":5,\"description\":\"С говядиной\",\"image\":\"https://res.cloudinary.com/sugrobov/image/upload/v1651932687/repos/momos/8dee5a92281746aa887d6f19cf9fdcc7.jpg\"},{\"id\":2,\"name\":\"Хинкали\",\"price\":3.5,\"description\":\"Со свининой\",\"image\":\"https://res.cloudinary.com/sugrobov/image/upload/v1651932687/repos/momos/50b583271fa0409fb3d8ffc5872e99bb.jpg\"},{\"id\":3,\"name\":\"Манты\",\"price\":2.75,\"description\":\"С мясом молодых бычков\",\"image\":\"https://storage.yandexcloud.net/momo-store-zateevant/momos/8b50f76f514a4ccaaacdcb832a1b3a2f.jpg\"},{\"id\":4,\"name\":\"Буузы\",\"price\":4,\"description\":\"С телятиной и луком\",\"image\":\"https://storage.yandexcloud.net/momo-store-zateevant/momos/788c073d83c14b3fa00675306dfb32b5.jpg\"},{\"id\":5,\"name\":\"Цзяоцзы\",\"price\":7.25,\"description\":\"С говядиной и свининой\",\"image\":\"https://storage.yandexcloud.net/momo-store-zateevant/momos/32cc88a33c3243a6a8838c034878c564.jpg\"},{\"id\":6,\"name\":\"Гедза\",\"price\":3.5,\"description\":\"С соевым мясом\",\"image\":\"https://storage.yandexcloud.net/momo-store-zateevant/momos/7685ad7e9e634a58a4c29120ac5a5ee1.jpg\"},{\"id\":7,\"name\":\"Дим-самы\",\"price\":2.65,\"description\":\"С уткой\",\"image\":\"https://storage.yandexcloud.net/momo-store-zateevant/momos/4bdaeab0ee1842dc888d87d4a435afdd.jpg\"},{\"id\":8,\"name\":\"Момо\",\"price\":5,\"description\":\"С бараниной\",\"image\":\"https://storage.yandexcloud.net/momo-store-zateevant/momos/f64dcea998e34278a0006e0a2b104710.jpg\"},{\"id\":9,\"name\":\"Вонтоны\",\"price\":4.1,\"description\":\"С креветками\",\"image\":\"https://res.cloudinary.com/sugrobov/image/upload/v1651932687/repos/momos/8dee5a92281746aa887d6f19cf9fdcc7.jpg\"},{\"id\":10,\"name\":\"Баоцзы\",\"price\":4.2,\"description\":\"С капустой\",\"image\":\"https://res.cloudinary.com/sugrobov/image/upload/v1651932687/repos/momos/50b583271fa0409fb3d8ffc5872e99bb.jpg\"},{\"id\":11,\"name\":\"Кундюмы\",\"price\":5.45,\"description\":\"С грибами\",\"image\":\"https://storage.yandexcloud.net/momo-store-zateevant/momos/8b50f76f514a4ccaaacdcb832a1b3a2f.jpg\"},{\"id\":12,\"name\":\"Курзе\",\"price\":3.25,\"description\":\"С крабом\",\"image\":\"https://storage.yandexcloud.net/momo-store-zateevant/momos/788c073d83c14b3fa00675306dfb32b5.jpg\"},{\"id\":13,\"name\":\"Бораки\",\"price\":4,\"description\":\"С говядиной и бараниной\",\"image\":\"https://storage.yandexcloud.net/momo-store-zateevant/momos/7685ad7e9e634a58a4c29120ac5a5ee1.jpg\"},{\"id\":14,\"name\":\"Равиоли\",\"price\":2.9,\"description\":\"С рикоттой\",\"image\":\"https://storage.yandexcloud.net/momo-store-zateevant/momos/4bdaeab0ee1842dc888d87d4a435afdd.jpg\"}]}\n"

		assert.NoError(t, err)
		assert.JSONEq(t, string(expectedJSON), w.Body.String())
	})

	t.Run("healthcheck", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/health", nil)
		w := httptest.NewRecorder()
		app.HealthcheckController(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
