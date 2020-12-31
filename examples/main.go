package main

import (
	"fmt"
	emoji "github.com/sujit-baniya/Go-Emoji-Utils"
)

func main() {
	str1 := "포기하고 싶은 순🤮간들 바🅱️로 그 순간   🤠빨리 '희망의🤧 스위치'😇👌🎍😍를 올리자. 찰칵! I like a :pizza: and :sushi:!!"
	str := emoji.Parse(str1)
	matches := emoji.FindAll(str)
	totalUniqueEmoji := len(matches)
	emojiRemoved := emoji.RemoveAll(str)
	fmt.Println(fmt.Sprintf("Total Emojis %d", totalUniqueEmoji))
	fmt.Println(emojiRemoved)

}
