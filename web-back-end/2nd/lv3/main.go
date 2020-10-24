package main

type VideoPage struct {
	Author   User
	Video    Video
	Likes    int
	Stars    int
	Coins    int
}

type User struct {
	Name	string
}

type Video struct {
	Name	string
}

func (videoPage *VideoPage) toLike() {
	videoPage.Likes++
}
func (videoPage *VideoPage) toStar() {
	videoPage.Stars++
}
func (videoPage *VideoPage) toCoin() {
	videoPage.Coins++
}
func (videoPage *VideoPage) oneClickThreeFollow() {
	videoPage.Likes++
	videoPage.Stars++
	videoPage.Coins++
}
func publishVideo(authorName string, videoName string) VideoPage{
	return VideoPage{
		Author: User{
			Name: authorName,
		},
		Video: Video{
			Name: videoName,
		},
	}
}
func main() {
	
}
