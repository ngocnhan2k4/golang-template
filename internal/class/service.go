package class

import (
	"Template/internal/entity"
	"Template/pkg/log"
	"context"
	"errors"
	"strconv"

	//"strconv"
	"time"

	"gorm.io/gorm"
)

type Service interface {
	Create(ctx context.Context, input CreateClassRequest) entity.Result
	Query(ctx context.Context, page, limit int, classId, semester, year *int) entity.Result
	// Get(ctx context.Context, id int) entity.Result
	Update(ctx context.Context, id string, input UpdateClassRequest) entity.Result
	Delete(ctx context.Context, id string) entity.Result
}

type Faculty struct {
	entity.Faculty
}

type CreateClassRequest struct {
	Id          string    `json:"classId"`
	Year        int       `json:"academicYear"`
	CourseId    string    `json:"courseId"`
	Semester    int       `json:"semester"`
	TeacherName string    `json:"teacherName"`
	MaxStudents int       `json:"maxStudents"`
	Room        string    `json:"room"`
	DayOfWeek   int       `json:"dayOfWeek"`
	StartTime   float64   `json:"startTime"`
	EndTime     float64   `json:"endTime"`
	Deadline    time.Time `json:"deadline"`
}

type UpdateClassRequest struct {
	Id          string    `json:"classId"`
	Year        int       `json:"academicYear"`
	CourseId    string    `json:"courseId"`
	Semester    int       `json:"semester"`
	TeacherName string    `json:"teacherName"`
	MaxStudents int       `json:"maxStudents"`
	Room        string    `json:"room"`
	DayOfWeek   int       `json:"dayOfWeek"`
	StartTime   float64   `json:"startTime"`
	EndTime     float64   `json:"endTime"`
	Deadline    time.Time `json:"deadline"`
}

type GetClassRequest struct {
	Id          string `json:"classId"`
	Year        int    `json:"academicYear"`
	CourseId    string `json:"courseId"`
	CourseName  string `json:"courseName"`
	Semester    int    `json:"semester"`
	TeacherName string `json:"teacherName"`
	MaxStudents int    `json:"maxStudents"`
	//current students
	Room      string    `json:"room"`
	DayOfWeek int       `json:"dayOfWeek"`
	StartTime float64   `json:"startTime"`
	EndTime   float64   `json:"endTime"`
	Deadline  time.Time `json:"deadline"`
}

type GetAllClassRequest struct {
	Data  []GetClassRequest `json:"data"`
	Total int               `json:"total"`
}

type service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

func (s service) Create(ctx context.Context, class CreateClassRequest) entity.Result {
	course, err := s.repo.GetCourse(ctx, class.CourseId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.Fail("COURSE_NOT_FOUND", "Khóa học không tồn tại.", nil)
	}
	err = s.repo.Create(ctx, entity.Class{
		AcademicYear: class.Year,
		CourseID:     course.ID,
		Semester:     class.Semester,
		TeacherName:  class.TeacherName,
		MaxStudents:  class.MaxStudents,
		Room:         class.Room,
		DayOfWeek:    class.DayOfWeek,
		StartTime:    class.StartTime,
		EndTime:      class.EndTime,
		DeadLine:     class.Deadline,
	})
	if err != nil {
		return entity.Fail("ADD_CLASS_FAILED", err.Error(), nil)
	}
	return entity.Ok(class, nil)
}

func (s service) Query(ctx context.Context, page, limit int, classId, semester, year *int) entity.Result {
	classes, err := s.repo.Query(ctx, page, limit, classId, semester, year)
	if err != nil {
		return entity.Fail("GET_CLASSES_FAILED", err.Error(), nil)
	}
	res := make([]GetClassRequest, len(classes))
	for i := range classes {
		res[i] = GetClassRequest{
			Id:          strconv.Itoa(classes[i].ID),
			Year:        classes[i].AcademicYear,
			CourseId:    strconv.Itoa(classes[i].CourseID),
			CourseName:  classes[i].Course.Name,
			Semester:    classes[i].Semester,
			TeacherName: classes[i].TeacherName,
			MaxStudents: classes[i].MaxStudents,
			Room:        classes[i].Room,
			DayOfWeek:   classes[i].DayOfWeek,
			StartTime:   classes[i].StartTime,
			EndTime:     classes[i].EndTime,
			Deadline:    classes[i].DeadLine,
		}
	}
	total := len(res)
	return entity.Ok(GetAllClassRequest{
		Data:  res,
		Total: total,
	}, nil)
}

// func (s service) Get(ctx context.Context, id int) entity.Result {
// 	courses, err := s.repo.Get(ctx, id)
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return entity.Fail("COURSE_NOT_FOUND", "Khóa học không tồn tại.", nil)
// 		}
// 		return entity.Fail("GET_COURSE_FAILED", err.Error(), nil)
// 	}
// 	var deletedAt *string
// 	if !courses.DeletedAt.IsZero() {
// 		deletedAt = new(string)
// 		*deletedAt = courses.DeletedAt.Format("2006-01-02 15:04:05")
// 	} else {
// 		deletedAt = nil
// 	}
// 	var requiredCourseId *string
// 	if courses.RequiredCourseId != nil {
// 		id := strconv.Itoa(*courses.RequiredCourseId)
// 		requiredCourseId = &id
// 	}
// 	res := GetCourseRequest{
// 		Id:               strconv.Itoa(courses.ID),
// 		CourseName:       entity.LocalizedName{Vi: courses.Name, En: courses.EngName},
// 		Credits:          courses.Credits,
// 		FacultyId:        strconv.Itoa(courses.FacultyId),
// 		FacultyName:      entity.LocalizedName{Vi: courses.Faculty.Name, En: courses.Faculty.EngName},
// 		Description:      entity.LocalizedName{Vi: courses.Description, En: courses.DescriptionEng},
// 		RequiredCourseId: requiredCourseId,
// 		CreatedAt:        courses.CreatedAt.Format("2006-01-02 15:04:05"),
// 		DeletedAt:        deletedAt,
// 	}
// 	return entity.Ok(res, nil)
// }

func (s service) Update(ctx context.Context, id string, classReq UpdateClassRequest) entity.Result {
	intID, _ := strconv.Atoi(id)
	class, err := s.repo.Get(ctx, intID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.Fail("CLASS_NOT_FOUND", "Lớp học không tồn tại.", nil)
	}
	courseId, _ := strconv.Atoi(classReq.CourseId)

	err = s.repo.Update(ctx, entity.Class{
		ID:           intID,
		AcademicYear: classReq.Year,
		CourseID:     courseId,
		Semester:     classReq.Semester,
		TeacherName:  classReq.TeacherName,
		MaxStudents:  classReq.MaxStudents,
		Room:         classReq.Room,
		DayOfWeek:    classReq.DayOfWeek,
		StartTime:    classReq.StartTime,
		EndTime:      classReq.EndTime,
		DeadLine:     classReq.Deadline,
		IsDeleted:    class.IsDeleted,
	})
	if err != nil {
		return entity.Fail("UPDATE_CLASS_FAILED", err.Error(), nil)
	}
	return entity.Ok(classReq, nil)
}

func (s service) Delete(ctx context.Context, id string) entity.Result {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Fail("COURSE_NOT_FOUND", "Lớp học không tồn tại.", nil)
		}
		return entity.Fail("DELETE_COURSE_FAILED", err.Error(), nil)
	}
	getClassRequest := GetClassRequest{
		Id: id,
	}
	return entity.Ok(getClassRequest, nil)
}
