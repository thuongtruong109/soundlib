package constants

const (
	DEBUG int8 = iota
	INFO
	WARNING
	ERROR
	FATAL

	LABEL
	DESC
	QUERY
	INPUT
)

const (
	GENRE_PATH             string = "database/genres.json"
	TRACK_PATH             string = "database/tracks.json"
	ARTIST_PATH            string = "database/artists.json"
	ALBUM_PATH             string = "database/albums.json"
	PLAYLIST_PATH          string = "database/playlists.json"
	TRACK_IN_ALBUM_PATH    string = "database/track_in_album.json"
	TRACK_IN_PLAYLIST_PATH string = "database/track_in_playlist.json"
)

const (
	APP_NAME      string = "SOUNDLIB"
	THANKYOU_TEXT string = "\n --- ♥ Thank you for using service!"
	AGAIN_TEXT    string = "\n ¤ Do you want to continue? (y/n): "
)