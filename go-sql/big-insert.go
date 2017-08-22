package main

import "time"

// begin insert sql OMIT
const (
	insertSQL = `
insert into DiaryItems(
		UserId, DiaryDate, MealID, Position, MealName, Version, DocumentId, DocumentVersion,
		UpdatedAt, EntryType, DatabaseID, ProductID, Barcode, Brand, Name, GS1Name, Company,
    	Keyword, ChoicePath, ChoiceTerms, ChoiceLevels, UserSearchText, Energy,
		Protein, TotalFat, SaturatedFat, Carbohydrate, Sugars, Sodium, Potassium,
		Calcium, Phosphorous, Fibre, EnergyAttributes, ProteinAttributes,
		TotalFatAttributes, SaturatedFatAttributes, CarbohydrateAttributes,
		SugarsAttributes, SodiumAttributes, PotassiumAttributes, CalciumAttributes,
		PhosphorousAttributes, FibreAttributes, ServingsQuantity, ServingsHasFraction,
		MeasureID, MeasureGrams, MeasureName, UseMilliLitreUnits, CustomServingsPerPackage,
		CustomMeasureGrams, CustomMeasureName, CustomDisplayCalories, ExerciseDuration,
		ExerciseEnergyBurnedPerMinute, ExerciseMets, ExerciseDisplayCalories,
		RecipeServingName, RecipeServingsGrams, CreateTimeStamp, ModifyTimeStamp,
		ConsumeTimeStamp
) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,
	?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
)

// end insert sql OMIT

func doInsert() {
	// begin insert example OMIT
	type User struct {
		UserID     int `sql:"primary key autoincrement"`
		GivenName  string
		FamilyName string
		DOB        time.Time
		Birthplace string
		// ... lots and lots of other columns here ...
		CreatedAt  time.Time
		ModifiedAt time.Time
	}

	user := &User{
		GivenName:  "Barack",
		FamilyName: "Obama",
		DOB:        time.Date(1961, 8, 4, 0, 0, 0, 0, time.UTC),
		Birthplace: "Hawaii",
		// ... lots and lots more fields here ...
	}

	err := schema.Exec(db, user, "insert into users({}) values({})") // HL
	// end insert example OMIT
}
