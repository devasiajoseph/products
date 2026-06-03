package item
import "time"
type Item struct {
	ItemID               int       `db:"itemid" json:"itemid"`
	Title                string    `db:"title" json:"title"`
	Description          string    `db:"description" json:"description"`
	Category             string    `db:"category" json:"category"`
	Price                int       `db:"price" json:"price"`
	DiscountPercentage   float64   `db:"discount_percentage" json:"discount_percentage"`
	Rating               float64   `db:"rating" json:"rating"`
	Stock                int       `db:"stock" json:"stock"`
	Brand                string    `db:"brand" json:"brand"`
	SKU                  string    `db:"sku" json:"sku"`
	Weight               int       `db:"weight" json:"weight"`
	Width                float64   `db:"width" json:"width"`
	Height               float64   `db:"height" json:"height"`
	Depth                float64   `db:"depth" json:"depth"`
	WarrantyInfo         string    `db:"warranty_info" json:"warranty_info"`
	ShippingInfo         string    `db:"shipping_info" json:"shipping_info"`
	AvailabilityStatus   string    `db:"availability_status" json:"availability_status"`
	ReturnPolicy         string    `db:"return_policy" json:"return_policy"`
	MinimumOrderQuantity int       `db:"minimum_order_quantity" json:"minimum_order_quantity"`
	CreatedAt            time.Time `db:"created_at" json:"created_at"`
	UpdatedAt            time.Time `db:"updated_at" json:"updated_at"`
	Barcode              string    `db:"barcode" json:"barcode"`
	QRCode               string    `db:"qr_code" json:"qr_code"`
	ThumbnailURL         string    `db:"thumbnail_url" json:"thumbnail_url"`
}
type Tag struct {
	TagID   int    `db:"tagid" json:"tagid"`
	TagName string `db:"tagname" json:"tagname"`
}
type ItemTag struct {
	ItemTags int `db:"item_tags" json:"item_tags"`
	ItemID   int `db:"item_id" json:"item_id"`
	TagID    int `db:"tag_id" json:"tag_id"`
}
type Review struct {
	ReviewID      int       `db:"review_id" json:"review_id"`
	Rating        int       `db:"rating" json:"rating"`
	Comment       string    `db:"comment" json:"comment"`
	Date          time.Time `db:"date" json:"date"`
	ReviewerName  string    `db:"reviewer_name" json:"reviewer_name"`
	ReviewerEmail string    `db:"reviewer_email" json:"reviewer_email"`
	ItemID        int       `db:"item_id" json:"item_id"`
}
type ItemImage struct {
	ImageID  int    `db:"image_id" json:"image_id"`
	ItemID   int    `db:"item_id" json:"item_id"`
	ImageURL string `db:"image_url" json:"image_url"`
}