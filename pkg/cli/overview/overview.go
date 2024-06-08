package overview

import (
	"fmt"

	"github.com/Pixium/MallData/pkg/db"
	"github.com/Pixium/MallData/pkg/types"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type plotDelegate struct {
	plot *types.Plot
}

func (d plotDelegate) Title() string {
	return fmt.Sprintf("%s by %s", d.plot.Name, d.plot.OwnerMinecraftName)
}

func (d plotDelegate) Description() string {
	return fmt.Sprintf("%d note(s)", len(d.plot.GetNotes()))
}

func (d plotDelegate) FilterValue() string {
	return d.Title()
}

type model struct {
	plots *[]types.Plot
	list  list.Model
}

func NewModel() model {
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Select plot to view"
	m := model{list: l}
	return m
}

func (m model) Init() tea.Cmd {
	return getPlots
}

type successMessage *[]types.Plot
type errorMessage error

func getPlots() tea.Msg {
	plots, err := db.GetPlots()
	if err != nil {
		return errorMessage(err)
	}

	return successMessage(&plots)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case successMessage:
		var listItems []list.Item
		for _, plot := range *msg {
			listItems = append(listItems, plotDelegate{
				plot: &plot,
			})
		}

		m.list.SetItems(listItems)
		m.plots = msg

	case errorMessage:
		log.Fatal(msg)
		return m, tea.Quit

	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	if m.list.Title != "" {
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	// historyStyle := lipgloss.NewStyle().
	// 	Align(lipgloss.Left).
	// 	Foreground(lipgloss.Color("#FAFAFA")).
	// 	Background(lipgloss.Color("#FFFFFF")).
	// 	Margin(1).
	// 	Padding(1, 2).
	// 	Width(10).
	// 	Height(5)

	// out = []string{
	// 	historyStyle.Render("test"),
	// }

	// log.Info("View")

	// return lipgloss.JoinHorizontal(lipgloss.Top, m.list.View())

	return docStyle.Render(m.list.View())
}
