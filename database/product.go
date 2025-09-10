package database

var productList []Product // empty slice of Product

type Product struct {
	// first letter capitalized means it is exported and can be accessed outside the package (Access modifier :Public)
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	ImageURL string `json:"imageUrl"` 
	
	/* what if we want to write small letter variable in json?
	we can use json tag to specify the name of the field in the JSON object
	Example: json:"id" means the field will be represented as "id" in the JSON object

	for example: ID int `json:"id"`

	we can also use json:"-" to ignore the field in the JSON object
	we can also use json:"omitempty" to omit the field if it is empty
	*/
}

func Store(p Product) Product{
	p.ID = len(productList) + 1 // set the ID of the new product
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func Get(productID int) *Product {
	for _, product := range productList {

		if product.ID == productID {
			return &product
		}
	}
	return nil
}

func Update(product Product) {
	for idx, p := range productList {

		if p.ID == product.ID {
			productList[idx] = product
		}
	}
}

func Delete(productID int) {

	var tempList []Product
	for _, p := range productList {

		if p.ID != productID {
			tempList = append(tempList,p)
		}
	}
	productList = tempList
}	

func init(){
	prod1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Fresh Orange",
		Price:       105,
		ImageURL:    "https://static.vecteezy.com/system/resources/previews/015/606/509/non_2x/sweet-orange-fruit-photo.jpg",
	}
	prod2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Deshi Apple",
		Price:       150,
		ImageURL:    "https://t3.ftcdn.net/jpg/12/65/25/16/360_F_1265251658_IhcbSYTueEiyDk39cuU00uyMGi7Os0ut.jpg",
	}
	prod3 := Product{
		ID: 		3,
		Title:       "Banana",
		Description: "Norshindir Banana",
		Price:       50,
		ImageURL:    "https://static.vecteezy.com/system/resources/previews/035/524/351/non_2x/bunch-banana-cartoon-illustration-fruit-and-food-concept-design-flat-style-isolated-white-background-clip-art-icon-design-vector.jpg",
	}
	prod4 := Product{
		ID:          4,	
		Title:       "Mango",
		Description: "Rajshahir Mango",	
		Price:       120,
		ImageURL:    "https://www.shutterstock.com/image-photo/red-mango-on-tree-green-600nw-2455963709.jpg",
	}
	prod5 := Product{
		ID:          5,
		Title:       "Pineapple",
		Description: "Modhupurer Pineapple",
		Price:       10,
		ImageURL:    "https://miro.medium.com/v2/resize:fit:1400/1*itympMq6N_UCCM4fPa5p8Q.jpeg",
	}
	prod6 := Product{
		ID:          6,	
		Title:       "Watermelon",
		Description: "Fresh Watermelon",
		Price:       80,
		ImageURL:    "https://whatscookingamerica.net/wp-content/uploads/2015/03/Watermelon-Sliced-Eqyptian4.jpg",
	}

	// append products to the productList slice
	productList = append(productList, prod1)
	productList = append(productList, prod2)
	productList = append(productList, prod3)	
	productList = append(productList, prod4)
	productList = append(productList, prod5)
	productList = append(productList, prod6) 
}