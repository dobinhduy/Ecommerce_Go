package routes
import(
	"/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/users/signup",controllers.SignUp())
	incomingRoutes.POST("/users/login",controllers.Login())
	incomingRoutes.POST("/users/productview",controllers.SearchProduct())
	incomingRoutes.POST("/users/search",controllers.SearchProductByQuery())
	incomingRoutes.POST("/admin/addproduct",controllers.ProductViewerAdmin())

}
