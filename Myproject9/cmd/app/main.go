package main

import (
    "Myproject9/internal/data"
    "Myproject9/internal/logic"
    "Myproject9/pkg/utils"
)

func main() {
    utils.Log("Приложение запущено!")
    data.LoadData()
    logic.ProcessData()
    utils.Log("Работа завершена успешно.")
}
