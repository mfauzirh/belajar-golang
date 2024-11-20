package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"eko", "kurniawan", "Khannedy"})
	_ = writer.Write([]string{"eko2", "kurniawan", "Khannedy"})
	_ = writer.Write([]string{"eko3", "kurniawan", "Khannedy"})

	writer.Flush()
}
