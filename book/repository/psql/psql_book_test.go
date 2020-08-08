package psql

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"regexp"
	"strconv"

	"github.com/stretchr/testify/assert"
	"go-clean-architecture/domain"
	"testing"
)

func TestFetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	gormMock, err := gorm.Open("postgres", db)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockArticles := []domain.Book{
		domain.Book{
			ID: 1, Title: "title 1", Author: "author 1",
		},
		domain.Book{
			ID: 2, Title: "title 2", Author: "author 2",
		},
		domain.Book{
			ID: 3, Title: "title 2", Author: "author 2",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "author"}).
		AddRow(mockArticles[0].ID, mockArticles[0].Title, mockArticles[0].Author).
		AddRow(mockArticles[1].ID, mockArticles[1].Title, mockArticles[1].Author).
		AddRow(mockArticles[2].ID, mockArticles[2].Title, mockArticles[2].Author)

	query := regexp.QuoteMeta(`SELECT * FROM "books"`)

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := psqlBookRepository{gormMock}
	list, err := a.Fetch(context.TODO())
	assert.NoError(t, err)
	//assert.Contains(t, "2", strconv.FormatUint(uint64(mockArticles[0].ID), 10),"testing compare")
	assert.Equal(t, "1", strconv.FormatUint(uint64(mockArticles[0].ID), 10))
	assert.Len(t, list, 3)
}
