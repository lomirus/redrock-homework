package main

import "fmt"

type VideoPage struct {
	id       string
	Author   User
	Video    Video
	Comments []Comment
}

type User struct {
	UID       int
	Name      string
	Signature string
	Avatar    string // 存储头像地址
	VIP       bool
	Level     int
	Followers []int // 存储UID
	Following []int // 存储UID
}
type Video struct {
	Name          string
	Time          int
	MaxResolution [2]int
	Size          int
}
type Comment struct {
	User    int // 存储UID
	Time    string
	Content string
}

func main() {
	var videoPage VideoPage = VideoPage{
		id: "1x54y1C727",
		Author: User{
			UID:       2177677,
			Name:      "出格字幕组",
			Signature: "一个人的字幕组，微博@一只出格君，哔哩哔哩私人号也是@一只出格君",
			Avatar:    "https://i2.hdslb.com/bfs/face/73f6f268277b865c6d0069d2da7edea0249f8310.jpg",
			VIP:       false,
			Level:     6,
			Followers: []int{},
			Following: []int{},
		},
		Video: Video{
			Name:          "【氰化欢乐秀】水晶球里觅真爱",
			Time:          73,
			MaxResolution: [2]int{1920, 1080},
			Size:          1024,
		},
		Comments: []Comment{
			{}, {}, {},
		},
	}
	fmt.Println(videoPage)
}
