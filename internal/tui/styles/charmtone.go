package styles

import (
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/charmbracelet/x/exp/charmtone"
)

func NewCharmtoneTheme() *Theme {
	t := &Theme{
		Name:   "zero",
		IsDark: true,

		// Blue/Teal color scheme for Zero
		Primary:   ParseHex("#0066CC"), // Primary blue
		Secondary: ParseHex("#0099CC"), // Light blue
		Tertiary:  ParseHex("#0088AA"), // Teal
		Accent:    ParseHex("#00CCCC"), // Bright teal

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
		Border:      ParseHex("#334155"), // Slate gray
		BorderFocus: ParseHex("#0066CC"), // Primary blue

		// Status
		Success: ParseHex("#00AA88"), // Teal green
		Error:   ParseHex("#CC4400"), // Red-orange
		Warning: ParseHex("#CCAA00"), // Golden yellow
		Info:    ParseHex("#0099CC"), // Light blue

		// Colors
		White: charmtone.Butter,

		BlueLight: ParseHex("#66B3FF"), // Light blue
		Blue:      ParseHex("#0066CC"), // Primary blue

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
