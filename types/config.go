package types

type Config struct {
	Settings       Settings       `toml:"settings"`
	Variables      Variables      `toml:"variables"`
	CustomCommands CustomCommands `toml:"custom_commands"`
}
type Settings struct {
	ServerName string `toml:"server_name"`
}
type Variables struct {
	String1 string `toml:"string1"`
	Int1    int    `toml:"int1"`
}
type Games struct {
	Name          string `toml:"name"`
	StatusCommand string `toml:"status_command"`
	StartCommand  string `toml:"start_command"`
	StopCommand   string `toml:"stop_command"`
}
type CustomCommands struct {
	Games []Games `toml:"games"`
}
