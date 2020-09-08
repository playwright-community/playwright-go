package playwright

type Mouse struct {
	channel *Channel
}

func newMouse(channel *Channel) *Mouse {
	return &Mouse{
		channel: channel,
	}
}

func (m *Mouse) Move(x float64, y float64, options ...MouseMoveOptions) error {
	_, err := m.channel.Send("mouseMove", map[string]interface{}{
		"x": x,
		"y": y,
	}, options)
	return err
}

func (m *Mouse) Down(options ...MouseDownOptions) error {
	_, err := m.channel.Send("mouseDown", options)
	return err
}

func (m *Mouse) Up(options ...MouseUpOptions) error {
	_, err := m.channel.Send("mouseUp", options)
	return err
}

func (m *Mouse) Click(x, y float64, options ...MouseClickOptions) error {
	_, err := m.channel.Send("mouseClick", map[string]interface{}{
		"x": x,
		"y": y,
	}, options)
	return err
}

func (m *Mouse) DblClick(x, y float64, options ...MouseDblclickOptions) error {
	var option MouseDblclickOptions
	if len(options) == 1 {
		option = options[0]
	}
	return m.Click(x, y, MouseClickOptions{
		ClickCount: Int(2),
		Button:     option.Button,
		Delay:      option.Delay,
	})
}

type Keyboard struct {
	channel *Channel
}

func newKeyboard(channel *Channel) *Keyboard {
	return &Keyboard{
		channel: channel,
	}
}

func (m *Keyboard) Down(key string) error {
	_, err := m.channel.Send("keyboardDown", map[string]interface{}{
		"key": key,
	})
	return err
}

func (m *Keyboard) Up(key string) error {
	_, err := m.channel.Send("keyboardUp", map[string]interface{}{
		"key": key,
	})
	return err
}

func (m *Keyboard) InsertText(text string) error {
	_, err := m.channel.Send("keyboardInsertText", map[string]interface{}{
		"text": text,
	})
	return err
}

func (m *Keyboard) Type(text string, options ...KeyboardTypeOptions) error {
	_, err := m.channel.Send("keyboardInsertText", map[string]interface{}{
		"text": text,
	}, options)
	return err
}

func (m *Keyboard) Press(key string, options ...KeyboardPressOptions) error {
	_, err := m.channel.Send("keyboardPress", map[string]interface{}{
		"key": key,
	}, options)
	return err
}
