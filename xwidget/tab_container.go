package xwidget

import (
	"fyne.io/fyne/v2"
	fyneContainer "fyne.io/fyne/v2/container"
	"sort"
)

type TabContainer struct {
	fyneContainer.AppTabs
	selectedElementIndex int
	tabsContent          map[string][]fyne.CanvasObject
	onElementSelected    func(element *fyne.CanvasObject)
	onElementUnselected  func(element *fyne.CanvasObject)
}

func NewTabContainer(
	tabsData map[string][]fyne.CanvasObject,
	onElementSelected func(element *fyne.CanvasObject),
	onElementUnselected func(element *fyne.CanvasObject)) *TabContainer {
	container := newTabContainer(tabsData, onElementSelected, onElementUnselected)
	container.ExtendBaseWidget(container)
	return container
}

// Should be used for dialog creation by any xwidget that embed this xwidget so it can properly extend fyne's BaseWidget
func newTabContainer(
	tabsData map[string][]fyne.CanvasObject,
	onElementSelected func(element *fyne.CanvasObject),
	onElementUnselected func(element *fyne.CanvasObject)) *TabContainer {
	container := &TabContainer{
		selectedElementIndex: 0,
		tabsContent:          tabsData,
		onElementSelected:    onElementSelected,
		onElementUnselected:  onElementUnselected,
	}
	container.AppTabs.Items = getTabsFromData(tabsData)
	container.selectElement(0)
	container.SelectTabIndex(0)
	return container
}

func getTabsFromData(tabsData map[string][]fyne.CanvasObject) []*fyneContainer.TabItem {
	var tabs []*fyneContainer.TabItem
	sortedTabsNames := getAlphabeticallySortedTabsNames(tabsData)
	for _, tabName := range sortedTabsNames {
		formItem := fyneContainer.NewTabItem(tabName, fyneContainer.NewVBox(tabsData[tabName]...))
		tabs = append(tabs, formItem)
	}
	return tabs
}

func getAlphabeticallySortedTabsNames(tabsData map[string][]fyne.CanvasObject) []string {
	sortedNames := make([]string, 0, len(tabsData))
	for tabName := range tabsData {
		sortedNames = append(sortedNames, tabName)
	}
	sort.Strings(sortedNames)
	return sortedNames
}

func (container *TabContainer) selectElement(num int) {
	if container.currentTabHasElements() {
		selectedElement := container.selectedElement()
		container.onElementUnselected(&selectedElement)
		selectedElement.Refresh()
		container.selectedElementIndex = num
		selectedElement = container.selectedElement()
		container.onElementSelected(&selectedElement)
		container.selectedElement().Refresh()
	}
}

func (container *TabContainer) currentTabHasElements() bool {
	currentTab := container.CurrentTab()
	if currentTab != nil {
		return len(container.tabsContent[currentTab.Text]) > 0
	}
	return false
}

func (container *TabContainer) selectedElement() fyne.CanvasObject {
	currentTabText := container.CurrentTab().Text
	return container.tabsContent[currentTabText][container.selectedElementIndex]
}

func (container *TabContainer) SelectNextTab() {
	currentTabIndex := container.CurrentTabIndex()
	if currentTabIndex >= len(container.Items())-1 {
		container.setTabTo(0)
	} else {
		container.setTabTo(currentTabIndex + 1)
	}
}

func (container *TabContainer) SelectPreviousTab() {
	currentTabIndex := container.CurrentTabIndex()
	if currentTabIndex == 0 {
		container.setTabTo(len(container.Items()) - 1)
	} else {
		container.setTabTo(currentTabIndex - 1)
	}
}

func (container *TabContainer) setTabTo(index int) {
	container.SelectTabIndex(index)
	container.selectElement(0)
}

func (container *TabContainer) Items() []*fyneContainer.TabItem {
	return container.AppTabs.Items
}
