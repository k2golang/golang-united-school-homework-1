package solution

import "github.com/kyokomi/emoji/v2"

func GetMessage() string {
	str := emoji.Sprint(":world_map:")
	return "Hello " + str + "!"
}
