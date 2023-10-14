package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Character struct {
	Name              string
	Age               int
	Gender            string
	Appearance        string
	Background        string
	Hometown          string
	PastEvents        string
	Personality       string
	GoalsDesires      string
	StrengthsWeaknesses string
	Relationships     string
	GrowthChange      string
	MotifsSymbols     string
	NameMeaning       string
}

func main() {
	// 作品名の入力
	fmt.Print("作品名を入力してください: ")
	var workName string
	fmt.Scan(&workName)

	// CSVファイル名を作成
	csvFileName := workName + ".csv"

	// CSVファイルを作成
	createCSVFileWithHeader(csvFileName)

	// ユーザーが新しいキャラクター情報を入力
	characterCount := 0 // キャラクターの数をカウントする変数
	for {
		characterCount++
		fmt.Printf("登場人物%d人目\n", characterCount)
		character := collectCharacterInfo()
		writeCharacterToCSV(character, csvFileName)

		fmt.Print("新しいキャラクターを追加しますか？ (y/n): ")
		var response string
		fmt.Scan(&response)
		if strings.ToLower(response) != "y" {
			break
		}
	}
}

func createCSVFileWithHeader(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("CSVファイルを作成できません: ", err)
		os.Exit(1)
	}
	defer file.Close()

	// ヘッダー行を書き込む
	header := []string{
		"名前",
		"年齢",
		"性別",
		"外見",
		"背景と経歴",
		"出身地",
		"過去の出来事",
		"性格",
		"目標と欲望",
		"強みと弱点",
		"関係性",
		"成長と変化",
		"モチーフと象徴",
		"名前の意味",
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write(header)
}

func collectCharacterInfo() Character {
	var character Character

	fmt.Print("名前: ")
	fmt.Scanln(&character.Name)

	fmt.Print("年齢: ")
	fmt.Scan(&character.Age)

	fmt.Print("性別: ")
	fmt.Scanln(&character.Gender)

	fmt.Print("外見: ")
	fmt.Scanln(&character.Appearance)

	fmt.Print("背景と経歴: ")
	fmt.Scanln(&character.Background)

	fmt.Print("出身地: ")
	fmt.Scanln(&character.Hometown)

	fmt.Print("過去の出来事: ")
	fmt.Scanln(&character.PastEvents)

	fmt.Print("性格: ")
	fmt.Scanln(&character.Personality)

	fmt.Print("目標と欲望: ")
	fmt.Scanln(&character.GoalsDesires)

	fmt.Print("強みと弱点: ")
	fmt.Scanln(&character.StrengthsWeaknesses)

	fmt.Print("主要登場人物との関係性: ")
	fmt.Scanln(&character.Relationships)

	fmt.Print("成長と変化: ")
	fmt.Scanln(&character.GrowthChange)

	fmt.Print("モチーフと象徴: ")
	fmt.Scanln(&character.MotifsSymbols)

	fmt.Print("名前の意味: ")
	fmt.Scanln(&character.NameMeaning)

	return character
}


func writeCharacterToCSV(character Character, fileName string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("CSVファイルを開けません: ", err)
		os.Exit(1)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	characterSlice := []string{
		character.Name,
		fmt.Sprintf("%d", character.Age),
		character.Gender,
		character.Appearance,
		character.Background,
		character.Hometown,
		character.PastEvents,
		character.Personality,
		character.GoalsDesires,
		character.StrengthsWeaknesses,
		character.Relationships,
		character.GrowthChange,
		character.MotifsSymbols,
		character.NameMeaning,
	}

	if err := writer.Write(characterSlice); err != nil {
		fmt.Println("CSVファイルに書き込めません: ", err)
		os.Exit(1)
	}
	fmt.Println("キャラクター情報をCSVファイルに書き込みました")
}
