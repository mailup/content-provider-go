package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mailup/content-provider-go/model"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.Default()
	router.GET("content/products/:email", ListRecommendedProducts)
	router.Run(":" + port)
}

// ListRecommendedProducts return a list of recommended products
func ListRecommendedProducts(ctx *gin.Context) {
	email := ctx.Param("email")
	content := &model.ExternalContent{User: GetUserProfile(email),
		RecommendedProducts: make([]model.Product, 3)}

	content.RecommendedProducts[0] = model.Product{
		Name:        "Simon Rattle Edition: The Second Viennese School",
		Currency:    "$",
		Price:       "10",
		Brand:       "Calssic Music",
		ContentLink: "https://www.amazon.it/Simon-Rattle-Second-Viennese-School/dp/B008PCHYFO/ref=sr_1_5?s=music&ie=UTF8&qid=1496754624&sr=1-5&keywords=classica",
		ImageLink:   "https://images-eu.ssl-images-amazon.com/images/I/51FhD2kajxL._SS500.jpg",
		Discount:    "0",
		ProductID:   "1",
		ProductLink: "https://www.amazon.it/Simon-Rattle-Second-Viennese-School/dp/B008PCHYFO/ref=sr_1_5?s=music&ie=UTF8&qid=1496754624&sr=1-5&keywords=classica",
		Text:        "The Second Viennese School of Simon Rattle edited by Warner Classics",
		Title:       "Simon Rattle",
		SubTitle:    "The Second Viennese School",
		Categories:  []model.Category{model.Category{CategoryID: "1", CategoryName: "Classica", CategoryLink: "https://www.amazon.it/s/ref=nb_sb_noss_2?__mk_it_IT=%C3%85M%C3%85%C5%BD%C3%95%C3%91&url=search-alias%3Dpopular&field-keywords=classica"}},
	}
	content.RecommendedProducts[1] = model.Product{
		Name:        "I Magnifici Della Musica Classica",
		Currency:    "$",
		Price:       "12",
		Brand:       "Calssic Music",
		ContentLink: "https://www.amazon.it/I-Magnifici-Della-Musica-Classica/dp/B016NGIGNC/ref=sr_1_3?s=music&ie=UTF8&qid=1496754624&sr=1-3&keywords=classica",
		ImageLink:   "https://images-eu.ssl-images-amazon.com/images/I/610pGjbgbgL._AC_US327_FMwebp_QL65_.jpg",
		Discount:    "0",
		ProductID:   "2",
		ProductLink: "https://www.amazon.it/I-Magnifici-Della-Musica-Classica/dp/B016NGIGNC/ref=sr_1_3?s=music&ie=UTF8&qid=1496754624&sr=1-3&keywords=classica",
		Text:        "I Magnifici Della Musica Classica edited by The Saifam Group",
		Title:       "I Magnifici Della Musica Classica",
		SubTitle:    "Mozart, Verdi, Beethoven, Rossini, Vivaldi",
		Categories:  []model.Category{model.Category{CategoryID: "1", CategoryName: "Classica", CategoryLink: "https://www.amazon.it/s/ref=nb_sb_noss_2?__mk_it_IT=%C3%85M%C3%85%C5%BD%C3%95%C3%91&url=search-alias%3Dpopular&field-keywords=classica"}},
	}
	content.RecommendedProducts[2] = model.Product{
		Name:        "Tchaikovsky Edition",
		Currency:    "$",
		Price:       "14",
		Brand:       "Calssic Music",
		ContentLink: "https://www.amazon.it/Tchaikovsky-Various-Artists/dp/B00YO3Z2MA/ref=sr_1_9?s=music&ie=UTF8&qid=1496754624&sr=1-9&keywords=classica",
		ImageLink:   "https://images-eu.ssl-images-amazon.com/images/I/518N66tnMGL._AC_US327_FMwebp_QL65_.jpg",
		Discount:    "10%",
		ProductID:   "3",
		ProductLink: "https://www.amazon.it/Tchaikovsky-Various-Artists/dp/B00YO3Z2MA/ref=sr_1_9?s=music&ie=UTF8&qid=1496754624&sr=1-9&keywords=classica",
		Text:        "Deutsche Grammophon,Box Classica,Romantico,",
		Title:       "Tchaikovsky Edition",
		SubTitle:    "Masterworks Edition",
		Categories:  []model.Category{model.Category{CategoryID: "1", CategoryName: "Classica", CategoryLink: "https://www.amazon.it/s/ref=nb_sb_noss_2?__mk_it_IT=%C3%85M%C3%85%C5%BD%C3%95%C3%91&url=search-alias%3Dpopular&field-keywords=classica"}},
	}

	ctx.JSON(http.StatusOK, content)
}

// GetUserProfile returns the profile of the user
func GetUserProfile(email string) *model.UserProfile {
	return &model.UserProfile{
		Email:           email,
		Address:         "Via Roma",
		City:            "Milan",
		Company:         "MailUp",
		DeliveryAddress: "Via Palermo",
		DeliveryCity:    "Cremona",
		FirstName:       "Mario",
		LastName:        "Rossi",
		ZipCode:         "26100",
	}
}
