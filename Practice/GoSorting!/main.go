package main

import (
	"encoding/json"
	"strconv"
	"strings"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("GoSorting!")

	label1 := widget.NewLabel("Введите массив чисел и выберите метод сортировки.")
	label2 := widget.NewLabel("Отсортированный массив:")
	entry := widget.NewEntry()
	answer := widget.NewLabel("")

	btn1 := widget.NewButton("Пузырёк", func() {
		arrS := strings.Split(entry.Text, ",")

		var arr []int
		for _, s := range arrS {
			num, err := strconv.Atoi(s)
			if err == nil {
				arr = append(arr, num)
			}
		}

		len := len(arr)
		for i := 0; i < len-1; i++ {
			for j := 0; j < len-i-1; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
		res, _ := json.Marshal(arr)

		answer.SetText(strings.Trim(string(res), ""))
	})

	btn2 := widget.NewButton("Выбором", func() {
		arrSs := strings.Split(entry.Text, ",")
		var arr []int
		for _, s := range arrSs {
			num, err := strconv.Atoi(s)
			if err == nil {
				arr = append(arr, num)
			}
		}

		len := len(arr)
		for i := 0; i < len-1; i++ {
			minIndex := i
			for j := i + 1; j < len; j++ {
				if arr[j] < arr[minIndex] {
					arr[j], arr[minIndex] = arr[minIndex], arr[j]
				}
			}
		}
		res, _ := json.Marshal(arr)

		answer.SetText(strings.Trim(string(res), ""))
	})

	btn3 := widget.NewButton("Вставки", func() {
		arrS := strings.Split(entry.Text, ",")
		var arr []int
		for _, s := range arrS {
			num, err := strconv.Atoi(s)
			if err == nil {
				arr = append(arr, num)
			}
		}

		len := len(arr)
		for i := 1; i < len; i++ {
			for j := 0; j < i; j++ {
				if arr[j] > arr[i] {
					arr[j], arr[i] = arr[i], arr[j]
				}
			}
		}
		res, _ := json.Marshal(arr)

		answer.SetText(strings.Trim(string(res), ""))
	})

	btn4 := widget.NewButton("Шелла", func() {
		arrS := strings.Split(entry.Text, ",")
		var arr []int
		for _, s := range arrS {
			num, err := strconv.Atoi(s)
			if err == nil {
				arr = append(arr, num)
			}
		}

		for d := int(len(arr) / 2); d > 0; d /= 2 {
			for j := d; j < len(arr); j++ {
				for k := j; k >= d && arr[k-d] > arr[k]; k -= d {
					arr[k], arr[k-d] = arr[k-d], arr[k]
				}
			}
		}
		res, _ := json.Marshal(arr)

		answer.SetText(strings.Trim(string(res), ""))
	})

	w.SetContent(container.NewVBox(
		label1,
		entry,
		btn1,
		btn2,
		btn3,
		btn4,
		label2,
		answer,
	))

	w.ShowAndRun()
}
