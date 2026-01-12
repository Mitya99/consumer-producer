package main

import (
	"consumer-producer/logger"
	"consumer-producer/ollamaclient" // my ollama
	"fmt"
	"os"

	"go.uber.org/zap"
)

func main() {

	//Loger Initialization
	logger.InitLogger()
	defer logger.CloseLogger()

	// 	Obtaining a brief summary of the retrieved set of strings (using LLM)
	var fifty_str string = "Гостиная Анны Павловны начала понемногу наполняться. Приехала высшая знать Петербурга, люди самые разнородные по возрастам и характерам, но одинаковые по обществу, в каком все жили; приехала дочь князя Василия, красавица Элен, заехавшая за отцом, чтобы с ним вместе ехать на праздник посланника. Она была в шифре и бальном платье. Приехала и известная, как la femme la plus séduisante de Pétersbourg [самая обворожительная женщина в Петербурге,], молодая, маленькая княгиня Болконская, прошлую зиму вышедшая замуж и теперь не выезжавшая в  большой свет по причине своей беременности, но ездившая еще на небольшие вечера. Приехал князь Ипполит, сын князя Василия, с Мортемаром, которого он представил; приехал и аббат Морио и многие другие."
	prompt := "Сделай краткое содержание текста, выдавая в ответ ничего лишнего кроме результата: " + fifty_str

	ollamaResp, err := ollamaclient.Generate(prompt)
	if err != nil {
		logger.Error("Ошибка при обращении к Ollama: %s\n", zap.String("err", err.Error()))
		os.Exit(1)
	}

	// Логирование результата
	logger.Info("Сгенерировано краткое содержание",
		zap.String("текст", fifty_str),
		zap.String("краткое_содержание", ollamaResp),
	)

	// Вывод результата
	fmt.Println(ollamaResp)
}
