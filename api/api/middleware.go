package api

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

const KeyQuery = "api:query"

func QueryParser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var s string
			q := DefaultQuery()
			s = c.QueryParam("offset")
			if offset, err := strconv.ParseInt(s, 10, 32); err == nil {
				q.Pagination.Offset = int(offset)
			}
			s = c.QueryParam("limit")
			if limit, err := strconv.ParseInt(s, 10, 32); err == nil {
				q.Pagination.Limit = int(limit)
			}
			if sortBy := c.QueryParam("sortBy"); sortBy != "" {
				q.OrderBy = sortBy
			}
			s = c.QueryParam("desc")
			if desc, err := strconv.ParseBool(s); err == nil {
				q.Descending = desc
			}
			q.Filter = Filter(c.QueryParam("filter"))

			// TODO: validate

			c.Set(KeyQuery, q)
			return next(c)
		}
	}
}

func QueryFromContext(c echo.Context) Query {
	tmp := c.Get(KeyQuery)
	if tmp == nil {
		return DefaultQuery()
	}
	if q, ok := tmp.(Query); ok {
		return q
	}
	return DefaultQuery()
}
