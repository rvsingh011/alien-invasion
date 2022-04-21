package simulation

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNewSimulation(t *testing.T) {
	type args struct {
		iterations   int
		alienNumbers int
		alienNames   string
		worldFile    string
		randomSeed   *rand.Rand
		logger       *zap.Logger
	}
	tests := []struct {
		name    string
		args    args
		want    *Simulation
		wantErr bool
	}{
		{
			name: "Create new simulations",
			args: args{
				iterations:   10,
				alienNumbers: 20,
				alienNames:   "./data/alien_names.txt",
			},
			wantErr: false,
			want:    &Simulation{Iterations: 10, AlienNames: "./data/alien_names.txt", NumberOfAliens: 20},
		},
		{
			name: "Create new simulations with world file",
			args: args{
				iterations:   10,
				alienNumbers: 20,
				alienNames:   "./data/alien_names.txt",
				worldFile:    "./data/word.txt",
			},
			wantErr: false,
			want: &Simulation{
				Iterations:     10,
				AlienNames:     "./data/alien_names.txt",
				NumberOfAliens: 20,
				WorldFile:      "./data/word.txt",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSimulation(tt.args.iterations, tt.args.alienNumbers, tt.args.alienNames, tt.args.worldFile, tt.args.randomSeed, tt.args.logger)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSimulation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, got.Iterations, tt.want.Iterations)
			assert.Equal(t, got.AlienNames, tt.want.AlienNames)
			assert.Equal(t, got.NumberOfAliens, tt.want.NumberOfAliens)
			assert.Equal(t, got.WorldFile, tt.want.WorldFile)
		})
	}
}

func TestSimulation_CreateWorld(t *testing.T) {
	type fields struct {
		Iterations       int
		WorldFile        string
		NumberOfAliens   int
		AlienNames       string
		World            map[string][]*City
		Aliens           []*Alien
		Cities           []*City
		AlienCityMapping map[string]string
		CityAlienMapping map[string][]string
		RandSeed         *rand.Rand
		logger           *zap.Logger
	}
	tests := []struct {
		name               string
		fields             fields
		wantErr            bool
		expectedCitylength int
		expectedCityMap    map[string]int
	}{
		{
			name: "Test with world example 1",
			fields: fields{
				WorldFile: "../data/world-example-1.txt",
				World:     make(map[string][]*City),
			},
			wantErr:            false,
			expectedCitylength: 6,
			expectedCityMap: map[string]int{
				"Foo":   3,
				"Bar":   3,
				"Baz":   1,
				"Qu-ux": 1,
				"Bee":   1,
				"Lee":   1,
			},
		},
		{
			name: "Test with world example 2",
			fields: fields{
				WorldFile: "../data/world-example-1.txt",
				World:     make(map[string][]*City),
			},
			wantErr:            false,
			expectedCitylength: 6,
			expectedCityMap: map[string]int{
				"Foo":   3,
				"Bar":   3,
				"Baz":   1,
				"Qu-ux": 1,
				"Bee":   1,
				"Lee":   1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sim := &Simulation{
				Iterations:       tt.fields.Iterations,
				WorldFile:        tt.fields.WorldFile,
				NumberOfAliens:   tt.fields.NumberOfAliens,
				AlienNames:       tt.fields.AlienNames,
				World:            tt.fields.World,
				Aliens:           tt.fields.Aliens,
				Cities:           tt.fields.Cities,
				AlienCityMapping: tt.fields.AlienCityMapping,
				CityAlienMapping: tt.fields.CityAlienMapping,
				RandSeed:         tt.fields.RandSeed,
				logger:           tt.fields.logger,
			}
			if err := sim.CreateWorld(); (err != nil) != tt.wantErr {
				t.Errorf("Simulation.CreateWorld() error = %v, wantErr %v", err, tt.wantErr)
			}
			// TODO: Convert into detailed assertion using automtion
			// validate the world creation
			assert.Equal(t, tt.expectedCitylength, len(sim.Cities))
			assert.Equal(t, tt.expectedCitylength, len(sim.World))

			for city, connectedCities := range tt.expectedCityMap {
				_, cityExists := sim.World[city]
				assert.Equal(t, true, cityExists)
				assert.Equal(t, connectedCities, len(sim.World[city]))
			}
		})
	}
}

