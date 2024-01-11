package main

import (
	"bufio"
	"fmt"
	"os"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err // Если по какой-то причине мы получили ошибку, проигнорируем эту итерацию
	}

	if info.IsDir() {
		return nil // Проигнорируем директории
	}

	file, err := os.Open(info.Name())
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		if len(s.Text()) > 0 { // возвращает true, пока файл не будет прочитан до конца
			fmt.Printf("%s\n", s.Text()) // s.Text() содержит данные, считанные на данной итерации
		}
	}
	return nil
}
