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
    // Num 1-9
    Num int
    Val string
    // Pic "O.png", "X.png", "blank.png"
    Pic string
}






func NewCell(num int) Cell {
    return Cell{Val: "o",Num:num,Pic:"blank.png"}
}




// Make the Cells for first time
// and CellList to contain them

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
var firstmove = true



func refresh() {
// refresh all the cells
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
    // check if anyone won?
    ch := has_anyone_won()
    if ch {
        fmt.Println("You won!")
        return
    }
    
    // If not, make your move 
    
    last_move := c.Num

    
    fmt.Println("Im registering that you clicked: ")
    fmt.Printf("%v\n",c)
    
    //initialize Next_move_square

    // get_next_move always returns 1
    // if this is not firstmove 
    if firstmove==false {
    
    // get the next move 
    Next_move_square := get_next_move(last_move)
    
    fmt.Println("\nnext move is calculated as: ")
    fmt.Println(Next_move_square)
    // change that square 
    Next_cell := CellList[Next_move_square-1] 
    Next_cell.Pic = "O.png"
    refresh()    
    ch := has_anyone_won()
    if ch {
        fmt.Println("You won!")
        return
    }

} else {
    // if it is the first move 
    firstmove = false 
    // choose random value for first move
    Next_cell := CellList[firstrand-1]
    Next_cell.Pic = "O.png"
    refresh() 
}

    
}


type WinCom struct {
    // winning combinations
    comb []int 
}


