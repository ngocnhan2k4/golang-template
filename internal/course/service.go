package course

import (
	"Template/internal/entity"
	"Template/pkg/log"
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type Service interface {
	Create(ctx context.Context, input CreateCourseRequest) entity.Result
	Query(ctx context.Context, page, limit int, facultyId, courseId *int, isDeleted *bool) entity.Result
	Get(ctx context.Context, id int) entity.Result
	Update(ctx context.Context, id string, input UpdateCourseRequest) entity.Result
	Delete(ctx context.Context, id string) entity.Result
}

type Faculty struct {
	entity.Faculty
}

type CreateCourseRequest struct {
	Id                   string `json:"courseId"`
	entity.LocalizedName `json:"courseName"`
	Credits              int    `json:"credits"`
	FacultyId            string `json:"facultyId"`
	Description          struct {
		Vi string `json:"vi"`
		En string `json:"en"`
	} `json:"description"`
	RequiredCourseId *string `json:"requiredCourseId"`
	CreatedAt        string  `json:"createdAt"`
}

type UpdateCourseRequest struct {
	CourseName  entity.LocalizedName `json:"courseName"`
	Credits     int                  `json:"credits"`
	FacultyId   string               `json:"facultyId"`
	Description entity.LocalizedName `json:"description"`
}

type GetCourseRequest struct {
	Id               string               `json:"courseId"`
	CourseName       entity.LocalizedName `json:"courseName"`
	Credits          int                  `json:"credits"`
	FacultyId        string               `json:"facultyId"`
	FacultyName      entity.LocalizedName `json:"facultyName"`
	Description      entity.LocalizedName `json:"description"`
	RequiredCourseId *string              `json:"requiredCourseId"`
	DeletedAt        *string              `json:"deletedAt"`
	CreatedAt        string               `json:"createdAt"`
}

type GetAllCourseRequest struct {
	Data  []GetCourseRequest `json:"data"`
	Total int                `json:"total"`
}

type service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

func (s service) Create(ctx context.Context, course CreateCourseRequest) entity.Result {
	facultyId, _ := strconv.Atoi(course.FacultyId)
	fmt.Println("RequiredCourseId:", course.RequiredCourseId)
	var requiredCourseId *int
	if course.RequiredCourseId == nil {
		requiredCourseId = nil
	} else {
		id, _ := strconv.Atoi(*course.RequiredCourseId)
		requiredCourseId = &id
		_, err := s.repo.Get(ctx, *requiredCourseId)
		if errors.Is(err, gorm.ErrRecordNotFound) && course.RequiredCourseId != nil {
			return entity.Fail("PRE_COURSE_NOT_FOUND", "Khóa học tiên quyết không tồn tại.", nil)
		}
	}

	err := s.repo.Create(ctx, entity.Course{
		EngName:          course.LocalizedName.En,
		Name:             course.LocalizedName.Vi,
		Credits:          course.Credits,
		FacultyId:        facultyId,
		Description:      course.Description.Vi,
		DescriptionEng:   course.Description.En,
		RequiredCourseId: requiredCourseId,
	})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.ConstraintName {
			case "courses_name_eng_key":
				return entity.Fail("DUPLICATE_COURSE_NAME", "Tên khóa học 'EN' đã tồn tại.", nil)
			case "courses_name_key":
				return entity.Fail("DUPLICATE_COURSE_NAME", "Tên khóa học 'VI' đã tồn tại.", nil)
			case "courses_pkey":
				return entity.Fail("ADD_COURSE_FAILED", "Mã khóa học đã tồn tại.", nil)
			case "courses_faculty_id_fkey":
				return entity.Fail("FACULTY_NOT_EXIST", "Khoa không tồn tại", nil)
			}
		}
		return entity.Fail("ADD_COURSE_FAILED", err.Error(), nil)
	}
	return entity.Ok(course, nil)
}

