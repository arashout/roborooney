package main

//go:generate echo '{"apiToken":"","channelId":""}' > config.json

func main() {
	robo := NewRobo()
	robo.Connect()
}
