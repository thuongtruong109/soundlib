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
	Gray string = "\033[37m"
)

const (
	GENRE_PATH    string = "database/genres.json"
	TRACK_PATH    string = "database/tracks.json"
	ARTIST_PATH   string = "database/artists.json"
	ALBUM_PATH    string = "database/albums.json"
	PLAYLIST_PATH string = "database/playlists.json"
)

const (
	CREATE_FAILED string = "create failed"
	GET_FAILED    string = "get failed"
	UPDATE_FAILED string = "update failed"
	DELETE_FAILED string = "delete failed"

	CREATE_SUCCESS string = "create successfully"
	GET_SUCCESS    string = "get successfully"
	UPDATE_SUCCESS string = "update successfully"
	DELETE_SUCCESS string = "delete successfully"
)

const (
	NOT_FOUND_DATA string = "not found data"
	EMPTY_DATA     string = "empty data"
)

const (
	APP_NAME      string = "SOUNDLIB"
	THANKYOU_TEXT string = "\n --- ♥ Thank you for using service!"
	AGAIN_TEXT    string = "\n ¤ Do you want to continue? (y/n): "
)