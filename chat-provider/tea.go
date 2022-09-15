package main

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/wish"
	"github.com/gliderlabs/ssh"
	"github.com/google/uuid"
	chat "github.com/jordan-rash/wasmcloud-chat/interface"
)

type (
	errMsg string
	strMsg string
	model  struct {
		msgs      *Stack
		users     *Stack
		content   string
		ready     bool
		chatport  viewport.Model
		userport  viewport.Model
		textinput textinput.Model
		err       error
	}
)

func (e errMsg) Error() string {
	return string(e)
}

func (s strMsg) Print() string {
	return string(s)
}

var (
	chatStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	localUserStyle = func() lipgloss.Style {
		return lipgloss.NewStyle()
	}()

	remoteUserStyle = func() lipgloss.Style {
		return lipgloss.NewStyle()
	}()

	usersStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		return chatStyle.Copy().BorderStyle(b)
	}()

	textStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		return lipgloss.NewStyle().
			Background(lipgloss.Color("ffff00")).
			BorderStyle(b).Padding(0, 1)
	}()
)

func GenGuid() string {
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return uuid
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	log.
		WithField("msg", msg).
		Print("[Update] processing message")

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			log.
				WithField("msg", msg).
				Debug("[tea.KeyMsg] processing message")

			if m.textinput.Value() == "/clear" {
				m.content = "/clear"
				m.textinput.SetValue("")
				return m, nil
			}

			n := chat.Msg{
				Owner: &localUser,
				Time:  time.Now().Format(time.RFC822),
				Value: m.textinput.Value(),
				Id:    uuid.NewString(),
			}

			m.textinput.SetValue("")
			return m.Update(n)

		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil

	case tea.WindowSizeMsg:
		mainHeight := lipgloss.Height(m.mainView())
		textHeight := lipgloss.Height(m.userView())
		verticalMarginHeight := mainHeight + textHeight

		if !m.ready {
			m.chatport = viewport.New(int(float64(msg.Width-3)*.75), msg.Height-verticalMarginHeight)
			m.userport = viewport.New(int(float64(msg.Width-3)*.25), msg.Height-verticalMarginHeight)

			m.chatport.YPosition = mainHeight
			m.userport.YPosition = textHeight

			m.chatport.SetContent(m.content)
			m.userport.SetContent(m.content)

			m.ready = true

		} else {
			m.chatport.Width = int(float64(msg.Width-3) * .75)
			m.userport.Width = int(float64(msg.Width-3) * .25)

			m.chatport.Height = msg.Height - verticalMarginHeight
			m.userport.Height = msg.Height - verticalMarginHeight
		}
	case chat.Msg:
		log.
			WithField("msg", msg).
			Print("processing message")
		switch msg.Owner.Name {
		case localUser.Name:
			sendDownLattice(msg, SEND_CHAT_MSG)
			m.msgs.Push(&msg)
		default:
			switch msg.Value {
			case "USERLEAVE":
				m.users.Delete(&msg)
			case "USERJOIN":
				m.users.Push(&msg)
			case "NOOP", "":
				return nil, nil
			default:
				m.msgs.Push(&msg)
			}
		}

		return m, nil
	}

	m.textinput, cmd = m.textinput.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}
	return lipgloss.JoinVertical(lipgloss.Left, m.mainView(), m.userView())
}

func (m model) mainView() string {

	switch m.content {
	case "/clear":
		m.content = ""
	default:
		m.content = ""
		for _, n := range m.msgs.Read() {
			m.content += localUserStyle.
				Foreground(lipgloss.Color(n.Owner.Color)).
				Render(n.Owner.Name) + "[" + n.Time + "]\n\t" + n.Value + "\n"
		}
	}

	m.chatport.SetContent(m.content)
	m.userport.SetContent(m.users.View())

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		chatStyle.Height(m.chatport.Height-2).Width(m.chatport.Width).Render(m.chatport.View()),
		usersStyle.Width(m.userport.Width).Render(m.userport.View()),
	)
}

func (m model) userView() string {
	return textStyle.
		Width(2 + m.chatport.Width + m.userport.Width).
		Render(m.textinput.View())
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	_, _, active := s.Pty()
	if !active {
		wish.Fatalln(s, "no active terminal, skipping")
		return nil, nil
	}

	ti := textinput.New()
	ti.Placeholder = "wasmCloud Chat"
	ti.Focus()
	ti.CharLimit = 200
	ti.Width = 100

	// user joins chat
	localUser.Name = s.User()
	localUser.Color = randomColor()
	msg := chat.Msg{
		Owner: &localUser,
		Value: "join",
		Id:    uuid.NewString(),
	}

	log.Debugf("Sending new user: %s", msg.Owner.Name)
	//sendDownLattice(msg, SEND_USER_UPDATE)

	m = model{
		content:   banner,
		textinput: ti,
		msgs:      NewStack(500),
		users:     NewStack(500),
		err:       nil,
	}

	m.users.Push(&chat.Msg{Owner: &localUser})
	return m, []tea.ProgramOption{tea.WithAltScreen()}
}
