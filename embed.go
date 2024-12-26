package root

import _ "embed"

//go:embed/assets/ascii-art.txt
var AsciiArt string
//go:embed/config/config.json
var ConfigFile string