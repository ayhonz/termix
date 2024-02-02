package buffer

type PieceTable struct {
	origin []rune
	add    []rune
	pieces []Piece
}

type Piece struct {
	pieceType string // "origin" or "add"
	start     int
	end       int
}

func NewPieceTable(origin []rune) *PieceTable {
	return &PieceTable{
		origin: origin,
	}
}

func (pt *PieceTable) Insert(r string, index int) {
	pt.add = []rune(r)

	if index == 0 {
		pt.pieces = append(pt.pieces, Piece{
			pieceType: "add",
			start:     index,
			end:       len(pt.add),
		}, Piece{
			pieceType: "origin",
			start:     0,
			end:       len(pt.origin),
		})
		return
	} else if index == len(pt.origin) {
		pt.pieces = append(pt.pieces, Piece{
			pieceType: "origin",
			start:     0,
			end:       len(pt.origin),
		}, Piece{
			pieceType: "add",
			start:     index,
			end:       index + len(pt.add),
		})
		return
	} else {
		addPiece := Piece{
			pieceType: "add",
			start:     index,
			end:       index + len(pt.add),
		}
		startPiece := Piece{
			pieceType: "origin",
			start:     0,
			end:       index,
		}
		endPiece := Piece{
			pieceType: "origin",
			start:     index,
			end:       len(pt.origin),
		}
		pt.pieces = append(pt.pieces, startPiece, addPiece, endPiece)

	}

}

func (pt *PieceTable) String() string {
	var s []rune
	for _, p := range pt.pieces {
		if p.pieceType == "origin" {
			e := pt.origin[p.start:p.end]
			s = append(s, e...)
		}
		if p.pieceType == "add" {
			s = append(s, pt.add...)
		}
	}

	return string(s)
}
