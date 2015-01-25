import QtQuick 2.0
import GoExtensions 1.0

Rectangle {
    id: page
    width: 500; height: 200
    color: "lightgray"

    Person {
        id: person
        name : "naeun"
    }

    Text {
        id: helloText
        text: "Hello " + person.name
        y: 30
        anchors.horizontalCenter: page.horizontalCenter
        font.pointSize: 24; font.bold: true
    }
}
