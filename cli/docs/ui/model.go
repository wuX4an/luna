package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

// Model representa el estado de nuestra aplicación TUI.
type Model struct {
	Viewport    viewport.Model
	TextInput   textinput.Model
	Ready       bool   // Indica si el viewport ha sido inicializado
	InputActive bool   // Controla si el text input está actualmente enfocado
	Content     string // El contenido mostrado en el viewport
}

// InitialModel devuelve un nuevo modelo inicializado.
func InitialModel() Model {
	ti := textinput.New()
	ti.Placeholder = "..."
	ti.CharLimit = 256
	ti.Width = 40

	// Carga el contenido de index.txt embebido por defecto
	initialContent := Manual.Index

	return Model{
		TextInput:   ti,
		InputActive: false, // El input está activo inicialmente
		Content:     initialContent,
	}
}

// Init inicializa el modelo, devolviendo un comando para Bubble Tea.
func (m Model) Init() tea.Cmd {
	return textinput.Blink // Comando para que el cursor del text input parpadee
}
