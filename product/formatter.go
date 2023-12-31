package product

import "strings"

type ProductFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	Price       	 int    `json:"price"`
	Discount       	 int    `json:"discount"`
	Discountwo       int    `json:"discountwo"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
	Perks            []string  `json:"perks"`

}

func FormatProduct(product Product) ProductFormatter {

	productFormatter := ProductFormatter{}


	var perks []string
	for _, perk := range strings.Split(product.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	productFormatter.ID = product.ID
	productFormatter.UserID = product.UserID
	productFormatter.Name = product.Name
	productFormatter.ShortDescription = product.ShortDescription
	productFormatter.Price = product.Price
	productFormatter.Discount = product.Discount
	productFormatter.Discountwo = product.Discountwo
	productFormatter.CurrentAmount = product.CurrentAmount
	productFormatter.Slug = product.Slug
	productFormatter.ImageURL = ""
	productFormatter.Perks = perks
	if len(product.ProductImages) > 0 {
		productFormatter.ImageURL = product.ProductImages[0].FileName

	}

	

	return productFormatter
}

func FormatProducts(products []Product) []ProductFormatter {

	productsFormatter := []ProductFormatter{}

	for _, product := range products {
		productFormatter := FormatProduct(product)
		productsFormatter = append(productsFormatter, productFormatter)
	}

	return productsFormatter
}

type ProductDetailFormatter struct {
	ID               int                      `json:"id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_description"`
	Description      string                   `json:"description"`
	ImageURL         string                   `json:"image_url"`
	Price            int                      `json:"price"`
	Discount	     int					  `json:"discount"`
	Discountwo       int                      `json:"discountwo"`
	Pricetwo         int                      `json:"pricetwo"`
	Pricetree        int                      `json:"pricetree"`
	CurrentAmount    int                      `json:"current_amount"`
	BackerCount      int                      `json:"backer_count"`
	UserID           int                      `json:"user_id"`
	Slug             string                   `json:"slug"`
	Perks            []string                 `json:"perks"`
	User             ProductUserFormatter    `json:"user"`
	Images           []ProductImageFormatter `json:"images"`
}

type ProductUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type ProductImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatProductDetail(product Product) ProductDetailFormatter {
	productDetailFormatter := ProductDetailFormatter{}

	productDetailFormatter.ID = product.ID
	productDetailFormatter.Name = product.Name
	productDetailFormatter.ShortDescription = product.ShortDescription
	productDetailFormatter.Description = product.Description
	productDetailFormatter.Price = product.Price
	productDetailFormatter.CurrentAmount = product.CurrentAmount
	productDetailFormatter.BackerCount = product.BackerCount
	productDetailFormatter.UserID = product.UserID
	productDetailFormatter.Slug = product.Slug
	productDetailFormatter.Discount = product.Discount
	productDetailFormatter.Discountwo = product.Discountwo
    productDetailFormatter.Pricetwo = product.Price - (product.Discount * product.Price)  / 100
	productDetailFormatter.Pricetree = product.Price - (product.Discountwo *  product.Price)  /100 
	productDetailFormatter.ImageURL = ""

	if len(product.ProductImages) > 0 {
		productDetailFormatter.ImageURL = product.ProductImages[0].FileName
	}

	var perks []string
	for _, perk := range strings.Split(product.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	productDetailFormatter.Perks = perks

	user := product.User

	productUserFormatter := ProductUserFormatter{}
	productUserFormatter.Name = user.Name
	productUserFormatter.ImageURL = user.AvatarFileName

	productDetailFormatter.User = productUserFormatter

	images := []ProductImageFormatter{}

	for _, image := range product.ProductImages {
		productImageFormatter := ProductImageFormatter{}
		productImageFormatter.ImageURL = image.FileName

		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		productImageFormatter.IsPrimary = isPrimary

		images = append(images, productImageFormatter)
	}

	productDetailFormatter.Images = images

	return productDetailFormatter
}
