package ui

import "supersonic/backend"

type PageName int

const (
	Blank PageName = iota
	Album
	Albums
	Artist
	Artists
	Playlist
	Playlists
)

type Route struct {
	Page PageName
	Arg  string
}

func AlbumsRoute(sortOrder backend.AlbumSortOrder) Route {
	return Route{Page: Albums, Arg: string(sortOrder)}
}

func ArtistRoute(artistID string) Route {
	return Route{Page: Artist, Arg: artistID}
}

func AlbumRoute(albumID string) Route {
	return Route{Page: Album, Arg: albumID}
}

type NavigationHandler interface {
	SetPage(Page)
}

type Router struct {
	App *backend.App
	Nav NavigationHandler
}

func NewRouter(app *backend.App, nav NavigationHandler) Router {
	return Router{
		App: app,
		Nav: nav,
	}
}

func (r Router) CreatePage(rte Route) Page {
	switch rte.Page {
	case Album:
		return NewAlbumPage(rte.Arg, r.App.LibraryManager, r.App.ImageManager, r.OpenRoute)
	case Albums:
		return NewAlbumsPage("Albums", rte.Arg, r.App.LibraryManager, r.App.ImageManager, r.OpenRoute)
	case Artist:
		return NewArtistPage(rte.Arg, r.App.ServerManager, r.App.ImageManager, r.OpenRoute)
	}
	return nil
}

func (r Router) OpenRoute(rte Route) {
	r.Nav.SetPage(r.CreatePage(rte))
}