func has_anyone_won() bool {

    wincomlist := make([]*WinCom,9,9)
    // winning combinations
    a1 := &WinCom{[]int{1,2,3}}
    a2 := &WinCom{[]int{1,4,7}}
    a3 := &WinCom{[]int{1,5,9}}
    a4 := &WinCom{[]int{2,5,8}}
    a5 := &WinCom{[]int{3,5,7}}
    a6 := &WinCom{[]int{3,6,9}}
    a7 := &WinCom{[]int{4,5,6}}
    a8 := &WinCom{[]int{7,8,9}}
    wincomlist[0] = a1
    wincomlist[1] = a2
    wincomlist[2] = a3
    wincomlist[3] = a4
    wincomlist[4] = a5
    wincomlist[5] = a6
    wincomlist[6] = a7
    wincomlist[7] = a8
    os := []int{}
    xs := []int{}
    for _, cel := range CellList {
        if cel.Pic == "O.png" {
            os = append(os,cel.Num)
    }
    if cel.Pic == "X.png" {
        xs = append(xs,cel.Num)
    }
}

fmt.Println("The Os are: ")

fmt.Println(os)
fmt.Println("\nThe Xs are: ")

fmt.Println(xs)

return false
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
// if it isnt the same square already clicked
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


func checkcell(cellnum int) int {
    fmt.Println("\n\nCellnum is : \n")
    fmt.Println(cellnum)
    ac2 := CellList[cellnum-1]
     if ac2.Pic == "blank.png" {
         return cellnum 
         }
    return 0 
     }



func get_winning_square(square1 int, square2 int) int {

    // given two square returning winning square
    ac := []int{square1, square2}

    switch ac[0] {
        case 1:
            switch ac[1]{
            case 2:
                ab := checkcell(3); if ab != 0 { return ab }
            case 3:
                ab := checkcell(2); if ab != 0 { return ab }

        case 7:
            ab := checkcell(4); if ab != 0 { return ab }

            case 4:
                ab := checkcell(7); if ab != 0 { return ab }

            case 5:
                ab := checkcell(9); if ab != 0 { return ab }
 
            
            case 9:
                ab := checkcell(5); if ab != 0 { return ab }
 
            }
        case 2:
            switch ac[1]{
            case 1:
                ab:= checkcell(3); if ab != 0 { return ab }
            case 3:
                ab:= checkcell(1); if ab != 0 { return ab }
            case 5:
                ab:= checkcell(8); if ab != 0 { return ab }
            case 8:
                ab:= checkcell(5); if ab != 0 { return ab }
            }
        case 3:
            switch ac[1]{
            case 1:
                ab:= checkcell(2); if ab != 0 { return ab }
            case 2:
                ab:= checkcell(1); if ab != 0 { return ab }
            case 5:
                ab:= checkcell(7); if ab != 0 { return ab }
            case 6:
                ab:= checkcell(9); if ab != 0 { return ab }
            case 9:
                ab:= checkcell(6); if ab != 0 { return ab }
            case 7:
                ab:= checkcell(5); if ab != 0 { return ab }
            }

        case 4:
            switch ac[1]{
            case 1:
                ab:= checkcell(7); if ab != 0 { return ab }
            case 6:
                ab:= checkcell(5); if ab != 0 { return ab }
            case 5:
                ab:= checkcell(6); if ab != 0 { return ab }
            case 7: 
                ab:= checkcell(1); if ab != 0 { return ab }
            }

        case 5:
            switch ac[1]{
            case 1:
                ab:= checkcell(9); if ab != 0 { return ab }
            case 9:
                ab:= checkcell(1); if ab != 0 { return ab }
            case 2:
                ab:= checkcell(8); if ab != 0 { return ab }
            case 8:
                ab:= checkcell(2); if ab != 0 { return ab }
            case 3:
                ab:= checkcell(7); if ab != 0 { return ab }
            case 7:
                ab:= checkcell(3); if ab != 0 { return ab }
            case 4:
                ab:= checkcell(6); if ab != 0 { return ab }
            case 6:
                ab:= checkcell(4); if ab != 0 { return ab }
            }

        case 6:
            switch ac[1]{
            case 3:
                ab:= checkcell(9); if ab != 0 { return ab }
            case 9:
                ab:= checkcell(3); if ab != 0 { return ab }
            case 5:
                ab:= checkcell(4); if ab != 0 { return ab }
            case 4:
                ab:= checkcell(5); if ab != 0 { return ab }
            }

        case 7:
            switch ac[1] {
            case 8:
                ab:= checkcell(9); if ab != 0 { return ab }
            case 9:
                ab:= checkcell(8); if ab != 0 { return ab }
            case 5:
                ab:= checkcell(3); if ab != 0 { return ab } 
            case 3:
                ab:= checkcell(5); if ab != 0 { return ab }
            case 1:
                ab:= checkcell(4); if ab != 0 { return ab }
            case 4:
                ab:= checkcell(1); if ab != 0 { return ab }

            }

        case 8:
            switch ac[1] {
            case 9:
                ab:= checkcell(7); if ab != 0 { return ab }
            case 7: 
            ab:= checkcell(9); if ab != 0 { return ab }

        case 2:
            ab:= checkcell(5); if ab != 0 { return ab }
        case 5:
            ab:= checkcell(2); if ab != 0 { return ab }
        }

    case 9:
        switch ac[1] {
        case 1:
            
            ab:= checkcell(5); if ab != 0 { return ab }
        case 5:
            ab:= checkcell(1); if ab != 0 { return ab }
        case 3:
            ab:= checkcell(6); if ab != 0 { return ab }
        case 6:
            ab:= checkcell(3); if ab != 0 { return ab }
        case 7:
            
            ab:= checkcell(8); if ab != 0 { return ab }
        case 8:
            ab:= checkcell(7); if ab != 0 { return ab }
        }
    }

return 1
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
    // To resize this we need to pass height/width vars 

    //engine.Context().SetVar("holder",Ajcc)
    //engine.Context().SetVar("holderwidth",Ajcc.Width) 

    component, err := engine.LoadFile("taq.qml")
	if err != nil {
		return err
	}
	window := component.CreateWindow(nil)
 	
   
    window.Show()
// 	
//     Ajcc.Height = 360
//     Ajcc.Width = 360
//     fmt.Println(Ajcc.Width)
// 
//     qml.Changed(&Ajcc,&Ajcc.Height)
//     qml.Changed(&Ajcc,&Ajcc.Width) 
// 
// 
    
    
    
    
    window.Wait()
	return nil
}


