package bst_models

import "net/http"

type Error struct {
	Code int
	CorrespondingHttpCode int
	Message string
}

func(e1 Error) Equals(e2 Error) bool {
	return e1.Code == e2.Code
}

// common errors 0-99
var (
	ErrorOK = Error {
		Code: 0,
		CorrespondingHttpCode: http.StatusOK,
		Message: "OK",
	}
	ErrorJwt = Error {
		Code: 1,
		CorrespondingHttpCode: http.StatusUnauthorized,
		Message: "invalid token supplied",
	}
	ErrorJwtProfile = Error {
		Code: 1,
		CorrespondingHttpCode: http.StatusUnauthorized,
		Message: "token did not contain correct profile structure",
	}

	ErrorDBConnection = Error {
		Code: 5,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "database connection failed",
	}

	ErrorBadHeader = Error {
		Code: 20,
		CorrespondingHttpCode: http.StatusBadRequest,
		Message: "eader did not contain correct contents",
	}
	ErrorBadBody = Error {
		Code: 21,
		CorrespondingHttpCode: http.StatusBadRequest,
		Message: "request body did not contain correct contents",
	}
	ErrorBadQuery = Error {
		Code: 22,
		CorrespondingHttpCode: http.StatusBadRequest,
		Message: "query did not contain correct contents",
	}

	ErrorStringSearch = Error {
		Code: 30,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to find required substring",
	}

	ErrorStringParse = Error {
		Code: 40,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to parse string",
	}

	ErrorTimeParse = Error {
		Code: 41,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to parse time",
	}

	ErrorTimeLocLoad = Error {
		Code: 50,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to load timezone",
	}

	ErrorCreateRequest = Error {
		Code: 60,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed creating http request",
	}

	ErrorClientRequest = Error {
		Code: 70,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "http client request failed",
	}
	ErrorClientResponse = Error {
		Code: 71,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "http client response failed",
	}
	ErrorRoundTrip = Error {
		Code: 72,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "round trip failed",
	}

	ErrorJsonDecode = Error {
		Code: 90,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "json decoding error",
	}
	ErrorJsonEncode = Error {
		Code: 91,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "json encoding error",
	}
)

// api errors 100-199
var (
	ErrorPrivateUser = Error {
		Code: 110,
		CorrespondingHttpCode: http.StatusForbidden,
		Message: "user does not allow data views",
	}
	ErrorUnknownUser = Error {
		Code: 111,
		CorrespondingHttpCode: http.StatusNotFound,
		Message: "user does not exist in this context",
	}

	ErrorNoUserCache = Error {
		Code: 120,
		CorrespondingHttpCode: http.StatusNotFound,
		Message: "no cache exists for requested user",
	}
	ErrorNoSuppliedUserCache = Error {
		Code: 121,
		CorrespondingHttpCode: http.StatusBadRequest,
		Message: "no user supplied for cache request",
	}
	ErrorWriteUserCache = Error {
		Code: 122,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "could not write profile changes to db",
	}

	ErrorWriteCookie = Error {
		Code: 130,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "could not write cookie to db",
	}

	ErrorReadCookie = Error {
		Code: 131,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "could not read cookie from db",
	}
	ErrorWriteWebUser = Error {
		Code: 132,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "could not write web user to db",
	}

	ErrorReadWebUser = Error {
		Code: 133,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "could not read web user from db",
	}

	ErrorBadRequest = Error {
		Code: 140,
		CorrespondingHttpCode: http.StatusBadRequest,
		Message: "failed to get resource",
	}

	ErrorGormDocument = Error {
		Code: 160,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to convert reader to gorm document",
	}
	ErrorGormSelector = Error {
		Code: 161,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to select required elementsG",
	}

	ErrorApiProfileDbRead = Error {
		Code: 180,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to read api profile in db",
	}

	ErrorApiProfileDbWrite = Error {
		Code: 190,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to write api profile to db",
	}
)

// eagate errors 200-299
var (
	ErrorNoCookie = Error {
		Code: 200,
		CorrespondingHttpCode: http.StatusUnauthorized,
		Message: "no eagate cookie for user",
	}
	ErrorBadCookie = Error {
		Code: 201,
		CorrespondingHttpCode: http.StatusUnauthorized,
		Message: "bad or expired eagate cookie",
	}
	ErrorNoEaUser = Error {
		Code: 202,
		CorrespondingHttpCode: http.StatusUnauthorized,
		Message: "no eagate user found",
	}

	ErrorLoginFailed = Error {
		Code: 210,
		CorrespondingHttpCode: http.StatusUnauthorized,
		Message: "ea login failed",
	}

	ErrorMd5CharacterMapping = Error {
		Code: 240,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "md5 does not match character",
	}
)

// ddr errors 300-399
var (
	ErrorDdrSongIds = Error {
		Code: 300,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to get ddr song ids from eagate",
	}
	ErrorDdrSongData = Error {
		Code: 301,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to get ddr song data from eagate",
	}
	ErrorDdrSongDifficulties = Error {
		Code: 302,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to get ddr song difficulties from eagate",
	}
	ErrorDdrStats = Error {
		Code: 303,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to get ddr stats from eagate",
	}
	ErrorDdrPlayerInfo = Error {
		Code: 304,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to get ddr player info from eagate",
	}

	ErrorDdrSongIdsDbWrite = Error {
		Code: 330,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to write ddr song ids to db",
	}
	ErrorDdrSongDataDbWrite = Error {
		Code: 331,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to write ddr song data to db",
	}
	ErrorDdrSongDifficultiesDbWrite = Error {
		Code: 332,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to write ddr song difficulties to db",
	}
	ErrorDdrStatsDbWrite = Error {
		Code: 333,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to write ddr stats to db",
	}
	ErrorDdrPlayerInfoDbWrite = Error {
		Code: 334,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to write ddr player info to db",
	}

	ErrorDdrSongIdsDbRead = Error {
		Code: 360,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to get ddr song ids from db",
	}
	ErrorDdrSongDataDbRead = Error {
		Code: 361,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to get ddr song data from db",
	}
	ErrorDdrSongDifficultiesDbRead = Error {
		Code: 362,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to read ddr song difficulties from db",
	}
	ErrorDdrStatsDbRead = Error {
		Code: 363,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to read ddr stats from db",
	}
	ErrorDdrPlayerInfoDbRead = Error {
		Code: 364,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to read ddr player info from db",
	}
)

// drs errors 400-499
var (

	ErrorDrsSongData = Error {
		Code: 401,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to get drs song data from eagate",
	}
	ErrorDrsPlayerInfo = Error {
		Code: 404,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to get drs player info from eagate",
	}

	ErrorDrsSongDataDbWrite = Error {
		Code: 431,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to write drs song data to db",
	}
	ErrorDrsPlayerInfoDbWrite = Error {
		Code: 434,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to write drs player info to db",
	}

	ErrorDrsSongDataDbRead = Error {
		Code: 431,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to get drs song data from db",
	}
	ErrorDrsPlayerInfoDbRead = Error {
		Code: 464,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to get drs player info from db",
	}
)