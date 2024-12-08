package routes

// mock controllers
// func mockCreateBook(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusCreated)
// 	w.Write([]byte("Book Created"))
// }

// func mockUpdateBook(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Book Updated"))
// }

// func mockDeleteBook(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusNoContent)
// }
// func mockGetBookById(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Book fetched"))
// }
// func mockGetBooks(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Books fetched"))
// }

// func TestRegisterBookstoreRoutes(t *testing.T) {
// 	mockHandlers := &controllers.BookstoreHandler{
// 		CreateBook:  mockCreateBook,
// 		UpdateBook:  mockUpdateBook,
// 		DeleteBook:  mockDeleteBook,
// 		GetBooks:    mockGetBooks,
// 		GetBookById: mockGetBookById,
// 	}

// 	r := mux.NewRouter()
// 	RegisterBookstoreRoutes(r, mockHandlers)

// 	tests := []struct {
// 		name           string
// 		method         string
// 		url            string
// 		expectedStatus int
// 		expectedBody   string
// 	}{
// 		{
// 			name:           "GET BOOKS route",
// 			method:         "GET",
// 			url:            "/books/",
// 			expectedStatus: http.StatusOK,
// 			expectedBody:   "Books fetched",
// 		},
// 		{
// 			name:           "GET BOOK BY ID route",
// 			method:         "GET",
// 			url:            "/books/1",
// 			expectedStatus: http.StatusOK,
// 			expectedBody:   "Book fetched",
// 		},
// 		{
// 			name:           "DELETE BOOK route",
// 			method:         "DELETE",
// 			url:            "/books/1",
// 			expectedStatus: http.StatusNoContent,
// 			expectedBody:   "",
// 		},
// 		{
// 			name:           "UPDATE BOOK route",
// 			method:         "PUT",
// 			url:            "/books/1",
// 			expectedStatus: http.StatusOK,
// 			expectedBody:   "Book Updated",
// 		},
// 		{
// 			name:           "CREATE BOOK route",
// 			method:         "POST",
// 			url:            "/books/",
// 			expectedStatus: http.StatusCreated,
// 			expectedBody:   "Book Created",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			req := httptest.NewRequest(tt.method, tt.url, nil)
// 			rec := httptest.NewRecorder()
// 			r.ServeHTTP(rec, req)

// 			if rec.Code != tt.expectedStatus {
// 				t.Errorf("Expected Status = %v; got  %v", tt.expectedStatus, rec.Code)
// 			}
// 			if rec.Body.String() != tt.expectedBody {
// 				t.Errorf("Expected body = %v; got %v", tt.expectedBody, rec.Body.String())
// 			}
// 			if contentTypeHeader := rec.Header().Get("Content-Type"); contentTypeHeader != "application/json" {
// 				t.Errorf("Expected application/json content-type header; got %v", contentTypeHeader)
// 			}
// 		})
// 	}
// }
