package controllers

import (
	"github.com/gofiber/fiber/v2"
	"majoo.id/satu/config"
	"majoo.id/satu/models"
)

type response_c struct {
	models.Pagination
	Data []models.OpsiC `json:"data"`
}

func OpsiC(c *fiber.Ctx) error {
	db := config.DB
	ctx := c.Locals("ctx").(*models.Context)

	// Parse pagination query
	var pagination models.Pagination
	c.QueryParser(&pagination)
	pagination.Format()

	var result []models.OpsiC
	start := "2021-11-01"
	end := "2021-11-30"

	// Get transactions with merchant name, and omzet
	if err := db.Raw(`WITH recursive cte AS (
		select ? as created_at
		union all
		select created_at + interval 1 day from cte where DATE(created_at) < ?)
 		select cte.created_at as date,merchants.merchant_name,
 		IF(DATE(cte.created_at) = DATE(transactions.created_at),sum(transactions.bill_total),0) as omzet
 		from cte LEFT JOIN transactions on DATE(transactions.created_at) <= cte.created_at
 		inner join merchants on merchants.id = transactions.merchant_id
 		inner join users on users.id = merchants.user_id
 		where users.id = ?
 		group by DATE(cte.created_at) order by cte.created_at asc LIMIT ?,?`, start, end, ctx.User.ID, pagination.Start, pagination.PageSize).Scan(&result).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Sent JSON to client
	return c.Status(fiber.StatusOK).JSON(response_c{Pagination: pagination, Data: result})
}
