package controllers

import (
	"github.com/gofiber/fiber/v2"
	"majoo.id/satu/config"
	"majoo.id/satu/models"
)

type response_d struct {
	models.Pagination
	Data []models.OpsiD `json:"data"`
}

func OpsiD(c *fiber.Ctx) error {
	db := config.DB
	ctx := c.Locals("ctx").(*models.Context)

	// Parse pagination query
	var pagination models.Pagination
	c.QueryParser(&pagination)
	pagination.Format()

	var result []models.OpsiD

	start := "2021-11-01"
	end := "2021-11-30"

	// Get transactions with merchant name, outlet_name, and omzet
	if err := db.Raw(`WITH recursive cte AS (
		select ? as date
		union all
		select date + interval 1 day from cte where DATE(date) < ?)
		select cte.date,merchants.merchant_name,outlets.outlet_name as outlet_name,
 		IF(DATE(cte.date) = DATE(transactions.created_at),sum(transactions.bill_total),0) as omzet
 		from cte LEFT JOIN transactions on DATE(transactions.created_at) <= cte.date
 		inner join merchants on merchants.id = transactions.merchant_id
 		inner join users on users.id = merchants.user_id
		inner join outlets on outlets.id = transactions.outlet_id AND merchants.id = outlets.merchant_id
 		where users.id = ?
 		group by DATE(cte.date) order by cte.date asc LIMIT ?,?`, start, end, ctx.User.ID, pagination.Start, pagination.PageSize).Scan(&result).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Sent JSON to client
	return c.Status(fiber.StatusOK).JSON(response_d{Pagination: pagination, Data: result})
}
