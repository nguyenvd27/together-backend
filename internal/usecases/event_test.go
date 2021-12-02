package usecases

import (
	"fmt"
	"reflect"
	"testing"
	"time"
	"together-backend/internal/models"
	"together-backend/internal/repositories"
	mock_repositories "together-backend/internal/repositories/mock"

	"github.com/golang/mock/gomock"
)

func Test_eventUsecase_CreateEventUsecase(t *testing.T) {
	t.Run("error title is empty", func(t *testing.T) {
		// prepare input
		imageUrl := []string{"image1", "image2"}
		startTime := time.Now()
		endTime := startTime.AddDate(0, 0, 1)
		reqBody := &ReqBodyEvent{
			Title:          "",
			Content:        "content",
			CreatedBy:      uint64(1),
			StartTime:      startTime,
			EndTime:        endTime,
			Location:       1,
			DetailLocation: "detail location",
		}

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		eventRepo := mock_repositories.NewMockEventRepo(mockCtrl)

		// output
		var expectedResult *models.Event = nil
		var expectedErr error = fmt.Errorf("title cannot be empty")

		uc := &eventUsecase{
			eventRepo:     eventRepo,
			userRepo:      nil,
			imageRepo:     nil,
			userEventRepo: nil,
		}
		got, err := uc.CreateEventUsecase(reqBody, imageUrl)

		if err.Error() != expectedErr.Error() {
			t.Errorf("eventUsecase.CreateEventUsecase() error = %v, wantErr %v", err, expectedErr)
			return
		}
		if !reflect.DeepEqual(got, expectedResult) {
			t.Errorf("eventUsecase.CreateEventUsecase() = %v, want %v", got, expectedResult)
		}
	})

	t.Run("error content is empty", func(t *testing.T) {
		// prepare input
		imageUrl := []string{"image1", "image2"}
		startTime := time.Now()
		endTime := startTime.AddDate(0, 0, 1)
		reqBody := &ReqBodyEvent{
			Title:          "title",
			Content:        "",
			CreatedBy:      uint64(1),
			StartTime:      startTime,
			EndTime:        endTime,
			Location:       1,
			DetailLocation: "detail location",
		}

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		eventRepo := mock_repositories.NewMockEventRepo(mockCtrl)

		// output
		var expectedResult *models.Event = nil
		var expectedErr error = fmt.Errorf("content cannot be empty")

		uc := &eventUsecase{
			eventRepo:     eventRepo,
			userRepo:      nil,
			imageRepo:     nil,
			userEventRepo: nil,
		}
		got, err := uc.CreateEventUsecase(reqBody, imageUrl)

		if err.Error() != expectedErr.Error() {
			t.Errorf("eventUsecase.CreateEventUsecase() error = %v, wantErr %v", err, expectedErr)
			return
		}
		if !reflect.DeepEqual(got, expectedResult) {
			t.Errorf("eventUsecase.CreateEventUsecase() = %v, want %v", got, expectedResult)
		}
	})

	t.Run("error created_by is empty", func(t *testing.T) {
		// prepare input
		imageUrl := []string{"image1", "image2"}
		startTime := time.Now()
		endTime := startTime.AddDate(0, 0, 1)
		reqBody := &ReqBodyEvent{
			Title:          "title",
			Content:        "content",
			CreatedBy:      uint64(0),
			StartTime:      startTime,
			EndTime:        endTime,
			Location:       1,
			DetailLocation: "detail location",
		}

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		eventRepo := mock_repositories.NewMockEventRepo(mockCtrl)

		// output
		var expectedResult *models.Event = nil
		var expectedErr error = fmt.Errorf("created_by cannot be empty")

		uc := &eventUsecase{
			eventRepo:     eventRepo,
			userRepo:      nil,
			imageRepo:     nil,
			userEventRepo: nil,
		}
		got, err := uc.CreateEventUsecase(reqBody, imageUrl)

		if err.Error() != expectedErr.Error() {
			t.Errorf("eventUsecase.CreateEventUsecase() error = %v, wantErr %v", err, expectedErr)
			return
		}
		if !reflect.DeepEqual(got, expectedResult) {
			t.Errorf("eventUsecase.CreateEventUsecase() = %v, want %v", got, expectedResult)
		}
	})

	t.Run("error start time is greater than end time", func(t *testing.T) {
		// prepare input
		imageUrl := []string{"image1", "image2"}
		startTime := time.Now()
		endTime := startTime.AddDate(0, 0, 1)
		reqBody := &ReqBodyEvent{
			Title:          "title",
			Content:        "content",
			CreatedBy:      uint64(1),
			StartTime:      endTime,
			EndTime:        startTime,
			Location:       1,
			DetailLocation: "detail location",
		}

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		eventRepo := mock_repositories.NewMockEventRepo(mockCtrl)

		// output
		var expectedResult *models.Event = nil
		var expectedErr error = fmt.Errorf("start time must be less than end time")

		uc := &eventUsecase{
			eventRepo:     eventRepo,
			userRepo:      nil,
			imageRepo:     nil,
			userEventRepo: nil,
		}
		got, err := uc.CreateEventUsecase(reqBody, imageUrl)

		if err.Error() != expectedErr.Error() {
			t.Errorf("eventUsecase.CreateEventUsecase() error = %v, wantErr %v", err, expectedErr)
			return
		}
		if !reflect.DeepEqual(got, expectedResult) {
			t.Errorf("eventUsecase.CreateEventUsecase() = %v, want %v", got, expectedResult)
		}
	})

	t.Run("error end time is less than current time", func(t *testing.T) {
		// prepare input
		imageUrl := []string{"image1", "image2"}
		timeNow := time.Now()
		startTime := timeNow.AddDate(0, 0, -2)
		endTime := timeNow.AddDate(0, 0, -1)
		reqBody := &ReqBodyEvent{
			Title:          "title",
			Content:        "content",
			CreatedBy:      uint64(1),
			StartTime:      startTime,
			EndTime:        endTime,
			Location:       1,
			DetailLocation: "detail location",
		}

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		eventRepo := mock_repositories.NewMockEventRepo(mockCtrl)

		// output
		var expectedResult *models.Event = nil
		var expectedErr error = fmt.Errorf("end time must be greater than current time")

		uc := &eventUsecase{
			eventRepo:     eventRepo,
			userRepo:      nil,
			imageRepo:     nil,
			userEventRepo: nil,
		}
		got, err := uc.CreateEventUsecase(reqBody, imageUrl)

		if err.Error() != expectedErr.Error() {
			t.Errorf("eventUsecase.CreateEventUsecase() error = %v, wantErr %v", err, expectedErr)
			return
		}
		if !reflect.DeepEqual(got, expectedResult) {
			t.Errorf("eventUsecase.CreateEventUsecase() = %v, want %v", got, expectedResult)
		}
	})

	t.Run("error when create event", func(t *testing.T) {
		// prepare input
		imageUrl := []string{"image1", "image2"}
		startTime := time.Now()
		endTime := startTime.AddDate(0, 0, 1)
		reqBody := &ReqBodyEvent{
			Title:          "title",
			Content:        "content",
			CreatedBy:      uint64(1),
			StartTime:      startTime,
			EndTime:        endTime,
			Location:       1,
			DetailLocation: "detail location",
		}

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		eventRepo := mock_repositories.NewMockEventRepo(mockCtrl)
		eventRepo.EXPECT().CreateEvent(reqBody.Title, reqBody.Content, imageUrl, reqBody.CreatedBy, reqBody.StartTime, reqBody.EndTime, reqBody.Location, reqBody.DetailLocation).Return(
			nil, fmt.Errorf("error"),
		)

		// output
		var expectedResult *models.Event = nil
		var expectedErr error = fmt.Errorf("error")

		uc := &eventUsecase{
			eventRepo:     eventRepo,
			userRepo:      nil,
			imageRepo:     nil,
			userEventRepo: nil,
		}
		got, err := uc.CreateEventUsecase(reqBody, imageUrl)

		if err.Error() != expectedErr.Error() {
			t.Errorf("eventUsecase.CreateEventUsecase() error = %v, wantErr %v", err, expectedErr)
			return
		}
		if !reflect.DeepEqual(got, expectedResult) {
			t.Errorf("eventUsecase.CreateEventUsecase() = %v, want %v", got, expectedResult)
		}
	})

	t.Run("pass", func(t *testing.T) {
		// prepare input
		imageUrl := []string{"image1", "image2"}
		startTime := time.Now()
		endTime := startTime.AddDate(0, 0, 1)
		reqBody := &ReqBodyEvent{
			Title:          "title",
			Content:        "content",
			CreatedBy:      uint64(1),
			StartTime:      startTime,
			EndTime:        endTime,
			Location:       1,
			DetailLocation: "detail location",
		}

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		eventRepo := mock_repositories.NewMockEventRepo(mockCtrl)
		eventRepo.EXPECT().CreateEvent(reqBody.Title, reqBody.Content, imageUrl, reqBody.CreatedBy, reqBody.StartTime, reqBody.EndTime, reqBody.Location, reqBody.DetailLocation).Return(
			&models.Event{
				Id:             uint(1),
				Title:          reqBody.Title,
				Content:        reqBody.Content,
				CreatedBy:      reqBody.CreatedBy,
				StartTime:      reqBody.StartTime,
				EndTime:        reqBody.EndTime,
				Location:       reqBody.Location,
				DetailLocation: reqBody.DetailLocation,
				CreatedAt:      startTime,
				EventImages:    []models.EventImage{},
				Comments:       nil,
				Users:          nil,
			}, nil,
		)

		// output
		expectedResult := &models.Event{
			Id:             uint(1),
			Title:          reqBody.Title,
			Content:        reqBody.Content,
			CreatedBy:      reqBody.CreatedBy,
			StartTime:      reqBody.StartTime,
			EndTime:        reqBody.EndTime,
			Location:       reqBody.Location,
			DetailLocation: reqBody.DetailLocation,
			CreatedAt:      startTime,
			EventImages:    []models.EventImage{},
			Comments:       nil,
			Users:          nil,
		}
		var expectedErr error = nil

		uc := &eventUsecase{
			eventRepo:     eventRepo,
			userRepo:      nil,
			imageRepo:     nil,
			userEventRepo: nil,
		}
		got, err := uc.CreateEventUsecase(reqBody, imageUrl)

		if (err != nil) || (expectedErr != nil) {
			t.Errorf("eventUsecase.CreateEventUsecase() error = %v, wantErr %v", err, expectedErr)
			return
		}
		if !reflect.DeepEqual(got, expectedResult) {
			t.Errorf("eventUsecase.CreateEventUsecase() = %v, want %v", got, expectedResult)
		}
	})
}

func TestNewEventUsecase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	type args struct {
		eventRepo     repositories.EventRepo
		userRepo      repositories.UserRepo
		imageRepo     repositories.ImageRepo
		userEventRepo repositories.UserEventRepo
	}
	tests := []struct {
		name string
		args args
		want EventUseCase
	}{
		{
			name: "case 1",
			args: args{
				eventRepo:     mock_repositories.NewMockEventRepo(mockCtrl),
				userRepo:      mock_repositories.NewMockUserRepo(mockCtrl),
				imageRepo:     mock_repositories.NewMockImageRepo(mockCtrl),
				userEventRepo: mock_repositories.NewMockUserEventRepo(mockCtrl),
			},
			want: &eventUsecase{
				eventRepo:     mock_repositories.NewMockEventRepo(mockCtrl),
				userRepo:      mock_repositories.NewMockUserRepo(mockCtrl),
				imageRepo:     mock_repositories.NewMockImageRepo(mockCtrl),
				userEventRepo: mock_repositories.NewMockUserEventRepo(mockCtrl),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEventUsecase(tt.args.eventRepo, tt.args.userRepo, tt.args.imageRepo, tt.args.userEventRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEventUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
