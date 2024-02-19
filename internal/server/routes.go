package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "website/components/scraper"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)
	// r.GET("/updateDatabase", s.updateDatabaseHandler)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

// func (s *Server) updateDatabaseHandler(c *gin.Context) {
// 	OmnibusList:=scraper.Scrape("https://www.instocktrades.com/newreleases")
// 	c.JSON(http.StatusOK, s.db.Health())
// }


// func CreateOmnibus(c*fiber.Ctx) error {
// 	omnibus := new(models.Omnibus)
// 	if err := c.BodyParser(omnibus); err != nil{
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}
// 	database.DB.Db.Create(&omnibus)
// 	return c.Status(200).JSON(omnibus)
// }

// func GetAllOmnibus(c*fiber.Ctx) error {
// 	omnibus := []models.Omnibus{}
// 	database.DB.Db.Find(&omnibus)

// 	return c.Status(200).JSON(omnibus)

// }

// func ScrapeAndCreate(c*fiber.Ctx) error{
// 	OmnibusList:=scraper.Scrape("https://www.instocktrades.com/newreleases")
// 	// jsonData, err := json.Marshal(OmnibusList)
// 	// if err != nil {
// 	// 	fmt.Println("Error marshaling JSON:", err)
// 	// 	return
// 	// }

// 	for _, omnibus := range OmnibusList {
// 		if err := database.DB.Db.Create(&omnibus).Error; err != nil {
// 			fmt.Printf("Error inserting %s into the database: %s\n", omnibus.Name, err.Error())
// 		}
// 	}

// 	return c.Status(200).JSON(fiber.Map{
// 		"message": "Data inserted into the database successfully",
// 	})
// }

