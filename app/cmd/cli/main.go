package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ==================== MODELS ====================

type Segmentation struct {
	ID               uint64    `gorm:"primaryKey"`
	UserID           int64     `gorm:"not null"`
	SegmentationType string    `gorm:"not null"`
	SegmentationName string    `gorm:"not null"`
	Data             any       `gorm:"type:jsonb;not null"`
	CreatedAt        time.Time `gorm:"not null"`
	UpdatedAt        time.Time `gorm:"not null"`
}

// ==================== USER SEGMENTATION ====================

type UserSegmentationsResponse struct {
	UserID        int64              `json:"user_id"`
	Segmentations []UserSegmentation `json:"segmentations"`
}

type UserSegmentation struct {
	Patients    []Patient   `json:"patients"`
	Specialties []Specialty `json:"specialties"`
	Drugs       []Drug      `json:"drugs"`
}

type Patient struct {
	Name string                 `json:"name"`
	Data map[string]interface{} `json:"data"`
}

type Specialty struct {
	Name string                 `json:"name"`
	Data map[string]interface{} `json:"data"`
}

type Drug struct {
	Name string                 `json:"name"`
	Data map[string]interface{} `json:"data"`
}

// ==================== REPOSITORY ====================

type Repository interface {
	Save(data interface{}) error
	GetDB() *gorm.DB
}

type SegmentationRepository struct {
	db *gorm.DB
}

func NewSegmentationRepository(db *gorm.DB) *SegmentationRepository {
	return &SegmentationRepository{db: db}
}

func (r *SegmentationRepository) Save(data interface{}) error {
	return r.db.Create(data).Error
}

func (r *SegmentationRepository) GetDB() *gorm.DB {
	return r.db
}

// ==================== SERVICE ====================

type SegmentationService struct {
	db *gorm.DB
}

func NewSegmentationService(db *gorm.DB) *SegmentationService {
	return &SegmentationService{db: db}
}

func (s *SegmentationService) ProcessRecord(record []string) error {
	repo := NewSegmentationRepository(s.db)

	if len(record) < 4 {
		return fmt.Errorf("invalid columns: %v", record)
	}

	userID, _ := strconv.ParseInt(strings.TrimSpace(record[0]), 10, 64)
	typ := strings.TrimSpace(record[1])
	name := strings.TrimSpace(record[2])
	dataStr := strings.TrimSpace(record[3])

	segmentation := map[string]interface{}{
		"user_id":           userID,
		"segmentation_type": typ,
		"segmentation_name": name,
		"data":              dataStr,
		"created_at":        time.Now(),
		"updated_at":        time.Now(),
	}

	err := repo.Save(segmentation)
	if err != nil {
		return err
	}

	return nil
}

// ==================== MAIN ====================

// Cenário:
// - Ler um arquivo CSV com 4 colunas: user_id, segment_type, segment_name e data
// - Este arquivo terá 1 milhão de linhas
// - O processamento deve ser performático e otimizado
// - Validar se os dados são válidos
// - Salvar no banco de dados
// - Se houver erro, mostrar ou salvar essa informação de alguma forma
func main() {
	db := getDB()
	service := NewSegmentationService(db)

	csvPath := "segmentations.csv"
	arquivo, err := os.Open(csvPath)
	if err != nil {
		log.Printf("open csv: %v", err)
		return
	}
	defer arquivo.Close()

	leitor := csv.NewReader(bufio.NewReader(arquivo))
	records, _ := leitor.ReadAll()

	errs := []error{}

	for _, record := range records {
		err := service.ProcessRecord(record)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		log.Fatalf("Processamento concluído com erros: %v", errs)
	}

	log.Printf("Processamento concluído sem erros")
}

func getDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"), &gorm.Config{})

	if err != nil {
		fmt.Println("Erro ao tentar se conectar com o banco de dados.", err)
		return nil
	}
	fmt.Println("Connexão com banco de dados com sucesso!")
	return db
}