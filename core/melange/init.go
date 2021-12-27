package melange

import (
	Melange "github.com/NethermindEth/MelangeBE/DataIngestor/configs"
	MelangeInit "github.com/NethermindEth/MelangeBE/DataIngestor"
)

func AppInit(path string) (app Melange.App) {
	app, _ = MelangeInit.Init(path)
	return
}