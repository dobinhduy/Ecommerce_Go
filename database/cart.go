package database



var(
	ErrCantFindProduct = errors.New("can't find the product")
	ErrCantDecodeProducts = errors.New("cant find the product")
	ErrUserIdNorValid = errors.New("this user is not valid")
	ErrCantUpdateUser = errors.New("cannot add this product to cart")
	ErrCantRemoveItemCart = errors.New("cannot remove this item from cart")
	ErrCantGetItem = errors.New("was unable to get the item form the cart")
	ErrCantBuyCartItem  = errors.New("cannot update the purchase")
)	

func AddProductToCard(){}
func RemoveCartItem(){}
func BuyItemFromCart(){}
func InstantBuyer(){}