func TestSimulation_CreateAliens(t *testing.T) {
	type fields struct {
		World          map[string][]*City
		Iterations     int
		WorldFile      string
		NumberOfAliens int
		AlienNamesFile string
		Aliens         []*Alien
		Cities         []*City
		RandSeed       *rand.Rand
		logger         *zap.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Create zero aliens",
			fields: fields{
				AlienNamesFile: "../data/alien_names.txt",
				NumberOfAliens: 0,
			},
			wantErr: false,
		},
		{
			name: "Create one aliens",
			fields: fields{
				AlienNamesFile: "../data/alien_names.txt",
				NumberOfAliens: 1,
			},
			wantErr: false,
		},
		{
			name: "Create ten aliens",
			fields: fields{
				AlienNamesFile: "../data/alien_names.txt",
				NumberOfAliens: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sim := &Simulation{
				World:          tt.fields.World,
				Iterations:     tt.fields.Iterations,
				WorldFile:      tt.fields.WorldFile,
				NumberOfAliens: tt.fields.NumberOfAliens,
				AlienNames:     tt.fields.AlienNamesFile,
				Aliens:         tt.fields.Aliens,
				Cities:         tt.fields.Cities,
				RandSeed:       tt.fields.RandSeed,
				logger:         tt.fields.logger,
			}
			if err := sim.CreateAliens(); (err != nil) != tt.wantErr {
				t.Errorf("Simulation.CreateAliens() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.fields.NumberOfAliens, len(tt.fields.Aliens))
			}
		})
	}
}

