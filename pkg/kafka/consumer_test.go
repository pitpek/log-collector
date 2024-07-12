package kafka_test

import (
	"context"
	"errors"
	consumer "logcollector/pkg/kafka"
	"testing"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockStorage представляет собой мок-версию интерфейса Storage
type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) InsertMessage(date time.Time, key, message string) error {
	args := m.Called(date, key, message)
	return args.Error(0)
}

// MockReader представляет собой мок-версию интерфейса Reader
type MockReader struct {
	mock.Mock
}

func (m *MockReader) ReadMessage(ctx context.Context) (kafka.Message, error) {
	args := m.Called(ctx)
	return args.Get(0).(kafka.Message), args.Error(1)
}

func (m *MockReader) Close() error {
	args := m.Called()
	return args.Error(0)
}

func TestConsumer_Start(t *testing.T) {
	mockStorage := new(MockStorage)
	mockReader := new(MockReader)
	consumer := consumer.NewConsumer(mockReader, mockStorage)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Мокируем поведение
	message := kafka.Message{
		Key:   []byte("test-key"),
		Value: []byte("test-message"),
	}
	mockReader.On("ReadMessage", ctx).Return(message, nil).Once()                              // Первый успешный вызов ReadMessage
	mockReader.On("ReadMessage", ctx).Return(kafka.Message{}, errors.New("read error")).Once() // Второй вызов с ошибкой
	mockReader.On("Close").Return(nil)                                                         // Мокируем Close

	// Мокируем InsertMessage только для первого успешного вызова ReadMessage
	mockStorage.On("InsertMessage", mock.Anything, "test-key", "test-message").Return(nil).Once()

	// Запускаем consumer в отдельной горутине
	go func() {
		err := consumer.Start(ctx)
		assert.Error(t, err, "expected error from consumer.Start")
	}()

	// Даем немного времени на выполнение
	time.Sleep(100 * time.Millisecond)
	cancel()

	// Проверяем, что мок-методы были вызваны ожидаемое количество раз
	mockReader.AssertExpectations(t)
	mockStorage.AssertExpectations(t)
}
