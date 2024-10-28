package main

import (
	"log"
	"net/http"

	"maxischmaxi/jstreams/account"
	"maxischmaxi/jstreams/account/v1/accountv1connect"
	"maxischmaxi/jstreams/assets"
	"maxischmaxi/jstreams/assets/v1/assetsv1connect"
	"maxischmaxi/jstreams/champions"
	"maxischmaxi/jstreams/champions/v1/championsv1connect"
	"maxischmaxi/jstreams/entries"
	"maxischmaxi/jstreams/entries/v1/entriesv1connect"
	"maxischmaxi/jstreams/items"
	"maxischmaxi/jstreams/items/v1/itemsv1connect"
	"maxischmaxi/jstreams/masteries"
	"maxischmaxi/jstreams/masteries/v1/masteriesv1connect"
	"maxischmaxi/jstreams/matches"
	"maxischmaxi/jstreams/matches/v1/matchesv1connect"
	"maxischmaxi/jstreams/summoner"
	"maxischmaxi/jstreams/summoner/v1/summonerv1connect"
	"maxischmaxi/jstreams/tier"
	"maxischmaxi/jstreams/tier/v1/tierv1connect"
	"maxischmaxi/jstreams/utils"
	"maxischmaxi/jstreams/version"
	"maxischmaxi/jstreams/version/v1/versionv1connect"

	"connectrpc.com/connect"
)

func main() {
	mux := http.NewServeMux()
	interceptors := connect.WithInterceptors(LoggingInterceptor())

	items := &items.ItemsServer{}
	path, handler := itemsv1connect.NewItemsServiceHandler(items, interceptors)
	utils.HandleServer(mux, path, handler)

	champions := &champions.ChampionsServer{}
	path, handler = championsv1connect.NewChampionsServiceHandler(champions, interceptors)
	utils.HandleServer(mux, path, handler)

	summoner := &summoner.SummonerServer{}
	path, handler = summonerv1connect.NewSummonerServiceHandler(summoner, interceptors)
	utils.HandleServer(mux, path, handler)

	matches := &matches.MatchesServer{}
	path, handler = matchesv1connect.NewMatchesServiceHandler(matches, interceptors)
	utils.HandleServer(mux, path, handler)

	entries := &entries.EntriesServer{}
	path, handler = entriesv1connect.NewEntriesServiceHandler(entries, interceptors)
	utils.HandleServer(mux, path, handler)

	assets := &assets.AssetsServer{}
	path, handler = assetsv1connect.NewAssetsServiceHandler(assets, interceptors)
	utils.HandleServer(mux, path, handler)

	tier := &tier.TierServer{}
	path, handler = tierv1connect.NewTierServiceHandler(tier, interceptors)
	utils.HandleServer(mux, path, handler)

	account := &account.AccountServer{}
	path, handler = accountv1connect.NewAccountServiceHandler(account, interceptors)
	utils.HandleServer(mux, path, handler)

	version := &version.VersionServer{}
	path, handler = versionv1connect.NewVersionServiceHandler(version, interceptors)
	utils.HandleServer(mux, path, handler)

	masteris := &masteries.MasteriesServer{}
	path, handler = masteriesv1connect.NewMasteriesServiceHandler(masteris, interceptors)
	utils.HandleServer(mux, path, handler)

	log.Fatal(http.ListenAndServe(
		"0.0.0.0:8080",
		mux,
	))
}
