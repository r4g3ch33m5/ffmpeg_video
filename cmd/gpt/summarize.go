package gpt

import (
	"context"
	"fmt"

	"github.com/r4g3ch33m5/ffmpeg_video/cmd/gpt/adapter"
	"github.com/urfave/cli/v3"
)

// summarizePrompt generates the prompt for summarization
func Prompt(subject, episode string) string {
	basePrompt := `
[you are summarizing %[1]s] 
[each episode contains multiple highlights]
[total duration do not exceed 30 seconds] 
[each highlights in format: <start of highlight timestamp in video>-<end of highlight timestamp in video>]
[answer in format: "Summary: [episode name] <summary of episode>\nHighlights: <list of highlight separate by newline>\n"]
Summarize the following content:
%[2]s
`
	return fmt.Sprintf(basePrompt, subject, episode)
}

// SummarizeCommand defines the command to summarize content using ChatGPT
var SummarizeCommand = &cli.Command{
	Name:    "summarize",
	Usage:   "Summarize content with specific instructions",
	Aliases: []string{"sum"},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "subject",
			Aliases:  []string{"s"},
			Usage:    "The subject being summarized",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "episode",
			Aliases:  []string{"e"},
			Usage:    "season and ep of content we want to summarize",
			Required: true,
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		subject := c.String("subject")
		arcs := c.String("episode")

		client := adapter.NewOpenAIClient(apiKey)
		fmt.Println(subject, arcs)
		result, err := client.ChatCompletion(Prompt(subject, arcs), "gpt-3.5-turbo", 1500)
		if err != nil {
			return fmt.Errorf("error during summarization: %v", err)
		}
		fmt.Println(result)
		return nil
	},
}
