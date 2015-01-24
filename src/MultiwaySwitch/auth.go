package MultiwaySwitch

import (
	"bytes"
	"code.google.com/p/log4go"
	"encoding/json"
	"github.com/Unknwon/goconfig"
	js "github.com/bitly/go-simplejson"
	"io"
	"net"
	"strconv"
)

type PackageType byte

const (
	AUTHORIZATION PackageType = iota
)

func serverTcp() {
	bind_addr := configServer("bind", FATAL)
	listener, err := net.Listen("tcp", bind_addr)
	if err != nil {
		logger.Critical("CANNOT bind address:", bind_addr, " -- ", err.Error())
		panic("CANNOT bind address")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("Error Accept: ", err.Error())
		}
		logger.Info("Accepted the Connection :", conn.RemoteAddr())
		go authServer(&conn)
	}
}

func authServer(conn *net.Conn) {
	defer conn.Close()
	receiveBufferSize, err := strconv.Atoi(configServer("recevie_buffer_size", FATAL))
	if err != nil {
		logger.Debug("no recevie_buffer_size field in config file")
	}
	receiveBufferSize = 1024
	buffer := make([]byte, receiveBufferSize)
	data_received := bytes.NewBuffer([]byte{})
	for {
		n_buffer, err := conn.Read(buffer)
		switch err {
		case nil:
			data_received.Write(buffer)
		case io.EOF:
			data_received.Write(buffer[:n_buffer])
			selectPackgeType(*data_received, conn)
			data_received.Reset()
		default:
			logger.Error(err.Error())
			break
		}
	}
}

func parseJsonFromBuffer(data bytes.Buffer) *js.Json {
	json, err := js.NewFromReader(bytes.NewReader(data.Bytes()))
	if err != nil {
		logger.Error("parsing json failed:", err.Error())
	}
	return json
}

func selectPackgeType(data bytes.Buffer, conn *net.Conn) {
	jsonData := parseJsonFromBuffer(data)
	droneId := jsonData.Get("drone_id").MustInt64()
	packageFormat := PackageType(jsonData.Get("type").MustInt())
	switch packageFormat {
	case AUTHORIZATION:
		response := authorization(droneId, jsonData)
		conn.Write(response)
	default:
		logger.Debug("package type not recognized")
	}
}

func authorization(id int64, jsonData *js.Json) []byte {
	secretKey := jsonData.Get("secret_key").MustString()
	if CheckSecretKey(id, secretKey) {
		token, err := FlushToken(id)
		if err != nil {
			logger.Warn(err.Error())
		}

	}
}
