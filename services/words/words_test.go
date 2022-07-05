package words

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/leomirandadev/improve-your-vocabulary/entities"
	"github.com/leomirandadev/improve-your-vocabulary/repositories"
	mocksWordsRepo "github.com/leomirandadev/improve-your-vocabulary/repositories/words/mocks"
	mocksCache "github.com/leomirandadev/improve-your-vocabulary/utils/cache/mocks"
	mocksLogger "github.com/leomirandadev/improve-your-vocabulary/utils/logger/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {

	// setup
	var userID uint64 = 123
	var wordID uint64 = 1

	cases := map[string]struct {
		input         entities.WordRequest
		expectedData  *entities.Word
		expectedError error
		prepareMocks  func(cache *mocksCache.MockCache, log *mocksLogger.MockLogger, wordsRepo *mocksWordsRepo.MockIRepository)
	}{
		"Create successfully": {
			input: entities.WordRequest{
				Word:   "deserve",
				UserID: userID,
			},
			expectedData: &entities.Word{
				ID:     wordID,
				Word:   "deserve",
				UserID: userID,
			},
			expectedError: nil,
			prepareMocks: func(cache *mocksCache.MockCache, log *mocksLogger.MockLogger, wordsRepo *mocksWordsRepo.MockIRepository) {
				wordsRepo.EXPECT().Create(gomock.Any(), entities.WordRequest{
					Word:   "deserve",
					UserID: userID,
				}).Times(1).Return(wordID, nil)

				wordsRepo.EXPECT().GetByID(gomock.Any(), wordID, userID).Times(1).Return(&entities.Word{
					ID:     wordID,
					Word:   "deserve",
					UserID: userID,
				}, nil)

				cache.EXPECT().Del(gomock.Any(), CACHE_GET_ALL_WORDS).Times(1).Return(true)
			},
		},
		"Error on insert": {
			input: entities.WordRequest{
				Word:   "deserve",
				UserID: userID,
			},
			expectedData:  nil,
			expectedError: errors.New("error on insert"),
			prepareMocks: func(cache *mocksCache.MockCache, log *mocksLogger.MockLogger, wordsRepo *mocksWordsRepo.MockIRepository) {

				err := errors.New("error on insert")

				wordsRepo.EXPECT().Create(gomock.Any(), entities.WordRequest{
					Word:   "deserve",
					UserID: userID,
				}).Times(1).Return(uint64(0), err)

				log.EXPECT().ErrorContext(gomock.Any(), "Word.Service.Create", err).Times(1)

			},
		},
		"Error on get by id": {
			input: entities.WordRequest{
				Word:   "deserve",
				UserID: userID,
			},
			expectedData:  nil,
			expectedError: errors.New("error get by id"),
			prepareMocks: func(cache *mocksCache.MockCache, log *mocksLogger.MockLogger, wordsRepo *mocksWordsRepo.MockIRepository) {

				err := errors.New("error get by id")

				wordsRepo.EXPECT().Create(gomock.Any(), entities.WordRequest{
					Word:   "deserve",
					UserID: userID,
				}).Times(1).Return(wordID, nil)

				wordsRepo.EXPECT().GetByID(gomock.Any(), wordID, userID).Times(1).
					Return(nil, err)

				log.EXPECT().ErrorContext(gomock.Any(), "Word.Service.GetByID", err).Times(1)

			},
		},
	}

	for name, caseTest := range cases {
		t.Run(name, func(t *testing.T) {
			ctr, ctx := gomock.WithContext(context.Background(), t)

			// mock data
			cache := mocksCache.NewMockCache(ctr)
			log := mocksLogger.NewMockLogger(ctr)
			wordsRepo := mocksWordsRepo.NewMockIRepository(ctr)

			// prepare mocks returns
			caseTest.prepareMocks(cache, log, wordsRepo)

			// initialize service
			words := New(&repositories.Container{Word: wordsRepo}, log, cache)

			// execute function
			data, err := words.Create(ctx, caseTest.input)

			// check if ok
			assert.Equal(t, data, caseTest.expectedData)
			assert.Equal(t, err, caseTest.expectedError)
		})
	}
}
