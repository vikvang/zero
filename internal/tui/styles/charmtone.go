package styles

import (
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/charmbracelet/x/exp/charmtone"
)

func NewCharmtoneTheme() *Theme {
	t := &Theme{
		Name:   "charmtone",
		IsDark: true,

		Primary:   charmtone.Charple,
		Secondary: charmtone.Dolly,
		Tertiary:  charmtone.Bok,
		Accent:    charmtone.Zest,

		// Backgrounds
		BgBase:        charmtone.Pepper,
		BgBaseLighter: charmtone.BBQ,
		BgSubtle:      charmtone.Charcoal,
		BgOverlay:     charmtone.Iron,

		// Foregrounds
		FgBase:      charmtone.Ash,
		FgMuted:     charmtone.Squid,
		FgHalfMuted: charmtone.Smoke,
		FgSubtle:    charmtone.Oyster,
		FgSelected:  charmtone.Salt,

		// Borders
		Border:      charmtone.Charcoal,
		BorderFocus: charmtone.Charple,

		// Status
		Success: charmtone.Guac,
		Error:   charmtone.Sriracha,
		Warning: charmtone.Zest,
		Info:    charmtone.Malibu,

		// Colors
		White: charmtone.Butter,

		BlueLight: charmtone.Sardine,
		Blue:      charmtone.Malibu,

		Yellow: charmtone.Mustard,
		Citron: charmtone.Citron,

		Green:      charmtone.Julep,
		GreenDark:  charmtone.Guac,
		GreenLight: charmtone.Bok,

		Red:      charmtone.Coral,
		RedDark:  charmtone.Sriracha,
		RedLight: charmtone.Salmon,
		Cherry:   charmtone.Cherry,
	}

	// Text selection.
	t.TextSelection = lipgloss.NewStyle().Foreground(charmtone.Salt).Background(charmtone.Charple)

	// LSP and MCP status.
	t.ItemOfflineIcon = lipgloss.NewStyle().Foreground(charmtone.Squid).SetString("●")
	t.ItemBusyIcon = t.ItemOfflineIcon.Foreground(charmtone.Citron)
	t.ItemErrorIcon = t.ItemOfflineIcon.Foreground(charmtone.Coral)
	t.ItemOnlineIcon = t.ItemOfflineIcon.Foreground(charmtone.Guac)

	t.YoloIconFocused = lipgloss.NewStyle().Foreground(charmtone.Oyster).Background(charmtone.Citron).Bold(true).SetString(" ! ")
	t.YoloIconBlurred = t.YoloIconFocused.Foreground(charmtone.Pepper).Background(charmtone.Squid)
	t.YoloDotsFocused = lipgloss.NewStyle().Foreground(charmtone.Zest).SetString(":::")
	t.YoloDotsBlurred = t.YoloDotsFocused.Foreground(charmtone.Squid)

	return t
}
