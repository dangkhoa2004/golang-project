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

	// Cấu hình CORS chuẩn xác để Frontend gọi không bị lỗi
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(corsConfig))

	// ==========================================
	// TIÊM PHỤ THUỘC (DEPENDENCY INJECTION)
	// ==========================================

	// 1. Khởi tạo Repository
	userRepo := repository.NewUserRepository(db)
	courseRepo := repository.NewCourseRepository(db)
	learningRepo := repository.NewLearningRepository(db)
	orderRepo := repository.NewOrderRepository(db)

	// 2. Khởi tạo Service
	userSvc := service.NewUserService(userRepo)
	courseSvc := service.NewCourseService(courseRepo)
	learningSvc := service.NewLearningService(learningRepo, courseRepo)
	orderSvc := service.NewOrderService(orderRepo, courseRepo)

	// 3. Khởi tạo Handler
	userHdl := handler.NewUserHandler(userSvc)
	courseHdl := handler.NewCourseHandler(courseSvc)
	learningHdl := handler.NewLearningHandler(learningSvc)
	orderHdl := handler.NewOrderHandler(orderSvc)

	// ==========================================
	// KHAI BÁO ROUTES
	// ==========================================
	api := r.Group("/api")

	// NHÓM 1: CÔNG KHAI (PUBLIC API)
	public := api.Group("/")
	{
		// 1. Auth (Public)
		public.POST("/auth/register", userHdl.Register)
		public.POST("/auth/login", userHdl.Login)
		public.POST("/auth/refresh-token", userHdl.RefreshToken) // Thêm mới

		// 2. Course Discovery
		public.GET("/categories", courseHdl.GetCategories)
		public.GET("/courses", courseHdl.GetCourses)
		public.GET("/courses/:courseId", courseHdl.GetCourseDetail)
		public.GET("/courses/:courseId/reviews", courseHdl.GetCourseReviews)
	}

	// NHÓM 2: BẢO MẬT (PRIVATE API - REQUIRE AUTH)
	protected := api.Group("/")
	protected.Use(middleware.RequireAuth())
	{
		// 1. Auth & User Profile
		protected.POST("/auth/logout", userHdl.Logout)
		protected.GET("/users/profile", userHdl.GetProfile)
		protected.PUT("/users/profile", userHdl.UpdateProfile)
		protected.PUT("/users/change-password", userHdl.ChangePassword)

		// 3. Cart & Checkout (Order)
		protected.GET("/cart", orderHdl.GetCart)
		protected.POST("/cart/add", orderHdl.AddToCart)
		protected.DELETE("/cart/remove/:courseId", orderHdl.RemoveFromCart)
		protected.POST("/cart/apply-coupon", orderHdl.ApplyCoupon)
		protected.POST("/orders", orderHdl.CreateOrder)
		protected.POST("/payments/process", orderHdl.ProcessPayment)
		protected.GET("/orders/history", orderHdl.GetOrderHistory)

		// 4. Learning Experience (Không gian học tập)
		learning := protected.Group("/")
		{
			learning.GET("/enrollments/my-courses", learningHdl.GetMyCourses)
			learning.GET("/courses/:courseId/content", learningHdl.GetCourseContent)
			learning.GET("/lessons/:lessonId", learningHdl.GetLessonDetail)
			learning.POST("/progress/track", learningHdl.TrackProgress)
			learning.PUT("/progress/complete/:lessonId", learningHdl.CompleteLesson)
			learning.GET("/progress/course/:courseId", learningHdl.GetCourseProgress)
		}

		// 5. Engagement (Tương tác)
		protected.GET("/courses/:courseId/qna", learningHdl.GetQnA)
		protected.POST("/courses/:courseId/qna", learningHdl.PostQuestion)
		protected.POST("/qna/:questionId/replies", learningHdl.PostReply)
		protected.POST("/courses/:courseId/reviews", learningHdl.PostReview)
		protected.GET("/notifications", userHdl.GetNotifications)

		// 6. Assessments & Certificates
		protected.GET("/quizzes/:quizId", learningHdl.GetQuiz)
		protected.POST("/quizzes/:quizId/submit", learningHdl.SubmitQuiz)
		protected.GET("/quizzes/:quizId/results", learningHdl.GetQuizResult)
		protected.GET("/certificates/:courseId", learningHdl.CheckAndIssueCertificate)
		protected.GET("/certificates/download/:certificateId", learningHdl.DownloadCertificate)
	}

	return r
}
