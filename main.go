package main

import(
	"github.com/therecipe/qt/widgets"
	"os"
)
func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Hello World!!")
	window.SetMinimumSize2(200, 200)

	layout := widgets.NewQVBoxLayout()

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)

	label := widgets.NewQLabel2("Hello World",nil,0)
	layout.AddWidget(label,0,0)

	window.SetCentralWidget(widget)

	window.Show()

	widgets.QApplication_Exec()
}
