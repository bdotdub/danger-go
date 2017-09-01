package danger

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

type DSL struct {
	Git    `json:"git"`
	GitHub `json:"github"`
}

type Results struct {
	Fails     []Violation `json:"fails"`
	Warnings  []Violation `json:"warnings"`
	Messages  []Violation `json:"messages"`
	Markdowns []string    `json:"markdowns"`
}

type Violation struct {
	Message string `json:"message"`
}

func (dr *Results) Fail(msg string) {
	dr.Fails = append(dr.Fails, Violation{msg})
}

func (dr *Results) Warn(msg string) {
	dr.Warnings = append(dr.Warnings, Violation{msg})
}

func (dr *Results) Message(msg string) {
	dr.Messages = append(dr.Messages, Violation{msg})
}

func (dr *Results) Markdown(msg string) {
	dr.Markdowns = append(dr.Markdowns, msg)
}

func (dr *Results) Flush(w io.Writer) error {
	b, err := json.Marshal(dr)
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	return err
}

func Danger() (DSL, Results, error) {
	var danger DSL

	results := Results{
		Fails:     []Violation{},
		Warnings:  []Violation{},
		Messages:  []Violation{},
		Markdowns: []string{},
	}

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return danger, results, err
	}

	err = json.Unmarshal(b, &danger)
	if err != nil {
		return danger, results, err
	}

	return danger, results, nil
}
