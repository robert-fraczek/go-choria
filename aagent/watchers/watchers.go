package watchers

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/choria-io/go-choria/aagent/watchers/execwatcher"
	"github.com/choria-io/go-choria/aagent/watchers/filewatcher"
	"github.com/pkg/errors"
)

type State int

const (
	Error State = iota
	Skipped
	Unchanged
	Changed
)

// Watcher is anything that can be used to watch the system for events
type Watcher interface {
	Name() string
	Type() string
	Run(context.Context, *sync.WaitGroup)
	NotifyStateChance()
	CurrentState() map[string]interface{}
	AnnounceInterval() time.Duration
}

// Machine is a Choria Machine
type Machine interface {
	Name() string
	State() string
	Directory() string
	Transition(t string, args ...interface{}) error
	NotifyWatcherState(string, map[string]interface{})
	Watchers() []*WatcherDef
	Debugf(name string, format string, args ...interface{})
	Infof(name string, format string, args ...interface{})
	Errorf(name string, format string, args ...interface{})
}

// Manager manages all the defined watchers in a specific machine
// implements machine.WatcherManager
type Manager struct {
	watchers map[string]Watcher
	machine  Machine
	sync.Mutex
}

func New() *Manager {
	return &Manager{
		watchers: make(map[string]Watcher),
	}
}

// SetMachine supplies the machine this manager will manage
func (m *Manager) SetMachine(t interface{}) (err error) {
	machine, ok := t.(Machine)
	if !ok {
		return errors.New("supplied machine does not implement watchers.Machine")
	}

	m.machine = machine

	return nil
}

func (m *Manager) configureWatchers() error {
	for _, w := range m.machine.Watchers() {
		w.ParseAnnounceInterval()

		switch w.Type {
		case "file":
			watcher, err := filewatcher.New(m.machine, w.Name, w.StateMatch, w.FailTransition, w.SuccessTransition, w.Interval, w.announceDuration, w.Properties)
			if err != nil {
				return errors.Wrapf(err, "could not create file watcher '%s'", w.Name)
			}

			m.watchers[w.Name] = watcher

		case "exec":
			watcher, err := execwatcher.New(m.machine, w.Name, w.StateMatch, w.FailTransition, w.SuccessTransition, w.Interval, w.announceDuration, w.Properties)
			if err != nil {
				return errors.Wrapf(err, "could not create exec watcher '%s'", w.Name)
			}

			m.watchers[w.Name] = watcher

		default:
			return fmt.Errorf("unknown watcher '%s'", w.Type)
		}
	}

	return nil
}

// Run starts all the defined watchers and periodically announce
// their state based on AnnounceInterval
func (m *Manager) Run(ctx context.Context, wg *sync.WaitGroup) error {
	m.Lock()
	defer m.Unlock()

	if m.machine == nil {
		return fmt.Errorf("manager requires a machine to manage")
	}

	err := m.configureWatchers()
	if err != nil {
		return errors.Wrap(err, "could not configure watchers")
	}

	for _, watcher := range m.watchers {
		wg.Add(1)
		go watcher.Run(ctx, wg)

		if watcher.AnnounceInterval() > 0 {
			go m.announceWatcherState(ctx, watcher)
		}
	}

	return nil
}

func (m *Manager) announceWatcherState(ctx context.Context, w Watcher) {
	announceTick := time.NewTicker(w.AnnounceInterval())

	for {
		select {
		case <-announceTick.C:
			m.machine.NotifyWatcherState(w.Name(), w.CurrentState())
		case <-ctx.Done():
			return
		}
	}
}

// NotifyStateChance implements machine.WatcherManager
func (m *Manager) NotifyStateChance() {
	m.Lock()
	defer m.Unlock()

	for _, watcher := range m.watchers {
		watcher.NotifyStateChance()
	}
}