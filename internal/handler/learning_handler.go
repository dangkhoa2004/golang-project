package handler

import (
	"golang-project/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LearningHandler struct {
	learningSvc service.LearningService
}

func NewLearningHandler(svc service.LearningService) *LearningHandler {
	return &LearningHandler{learningSvc: svc}
}

// GET /api/enrollments/my-courses
func (h *LearningHandler) GetMyCourses(c *gin.Context) {
	studentID, _ := GetUserIDFromContext(c)
	enrollments, err := h.learningSvc.GetMyCourses(studentID)
	if err != nil {
		SendError(c, http.StatusInternalServerError, "Lỗi lấy danh sách khóa học")
		return
	}
	SendSuccess(c, http.StatusOK, "Thành công", enrollments)
}

// GET /api/lessons/:lessonId
func (h *LearningHandler) GetLessonDetail(c *gin.Context) {
	studentID, _ := GetUserIDFromContext(c)
	lessonID, err := GetParamInt(c, "lessonId")
	if err != nil {
		SendError(c, http.StatusBadRequest, "ID bài học không hợp lệ")
		return
	}

	// Chú ý: Ở đây bạn cần truyền courseID vào từ URL hoặc tìm từ LessonID.
	// Hiện tại truyền tạm courseID = 1 để biên dịch thành công.
	lesson, err := h.learningSvc.GetLessonSecureContent(studentID, 1, lessonID)
	if err != nil {
		SendError(c, http.StatusForbidden, err.Error())
		return
	}
	SendSuccess(c, http.StatusOK, "Tải nội dung bài học thành công", lesson)
}

// POST /api/progress/track
func (h *LearningHandler) TrackProgress(c *gin.Context) {
	studentID, _ := GetUserIDFromContext(c)
	var req struct {
		CourseID       uint `json:"course_id"`
		LessonID       uint `json:"lesson_id"`
		WatchedSeconds int  `json:"watched_seconds"`
	}

	if err := BindJSON(c, &req); err != nil {
		SendError(c, http.StatusBadRequest, "Dữ liệu không hợp lệ")
		return
	}

	err := h.learningSvc.TrackVideoProgress(studentID, req.CourseID, req.LessonID, req.WatchedSeconds)
	if err != nil {
		SendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	SendSuccess(c, http.StatusOK, "Đã lưu vị trí video", nil)
}

// PUT /api/progress/complete/:lessonId
func (h *LearningHandler) CompleteLesson(c *gin.Context) {
	studentID, _ := GetUserIDFromContext(c)
	lessonID, _ := GetParamInt(c, "lessonId")

	// Giả định CourseID = 1, bạn cần truyền từ Frontend lên
	err := h.learningSvc.MarkLessonComplete(studentID, 1, lessonID)
	if err != nil {
		SendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	SendSuccess(c, http.StatusOK, "Đã hoàn thành bài học", nil)
}

// ==========================================
// CÁC HÀM PLACEHOLDER (Để Router không báo lỗi)
// ==========================================
func (h *LearningHandler) GetCourseContent(c *gin.Context) {
	SendSuccess(c, http.StatusOK, "Nội dung lộ trình", nil)
}
func (h *LearningHandler) GetCourseProgress(c *gin.Context) {
	SendSuccess(c, http.StatusOK, "Tiến độ: 45%", nil)
}
func (h *LearningHandler) GetQnA(c *gin.Context) {
	SendSuccess(c, http.StatusOK, "Danh sách Q&A", nil)
}
func (h *LearningHandler) PostQuestion(c *gin.Context) {
	SendSuccess(c, http.StatusOK, "Đăng câu hỏi", nil)
}
func (h *LearningHandler) PostReply(c *gin.Context) {
	SendSuccess(c, http.StatusOK, "Trả lời", nil)
}
func (h *LearningHandler) PostReview(c *gin.Context) {
	SendSuccess(c, http.StatusOK, "Đánh giá", nil)
}
func (h *LearningHandler) GetQuiz(c *gin.Context) {
	SendSuccess(c, http.StatusOK, "Câu hỏi Quiz", nil)
}
func (h *LearningHandler) SubmitQuiz(c *gin.Context) {
	SendSuccess(c, http.StatusOK, "Nộp bài", nil)
}
func (h *LearningHandler) GetQuizResult(c *gin.Context) {
	SendSuccess(c, http.StatusOK, "Kết quả Quiz", nil)
}
func (h *LearningHandler) CheckAndIssueCertificate(c *gin.Context) {
	SendSuccess(c, http.StatusOK, "Đủ điều kiện", nil)
}
func (h *LearningHandler) DownloadCertificate(c *gin.Context) {
	SendSuccess(c, http.StatusOK, "Link tải", nil)
}
