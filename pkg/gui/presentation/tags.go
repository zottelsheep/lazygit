package presentation

import (
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/gui/presentation/icons"
	"github.com/jesseduffield/lazygit/pkg/gui/style"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	"github.com/jesseduffield/lazygit/pkg/i18n"
	"github.com/jesseduffield/lazygit/pkg/theme"
	"github.com/jesseduffield/lazygit/pkg/utils"
	"github.com/samber/lo"
)

func GetTagListDisplayStrings(
	tags []*models.Tag,
	getItemOperation func(item types.HasUrn) types.ItemOperation,
	diffName string,
	tr *i18n.TranslationSet,
) [][]string {
	return lo.Map(tags, func(tag *models.Tag, _ int) []string {
		diffed := tag.Name == diffName
		return getTagDisplayStrings(tag, getItemOperation(tag), diffed, tr)
	})
}

// getTagDisplayStrings returns the display string of branch
func getTagDisplayStrings(t *models.Tag, itemOperation types.ItemOperation, diffed bool, tr *i18n.TranslationSet) []string {
	textStyle := theme.DefaultTextColor
	if diffed {
		textStyle = theme.DiffTerminalColor
	}
	res := make([]string, 0, 2)
	if icons.IsIconEnabled() {
		res = append(res, textStyle.Sprint(icons.IconForTag(t)))
	}
	descriptionColor := style.FgYellow
	descriptionStr := descriptionColor.Sprint(t.Description())
	itemOperationStr := itemOperationToString(itemOperation, tr)
	if itemOperationStr != "" {
		descriptionStr = style.FgCyan.Sprint(itemOperationStr+" "+utils.Loader()) + " " + descriptionStr
	}
	res = append(res, textStyle.Sprint(t.Name), descriptionStr)
	return res
}
