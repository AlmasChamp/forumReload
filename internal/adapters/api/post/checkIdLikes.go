package post

import (
	"strconv"
	"strings"
)

func CheckIdLikes(commLike string, commDislike string) (string, int) {
	if commLike == "" && commDislike == "" {
		return "", 0
	}
	like := strings.Split(commLike, " ")
	dislike := strings.Split(commDislike, " ")
	if like[0] != "" {
		idPost, _ := strconv.Atoi(like[1])
		return like[0], idPost
	}
	idPost, _ := strconv.Atoi(dislike[1])
	return dislike[0], idPost
}
