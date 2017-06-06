package model

// ExternalContent is the content provided by external APIs
type ExternalContent struct {
	Flow                *FlowControl
	Title               string
	SubTitle            string
	Text                string
	ImageLink           string
	ContentLink         string
	User                *UserProfile
	RecommendedProducts []Product
	PurchaseHistory     []Order
	BlogPosts           []BlogPost
	DynamicFields       map[string]string
}

// FlowControl defines flow behavioral elements
type FlowControl struct {
	// Do not send messages
	DoNotSend bool
	// Do not cache profile data
	DoNotCache bool
	// Profile data TTL
	TTL int64
}

// UserProfile defines the user profile fields
type UserProfile struct {
	Title            string
	SubTitle         string
	Text             string
	ImageLink        string
	ContentLink      string
	FirstName        string
	LastName         string
	Gender           string
	BirthYear        string
	BirthMonth       string
	BirthDay         string
	Address          string
	Address2         string
	City             string
	ZipCode          string
	Province         string
	Region           string
	State            string
	Nation           string
	JobTitle         string
	Company          string
	Email            string
	Prefix           string
	Telephone        string
	DeliveryAddress  string
	DeliveryAddress2 string
	DeliveryCity     string
	DeliveryZipCode  string
	DeliveryProvince string
	DeliveryRegion   string
	DeliveryState    string
	DeliveryNation   string
	Engagement       string
	WishList         []Product
	AbandonedBasket  []Product
	Tags             []string
	ProfileLink      string
}

// Category of a product.
type Category struct {
	Title        string
	SubTitle     string
	Text         string
	ImageLink    string
	ContentLink  string
	CategoryID   string
	CategoryName string
	CategoryLink string
}

// Product defines the product content fields
type Product struct {
	Title       string
	SubTitle    string
	Text        string
	ImageLink   string
	ContentLink string
	ProductID   string
	Name        string
	Categories  []Category
	Brand       string
	VariantID   string
	UnitCost    string
	Discount    string
	Price       string
	Currency    string
	ProductLink string
}

// Order defines the order content fields
type Order struct {
	Title             string
	SubTitle          string
	Text              string
	ImageLink         string
	ContentLink       string
	OrderID           string
	OrderCreationDate string
	OrderState        string
	OrderShippedDate  string
	TotalOrdered      string
	TotalCost         string
	Discount          string
	OrderAmount       string
	Currency          string
	InvoiceNumber     string
	InvoiceDate       string
	OrderLink         string
	BasketLines       []Product
}

// Comment defines the comment content fields
type Comment struct {
	Title       string
	SubTitle    string
	Text        string
	ImageLink   string
	ContentLink string
	CommentID   string
	CommentDate string
	Author      string
}

// BlogPost defines the post content fields
type BlogPost struct {
	Title       string
	SubTitle    string
	Text        string
	ImageLink   string
	ContentLink string
	PostID      string
	PostDate    string
	Author      string
	Comments    []Comment
}
