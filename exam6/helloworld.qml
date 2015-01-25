import QtQuick 2.0
import GoExtensions 1.0

Rectangle {
    id: page
    width: 500; height: 200
    color: "lightgray"
    
    GoRect {
        x: 60; y: 60; width: 100; height: 100
		
		/* NumberAnimation을 순차적으로 계속해서 실행 */
        SequentialAnimation on x {
            loops: Animation.Infinite
            NumberAnimation { from: 60; to: 320; duration: 4000; easing.type: Easing.InOutQuad }
            NumberAnimation { from: 320; to: 60; duration: 4000; easing.type:Easing.InOutQuad }
        }
    }
}

