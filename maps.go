package main

import "fmt"

type bookmarkMap = map[string]string

func maps() {
	bookmarks := bookmarkMap{}
Menu:
	for {
		userAnswer := getUserAnswer()

		switch userAnswer {
		case "1":
			printBookmarks(bookmarks)
		case "2":
			addBookmark(bookmarks)
		case "3":
			deleteBookmark(bookmarks)
		default:
			break Menu
		}
	}
}

func printBookmarks(bookmarks bookmarkMap) {
	if len(bookmarks) == 0 {
		fmt.Println("Пока нет закладок")
	}

	for key, value := range bookmarks {
		fmt.Println(key, ":", value)
	}
}

func getUserAnswer() string {
	var userAnswer string

	fmt.Println("Меню")
	fmt.Println("1 - вывести список закладок")
	fmt.Println("2 - добавить закладку")
	fmt.Println("3 - удалить закладку")
	fmt.Println("4 - выход")

	fmt.Scan(&userAnswer)

	return userAnswer
}

func createBookmark() (string, string) {
	var bookmarkName string
	var bookmarkLink string

	fmt.Println("Введите название закладки: ")
	fmt.Scan(&bookmarkName)
	fmt.Println("Введите ссылку: ")
	fmt.Scan(&bookmarkLink)

	return bookmarkName, bookmarkLink
}

func addBookmark(bookmarks bookmarkMap) {
	newBookmarkName, newBookmarkLink := createBookmark()
	bookmarks[newBookmarkName] = newBookmarkLink

	fmt.Println("Закладка добавлена!")
}

func deleteBookmark(bookmarks bookmarkMap) {
	var bookmarkName string

	fmt.Print("Введите название закладки, которую нужно удалить: ")
	fmt.Scan(&bookmarkName)

	delete(bookmarks, bookmarkName)

	fmt.Printf("Закладка удалена: %v", bookmarks)
}
