import QtQuick 2.0
import QtQuick.Controls 1.1

Rectangle {
    id: page
    width: 500; height: 200
    color: "lightgray"

    Text {
        id: helloText
        text: person.name
        y: 30
        anchors.horizontalCenter: page.horizontalCenter
        font.pointSize: 24; font.bold: true
    }

    Button {
        id: btn
        text: "click"
        onClicked: {
            person.name = "gopher"
            helloText.text = person.name
        }
    }
}