func TestSimulation_Start(t *testing.T) {
	type fields struct {
		Iterations       int
		WorldFile        string
		NumberOfAliens   int
		AlienNames       string
		World            map[string][]*City
		Aliens           []*Alien
		Cities           []*City
		AlienCityMapping map[string]string
		CityAlienMapping map[string][]string
		RandSeed         *rand.Rand
		logger           *zap.Logger
	}
	tests := []struct {
		name                 string
		fields               fields
		wantErr              bool
		wantAlienCityMapping map[string]string
		wantCityAlienMapping map[string][]string
	}{
		{
			name:                 "Run Single Iteration",
			wantAlienCityMapping: map[string]string{"Alien1": "Bar", "Alien3": "Lee"},
			wantCityAlienMapping: map[string][]string{
				"Bar": {"Alien1"},
				"Lee": {"Alien3"},
				"Mee": {},
			},
			fields: fields{
				World: map[string][]*City{
					"Foo": {
						NewCityWithDirections("Lee", "north"),
						NewCityWithDirections("Bar", "south"),
						NewCityWithDirections("Mee", "west"),
					},
					"Lee": {
						NewCityWithDirections("Foo", "south"),
					},
					"Bar": {
						NewCityWithDirections("Foo", "north"),
					},
					"Mee": {
						NewCityWithDirections("Foo", "east"),
					},
				},
				Iterations:       1,
				Cities:           []*City{NewCity("Foo"), NewCity("Bar"), NewCity("Lee"), NewCity("Mee")},
				RandSeed:         rand.New(rand.NewSource(3)),
				Aliens:           []*Alien{NewAlien("Alien0"), NewAlien("Alien1"), NewAlien("Alien2"), NewAlien("Alien3")},
				CityAlienMapping: make(map[string][]string),
				AlienCityMapping: make(map[string]string),
			},
			wantErr: false,
		},
		{
			name: "Run two predictable Iterations",
			wantAlienCityMapping: map[string]string{
				"Alien4": "Berlin",
				"Alien9": "Mee",
			},
			wantCityAlienMapping: map[string][]string{
				"Bar":       {},
				"Lee":       {},
				"Mee":       {"Alien9"},
				"Berlin":    {"Alien4"},
				"Moscow":    {},
				"Tokyo":     {},
				"Bangalore": {},
			},
			fields: fields{
				World: map[string][]*City{
					"Foo": {
						NewCityWithDirections("Lee", "north"),
						NewCityWithDirections("Bar", "south"),
						NewCityWithDirections("Mee", "west"),
						NewCityWithDirections("Moscow", "east"),
					},
					"Lee": {
						NewCityWithDirections("Foo", "south"),
						NewCityWithDirections("Berlin", "east"),
					},
					"Bar": {
						NewCityWithDirections("Foo", "north"),
						NewCityWithDirections("Delhi", "south"),
					},
					"Mee": {
						NewCityWithDirections("Foo", "east"),
					},
					"Delhi": {
						NewCityWithDirections("Bar", "north"),
						NewCityWithDirections("Bangalore", "east"),
					},
					"Berlin": {
						NewCityWithDirections("Lee", "west"),
						NewCityWithDirections("Moscow", "south"),
					},
					"Moscow": {
						NewCityWithDirections("Berlin", "north"),
						NewCityWithDirections("Tokyo", "south"),
					},
					"Tokyo": {
						NewCityWithDirections("Moscow", "north"),
						NewCityWithDirections("Bangalore", "south"),
					},
					"Bangalore": {
						NewCityWithDirections("Tokyo", "north"),
						NewCityWithDirections("Delhi", "west"),
					},
				},
				Iterations: 2,
				Cities: []*City{
					NewCity("Foo"),
					NewCity("Lee"),
					NewCity("Bar"),
					NewCity("Mee"),
					NewCity("Delhi"),
					NewCity("Berlin"),
					NewCity("Moscow"),
					NewCity("Tokyo"),
					NewCity("Bangalore"),
				},
				RandSeed: rand.New(rand.NewSource(3)),
				Aliens: []*Alien{
					NewAlien("Alien0"),
					NewAlien("Alien1"),
					NewAlien("Alien2"),
					NewAlien("Alien3"),
					NewAlien("Alien4"),
					NewAlien("Alien5"),
					NewAlien("Alien6"),
					NewAlien("Alien7"),
					NewAlien("Alien8"),
					NewAlien("Alien9"),
				},
				CityAlienMapping: make(map[string][]string),
				AlienCityMapping: make(map[string]string),
			},
			wantErr: false,
		},
		{
			name: "Run two predictable Iterations",
			wantAlienCityMapping: map[string]string{
				"Alien4": "Moscow",
				"Alien9": "Mee",
			},
			wantCityAlienMapping: map[string][]string{
				"Bar":       {},
				"Lee":       {},
				"Mee":       {"Alien9"},
				"Berlin":    {},
				"Moscow":    {"Alien4"},
				"Tokyo":     {},
				"Bangalore": {},
			},
			fields: fields{
				World: map[string][]*City{
					"Foo": {
						NewCityWithDirections("Lee", "north"),
						NewCityWithDirections("Bar", "south"),
						NewCityWithDirections("Mee", "west"),
						NewCityWithDirections("Moscow", "east"),
					},
					"Lee": {
						NewCityWithDirections("Foo", "south"),
						NewCityWithDirections("Berlin", "east"),
					},
					"Bar": {
						NewCityWithDirections("Foo", "north"),
						NewCityWithDirections("Delhi", "south"),
					},
					"Mee": {
						NewCityWithDirections("Foo", "east"),
					},
					"Delhi": {
						NewCityWithDirections("Bar", "north"),
						NewCityWithDirections("Bangalore", "east"),
					},
					"Berlin": {
						NewCityWithDirections("Lee", "west"),
						NewCityWithDirections("Moscow", "south"),
					},
					"Moscow": {
						NewCityWithDirections("Berlin", "north"),
						NewCityWithDirections("Tokyo", "south"),
					},
					"Tokyo": {
						NewCityWithDirections("Moscow", "north"),
						NewCityWithDirections("Bangalore", "south"),
					},
					"Bangalore": {
						NewCityWithDirections("Tokyo", "north"),
						NewCityWithDirections("Delhi", "west"),
					},
				},
				Iterations: 3,
				Cities: []*City{
					NewCity("Foo"),
					NewCity("Lee"),
					NewCity("Bar"),
					NewCity("Mee"),
					NewCity("Delhi"),
					NewCity("Berlin"),
					NewCity("Moscow"),
					NewCity("Tokyo"),
					NewCity("Bangalore"),
				},
				RandSeed: rand.New(rand.NewSource(3)),
				Aliens: []*Alien{
					NewAlien("Alien0"),
					NewAlien("Alien1"),
					NewAlien("Alien2"),
					NewAlien("Alien3"),
					NewAlien("Alien4"),
					NewAlien("Alien5"),
					NewAlien("Alien6"),
					NewAlien("Alien7"),
					NewAlien("Alien8"),
					NewAlien("Alien9"),
				},
				CityAlienMapping: make(map[string][]string),
				AlienCityMapping: make(map[string]string),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sim := &Simulation{
				Iterations:       tt.fields.Iterations,
				WorldFile:        tt.fields.WorldFile,
				NumberOfAliens:   tt.fields.NumberOfAliens,
				AlienNames:       tt.fields.AlienNames,
				World:            tt.fields.World,
				Aliens:           tt.fields.Aliens,
				Cities:           tt.fields.Cities,
				AlienCityMapping: tt.fields.AlienCityMapping,
				CityAlienMapping: tt.fields.CityAlienMapping,
				RandSeed:         tt.fields.RandSeed,
				logger:           tt.fields.logger,
			}
			if err := sim.Start(); (err != nil) != tt.wantErr {
				t.Errorf("Simulation.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.wantAlienCityMapping, sim.AlienCityMapping)
			assert.Equal(t, tt.wantCityAlienMapping, sim.CityAlienMapping)
		})
	}
}

func TestSimulation_fight(t *testing.T) {
	type fields struct {
		Iterations       int
		WorldFile        string
		NumberOfAliens   int
		AlienNames       string
		World            map[string][]*City
		Aliens           []*Alien
		Cities           []*City
		AlienCityMapping map[string]string
		CityAlienMapping map[string][]string
		RandSeed         *rand.Rand
		logger           *zap.Logger
	}
	tests := []struct {
		name            string
		fields          fields
		destroyedCity   []string
		destroyedAliens []string
	}{
		{
			name: "Test alien reach same city",
			fields: fields{
				World: map[string][]*City{
					"Foo": {
						NewCityWithDirections("Lee", "north"),
						NewCityWithDirections("Bar", "south"),
						NewCityWithDirections("Mee", "west"),
					},
					"Lee": {
						NewCityWithDirections("Foo", "south"),
					},
					"Bar": {
						NewCityWithDirections("Foo", "north"),
					},
					"Mee": {
						NewCityWithDirections("Foo", "east"),
					},
				},
				Aliens: []*Alien{NewAlien("Alien1"), NewAlien("Alien2")},
				CityAlienMapping: map[string][]string{
					"Foo": {"Alien1", "Alien2"},
					"Lee": {},
					"Bar": {},
					"Mee": {},
				},
				AlienCityMapping: map[string]string{"Alien1": "Foo", "Alien2": "Foo"},
				Cities:           []*City{NewCity("Foo"), NewCity("Lee"), NewCity("Bar"), NewCity("Mee")},
			},
			destroyedCity:   []string{"Foo"},
			destroyedAliens: []string{"Alien1", "Alien2"},
		},
		{
			name: "Test alien reach different city",
			fields: fields{
				World: map[string][]*City{
					"Foo": {
						NewCityWithDirections("Lee", "north"),
						NewCityWithDirections("Bar", "south"),
						NewCityWithDirections("Mee", "west"),
					},
					"Lee": {
						NewCityWithDirections("Foo", "south"),
					},
					"Bar": {
						NewCityWithDirections("Foo", "north"),
					},
					"Mee": {
						NewCityWithDirections("Foo", "east"),
					},
				},
				Aliens: []*Alien{NewAlien("Alien1"), NewAlien("Alien2")},
				CityAlienMapping: map[string][]string{
					"Foo": {"Alien1"},
					"Lee": {"Alien2"},
					"Bar": {},
					"Mee": {},
				},
				AlienCityMapping: map[string]string{"Alien1": "Foo", "Alien2": "Lee"},
				Cities:           []*City{NewCity("Foo"), NewCity("Lee"), NewCity("Bar"), NewCity("Mee")},
			},
			destroyedCity:   []string{},
			destroyedAliens: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sim := &Simulation{
				Iterations:       tt.fields.Iterations,
				WorldFile:        tt.fields.WorldFile,
				NumberOfAliens:   tt.fields.NumberOfAliens,
				AlienNames:       tt.fields.AlienNames,
				World:            tt.fields.World,
				Aliens:           tt.fields.Aliens,
				Cities:           tt.fields.Cities,
				AlienCityMapping: tt.fields.AlienCityMapping,
				CityAlienMapping: tt.fields.CityAlienMapping,
				RandSeed:         tt.fields.RandSeed,
				logger:           tt.fields.logger,
			}
			sim.fight()
			if len(tt.destroyedCity) == 0 {
				for alien, city := range tt.fields.AlienCityMapping {
					assert.Contains(t, sim.World, city)
					assert.Contains(t, sim.AlienCityMapping, alien)
				}
			} else {
				for _, city := range tt.destroyedCity {
					assert.NotContains(t, sim.Cities, city)
					assert.NotContains(t, sim.World, city)
				}
				for _, alien := range tt.destroyedAliens {
					assert.NotContains(t, sim.AlienCityMapping, alien)
				}
			}

		})
	}
}

func TestSimulation_burryDeadAliens(t *testing.T) {
	type fields struct {
		Iterations       int
		WorldFile        string
		NumberOfAliens   int
		AlienNames       string
		World            map[string][]*City
		Aliens           []*Alien
		Cities           []*City
		AlienCityMapping map[string]string
		CityAlienMapping map[string][]string
		RandSeed         *rand.Rand
		logger           *zap.Logger
	}
	type args struct {
		deadAliens []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test no Alien dead",
			fields: fields{
				Aliens:           []*Alien{NewAlien("Alien1"), NewAlien("Alien2")},
				AlienCityMapping: map[string]string{"Alien1": "Foo", "Alien2": "Lee"},
			},
		},
		{
			name: "Test alien 1 is dead",
			fields: fields{
				Aliens:           []*Alien{NewAlien("Alien1"), NewAlien("Alien2")},
				AlienCityMapping: map[string]string{"Alien1": "Foo", "Alien2": "Lee"},
			},
			args: args{deadAliens: []string{"Alien1"}},
		},
		{
			name: "Test alien 2 is dead",
			fields: fields{
				Aliens:           []*Alien{NewAlien("Alien1"), NewAlien("Alien2")},
				AlienCityMapping: map[string]string{"Alien1": "Foo", "Alien2": "Lee"},
			},
			args: args{deadAliens: []string{"Alien2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sim := &Simulation{
				Iterations:       tt.fields.Iterations,
				WorldFile:        tt.fields.WorldFile,
				NumberOfAliens:   tt.fields.NumberOfAliens,
				AlienNames:       tt.fields.AlienNames,
				World:            tt.fields.World,
				Aliens:           tt.fields.Aliens,
				Cities:           tt.fields.Cities,
				AlienCityMapping: tt.fields.AlienCityMapping,
				CityAlienMapping: tt.fields.CityAlienMapping,
				RandSeed:         tt.fields.RandSeed,
				logger:           tt.fields.logger,
			}
			sim.burryDeadAliens(tt.args.deadAliens)
			// validate the deletion of aliens
			for idx := range tt.args.deadAliens {
				// check alien removed fro city mapping
				assert.NotContains(t, sim.AlienCityMapping, tt.args.deadAliens[idx])
				// check dead alien should not exist in the alien list
				assert.NotContains(t, sim.Aliens, tt.args.deadAliens[idx])
			}
		})
	}
}

