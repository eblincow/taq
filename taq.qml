import QtQuick 2.0

Item {

width: 25
height: 24



Rectangle {
    width: 360
    height: 360
    color: "white"
    

Column {

spacing: 2
y: 10
x: 10

    Row {
    y: 10
    x: 5
    spacing:1

    Rectangle {
    width: 110
    height: 110
    color: "white"
    border.color: "black"
    border.width: 2
    
    Image {
            //id: cell1pic 
             width: 85
             height: 85
             
             anchors.centerIn: parent
             //property string imageName: cell1.getpic()

             source: cell1.pic
             fillMode: Image.PreserveAspectFit
         }
    
    MouseArea {
         anchors.fill: parent
         onClicked: cell1.clicked()
     }
    
    }

   
   
    Rectangle {
    width: 110
    height: 110
    color: "white"
    border.color: "black"
    border.width: 2
    
    Image {
            //id: cell1pic 
             width: 85
             height: 85
             
             anchors.centerIn: parent
             //property string imageName: cell1.getpic()

             source: cell2.pic
             fillMode: Image.PreserveAspectFit
         }
    
    MouseArea {
         anchors.fill: parent
         onClicked: cell2.clicked()
     }
}
   
   
  Rectangle {
    width: 110
    height: 110
    color: "white"
    border.color: "black"
    border.width: 2
    
    Image {
            //id: cell1pic 
             width: 85
             height: 85
             
             anchors.centerIn: parent
             //property string imageName: cell1.getpic()

             source: cell3.pic
             fillMode: Image.PreserveAspectFit
         }
    
    MouseArea {
         anchors.fill: parent
         onClicked: cell3.clicked()
     }
 
 }

   
   
   
   
 }

    Row {
    y: 10
    x: 5
    spacing:1

    Rectangle {
    width: 110
    height: 110
    color: "white"
    border.color: "black"
    border.width: 2
    
    Image {
            //id: cell1pic 
             width: 85
             height: 85
             
             anchors.centerIn: parent
             //property string imageName: cell1.getpic()

             source: cell4.pic
             fillMode: Image.PreserveAspectFit
         }
    
    MouseArea {
         anchors.fill: parent
         onClicked: cell4.clicked()
     }
    
    }

   
   
    Rectangle {
    width: 110
    height: 110
    color: "white"
    border.color: "black"
    border.width: 2
    
    Image {
            //id: cell1pic 
             width: 85
             height: 85
             
             anchors.centerIn: parent
             //property string imageName: cell1.getpic()

             source: cell5.pic
             fillMode: Image.PreserveAspectFit
         }
    
    MouseArea {
         anchors.fill: parent
         onClicked: cell5.clicked()
     }

  } 
   
  Rectangle {
    width: 110
    height: 110
    color: "white"
    border.color: "black"
    border.width: 2
    
    Image {
            //id: cell1pic 
             width: 85
             height: 85
             
             anchors.centerIn: parent
             //property string imageName: cell1.getpic()

             source: cell6.pic
             fillMode: Image.PreserveAspectFit
         }
    
    MouseArea {
         anchors.fill: parent
         onClicked: cell6.clicked()
     }
  
   
   
   
   
 }
}
 
    Row {
    y: 10
    x: 5
    spacing:1

    Rectangle {
    width: 110
    height: 110
    color: "white"
    border.color: "black"
    border.width: 2
    
    Image {
            //id: cell1pic 
             width: 85
             height: 85
             
             anchors.centerIn: parent
             //property string imageName: cell1.getpic()

             source: cell7.pic
             fillMode: Image.PreserveAspectFit
         }
    
    MouseArea {
         anchors.fill: parent
         onClicked: cell7.clicked()
     }
    
    }

   
   
    Rectangle {
    width: 110
    height: 110
    color: "white"
    border.color: "black"
    border.width: 2
    
    Image {
            //id: cell1pic 
             width: 85
             height: 85
             
             anchors.centerIn: parent
             //property string imageName: cell1.getpic()

             source: cell8.pic
             fillMode: Image.PreserveAspectFit
         }
    
    MouseArea {
         anchors.fill: parent
         onClicked: cell8.clicked()
     }

  }

   
  Rectangle {
    width: 110
    height: 110
    color: "white"
    border.color: "black"
    border.width: 2
    
    Image {
            //id: cell1pic 
             width: 85
             height: 85
             
             anchors.centerIn: parent
             //property string imageName: cell1.getpic()

             source: cell9.pic
             fillMode: Image.PreserveAspectFit
         }
    
    MouseArea {
         anchors.fill: parent
         onClicked: cell9.clicked()
     }
  
   
   
   
   
 }
}
}
}
}
