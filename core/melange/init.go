package melange

import (
	"github.com/NethermindEth/MelangeBE/ABIDecoder/logDecoder"
	MelangeInit "github.com/NethermindEth/MelangeBE/DataIngestor"
	Melange "github.com/NethermindEth/MelangeBE/DataIngestor/configs"
)

func AppInit(path string) (app Melange.App) {
	app, _ = MelangeInit.Init(path)
	store := make(logDecoder.ABIStore)
	app.ABIStore = &store
	return
}