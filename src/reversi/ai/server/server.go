package main

import(
  "encoding/json"
  "net/http"
  "net/url"
  "fmt"
  "reversi/ai"
  "reversi/game/board"
  "reversi/game/cell"
  "time"
  "strconv"
  "io"
)

func reversiAiBridge(w http.ResponseWriter, r *http.Request){

  query := r.URL.Query()

  var requestBoard board.Board
  var requestCellType int
  var resolvedCellChange cell.Cell
  var serializedCellChange []byte
  var err error

  if requestCellType, err = getCellTypeFromQuery(query); (err != nil) {
    returnBadRequestResponse(w, "Invalid Cell Type")
    return
  }

  if requestBoard, err = getBoardFromRequestBody(r.Body); (err != nil) {
    returnBadRequestResponse(w, "Invalid Board")
    return
  }

  if resolvedCellChange, err = ai.GetBestCellChangeInTime(requestBoard, uint8(requestCellType), time.Second); (err != nil){
    returnBadRequestResponse(w, "Unexpected Error occured")
    return
  }

  if serializedCellChange, err = json.Marshal(resolvedCellChange); (err != nil){
    returnBadRequestResponse(w, "Serialization Error occured")
    return
  }

  fmt.Fprintf(w, "%s", serializedCellChange)

}

func returnBadRequestResponse(w http.ResponseWriter, message string) {
 	w.WriteHeader(http.StatusBadRequest)
 	w.Write([]byte(message))
 }

func getBoardFromRequestBody(body io.ReadCloser) (board.Board, error){
  defer body.Close()
  decoder := json.NewDecoder(body)
  var rBoard board.Board
  err := decoder.Decode(&rBoard)
  return rBoard, err
}

func getCellTypeFromQuery(query url.Values) (int, error) {
	queryCellType := query.Get("type")
	return strconv.Atoi(queryCellType)
}

func main() {
	http.HandleFunc("/", reversiAiBridge)
	http.ListenAndServe(":8080", nil)
}
