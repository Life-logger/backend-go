package categories

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql" // MySQL 드라이버를 포함
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Open("mysql", "user:password@/init_db") // DB 연결 초기화
	if err != nil {
		log.Fatalln(err)
	}
}

type Categories struct {
	CategoriesId        int    `db:"categories_id"`
	UserId              int    `db:"user_id"`
	Color               string `db:"color"`
	CategoriesTitle     string `db:"categories_title"`
	CategoriesTotalTime int    `db:"categories_total_time"`
}

// 카테고리 추가
func AddCategory(category Categories) error {
	if db == nil {
		return fmt.Errorf("DB connection is not initialized")
	}

	query := `
        INSERT INTO Categories (categories_id, user_id, color, categories_title, categories_total_time) 
        VALUES (?, ?, ?, ?, ?)
    `

	_, err := db.Exec(query, category.CategoriesId, category.UserId, category.Color, category.CategoriesTitle, category.CategoriesTotalTime)
	if err != nil {
		return fmt.Errorf("failed to insert data: %v", err)
	}

	fmt.Printf("Added Category: CategoriesId: %d, UserId: %d, Color: %s, CategoriesTitle: %s, CategoriesTotalTime: %d\n",
		category.CategoriesId, category.UserId, category.Color, category.CategoriesTitle, category.CategoriesTotalTime)
	return nil
}

// 카테고리 수정
func UpdateCategory(category Categories) error {
	if db == nil {
		return fmt.Errorf("DB connection is not initialized")
	}

	query := `
        UPDATE Categories 
        SET color = ?, categories_title = ?
        WHERE categories_id = ?
    `

	_, err := db.Exec(query, category.Color, category.CategoriesTitle)
	if err != nil {
		return fmt.Errorf("failed to update data: %v", err)
	}

	fmt.Printf("Updated Category: CategoriesId: %d, UserId: %d, Color: %s, CategoriesTitle: %s, CategoriesTotalTime: %d\n",
		category.CategoriesId, category.UserId, category.Color, category.CategoriesTitle, category.CategoriesTotalTime)
	return nil
}

// 카테고리 삭제
func DeleteCategory(categoriesId int) error {
	if db == nil {
		return fmt.Errorf("DB connection is not initialized")
	}

	query := `
        DELETE FROM Categories 
        WHERE categories_id = ?
    `

	_, err := db.Exec(query, categoriesId)
	if err != nil {
		return fmt.Errorf("failed to delete data: %v", err)
	}

	fmt.Printf("Deleted Category with CategoriesId: %d\n", categoriesId)
	return nil
}

// 카테고리 조회
func GetCategory(categoriesId int) (*Categories, error) {
	if db == nil {
		return nil, fmt.Errorf("DB connection is not initialized")
	}

	var category Categories
	query := `
        SELECT categories_id, user_id, color, categories_title, categories_total_time
        FROM Categories
        WHERE categories_id = ?
    `

	err := db.Get(&category, query, categoriesId)
	if err != nil {
		return nil, fmt.Errorf("failed to get data: %v", err)
	}

	return &category, nil
}

// 카테고리 추가 핸들러
func addCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var newCategory Categories

	// JSON 바디 파싱
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// 카테고리 추가
	err = AddCategory(newCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 성공 응답
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCategory)
}

// 카테고리 수정 핸들러
func updateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var category Categories

	// JSON 바디 파싱
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// 카테고리 수정
	err = UpdateCategory(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 성공 응답
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(category)
}

// 카테고리 삭제 핸들러
func deleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = DeleteCategory(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 성공 응답
	w.WriteHeader(http.StatusNoContent)
}

// HTTP 서버 실행 및 핸들러 등록
func main() {
	http.HandleFunc("/categories/add", addCategoryHandler)
	http.HandleFunc("/categories/update", updateCategoryHandler)
	http.HandleFunc("/categories/delete", deleteCategoryHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// 예제 함수: 실제 사용 시 주석 처리하거나 별도로 테스트 코드로 사용
func ExampleCategoriesUsage() {
	// 카테고리 추가
	newCategory := Categories{
		UserId:              1,                 // 예시 사용자 ID
		Color:               "#FFF",            // 프론트에서 받아와야 하는 값
		CategoriesTitle:     "Sample Category", // 프론트에서 받아와야 하는 값
		CategoriesTotalTime: 0,                 // 처음엔 0으로 설정
	}

	// 카테고리 추가 요청
	data, err := json.Marshal(newCategory)
	if err != nil {
		log.Fatalf("Error marshalling new category: %v", err)
	}

	resp, err := http.Post("http://localhost:8080/categories/add", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Error adding category: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Fatalf("Failed to add category: %s", body)
	}

	fmt.Println("Category added successfully")

	// 카테고리 조회
	categoryId := 1 // 예시 카테고리 ID
	resp, err = http.Get(fmt.Sprintf("http://localhost:8080/categories/get?id=%d", categoryId))
	if err != nil {
		log.Fatalf("Error getting category: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Fatalf("Failed to get category: %s", body)
	}

	var category Categories
	err = json.NewDecoder(resp.Body).Decode(&category)
	if err != nil {
		log.Fatalf("Error decoding category: %v", err)
	}

	fmt.Printf("Retrieved Category: %+v\n", category)

	// 카테고리 수정
	category.Color = "Red" // 변경할 값
	data, err = json.Marshal(category)
	if err != nil {
		log.Fatalf("Error marshalling updated category: %v", err)
	}

	req, err := http.NewRequest(http.MethodPut, "http://localhost:8080/categories/update", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Error creating request to update category: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error updating category: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Fatalf("Failed to update category: %s", body)
	}

	fmt.Println("Category updated successfully")

	// 카테고리 삭제
	req, err = http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/categories/delete?id=%d", categoryId), nil)
	if err != nil {
		log.Fatalf("Error creating request to delete category: %v", err)
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error deleting category: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Fatalf("Failed to delete category: %s", body)
	}

	fmt.Println("Category deleted successfully")
}
