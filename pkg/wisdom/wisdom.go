package wisdom

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/tabwriter"
	"unicode"

	"github.com/google/uuid"

	"code.dwrz.net/src/pkg/log"
	"code.dwrz.net/src/pkg/store"
	"code.dwrz.net/src/pkg/text"
	"code.dwrz.net/src/pkg/wisdom/quote"
)

const coll = "data"

type Parameters struct {
	Log  *log.Logger
	Path string
	Wrap int
}

type Wisdom struct {
	log   *log.Logger
	store *store.Store
	wrap  int
}

func New(p Parameters) (*Wisdom, error) {
	store, err := store.New(store.Parameters{
		Log:  p.Log,
		Path: p.Path,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create store: %v", err)
	}

	if err := store.NewCollection(coll); err != nil {
		return nil, fmt.Errorf("failed to create collection: %v", err)
	}

	return &Wisdom{
		log:   p.Log,
		store: store,
		wrap:  p.Wrap,
	}, nil
}

func (w *Wisdom) Command(args []string) error {
	// Show a random quote by default.
	if len(args) == 0 {
		d, err := w.store.Collection(coll).Random()
		if err != nil {
			return fmt.Errorf("failed to load quote: %v", err)
		}

		var q quote.Quote
		if err := d.Unmarshal(&q); err != nil {
			return fmt.Errorf(
				"failed to unmarshal document: %v", err,
			)
		}

		fmt.Println(q.Render(w.wrap))

		return nil
	}

	command, rest := args[0], args[1:]
	switch command {
	case "add":
		if err := w.add(); err != nil {
			return fmt.Errorf("failed to add quote: %v", err)
		}

	case "edit":
		if err := w.edit(rest); err != nil {
			return fmt.Errorf("failed to edit quote: %v", err)
		}

	case "show":
		if err := w.show(rest); err != nil {
			return fmt.Errorf("failed to show quote: %v", err)
		}

	case "list":
		if err := w.list(); err != nil {
			return fmt.Errorf("failed to list quotes: %v", err)
		}

	case "remove":
		if err := w.remove(rest); err != nil {
			return fmt.Errorf("failed to remove quote: %v", err)
		}

	default:
		return fmt.Errorf("unrecognized command: %v", command)
	}

	return nil
}

func (w *Wisdom) add() error {
	var (
		in  = bufio.NewReader(os.Stdin)
		str strings.Builder
		q   = quote.Quote{
			Tags: map[string]struct{}{},
		}
	)

	// Text
	fmt.Print("Text (Ctrl-D for EOF):")
	if _, err := io.Copy(&str, os.Stdin); err != nil {
		if err != io.EOF {
			return fmt.Errorf("failed to read: %v", err)
		}
	}
	q.Text = strings.TrimRightFunc(str.String(), unicode.IsSpace)
	if q.Text == "" {
		return fmt.Errorf("missing text")
	}
	str.Reset()

	// Author
	fmt.Print("Author: ")
	line, err := in.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read: %v", err)
	}
	q.Author = strings.TrimSpace(line)

	// Source
	fmt.Print("Source: ")
	line, err = in.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read: %v", err)
	}
	q.Source = strings.TrimSpace(line)

	// Tags
	fmt.Print("Tags (comma delimited): ")
	line, err = in.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read: %v", err)
	}
	for _, t := range strings.Split(strings.TrimSpace(line), ",") {
		if t == "" {
			continue
		}
		q.Tags[t] = struct{}{}
	}

	// Comment
	fmt.Print("Comment (Ctrl-D for EOF): ")
	if _, err := io.Copy(&str, os.Stdin); err != nil {
		if err != io.EOF {
			return fmt.Errorf("failed to read: %v", err)
		}
	}
	q.Comment = strings.TrimSpace(str.String())
	str.Reset()

	if _, err := w.store.Collection(coll).Create(
		uuid.NewString(), q,
	); err != nil {
		return fmt.Errorf("failed to create: %v", err)
	}

	return nil
}

func (w *Wisdom) edit(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing id")
	}

	editor := strings.Split(os.Getenv("EDITOR"), " ")
	if len(editor) == 0 {
		return fmt.Errorf("missing $EDITOR")
	}

	path, err := exec.LookPath(editor[0])
	if err != nil {
		return fmt.Errorf(
			"failed to find $EDITOR %s: %w", editor, err,
		)
	}

	for _, id := range args {
		d, err := w.store.Collection(coll).FindId(id)
		if err != nil {
			return fmt.Errorf("failed to load quote: %v", err)
		}

		temp := filepath.Join(os.TempDir(), "wisdom", id)
		if err := os.WriteFile(temp, d.Data, 0600); err != nil {
			return err
		}

		cmd := exec.Cmd{
			Path:   path,
			Args:   append(editor, temp),
			Stdin:  os.Stdin,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}
		if err := cmd.Run(); err != nil {
			return err
		}

		data, err := os.ReadFile(temp)
		if err != nil {
			return fmt.Errorf(
				"failed to read file %s: %w", temp, err,
			)
		}

		var q = &quote.Quote{}
		if err := json.Unmarshal(data, q); err != nil {
			return fmt.Errorf(
				"failed to json unmarshal %s: %w", temp, err,
			)
		}

		if err := os.Remove(temp); err != nil {
			return fmt.Errorf("failed to remove temp file: %v", err)
		}

		d, err = w.store.Collection(coll).Create(uuid.NewString(), q)
		if err != nil {
			return fmt.Errorf(
				"failed to create new document: %v", err,
			)
		}

		if err := w.store.Collection(coll).Delete(id); err != nil {
			return fmt.Errorf(
				"failed to delete original document: %v", err,
			)
		}
	}

	return nil
}

func (w *Wisdom) list() error {
	docs, err := w.store.Collection(coll).All()
	if err != nil {
		return fmt.Errorf("failed to load quotes: %v", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(tw, "Id\tText\tAuthor\tSource\t")
	for _, d := range docs {
		var q = &quote.Quote{}

		if err := d.Unmarshal(q); err != nil {
			w.log.Error.Printf(
				"failed to unmarshal document %s: %v",
				d.Id, err,
			)
			continue
		}

		var t = q.Text
		if w.wrap > 0 {
			t = text.Truncate(q.Text, w.wrap-1) + "…"
		}

		fmt.Fprintf(
			tw, "%s\t%s\t%s\t%s\t\n",
			d.Id,
			strings.ReplaceAll(t, "\n", " "),
			q.Author,
			q.Source,
		)
	}
	tw.Flush()

	return nil
}

func (w *Wisdom) remove(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing quote filename")
	}

	for _, name := range args {
		if err := w.store.Collection(coll).Delete(name); err != nil {
			return fmt.Errorf("failed to delete %s: %v", name, err)
		}
	}

	return nil
}

func (w *Wisdom) show(args []string) error {
	for _, id := range args {
		d, err := w.store.Collection(coll).FindId(id)
		if err != nil {
			return fmt.Errorf("failed to load quotes: %v", err)
		}

		var q = &quote.Quote{}
		if err := d.Unmarshal(q); err != nil {
			return fmt.Errorf(
				"failed to unmarshal document %s: %v",
				d.Id, err,
			)
		}

		fmt.Println(q.Render(w.wrap))
	}

	return nil
}
