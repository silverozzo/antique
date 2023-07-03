package deck

import "antique/handy"

type AssetCard struct {
	handy.CardInterface
}

type AssetDeck struct {
	handy.DeckInterface
}

func newAsset(assetName string) *AssetCard {
	return &AssetCard{
		handy.NewCard(handy.Side{AssetNameTag: assetName}, handy.Side{}),
	}
}

func NewAssetDeck() *AssetDeck {
	cards := []handy.CardInterface{
		newAsset("Револьвер"),
		newAsset("Винтовка"),
		newAsset("Динамит"),
		newAsset("Топор"),
	}

	deck := AssetDeck{
		handy.NewDeck(cards),
	}

	return &deck
}
