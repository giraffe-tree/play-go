package sntp

import "github.com/beevik/ntp"

func main() {

	time, _ := ntp.Time("192.168.2.112")

	print(time.String())

}