func TestSimulation_removeDestroyedCities(t *testing.T) {
	type fields struct {
		Iterations       int
		WorldFile        string
		NumberOfAliens   int
		AlienNames       string
		World            map[string][]*City
		Aliens           []*Alien
		Cities           []*City
		AlienCityMapping map[string]string
		CityAlienMapping map[string][]string
		RandSeed         *rand.Rand
		logger           *zap.Logger
	}
	type args struct {
		destoyedCities []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test ",
			fields: fields{
				World: map[string][]*City{
					"Foo": {
						NewCityWithDirections("Lee", "north"),
						NewCityWithDirections("Bar", "south"),
						NewCityWithDirections("Mee", "west"),
					},
					"Lee": {
						NewCityWithDirections("Foo", "south"),
					},
					"Bar": {
						NewCityWithDirections("Foo", "north"),
					},
					"Mee": {
						NewCityWithDirections("Foo", "east"),
					},
				},
				Cities: []*City{NewCity("Foo"), NewCity("Lee"), NewCity("Bar"), NewCity("Mee")},
				CityAlienMapping: map[string][]string{
					"Foo": {},
					"Lee": {},
					"Bar": {},
					"Mee": {},
				},
			},
			args: args{[]string{"Foo", "Bar"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sim := &Simulation{
				Iterations:       tt.fields.Iterations,
				WorldFile:        tt.fields.WorldFile,
				NumberOfAliens:   tt.fields.NumberOfAliens,
				AlienNames:       tt.fields.AlienNames,
				World:            tt.fields.World,
				Aliens:           tt.fields.Aliens,
				Cities:           tt.fields.Cities,
				AlienCityMapping: tt.fields.AlienCityMapping,
				CityAlienMapping: tt.fields.CityAlienMapping,
				RandSeed:         tt.fields.RandSeed,
				logger:           tt.fields.logger,
			}
			sim.removeDestroyedCities(tt.args.destoyedCities)
			for _, city := range tt.args.destoyedCities {
				// city should be removed from the ciy mapping
				assert.NotContains(t, sim.CityAlienMapping, city)
				assert.NotContains(t, sim.Cities, city)
				assert.NotContains(t, sim.World, city)
				for _, Links := range sim.World {
					assert.NotContains(t, Links, city)
				}
			}

		})
	}
}

