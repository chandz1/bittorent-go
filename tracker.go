package main

import (
	"crypto/sha1"

	"github.com/marksamman/bencode"
)

type tracker struct {
    info_hash   string  // A 20 byte hash from the torrent file
    peer_id     string  // A 20 byte hash randomly generated peer_id
    port        int     // Specifying the port the client is listening on
    uploaded    int     // The total amount uploaded
    downloaded  int     // The total amount downloaded
    left        int     // The number of bytes the client has to download
    compact     int     // Setting to 0 or 1 indicating a compact response
    ip          int     // OPTIONAL: Specifying the client's true IP address
    numwant     int     // OPTIONAL: Number of peers client wants to receive from
    key         int     // OPTIONAL: An additional identification for client
    event       string  // OPTIONAL: Event set (started, stopped, completed) if specified
    trackerId   int     // OPTIONAL: If pervious announce contained a trackerid it will be set
}


func getHash(decodedFile map[string]interface{}) [20]byte {
    info_Encoded := bencode.Encode(decodedFile)

    sha1_decoded := sha1.Sum(info_Encoded)

    return sha1_decoded
}

