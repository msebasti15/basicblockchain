package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"

	"basicblockchain/internal/model"
)

func RunServer(blockChain *model.BlockChain) error {
	serverMux := makeMuxRouter(blockChain)
	httpAddr := os.Getenv("PORT")

	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        serverMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func makeMuxRouter(blockChain *model.BlockChain) http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain(blockChain)).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlock(blockChain)).Methods("POST")
	return muxRouter
}

func handleGetBlockchain(blockChain *model.BlockChain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bytes, err := json.MarshalIndent(blockChain, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		io.WriteString(w, string(bytes))
	}
}

func handleWriteBlock(blockChain *model.BlockChain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var m Message

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&m); err != nil {
			respondWithJSON(w, r, http.StatusBadRequest, r.Body)
			return
		}
		defer r.Body.Close()

		newBlock := model.NewBlock(m.Data, blockChain.Blocks[len(blockChain.Blocks)-1])

		if model.IsBlockValid(newBlock, blockChain.Blocks[len(blockChain.Blocks)-1]) {
			newBlockchain := append(blockChain.Blocks, newBlock)
			blockChain.ReplaceChain(newBlockchain)
			spew.Dump(blockChain)
		}

		respondWithJSON(w, r, http.StatusCreated, newBlock)

	}

}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}