func TestSimulation_deleteCityFromWorldMap(t *testing.T) {
	type fields struct {
		Iterations       int
		WorldFile        string
		NumberOfAliens   int
		AlienNames       string
		World            map[string][]*City
		Aliens           []*Alien
		Cities           []*City
		AlienCityMapping map[string]string
		CityAlienMapping map[string][]string
		RandSeed         *rand.Rand
		logger           *zap.Logger
	}
	type args struct {
		city string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test delete connected city from world map",
			fields: fields{
				World: map[string][]*City{
					"Foo": {
						NewCityWithDirections("Lee", "north"),
						NewCityWithDirections("Bar", "south"),
						NewCityWithDirections("Mee", "west"),
					},
					"Lee": {
						NewCityWithDirections("Foo", "south"),
					},
					"Bar": {
						NewCityWithDirections("Foo", "north"),
					},
					"Mee": {
						NewCityWithDirections("Foo", "east"),
					},
				},
			},
			args: args{"Foo"},
		},
		{
			name: "Test delete sparse connected city from world map",
			fields: fields{
				World: map[string][]*City{
					"Foo": {
						NewCityWithDirections("Lee", "north"),
						NewCityWithDirections("Bar", "south"),
						NewCityWithDirections("Mee", "west"),
					},
					"Lee": {
						NewCityWithDirections("Foo", "south"),
					},
					"Bar": {
						NewCityWithDirections("Foo", "north"),
					},
					"Mee": {
						NewCityWithDirections("Foo", "east"),
					},
				},
			},
			args: args{"Lee"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sim := &Simulation{
				Iterations:       tt.fields.Iterations,
				WorldFile:        tt.fields.WorldFile,
				NumberOfAliens:   tt.fields.NumberOfAliens,
				AlienNames:       tt.fields.AlienNames,
				World:            tt.fields.World,
				Aliens:           tt.fields.Aliens,
				Cities:           tt.fields.Cities,
				AlienCityMapping: tt.fields.AlienCityMapping,
				CityAlienMapping: tt.fields.CityAlienMapping,
				RandSeed:         tt.fields.RandSeed,
				logger:           tt.fields.logger,
			}
			sim.deleteCityFromWorldMap(tt.args.city)
			// delete the city and then
			assert.NotContains(t, sim.World, tt.args.city)
			for _, Links := range sim.World {
				assert.NotContains(t, Links, tt.args.city)
			}
		})
	}
}

