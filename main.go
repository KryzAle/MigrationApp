package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	adaptersfromproto "testfb/adaptersFromProto"
	"testfb/db"
	"testfb/utils"
	"time"

	"gorm.io/gorm"
)

type TemporalMigrationRepository struct {
	connection *gorm.DB
}

func (TemporalRepository *TemporalMigrationRepository) Init() *TemporalMigrationRepository {
	dbConnection := utils.GetConnection()
	return &TemporalMigrationRepository{
		connection: dbConnection,
	}
}

func main() {
	temporalRepository := new(TemporalMigrationRepository).Init()
	//temporalRepository.connection.AutoMigrate(&db.UserOrm{}, &db.SkillOrm{}, &db.CategoryOrm{}, &db.UserSkillOrm{})
	temporalRepository.migrateCategoriesAndSkills()
	temporalRepository.migrateUserSkills()
}

func getJsonFile(fileName string) ([]byte, error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue, nil
}

func (TemporalRepository *TemporalMigrationRepository) migrateCategoriesAndSkills() {
	var technologiesFromProto adaptersfromproto.TechnologiesFromProtoAdapter
	var categories = []string{"languages", "databases", "cloud_platforms", "web_frameworks", "other_frameworks_and_libraries", "other_tools", "operating_systems", "quality_assurance_tools"}
	technologiesBytes, err := getJsonFile("technologies.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = json.Unmarshal(technologiesBytes, &technologiesFromProto)
	if err != nil {
		fmt.Println(err)
	}
	for _, category := range categories {
		categoryId, err := TemporalRepository.CreateCategory(category)
		if err != nil {
			panic(err)
		}
		for _, technology := range technologiesFromProto.Technologies {
			if technology.Category == category {
				TemporalRepository.CreateSkill(technology.Label, categoryId)
			}
		}
	}
}

func (TemporalRepository *TemporalMigrationRepository) migrateUserSkills() {
	var userSkillsFromProto adaptersfromproto.UserSkillsFromProtoAdapter
	userSkillsBytes, err := getJsonFile("userSkills.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = json.Unmarshal(userSkillsBytes, &userSkillsFromProto)
	if err != nil {
		fmt.Println(err)
	}

	for _, userSkill := range userSkillsFromProto.Data {
		TemporalRepository.CreateUser(userSkill.Email, userSkill.English)
		for _, skill := range userSkill.UserSkills {
			skillId, err := TemporalRepository.GetSkillIdByName(skill.SkillName)
			if err != nil {
				panic(err)
			}
			TemporalRepository.CreateUserSkill(userSkill.Email, skillId, skill.Level)
		}
	}
}

func (TemporalRepository *TemporalMigrationRepository) GetSkillIdByName(skillName string) (uint, error) {

	skill := db.SkillOrm{}

	id := TemporalRepository.connection.Where("name = ?", skillName).First(&skill)
	if id.Error != nil {
		return 0, errors.New("no skills found")
	}
	return skill.ID, nil
}
func (TemporalRepository *TemporalMigrationRepository) CreateSkill(skillName string, categoryId uint) error {

	skill := db.SkillOrm{
		Name:      skillName,
		CategorId: categoryId,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	result := TemporalRepository.connection.Create(&skill)

	if result.Error != nil {
		return errors.New("error inserting new skill")
	}

	return nil
}
func (TemporalRepository *TemporalMigrationRepository) CreateCategory(categoryName string) (uint, error) {

	category := db.CategoryOrm{
		Name:      categoryName,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	result := TemporalRepository.connection.Create(&category)

	if result.Error != nil {
		return 0, errors.New("error inserting new category")
	}

	return category.ID, nil
}
func (TemporalRepository *TemporalMigrationRepository) CreateUser(email string, englishLevel uint) error {

	user := db.UserOrm{
		Email:        email,
		EnglishLevel: englishLevel,
	}

	result := TemporalRepository.connection.Create(&user)

	if result.Error != nil {
		return errors.New("error inserting new user")
	}

	return nil
}
func (TemporalRepository *TemporalMigrationRepository) CreateUserSkill(userEmail string, skillId uint, experience uint) error {

	userSkill := db.UserSkillOrm{
		UserEmail:  userEmail,
		SkillID:    skillId,
		Experience: experience,
	}

	result := TemporalRepository.connection.Create(&userSkill)

	if result.Error != nil {
		return errors.New("error inserting new category")
	}

	return nil
}
