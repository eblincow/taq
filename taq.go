package main

import (
	"fmt"
	"github.com/niemeyer/qml"
	"os"
    "math/rand"
    "time"
)



func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

var firstrand = random(1,9)



type Cell struct {
    Num int
    Val string
    Pic string
}


type BigCellList struct {
    list []Cell
}



func (b *BigCellList) GetCell(index int) *Cell {
    return &b.list[index+1]
}



func NewCell(num int) Cell {
    return Cell{Val: "o",Num:num,Pic:"blank.png"}
}

// Make the Cells for first time
	//var CellList []Cell
    var CellList = make([]*Cell,9)

var Cell1 = NewCell(1)
var Cell2 = NewCell(2)
var Cell3 = NewCell(3)
var Cell4 = NewCell(4)
var Cell5 = NewCell(5)
var Cell6 = NewCell(6)
var Cell7 = NewCell(7)
var Cell8 = NewCell(8)
var Cell9 = NewCell(9)
var counter = 0



func refresh() {
qml.Changed(&Cell1,&Cell1.Pic)
qml.Changed(&Cell2,&Cell2.Pic)
    qml.Changed(&Cell3,&Cell3.Pic)
    qml.Changed(&Cell4,&Cell4.Pic)
    qml.Changed(&Cell5,&Cell5.Pic)
qml.Changed(&Cell6,&Cell6.Pic)
qml.Changed(&Cell7,&Cell7.Pic)
qml.Changed(&Cell8,&Cell8.Pic)
qml.Changed(&Cell9,&Cell9.Pic)

}



func (c *Cell) Clicked() {
    c.Pic = "X.png"
    refresh()

    //qml.Changed(c, &c.Pic)
    // this worx -hallelujah

    
    fmt.Println("Im registering that you clicked: ")
    fmt.Printf("%v",c)
    counter++
    // get_next_move always returns 1
    Next_move_square := c 
    
    if counter > 0 {

    //ac := get_next_move(c.Num)

    fmt.Println("next move is calculated as: ")
    fmt.Println(Next_move_square.Num)
    Next_move_square.Pic = "O.png"
    
    qml.Changed(&Cell2,&Cell2.Pic)

} else {
    Next_move_square := CellList[firstrand-1]
    Next_move_square.Pic = "O.png"
    
}

    
    qml.Changed(&Cell2, &Cell2.Pic)
}




func get_next_move(cellnum int) int {

// check if two squares owned in a row
// check neighboring squares values
// - get neighboring squares
// - two match?
// - if not, random
// - if yes, hit the third

// squares to protect for each marked square
// these are all illegal things
one := []int{2,3,4,5,7,9} 
two := []int{1,3,5,8} 
three := []int{1,2,5,6,7,9} 
four := []int{1,5,6,7} 
five := []int{1,2,3,4,6,7,8,9}
six := []int{3,4,5,9} 
seven := []int{1,3,4,5,8,9} 
eight := []int{2,5,7,9} 
nine := []int{1,3,5,6,7,8}

list := []int{}


switch cellnum {
        case 1: list = one 
        case 2: list = two
        case 3: list = three 
        case 4: list = four 
        case 5: list = five  
        case 6: list = six 
        case 7: list = seven 
        case 8: list = eight 
        case 9: list = nine  
}
//s := make([]int,9)
// if a cell is marked
// return the neighboring square
nextmove := 0

for _, square := range list {
    if square != cellnum {
// if it isnt the same square he clicked
    ab := CellList[cellnum-1]    
if ab.Pic=="X.png" {
    // e.g. if the cell is marked
    nextmove := get_winning_square(square, cellnum)    
    return nextmove    
}
}

}
return nextmove
}


func get_winning_square(square1 int, square2 int) int {

    // given two square returning winning square
    ac := []int{square1, square2}

    switch ac[0] {
        case 1:
            switch ac[1]{
            case 2:
                return 3
            case 3:
                return 2
            case 7:
                return 4
            case 4:
                return 7
            case 5:
                return 9
            case 9:
                return 5
            
            }
        case 2:
            switch ac[1]{
            case 1:
                return 3
            case 3:
                return 1
            case 5:
                return 8
            case 8:
                return 5
            }
        case 3:
            switch ac[1]{
            case 1:
                return 2
            case 2:
                return 1
            case 5:
                return 7
            case 6:
                return 9
            case 9:
                return 6
            case 7:
                return 5
            }

        case 4:
            switch ac[1]{
            case 1:
                return 7
            case 6:
                return 5
            case 5:
                return 6
            case 7: 
                return 1
            }

        case 5:
            switch ac[1]{
            case 1:
                return 9
            case 9:
                return 1
            case 2:
                return 8
            case 8:
                return 2
            case 3:
                return 7
            case 7:
                return 3
            case 4:
                return 6
            case 6:
                return 4
            }

        case 6:
            switch ac[1]{
            case 3:
                return 9
            case 9:
                return 3
            case 5:
                return 4
            case 4:
                return 5
            }

        case 7:
            switch ac[1] {
            case 8:
                return 9
            case 9:
                return 8
            case 5:
                return 3 
            case 3:
                return 5
            case 1:
                return 4
            case 4:
                return 1

            }

        case 8:
            switch ac[1] {
            case 9:
                return 7
            case 7: 
            return 9

        case 2:
            return 5
        case 5:
            return 2
        }

    case 9:
        switch ac[1] {
        case 1:
            return 5
        case 5:
            return 1
        case 3:
            return 6
        case 6:
            return 3
        case 7:
            return 8
        case 8:
            return 7
        }
    }

return 0
}





func main() {

    CellList[0]=&Cell1
    CellList[1]=&Cell2
    CellList[2]=&Cell3
    CellList[3]=&Cell4
    CellList[4]=&Cell5
    CellList[5]=&Cell6
    CellList[6]=&Cell7
    CellList[7]=&Cell8
    CellList[8]=&Cell9


    
    
    
    if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
    
	
    
    
    qml.Init(nil)
	engine := qml.NewEngine()
    //Cells := BigCellList{list:CellList}
    // engine.Context().SetVar("Cells", &Cells)
    engine.Context().SetVar("cell1", &Cell1)
    engine.Context().SetVar("cell2",&Cell2)
    engine.Context().SetVar("cell3",&Cell3)
    engine.Context().SetVar("cell4",&Cell4)
    engine.Context().SetVar("cell5",&Cell5)
    engine.Context().SetVar("cell6",&Cell6)
    engine.Context().SetVar("cell7",&Cell7)
    engine.Context().SetVar("cell8",&Cell8)
    engine.Context().SetVar("cell9",&Cell9)
    //engine.Context().SetVar("clicked",&Clicked())

    
    
    component, err := engine.LoadFile("taq.qml")
	if err != nil {
		return err
	}
	window := component.CreateWindow(nil)
    
    window.Show()
	window.Wait()
	return nil
}