func TestSimulation_isNextIterationRequired(t *testing.T) {
	type fields struct {
		Iterations       int
		WorldFile        string
		NumberOfAliens   int
		AlienNames       string
		World            map[string][]*City
		Aliens           []*Alien
		Cities           []*City
		AlienCityMapping map[string]string
		CityAlienMapping map[string][]string
		RandSeed         *rand.Rand
		logger           *zap.Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Next round is required",
			fields: fields{Aliens: make([]*Alien, 10), Cities: make([]*City, 10)},
			want:   true,
		},
		{
			name:   "Next round is not required",
			fields: fields{Aliens: make([]*Alien, 0), Cities: make([]*City, 10)},
			want:   false,
		},
		{
			name:   "Next round is not required",
			fields: fields{Aliens: make([]*Alien, 10), Cities: make([]*City, 0)},
			want:   false,
		},
		{
			name:   "Next round is not required",
			fields: fields{Aliens: make([]*Alien, 0), Cities: make([]*City, 0)},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sim := &Simulation{
				Iterations:       tt.fields.Iterations,
				WorldFile:        tt.fields.WorldFile,
				NumberOfAliens:   tt.fields.NumberOfAliens,
				AlienNames:       tt.fields.AlienNames,
				World:            tt.fields.World,
				Aliens:           tt.fields.Aliens,
				Cities:           tt.fields.Cities,
				AlienCityMapping: tt.fields.AlienCityMapping,
				CityAlienMapping: tt.fields.CityAlienMapping,
				RandSeed:         tt.fields.RandSeed,
				logger:           tt.fields.logger,
			}
			if got := sim.isNextIterationRequired(); got != tt.want {
				t.Errorf("Simulation.isNextIterationRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimulation_prepareAttack(t *testing.T) {
	type fields struct {
		Iterations       int
		WorldFile        string
		NumberOfAliens   int
		AlienNames       string
		World            map[string][]*City
		Aliens           []*Alien
		Cities           []*City
		AlienCityMapping map[string]string
		CityAlienMapping map[string][]string
		RandSeed         *rand.Rand
		logger           *zap.Logger
	}
	tests := []struct {
		name     string
		fields   fields
		wantCity []string
	}{
		{
			name: "Test-1",
			fields: fields{
				Cities:           []*City{NewCity("Foo"), NewCity("Bar"), NewCity("Lee"), NewCity("Mee")},
				RandSeed:         rand.New(rand.NewSource(3)),
				Aliens:           []*Alien{NewAlien("Alien0"), NewAlien("Alien1"), NewAlien("Alien2"), NewAlien("Alien3")},
				CityAlienMapping: make(map[string][]string),
				AlienCityMapping: make(map[string]string),
			},
			wantCity: []string{"Foo", "Bar", "Foo", "Lee"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sim := &Simulation{
				Iterations:       tt.fields.Iterations,
				WorldFile:        tt.fields.WorldFile,
				NumberOfAliens:   tt.fields.NumberOfAliens,
				AlienNames:       tt.fields.AlienNames,
				World:            tt.fields.World,
				Aliens:           tt.fields.Aliens,
				Cities:           tt.fields.Cities,
				AlienCityMapping: tt.fields.AlienCityMapping,
				CityAlienMapping: tt.fields.CityAlienMapping,
				RandSeed:         tt.fields.RandSeed,
				logger:           tt.fields.logger,
			}
			sim.prepareAttack()
			// validate if right city are choosen by the alien
			// with random seed as 3, the order should be 0, 1, 0, 2
			for idx := range sim.Aliens {
				assert.Equal(t, tt.wantCity[idx], sim.AlienCityMapping[sim.Aliens[idx].Name])
			}
		})
	}
}

func TestSimulation_runNextRoundOfAttack(t *testing.T) {
	type fields struct {
		Iterations       int
		WorldFile        string
		NumberOfAliens   int
		AlienNames       string
		World            map[string][]*City
		Aliens           []*Alien
		Cities           []*City
		AlienCityMapping map[string]string
		CityAlienMapping map[string][]string
		RandSeed         *rand.Rand
		logger           *zap.Logger
	}
	tests := []struct {
		name               string
		fields             fields
		alienNextCity      []string
		alienCurrentCities []string
	}{
		{
			name: "Test alien moves",
			fields: fields{
				Aliens: []*Alien{NewAlien("Alien1"), NewAlien("Alien2"), NewAlien("Alien3")},
				World: map[string][]*City{
					"Foo": {
						NewCityWithDirections("Lee", "north"),
						NewCityWithDirections("Bar", "south"),
						NewCityWithDirections("Mee", "west"),
					},
				},
				AlienCityMapping: map[string]string{"Alien1": "Foo", "Alien2": "Foo", "Alien3": "Foo"},
				CityAlienMapping: map[string][]string{
					"Foo": {"Alien1", "Alien2", "Alien3"},
					"Lee": {},
					"Bar": {},
					"Mee": {},
				},
				RandSeed: rand.New(rand.NewSource(3)),
			},
			// since the random seed 3 with max range 3 generate number in sequence 0, 1, 0, 2, next city can be predicted
			alienNextCity:      []string{"Lee", "Bar", "Lee", "Mee"},
			alienCurrentCities: []string{"Foo", "Foo", "Foo"},
		},
		{
			name: "Test alien moves and some stays",
			fields: fields{
				Aliens: []*Alien{NewAlien("Alien1"), NewAlien("Alien2"), NewAlien("Alien3"), NewAlien("Alien4")},
				World: map[string][]*City{
					"Foo": {
						NewCityWithDirections("Lee", "north"),
						NewCityWithDirections("Bar", "south"),
						NewCityWithDirections("Mee", "west"),
					},
				},
				AlienCityMapping: map[string]string{"Alien1": "Foo", "Alien2": "Foo", "Alien3": "Foo", "Alien4": "Foo"},
				CityAlienMapping: map[string][]string{
					"Foo": {"Alien1", "Alien2", "Alien3", "Alien4"},
					"Lee": {},
					"Bar": {},
					"Mee": {},
				},
				RandSeed: rand.New(rand.NewSource(4)),
			},
			// since the random seed 4 with max range 3 generate number in sequence 0, 1, 0, 2, next city can be predicted
			alienNextCity:      []string{"Bar", "Lee", "Bar", "Foo"},
			alienCurrentCities: []string{"Foo", "Foo", "Foo", "Foo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sim := &Simulation{
				Iterations:       tt.fields.Iterations,
				WorldFile:        tt.fields.WorldFile,
				NumberOfAliens:   tt.fields.NumberOfAliens,
				AlienNames:       tt.fields.AlienNames,
				World:            tt.fields.World,
				Aliens:           tt.fields.Aliens,
				Cities:           tt.fields.Cities,
				AlienCityMapping: tt.fields.AlienCityMapping,
				CityAlienMapping: tt.fields.CityAlienMapping,
				RandSeed:         tt.fields.RandSeed,
				logger:           tt.fields.logger,
			}
			sim.runNextRoundOfAttack()
			for idx, alien := range tt.fields.Aliens {
				// Check if alien reached the next city as expected
				assert.Contains(t, sim.CityAlienMapping[tt.alienNextCity[idx]], alien.Name)
				if tt.alienNextCity[idx] != tt.alienCurrentCities[idx] {
					assert.NotContains(t, sim.CityAlienMapping[tt.alienCurrentCities[idx]], alien.Name)
				}
			}
		})
	}
}

func TestSimulation_EndAndConclude(t *testing.T) {
	type fields struct {
		Iterations       int
		WorldFile        string
		NumberOfAliens   int
		AlienNames       string
		World            map[string][]*City
		Aliens           []*Alien
		Cities           []*City
		AlienCityMapping map[string]string
		CityAlienMapping map[string][]string
		RandSeed         *rand.Rand
		logger           *zap.Logger
	}
	tests := []struct {
		name       string
		fields     fields
		wantString string
	}{
		{
			name: "Test1",
			fields: fields{
				World: map[string][]*City{
					"Foo": {NewCityWithDirections("Mee", "north"), NewCityWithDirections("Lee", "south")},
				},
			},
			wantString: "Foo north=Mee south=Lee\n",
		},
		{
			name: "Test2",
			fields: fields{
				World: map[string][]*City{},
			},
			wantString: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sim := &Simulation{
				Iterations:       tt.fields.Iterations,
				WorldFile:        tt.fields.WorldFile,
				NumberOfAliens:   tt.fields.NumberOfAliens,
				AlienNames:       tt.fields.AlienNames,
				World:            tt.fields.World,
				Aliens:           tt.fields.Aliens,
				Cities:           tt.fields.Cities,
				AlienCityMapping: tt.fields.AlienCityMapping,
				CityAlienMapping: tt.fields.CityAlienMapping,
				RandSeed:         tt.fields.RandSeed,
				logger:           tt.fields.logger,
			}
			assert.Equal(t, tt.wantString, sim.EndAndConclude())
		})
	}
}