func (s service) Query(ctx context.Context, page, limit int, facultyId, courseId *int, isDeleted *bool) entity.Result {
	courses, err := s.repo.Query(ctx, page, limit, facultyId, courseId, isDeleted)
	if err != nil {
		return entity.Fail("GET_COURSES_FAILED", err.Error(), nil)
	}
	res := make([]GetCourseRequest, len(courses))
	for i := range courses {
		var deletedAt *string
		if !courses[i].DeletedAt.IsZero() {
			deletedAt = new(string)
			*deletedAt = courses[i].DeletedAt.Format("2006-01-02 15:04:05")
		} else {
			deletedAt = nil
		}
		var requiredCourseId *string
		if courses[i].RequiredCourseId != nil {
			id := strconv.Itoa(*courses[i].RequiredCourseId)
			requiredCourseId = &id
		}
		res[i] = GetCourseRequest{
			Id:               strconv.Itoa(courses[i].ID),
			CourseName:       entity.LocalizedName{Vi: courses[i].Name, En: courses[i].EngName},
			Credits:          courses[i].Credits,
			FacultyId:        strconv.Itoa(courses[i].FacultyId),
			FacultyName:      entity.LocalizedName{Vi: courses[i].Faculty.Name, En: courses[i].Faculty.EngName},
			Description:      entity.LocalizedName{Vi: courses[i].Description, En: courses[i].DescriptionEng},
			RequiredCourseId: requiredCourseId,
			CreatedAt:        courses[i].CreatedAt.Format("2006-01-02 15:04:05"),
			DeletedAt:        deletedAt,
		}
	}
	total := len(res)
	return entity.Ok(GetAllCourseRequest{
		Data:  res,
		Total: total,
	}, nil)
}

func (s service) Get(ctx context.Context, id int) entity.Result {
	courses, err := s.repo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Fail("COURSE_NOT_FOUND", "Khóa học không tồn tại.", nil)
		}
		return entity.Fail("GET_COURSE_FAILED", err.Error(), nil)
	}
	var deletedAt *string
	if !courses.DeletedAt.IsZero() {
		deletedAt = new(string)
		*deletedAt = courses.DeletedAt.Format("2006-01-02 15:04:05")
	} else {
		deletedAt = nil
	}
	var requiredCourseId *string
	if courses.RequiredCourseId != nil {
		id := strconv.Itoa(*courses.RequiredCourseId)
		requiredCourseId = &id
	}
	res := GetCourseRequest{
		Id:               strconv.Itoa(courses.ID),
		CourseName:       entity.LocalizedName{Vi: courses.Name, En: courses.EngName},
		Credits:          courses.Credits,
		FacultyId:        strconv.Itoa(courses.FacultyId),
		FacultyName:      entity.LocalizedName{Vi: courses.Faculty.Name, En: courses.Faculty.EngName},
		Description:      entity.LocalizedName{Vi: courses.Description, En: courses.DescriptionEng},
		RequiredCourseId: requiredCourseId,
		CreatedAt:        courses.CreatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:        deletedAt,
	}
	return entity.Ok(res, nil)
}

func (s service) Update(ctx context.Context, id string, courseReq UpdateCourseRequest) entity.Result {
	intID, _ := strconv.Atoi(id)
	facultyId, _ := strconv.Atoi(courseReq.FacultyId)
	course, err := s.repo.Get(ctx, intID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.Fail("ADD_COURSE_FAILED", "Khóa học không tồn tại.", nil)
	}
	if _, err := s.repo.GetFaculty(ctx, courseReq.FacultyId); errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.Fail("ADD_COURSE_FAILED", "Khoa không tồn tại.", nil)
	}
	course.EngName = courseReq.CourseName.En
	course.Name = courseReq.CourseName.Vi
	course.Credits = courseReq.Credits
	course.FacultyId = facultyId
	course.Description = courseReq.Description.Vi
	course.DescriptionEng = courseReq.Description.En

	err = s.repo.Update(ctx, course)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.ConstraintName {
			case "courses_name_eng_key":
				return entity.Fail("DUPLICATE_COURSE_NAME", "Tên khóa 'EN' đã tồn tại.", nil)
			case "courses_name_key":
				return entity.Fail("DUPLICATE_COURSE_NAME", "Tên khóa 'VI' đã tồn tại.", nil)
			}
		}
		return entity.Fail("UPDATE_COURSE_FAILED", err.Error(), nil)
	}
	return entity.Ok(courseReq, nil)
}

func (s service) Delete(ctx context.Context, id string) entity.Result {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Fail("COURSE_NOT_FOUND", "Khóa học không tồn tại.", nil)
		}
		return entity.Fail("DELETE_COURSE_FAILED", "Course already has class. Status will be changed to Deactivated", nil)
	}
	return entity.Ok(id, nil)
}
