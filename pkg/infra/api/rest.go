package api

import (
	"hex/pkg/domain"
	"hex/pkg/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RestAPI struct {
	service domain.ProductService
}

func NewRestAPI(svc domain.ProductService) *RestAPI {
	return &RestAPI{
		service: svc,
	}
}

func (api *RestAPI) Run() {
	e := echo.New()
	ctx := e.AcquireContext()

	e.GET("/product", handleAll(api, ctx))
	e.GET("/product/:id", handleFind(api, ctx))
	e.POST("/product", handleCreate(api, ctx))
	e.DELETE("/product/:id", handleDelete(api, ctx))

	e.Logger.Fatal(e.Start(":1323"))
}

func handleAll(api *RestAPI, c echo.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		pp, err := api.service.All()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, pp)
	}
}

func handleFind(api *RestAPI, c echo.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		p, err := api.service.Find(id)

		if err != nil {
			return c.JSON(http.StatusNotFound, err)
		}

		return c.JSON(http.StatusOK, p)
	}

}

func handleCreate(api *RestAPI, c echo.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := new(types.Product)

		if err := c.Bind(p); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		product := types.ProductDTO{
			Id:    p.Id,
			Name:  p.Name,
			Brand: p.Brand,
			Price: p.Price,
		}

		err := api.service.Add(types.Product(product))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, nil)
	}
}

func handleDelete(api *RestAPI, c echo.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		err := api.service.Delete(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, nil)
	}

}
