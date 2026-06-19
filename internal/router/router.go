package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"golang-project/internal/handler"
	"golang-project/internal/middleware"
	"golang-project/internal/repository"
	"golang-project/internal/service"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// 1. Cấu hình CORS chuẩn xác (Sửa lỗi Blocked by CORS policy)
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(corsConfig))

	// 2. Tiêm phụ thuộc (Dependency Injection)
	// --- Repo ---
	userRepo := repository.NewUserRepository(db)
	courseRepo := repository.NewCourseRepository(db)

	// --- Service ---
	userSvc := service.NewUserService(userRepo)
	courseSvc := service.NewCourseService(courseRepo)

	// --- Handler ---
	userHdl := handler.NewUserHandler(userSvc)
	courseHdl := handler.NewCourseHandler(courseSvc)

	// 3. Khai báo Routes
	api := r.Group("/api")

	// ---> 3.1. Nhóm API CÔNG KHAI
	public := api.Group("/")
	{
		public.POST("/register", userHdl.Register)
		public.POST("/login", userHdl.Login)
		public.GET("/courses", courseHdl.GetHomeCourses)
	}

	// ---> 3.2. Nhóm API BẢO MẬT
	protected := api.Group("/")
	protected.Use(middleware.RequireAuth())
	{
		// ĐÃ THÊM ROUTE LẤY PROFILE TẠI ĐÂY
		protected.GET("/user/profile", userHdl.GetProfile)

		// API Demo Checkout
		protected.POST("/checkout", func(c *gin.Context) {
			userID, _ := c.Get("user_id")
			c.JSON(200, gin.H{
				"message":    "Thanh toán thành công",
				"student_id": userID,
			})
		})
	}

	return r
}
