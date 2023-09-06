package lead

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rohannair11/go-crm/database"
)

type Lead struct {
	gorm.Model
	Name    string
	Company string
	Emali   string
	Phone   int
}

func getLeads(c *fiber.Ctx) {

}
func getLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)

}
func NewLead() {

}
func DeleteLead() {

}
